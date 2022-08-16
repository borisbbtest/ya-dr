package middleware

import (
	"net/http"

	"github.com/borisbbtest/ya-dr/internal/handlers"
	"github.com/borisbbtest/ya-dr/internal/storage"
	"github.com/sirupsen/logrus"
)

var log = logrus.WithField("context", "service_system_loyalty")

type WrapperMiddleware struct {
	Session *storage.SessionHTTP
}

func (hook *WrapperMiddleware) MiddleSetSessionCookie(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if handlers.IsUserAuthed(r, hook.Session) {
			next.ServeHTTP(w, r)
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Login failed"))
	})
}
