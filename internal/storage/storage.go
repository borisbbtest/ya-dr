package storage

import (
	"github.com/borisbbtest/ya-dr/internal/model"
	"github.com/sirupsen/logrus"
)

type SessionHTTP struct {
	DBSession map[string]model.DataUser
}
type Storage interface {
	PutUser(v model.DataUser) (string, error)
	GetUser(u model.DataUser) (model.DataUser, error)
	Close()
}

var log = logrus.WithField("context", "system_loyalty")
