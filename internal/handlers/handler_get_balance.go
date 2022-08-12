package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/borisbbtest/ya-dr/internal/tools"
)

type responseJSON struct {
	Current   float32 `json:"current"`
	Withdrawn int     `json:"withdrawn"`
}

func (hook *WrapperHandler) GetJSONBalanceHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

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

	withdrawn, err := hook.Storage.GetWithdrawCount(currentPerson)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response := &responseJSON{
		Current:   balance,
		Withdrawn: withdrawn,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Info(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
