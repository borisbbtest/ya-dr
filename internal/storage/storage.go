package storage

import (
	"github.com/borisbbtest/ya-dr/internal/model"
	"github.com/sirupsen/logrus"
)

type StoreDBLocal struct {
	DBLocal  map[string]model.DataUsers
	ListUser map[string][]string
}
type Storage interface {
	PutUser(v model.DataUsers) (string, error)
	GetUser(u model.DataUsers) (model.DataUsers, error)
	Close()
}

var log = logrus.WithField("context", "system_loyalty")
