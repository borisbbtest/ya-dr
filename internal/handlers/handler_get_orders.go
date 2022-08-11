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
		w.Write([]byte("Внутренняя ошибка сервера;"))
		return
	}

	arrOrders, err := hook.Storage.GetOrders(currentPerson)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if len(arrOrders) < 0 {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	log.Info(arrOrders)
	json.NewEncoder(w).Encode(arrOrders)

	log.Println("Post handler")
}
