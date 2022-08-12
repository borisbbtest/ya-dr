package postgres

import (
	"context"
)

const (
	keyPostgresSelectBalance          = "pgsql.select.tb.balance"
	keyPostgresSelectWithdrawCount    = "pgsql.select.tb.withdraw.count"
	keyPostgresSelectSumWithdrawCount = "pgsql.select.tb.sum.withdraw.count"
)

func (p *Plugin) selectBalanceHandler(conn *postgresConn, key string, params []interface{}) (interface{}, error) {

	var buff float32
	query := `SELECT sum("Accrual") from "Orders" where "Person" = $1;`

	err := conn.postgresPool.QueryRow(context.Background(), query, params...).Scan(&buff)
	if err != nil {
		log.Error("Error selectBalanceHandler", err)
		return 0, err
	}

	return buff, nil
}

func (p *Plugin) selectSumWithdrawHandler(conn *postgresConn, key string, params []interface{}) (interface{}, error) {

	var buff float32
	query := `SELECT sum("Sum") from "Wallet" where "Person" = $1;`

	err := conn.postgresPool.QueryRow(context.Background(), query, params...).Scan(&buff)
	if err != nil {
		log.Error("Error selectSumWithdrawHandler", err)
		return 0, err
	}

	return buff, nil
}

func (p *Plugin) selectWithdrawCountHandler(conn *postgresConn, key string, params []interface{}) (interface{}, error) {

	var buff int
	query := `SELECT count("Sum") from "Wallet" where "Person" = $1 AND "Sum" < 0;`

	err := conn.postgresPool.QueryRow(context.Background(), query, params...).Scan(&buff)
	if err != nil {
		log.Error("selectWithdrawCountHandler", err)
		return 0, err
	}

	return buff, nil
}
