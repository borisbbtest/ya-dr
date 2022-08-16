package handlers

import (
	"encoding/json"
	"net/http"
)

func (hook *WrapperHandler) GetJSONOrdersHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	currentPerson, err := GetLogin(r, hook.Session)
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
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// sort.Slice(arrOrders, func(i, j int) bool {
	// 	return arrOrders[i].UploadedAt.Before(arrOrders[j].UploadedAt)
	// })

	log.Info("PostOrderHandler  ", arrOrders)
	if err := json.NewEncoder(w).Encode(arrOrders); err != nil {
		log.Info(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	log.Println("Post handler")
}
