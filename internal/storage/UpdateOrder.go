package storage

import "github.com/borisbbtest/ya-dr/internal/model"

func (hook *StoreDBinPostgreSQL) UpdateOrder(v *model.DataOrder) (string, error) {
	buff := []interface{}{v.Number, v.Status, v.Accrual}
	res, err := hook.pgp.NewDBConn("pgsql.update.tb.order", []string{}, hook.connStr, buff)
	if err != nil {
		return "", err
	}
	return res.(string), err
}
