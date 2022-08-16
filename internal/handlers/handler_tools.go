package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/borisbbtest/ya-dr/internal/model"
	"github.com/borisbbtest/ya-dr/internal/storage"
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
				_, err = hook.Storage.UpdateBalance(&model.DataBalance{
					Person:         currentUser,
					CurrentAccrual: order.Accrual,
					Withdrawn:      &x,
				})
				if err != nil {
					log.Error(err)
				}
			}
			//hook.Storage.PutWithdraw(model.Wallet{Person: currentUser, Order: orderNumber, Sum: *order.Accrual})
			return
		}
	}

}

const (
	AuthCookieKey = "SystemLoyalty"
)

var (
	ErrUnauthorized = errors.New("user is unauthorized")
)

func AddCookie(w http.ResponseWriter, r *http.Request, name, value string, ttl time.Duration) (res time.Time, err error) {
	expire := time.Now().Add(ttl)
	cookie := http.Cookie{
		Name:    name,
		Value:   value,
		Expires: expire,
	}
	http.SetCookie(w, &cookie)
	return expire, err
}

func IsUserAuthed(r *http.Request, session *storage.SessionHTTP) bool {
	cookie, err := r.Cookie(AuthCookieKey)
	if err != nil {
		return false
	}
	if session.DBSession[cookie.Value].SessionExpiredAt.Before(time.Now()) {
		return false
	}
	return true
}

func GetLogin(r *http.Request, session *storage.SessionHTTP) (int, error) {
	if IsUserAuthed(r, session) {
		cookie, err := r.Cookie(AuthCookieKey)
		if err != nil {
			log.Error(err)
		}
		return session.DBSession[cookie.Value].ID, nil
	}
	return -1, ErrUnauthorized
}

func GetMutex(r *http.Request, session *storage.SessionHTTP) (*sync.Mutex, error) {
	if IsUserAuthed(r, session) {
		cookie, err := r.Cookie(AuthCookieKey)
		if err != nil {
			log.Error(err)
		}
		return session.DBSession[cookie.Value].LocalMutex, nil
	}
	return nil, ErrUnauthorized
}
