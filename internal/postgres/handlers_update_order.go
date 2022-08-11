package postgres

import (
	"context"
)

const (
	keyPostgresUpdateOrder = "pgsql.update.tb.order"
)

func (p *Plugin) updateOrderHandler(conn *postgresConn, key string, params []interface{}) (interface{}, error) {

	var err error
	query := ` 	UPDATE  public."Orders"
				SET "Status"= $2 , "Accrual" = $3
				WHERE "Number" = $1;
			`

	if _, err = conn.postgresPool.Exec(context.Background(), query); err != nil {
		log.Error("updateOrderHandler ", err)
		return "didn't update ", err
	}

	return nil, nil
}
