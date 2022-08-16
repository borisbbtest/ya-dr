package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/borisbbtest/ya-dr/internal/config"
	"github.com/borisbbtest/ya-dr/internal/model"
	"github.com/borisbbtest/ya-dr/internal/storage"
	"github.com/borisbbtest/ya-dr/internal/tools"
	"github.com/sirupsen/logrus"
)

var log = logrus.WithField("context", "system_loyalty")

type WrapperHandler struct {
	ServerConf *config.MainConfig
	Storage    storage.Storage
	Session    *storage.SessionHTTP
}

func (hook *WrapperHandler) PostJSONLoginHandler(w http.ResponseWriter, r *http.Request) {

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error;"))
		return
	}
	log.Info("PostJSONHandler")
	defer r.Body.Close()

	var m model.DataUser
	if err := json.Unmarshal(bytes, &m); err != nil {
		log.Errorf("body error: %v \n error decoding message: %v", string(bytes), err)
		http.Error(w, "request body is not valid json", 400)
		return
	}

	// 200 — пользователь успешно аутентифицирован;
	// 400 — неверный формат запроса;
	// 401 — неверная пара логин/пароль;
	// 500 — внутренняя ошибка сервера.
	user, err := hook.Storage.GetUser(m)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
	}
	if tools.Equal(&user, &m) {
		tmp, _ := tools.GetID()
		str := fmt.Sprintf("%x", tmp)
		time, err := tools.AddCookie(w, r, tools.AuthCookieKey, str, 30*time.Minute)
		if err != nil {
			log.Error("Didn't set cooke", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error;"))
			return
		}
		user.SessionExpiredAt = time
		hook.Session.DBSession[str] = user
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}

	log.Println("Post handler")
}
