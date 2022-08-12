package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/borisbbtest/ya-dr/internal/tools"
)

type responseJSON struct {
	Current   float32 `json:"current"`
	Withdrawn float32 `json:"withdrawn"`
}

func (hook *WrapperHandler) GetJSONBalanceHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	currentPerson, err := tools.GetLogin(r, hook.Session)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error;"))
		return
	}

	balanceAccrual, err := hook.Storage.GetBalanceAccrual(currentPerson)
	if err != nil {
		// w.WriteHeader(http.StatusInternalServerError)
		// return
		log.Error("GetJSONBalanceHandler - ", err)
	}

	balanceWallet, err := hook.Storage.GetBalanceWallet(currentPerson)
	if err != nil {
		// w.WriteHeader(http.StatusInternalServerError)
		// return
		log.Error("GetJSONBalanceHandler - ", err)
	}

	//sumBalance := balanceAccrual - balanceWallet
	withdrawn, err := hook.Storage.GetWithdrawCount(currentPerson)
	if err != nil {
		log.Error("GetJSONBalanceHandler - ", err)
		//w.WriteHeader(http.StatusInternalServerError)
		//return
	}
	log.Info(balanceAccrual, "----", balanceWallet, "----", withdrawn)
	response := &responseJSON{
		Current:   balanceAccrual,
		Withdrawn: balanceWallet,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Info(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
