package tools

import (
	"errors"
	"net/http"
	"time"

	"github.com/borisbbtest/ya-dr/internal/storage"
	"github.com/sirupsen/logrus"
)

const (
	AuthCookieKey = "SystemLoyalty"
)

var log = logrus.WithField("context", "service_system_loyalty")

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
		cookie, _ := r.Cookie(AuthCookieKey)
		return session.DBSession[cookie.Value].ID, nil
	}
	return -1, ErrUnauthorized
}
