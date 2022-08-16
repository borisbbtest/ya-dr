package storage

import "github.com/borisbbtest/ya-dr/internal/model"

func (hook *StoreDBinPostgreSQL) PutWithdraw(v model.Wallet) (string, error) {
	buff := []interface{}{v.Order, v.Person, v.Sum}
	res, err := hook.pgp.NewDBConn("pgsql.insert.tb.withdraw", []string{}, hook.connStr, buff)
	if err != nil {
		return "", err
	}
	return res.(string), err
}
