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
	Close()
}

var log = logrus.WithField("context", "service_short_url")
