package storage

import "github.com/borisbbtest/ya-dr/internal/model"

func (hook *StoreDBinPostgreSQL) GetUser(k model.DataUser) (model.DataUser, error) {

	buff := []interface{}{k.Login}
	res, err := hook.pgp.NewDBConn("pgsql.select.tb.user", []string{}, hook.connStr, buff)
	if err != nil {
		log.Error("pgsql.select.tb.user", err)
		return model.DataUser{}, err
	}

	return res.(model.DataUser), nil
}
