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

func (hook *StoreDBinPostgreSQL) PutUser(v model.DataUsers) (string, error) {
	buff := []interface{}{v.Login, v.Password}
	res, err := hook.pgp.NewDBConn("pgsql.insert.tb.users", []string{}, hook.connStr, buff)
	if err != nil {
		return "", err
	}
	return res.(string), err
}

func (hook *StoreDBinPostgreSQL) GetUser(k string) (model.DataUsers, error) {

	buff := []interface{}{k}
	res, err := hook.pgp.NewDBConn("pgsql.select.tb.user", []string{}, hook.connStr, buff)
	if err != nil {
		log.Error("pgsql.select.tb.user", err)
		return model.DataUsers{}, err
	}

	return res.(model.DataUsers), nil
}
