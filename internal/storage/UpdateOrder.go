package storage

import (
	"context"

	"github.com/borisbbtest/ya-dr/internal/model"
)

func (hook *StoreDBinPostgreSQL) UpdateOrder(v *model.DataOrder) (string, error) {

	var err error
	query := ` 	UPDATE  public."Orders"
				SET "Status"= $2 , "Accrual" = $3
				WHERE "Number" = $1;
			`
	conn, err := hook.pgp.NewConn()
	if err != nil {
		log.Error("selectOrdersHandler - c: ", err)
		return "", err
	}
	if _, err = conn.PostgresPool.Exec(context.Background(), query, v.Number, v.Status, v.Accrual); err != nil {
		log.Error("updateOrderHandler ", err)
		return "didn't update ", err
	}

	return "", nil
}
