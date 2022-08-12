package handlers

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/borisbbtest/ya-dr/internal/model"
	"github.com/borisbbtest/ya-dr/internal/tools"
)

func (hook *WrapperHandler) GetJSONWithdrawHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	var reader io.Reader

	if r.Header.Get(`Content-Encoding`) == `gzip` {
		gz, err := gzip.NewReader(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		reader = gz
		defer gz.Close()
	} else {
		reader = r.Body
	}

	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatalln(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}
	orderNumber := string(bytes)

	if !tools.IsValid(orderNumber) {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("order number has already been uploaded by this user"))
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

	if balance < m.Sum {
		w.WriteHeader(http.StatusPaymentRequired)
		return
	}

	m.Person = currentPerson
	m.Sum = m.Sum * -1
	if _, err := hook.Storage.PutWithdraw(m); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}