package storage

import "github.com/borisbbtest/ya-dr/internal/model"

func (hook *StoreDBinPostgreSQL) GetOrders(k int) ([]model.DataOrder, error) {

	buff := []interface{}{k}
	res, err := hook.pgp.NewDBConn("pgsql.select.tb.orders", []string{}, hook.connStr, buff)
	if err != nil {
		log.Error("pgsql.select.tb.orders", err)
		return nil, err
	}

	return res.([]model.DataOrder), nil
}
