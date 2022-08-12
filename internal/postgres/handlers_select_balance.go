package postgres

import (
	"context"

	"github.com/borisbbtest/ya-dr/internal/model"
)

const (
	keyPostgresSelectBalance       = "pgsql.select.tb.balance"
	keyPostgresSelectWithdrawCount = "pgsql.select.tb.withdraw.count"
)

func (p *Plugin) selectBalanceHandler(conn *postgresConn, key string, params []interface{}) (interface{}, error) {

	var buff float32
	query := `SELECT sum("Sum") from "Wallet" where "Person" = $1;`

	err := conn.postgresPool.QueryRow(context.Background(), query, params...).Scan(&buff)
	if err != nil {
		log.Error("Error selectBalanceHandler", err)
		return 0, err
	}

	return buff, nil
}

func (p *Plugin) selectWithdrawCountHandler(conn *postgresConn, key string, params []interface{}) (interface{}, error) {

	buff := model.DataUser{}
	query := `SELECT count("Sum") from "Wallet" where "Person" = $1 AND "Sum" < 0;`

	err := conn.postgresPool.QueryRow(context.Background(), query, params[0]).Scan(&buff.Login, &buff.Password)
	if err != nil {
		log.Error("selectWithdrawCountHandler", err)
		return 0, err
	}

	return buff, nil
}
