package storage

import "github.com/borisbbtest/ya-dr/internal/model"

func (hook *StoreDBinPostgreSQL) GetWithdrawals(k int) ([]model.Wallet, error) {

	buff := []interface{}{k}
	res, err := hook.pgp.NewDBConn("pgsql.select.tb.withdrawals", []string{}, hook.connStr, buff)
	if err != nil {
		log.Error("pgsql.select.tb.withdrawals", err)
		return nil, err
	}

	return res.([]model.Wallet), nil
}
