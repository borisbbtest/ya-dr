package postgres

import (
	"context"
)

const (
	keyPostgresSelectBalance       = "pgsql.select.tb.balance"
	keyPostgresSelectWithdrawCount = "pgsql.select.tb.withdraw.count"
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

func (p *Plugin) selectWithdrawCountHandler(conn *postgresConn, key string, params []interface{}) (interface{}, error) {

	var buff int
	query := `SELECT count("Accrual") from "Orders" where "Person" = $1 AND "Accrual" < 0;`

	err := conn.postgresPool.QueryRow(context.Background(), query, params...).Scan(&buff)
	if err != nil {
		log.Error("selectWithdrawCountHandler", err)
		return 0, err
	}

	return buff, nil
}
