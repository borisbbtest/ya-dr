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
		w.Write([]byte("Неверный формат номера заказа;"))
		return
	}
	currentPerson, err := tools.GetLogin(r, hook.Session)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Внутренняя ошибка сервера;"))
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
		w.Write([]byte("номер заказа уже был загружен этим пользователем"))
		return
	}

	if err == nil && res != currentPerson {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("номер заказа уже был загружен другим пользователем"))
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("новый номер заказа принят в обработку"))

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

	// if err := s.Orders().UpdateStatus(order); err != nil {
	// 	return
	// }
}
