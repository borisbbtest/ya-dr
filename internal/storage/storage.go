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
	PutOrder(v model.DataOrder) (int, error)
	UpdateOrder(v *model.DataOrder) (string, error)
	GetOrders(k int) ([]model.DataOrder, error)

	GetBalance(v int) (model.DataBalance, error)

	PutWithdraw(v model.Wallet) (string, error)
	GetWithdrawals(k int) ([]model.Wallet, error)
	UpdateBalance(v *model.DataBalance) (string, error)

	Close()
}

var log = logrus.WithField("context", "system_loyalty")
