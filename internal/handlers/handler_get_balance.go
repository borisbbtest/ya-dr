package handlers

import (
	"encoding/json"
	"net/http"
)

func (hook *WrapperHandler) GetJSONBalanceHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	currentPerson, err := GetLogin(r, hook.Session)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error;"))
		return
	}
	log.Info(currentPerson)
	balanceAccrual, err := hook.Storage.GetBalance(currentPerson)
	if err != nil {
		// w.WriteHeader(http.StatusInternalServerError)
		// return
		log.Error("GetJSONBalanceHandler - ", err)
	}

	// if balanceAccrual.Withdrawn == nil {
	// 	*balanceAccrual.Withdrawn = 0
	// }

	log.Info("----", *balanceAccrual.CurrentAccrual, "----", balanceAccrual.Withdrawn)

	if err := json.NewEncoder(w).Encode(balanceAccrual); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Error"))
		log.Info(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
