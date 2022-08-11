package handlers

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/borisbbtest/ya-dr/internal/model"
	"github.com/borisbbtest/ya-dr/internal/tools"
)

func (hook *WrapperHandler) PostOrderHandler(w http.ResponseWriter, r *http.Request) {

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

	defer r.Body.Close()

	orderNumber := string(bytes)

	if !tools.IsValid(orderNumber) {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("order number has already been uploaded by this user"))
		return
	}
	currentPerson, err := tools.GetLogin(r, hook.Session)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error;"))
		return
	}

	var order = model.DataOrder{
		Number: orderNumber,
		Status: model.StatusNew,
		Person: currentPerson,
	}

	res, err := hook.Storage.PutOrder(order)
	if res == currentPerson {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("order number has already been uploaded by this user"))
		return
	}

	if err == nil && res != currentPerson {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("the order number has already been uploaded by another user"))
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("new order number accepted for processing"))

	//	go hook.calculateLoyaltySystem(orderNumber)

	log.Print(res)
	log.Println("Post handler")
}

func (hook *WrapperHandler) calculateLoyaltySystem(orderNumber string) {
	link := fmt.Sprintf("%s/api/orders/%s", hook.ServerConf.AccrualSystemAddress, orderNumber)
	req, err := http.NewRequest(http.MethodGet, link, nil)
	if err != nil {
		return
	}

	bytes, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Errorf("error get data: %v", err)
	}
	var order *model.DataOrder
	if err := json.NewDecoder(bytes.Body).Decode(&order); err != nil {
		log.Errorf("error decoding message: %v", err)
		return
	}
	defer bytes.Body.Close()
	// if err := s.Orders().UpdateStatus(order); err != nil {
	// 	return
	// }
}
