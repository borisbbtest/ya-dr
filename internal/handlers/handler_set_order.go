package handlers

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

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

	go hook.calculateLoyaltySystem(orderNumber)

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("new order number accepted for processing"))

	log.Print(res)
	log.Println("Post handler")
}

func (hook *WrapperHandler) calculateLoyaltySystem(orderNumber string) {
	link := fmt.Sprintf("%s/api/orders/%s", hook.ServerConf.AccrualSystemAddress, orderNumber)
	log.Info("calculateLoyaltySystem", link)

	req, err := http.NewRequest(http.MethodGet, link, nil)
	if err != nil {
		return
	}
	for {
		bytes, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Errorf("error get data: %v", err)
		}

		if bytes.StatusCode == http.StatusTooManyRequests {
			time.Sleep(60 * time.Millisecond)
		}
		if bytes.StatusCode != http.StatusOK {
			bytes.Body.Close()
			continue
		}
		log.Info(" calculateLoyaltySystem again ---  ", bytes.Status, bytes.Header)
		var order *model.DataOrder
		//	b, err := io.ReadAll(bytes.Body)
		//	log.Info("J_____", string(b))

		if err := json.NewDecoder(bytes.Body).Decode(&order); err != nil {
			log.Errorf("calculateLoyaltySystem  -  error decoding message: %v", err)
			bytes.Body.Close()
			continue
		}
		log.Info(order)
		bytes.Body.Close()
		if _, err := hook.Storage.UpdateOrder(order); err != nil {
			log.Errorf("calculateLoyaltySystem  -  DB : %v", err)
			continue
		}
		if order.Status == "PROCESSED" || order.Status == "INVALID" {
			return
		}
	}

}
