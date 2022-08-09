package tools

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

var log = logrus.WithField("context", "service_system_loyalty")

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
