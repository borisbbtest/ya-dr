package storage

import "github.com/borisbbtest/ya-dr/internal/model"

func (hook *StoreDBinPostgreSQL) PutUser(v model.DataUser) (string, error) {
	buff := []interface{}{v.Login, v.Password}
	res, err := hook.pgp.NewDBConn("pgsql.insert.tb.user", []string{}, hook.connStr, buff)
	if err != nil {
		log.Info(res)
		hook.pgp.NewDBConn("pgsql.insert.tb.balance", []string{}, hook.connStr, buff)
		return "", err
	}

	return res.(string), err
}
