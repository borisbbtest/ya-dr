package storage

import (
	"github.com/borisbbtest/ya-dr/internal/model"
	"github.com/borisbbtest/ya-dr/internal/postgres"
)

type StoreDBinPostgreSQL struct {
	pgp     postgres.Plugin
	connStr string
}

func NewPostgreSQLStorage(connStr string) (res *StoreDBinPostgreSQL, err error) {
	res = &StoreDBinPostgreSQL{}
	res.connStr = connStr
	res.pgp.Start()
	_, err = res.pgp.NewDBConn("pgsql.create.db.loyalty.system.url", []string{}, connStr, []interface{}{})
	if err != nil {
		log.Error("pgsql.create.db.loyalty.system.url", err)
	}
	return
}

func (hook *StoreDBinPostgreSQL) Close() {
	hook.pgp.Stop()
}

func (hook *StoreDBinPostgreSQL) PutUser(v model.DataUser) (string, error) {
	buff := []interface{}{v.Login, v.Password}
	res, err := hook.pgp.NewDBConn("pgsql.insert.tb.user", []string{}, hook.connStr, buff)
	if err != nil {
		return "", err
	}
	return res.(string), err
}

func (hook *StoreDBinPostgreSQL) PutOrder(v model.DataOrder) (string, error) {
	buff := []interface{}{v.Number, v.Status, v.Person}
	res, err := hook.pgp.NewDBConn("pgsql.insert.tb.order", []string{}, hook.connStr, buff)
	if err != nil {
		return "", err
	}
	return res.(string), err
}

func (hook *StoreDBinPostgreSQL) UpdateOrder(v *model.DataOrder) (string, error) {
	buff := []interface{}{v.Number, v.Status, v.Accrual}
	res, err := hook.pgp.NewDBConn("pgsql.update.tb.order", []string{}, hook.connStr, buff)
	if err != nil {
		return "", err
	}
	return res.(string), err
}

func (hook *StoreDBinPostgreSQL) GetUser(k model.DataUser) (model.DataUser, error) {

	buff := []interface{}{k.Login}
	res, err := hook.pgp.NewDBConn("pgsql.select.tb.user", []string{}, hook.connStr, buff)
	if err != nil {
		log.Error("pgsql.select.tb.user", err)
		return model.DataUser{}, err
	}

	return res.(model.DataUser), nil
}
func (hook *StoreDBinPostgreSQL) GetOrders(k string) ([]model.DataOrder, error) {

	buff := []interface{}{k}
	res, err := hook.pgp.NewDBConn("pgsql.select.tb.orders", []string{}, hook.connStr, buff)
	if err != nil {
		log.Error("pgsql.select.tb.orders", err)
		return nil, err
	}

	return res.([]model.DataOrder), nil
}
