package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"github.com/borisbbtest/ya-dr/internal/model"
	"github.com/borisbbtest/ya-dr/internal/tools"
)

func (hook *WrapperHandler) PostJSONRegisterHandler(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Info("PostJSONHandler")
	defer r.Body.Close()

	var m model.DataUser
	if err := json.Unmarshal(bytes, &m); err != nil {
		log.Errorf("body error: %v", string(bytes))
		log.Errorf("error decoding message: %v", err)
		http.Error(w, "request body is not valid json", 400)
		return
	}

	// 200 — пользователь успешно зарегистрирован и аутентифицирован;
	// 400 — неверный формат запроса;
	// 409 — логин уже занят;
	// 500 — внутренняя ошибка сервера.
	m.Password, err = tools.HashPassword(m.Password)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error;"))
		return
	}
	user, err := hook.Storage.PutUser(m)
	if err != nil {
		log.Error(err)
	}
	if user != "" {
		w.WriteHeader(http.StatusConflict)
		return
	} else {
		u, err := hook.Storage.GetUser(m)
		if err != nil {
			log.Error(err)
		}
		tmp, err := tools.GetID()
		if err != nil {
			log.Error(err)
		}
		str := fmt.Sprintf("%x", tmp)
		time, err := AddCookie(w, r, AuthCookieKey, str, 30*time.Minute)
		if err != nil {
			log.Error(err)
		}
		u.SessionExpiredAt = time
		u.LocalMutex = &sync.Mutex{}
		hook.Session.DBSession[str] = u
		w.WriteHeader(http.StatusOK)
	}

	log.Println("Post handler", user)
}
