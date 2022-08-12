package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/borisbbtest/ya-dr/internal/tools"
)

func (hook *WrapperHandler) GetJSONBalanceHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	currentPerson, err := tools.GetLogin(r, hook.Session)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error;"))
		return
	}

	balanceAccrual, err := hook.Storage.GetBalance(currentPerson)
	if err != nil {
		// w.WriteHeader(http.StatusInternalServerError)
		// return
		log.Error("GetJSONBalanceHandler - ", err)
	}

	log.Info("----", *balanceAccrual.CurrentAccrual, "----", *balanceAccrual.Withdrawn)

	if err := json.NewEncoder(w).Encode(balanceAccrual); err != nil {
		log.Info(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
