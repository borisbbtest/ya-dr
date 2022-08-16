package storage

import "github.com/borisbbtest/ya-dr/internal/model"

func (hook *StoreDBinPostgreSQL) UpdateBalance(v *model.DataBalance) (string, error) {
	buff := []interface{}{v.Person, v.CurrentAccrual, v.Withdrawn}
	res, err := hook.pgp.NewDBConn("pgsql.update.tb.balance", []string{}, hook.connStr, buff)
	if err != nil {
		return "", err
	}
	return res.(string), err
}
