package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/borisbbtest/ya-dr/internal/config"
	"github.com/borisbbtest/ya-dr/internal/storage"
	"github.com/borisbbtest/ya-dr/internal/tools"
	"github.com/sirupsen/logrus"
)

var log = logrus.WithField("context", "service_short_url")

type WrapperHandler struct {
	ServerConf *config.ServiceShortURLConfig
	Storage    storage.Storage
	UserID     string
}

func AddCookie(w http.ResponseWriter, r *http.Request, name, value string, ttl time.Duration) (res string, err error) {
	ck, err := r.Cookie(name)
	if err == nil {
		res = ck.Value
		return
	}
	log.Info("Cant find cookie : set cooke")
	expire := time.Now().Add(ttl)
	cookie := http.Cookie{
		Name:    name,
		Value:   value,
		Expires: expire,
	}
	http.SetCookie(w, &cookie)
	return value, err
}

func (hook *WrapperHandler) MidSetCookie(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tmp, _ := tools.GetID()
		hook.UserID, _ = AddCookie(w, r, "ShortURL", fmt.Sprintf("%x", tmp), 30*time.Minute)
		next.ServeHTTP(w, r)
	})
}
