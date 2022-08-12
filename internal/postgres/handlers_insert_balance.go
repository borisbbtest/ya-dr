package postgres

import (
	"context"
)

const (
	keyPostgresInsertBalance = "pgsql.insert.tb.balance"
)

func (p *Plugin) insertBalanceHandler(conn *postgresConn, key string, params []interface{}) (interface{}, error) {

	query := `INSERT INTO "Balance" ("Person", "Withdrawn" , "CurrentAccrual" ) VALUES ($1,0, 0);`

	if _, err := conn.postgresPool.Exec(context.Background(), query, params[0]); err != nil {
		log.Info("insertBalanceHandler --- ", err)
		return "insertBalanceHandler ", err
	}

	return "", nil
}