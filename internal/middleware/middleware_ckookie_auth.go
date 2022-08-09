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

var log = logrus.WithField("context", "service_system_loyalty")

type WrapperMiddlewareHandler struct {
	ServerConf *config.MainConfig
	Storage    storage.Storage
	UserID     string
}

func MidSetCookie(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tmp, _ := tools.GetID()
		tools.AddCookie(w, r, "SystemLoyalty", fmt.Sprintf("%x", tmp), 30*time.Minute)
		next.ServeHTTP(w, r)
	})
}
