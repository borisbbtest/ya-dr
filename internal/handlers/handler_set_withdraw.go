package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/borisbbtest/ya-dr/internal/model"
	"github.com/borisbbtest/ya-dr/internal/tools"
)

func (hook *WrapperHandler) GetJSONWithdrawHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}

	defer r.Body.Close()

	var m model.Wallet
	if err := json.Unmarshal(bytes, &m); err != nil {
		log.Errorf("body error: %v", string(bytes))
		log.Errorf("error decoding message: %v", err)
		http.Error(w, "request body is not valid json", 400)
		return
	}

	if !tools.IsValid(m.Order) {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("order number has already been uploaded by this user"))
		return
	}

	currentPerson, err := tools.GetLogin(r, hook.Session)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error;"))
		return
	}

	balance, err := hook.Storage.GetBalance(currentPerson)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if m.Sum < 0 {
		w.WriteHeader(http.StatusBadRequest)
	}

	if *balance.CurrentAccrual < m.Sum {
		w.WriteHeader(http.StatusPaymentRequired)
		return
	}

	m.Person = currentPerson
	withdrawn := m.Sum
	currentaccrual := m.Sum * -1.0
	if _, err := hook.Storage.PutWithdraw(m); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	log.Info("GetJSONWithdrawHandler ---->", withdrawn)
	if _, err := hook.Storage.UpdateBalance(&model.DataBalance{
		Withdrawn:      &withdrawn,
		CurrentAccrual: &currentaccrual,
		Person:         currentPerson,
	}); err != nil {
		log.Info("GetJSONWithdrawHandler", err)
	}
	w.WriteHeader(http.StatusOK)
}
