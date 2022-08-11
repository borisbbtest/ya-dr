package handlers

import (
	"encoding/json"
	"net/http"
	"sort"

	"github.com/borisbbtest/ya-dr/internal/tools"
)

func (hook *WrapperHandler) GetJSONOrdersHandler(w http.ResponseWriter, r *http.Request) {

	currentPerson, err := tools.GetLogin(r, hook.Session)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Error"))
		return
	}

	arrOrders, err := hook.Storage.GetOrders(currentPerson)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if len(arrOrders) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	sort.Slice(arrOrders, func(i, j int) bool {
		return arrOrders[i].UploadedAt.Before(arrOrders[j].UploadedAt)
	})

	log.Info(arrOrders)
	if err := json.NewEncoder(w).Encode(arrOrders); err != nil {
		log.Info("Error ----- ", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	log.Println("Post handler")
}
