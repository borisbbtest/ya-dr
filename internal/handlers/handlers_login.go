package handlers

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/borisbbtest/ya-dr/internal/config"
	"github.com/borisbbtest/ya-dr/internal/model"
	"github.com/borisbbtest/ya-dr/internal/storage"
	"github.com/sirupsen/logrus"
)

var log = logrus.WithField("context", "system_loyalty")

type WrapperHandler struct {
	ServerConf *config.MainConfig
	Storage    storage.Storage
	UserID     string
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

	var m model.DataUsers
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
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusCreated)
	}

	log.Println("Post handler")
}
