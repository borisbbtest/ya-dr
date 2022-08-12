package handlers

import (
	"encoding/json"
	"net/http"
	"sort"

	"github.com/borisbbtest/ya-dr/internal/tools"
)

func (hook *WrapperHandler) GetJSONWithdrawalsHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	currentPerson, err := tools.GetLogin(r, hook.Session)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Error"))
		return
	}

	withdrawals, err := hook.Storage.GetWithdrawals(currentPerson)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if len(withdrawals) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	sort.Slice(withdrawals, func(i, j int) bool {
		return withdrawals[i].ProcessedAt.Before(withdrawals[j].ProcessedAt)
	})

	log.Info(withdrawals)
	if err := json.NewEncoder(w).Encode(withdrawals); err != nil {
		log.Info(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	log.Println("Post handler")
}
