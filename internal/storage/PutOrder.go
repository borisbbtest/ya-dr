package storage

import "github.com/borisbbtest/ya-dr/internal/model"

func (hook *StoreDBinPostgreSQL) PutOrder(v model.DataOrder) (int, error) {
	buff := []interface{}{v.Number, v.Status, v.Person}
	res, err := hook.pgp.NewDBConn("pgsql.insert.tb.order", []string{}, hook.connStr, buff)
	if err != nil {
		return -1, err
	}
	return res.(int), err
}
