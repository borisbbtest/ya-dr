package storage

import (
	"context"

	"github.com/borisbbtest/ya-dr/internal/model"
)

func (hook *StoreDBinPostgreSQL) PutWithdraw(v model.Wallet) (string, error) {

	query := `INSERT INTO "Wallet" ("Order","Person", "Sum" , "Uploaded_at") VALUES ($1, $2, $3, NOW());`
	conn, err := hook.pgp.NewConn()
	if err != nil {
		log.Error("selectOrdersHandler - c: ", err)
		return "", err
	}

	if _, err := conn.PostgresPool.Exec(context.Background(), query, v.Order, v.Person, v.Sum); err != nil {
		log.Info("insertWithdrawHandler --- ", err)
		return "insertWithdrawHandler ", err
	}

	return "", nil
}
