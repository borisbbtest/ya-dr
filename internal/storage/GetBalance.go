package storage

import (
	"context"

	"github.com/borisbbtest/ya-dr/internal/model"
)

func (hook *StoreDBinPostgreSQL) GetBalance(v int) (model.DataBalance, error) {

	var res model.DataBalance
	query := `SELECT "CurrentAccrual", "Withdrawn"  from "Balance" where "Person" = $1 ;`
	conn, err := hook.pgp.NewConn()
	conn.PostgresPool.QueryRow(context.Background(), query, v).Scan(&res.CurrentAccrual, &res.Withdrawn)
	if err != nil {
		log.Error("Error selectBalanceHandler", err)
		return model.DataBalance{}, err
	}

	return res, nil
}
