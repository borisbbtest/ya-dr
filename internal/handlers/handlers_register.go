package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
	user, _ := hook.Storage.PutUser(m)
	if user != "" {
		w.WriteHeader(http.StatusConflict)
	} else {
		u, _ := hook.Storage.GetUser(m)
		tmp, _ := tools.GetID()
		str := fmt.Sprintf("%x", tmp)
		time, _ := tools.AddCookie(w, r, tools.AuthCookieKey, str, 30*time.Minute)
		u.SessionExpiredAt = time
		hook.Session.DBSession[str] = u
		w.WriteHeader(http.StatusOK)
	}

	log.Println("Post handler", user)
}
