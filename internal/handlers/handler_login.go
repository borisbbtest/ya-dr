package handlers

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
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
		time, _ := tools.AddCookie(w, r, tools.AuthCookieKey, str, 30*time.Minute)
		user.SessionExpiredAt = time
		hook.Session.DBSession[str] = user
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}

	log.Println("Post handler")
}