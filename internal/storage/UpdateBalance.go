package storage

import (
	"context"

	"github.com/borisbbtest/ya-dr/internal/model"
)

func (hook *StoreDBinPostgreSQL) UpdateBalance(v *model.DataBalance) (string, error) {
	var err error
	query := ` 	UPDATE  public."Balance"
				SET
					"CurrentAccrual"="CurrentAccrual" + $2 ,
					"Withdrawn" = "Withdrawn"+ $3
				WHERE "Person" = $1;
			`
	conn, err := hook.pgp.NewConn()
	if err != nil {
		log.Error("selectOrdersHandler - c: ", err)
		return "", err
	}
	if _, err = conn.PostgresPool.Exec(context.Background(), query, v.Person, v.CurrentAccrual, v.Withdrawn); err != nil {
		log.Error("updateBalanceHandler ", err)
		return "didn't update ", err
	}

	return "", nil
}
