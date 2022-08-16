package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/borisbbtest/ya-dr/internal/model"
)

func (hook *WrapperHandler) calculateLoyaltySystem(orderNumber string, currentUser int) {
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

		log.Info(" --- ", order, " --- ", *order.Accrual)

		order.Number = orderNumber
		bytes.Body.Close()
		log.Info(order)
		if _, err := hook.Storage.UpdateOrder(order); err != nil {
			log.Errorf("calculateLoyaltySystem  -  DB : %v", err)
			continue
		}
		if order.Status == "PROCESSED" || order.Status == "INVALID" {
			if *order.Accrual > 0 {
				x := float32(0)
				hook.Storage.UpdateBalance(&model.DataBalance{
					Person:         &currentUser,
					CurrentAccrual: order.Accrual,
					Withdrawn:      &x,
				})
			}
			//hook.Storage.PutWithdraw(model.Wallet{Person: currentUser, Order: orderNumber, Sum: *order.Accrual})
			return
		}
	}

}
