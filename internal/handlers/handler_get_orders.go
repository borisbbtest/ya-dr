package handlers

import (
	"encoding/json"
	"net/http"

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
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if len(arrOrders) == 0 {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	log.Info(arrOrders)
	if err := json.NewEncoder(w).Encode(arrOrders); err != nil {
		log.Info(err)
		return
	}
	log.Println("Post handler")
}
