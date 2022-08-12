package postgres

import (
	"context"
)

const (
	keyPostgresUpdateBalance = "pgsql.update.tb.balance"
)

func (p *Plugin) updateBalanceHandler(conn *postgresConn, key string, params []interface{}) (interface{}, error) {

	var err error
	query := ` 	UPDATE  public."Balance"
				SET
					"CurrentAccrual"="CurrentAccrual" + $2 ,
					"Withdrawn" = "Withdrawn"+ $3
				WHERE "Person" = $1;
			`

	if _, err = conn.postgresPool.Exec(context.Background(), query, params...); err != nil {
		log.Error("updateOrderBalance ", err)
		return "didn't update ", err
	}

	return "", nil
}
