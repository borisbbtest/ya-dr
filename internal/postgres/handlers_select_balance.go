package postgres

import (
	"context"

	"github.com/borisbbtest/ya-dr/internal/model"
)

const (
	keyPostgresSelectBalance = "pgsql.select.tb.balance"
)

func (p *Plugin) selectBalanceHandler(conn *postgresConn, key string, params []interface{}) (interface{}, error) {

	var buff model.DataBalance
	query := `SELECT "CurrentAccrual", "Withdrawn"  from "Balance" where "Person" = $1 ;`

	err := conn.PostgresPool.QueryRow(context.Background(), query, params...).Scan(&buff.CurrentAccrual, &buff.Withdrawn)
	if err != nil {
		log.Error("Error selectBalanceHandler", err)
		return 0, err
	}

	return buff, nil
}
