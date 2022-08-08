package storage

import (
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
