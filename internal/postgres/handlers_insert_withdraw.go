package postgres

import (
	"context"
)

const (
	keyPostgresInsertWithdraw = "pgsql.insert.tb.withdraw"
)

func (p *Plugin) insertWithdrawHandler(conn *postgresConn, key string, params []interface{}) (interface{}, error) {

	query := `INSERT INTO "Wallet" ("Order","Person", "Sum" , "Uploaded_at") VALUES ($1, $2, $3, NOW());`

	if _, err := conn.postgresPool.Exec(context.Background(), query, params...); err != nil {
		log.Info("insertWithdrawHandler --- ", err)
		return "insertWithdrawHandler ", err
	}

	return "", nil
}
