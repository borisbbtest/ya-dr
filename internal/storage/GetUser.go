package storage

import (
	"context"

	"github.com/borisbbtest/ya-dr/internal/model"
)

func (hook *StoreDBinPostgreSQL) GetUser(k model.DataUser) (model.DataUser, error) {
	buff := model.DataUser{}
	query := `SELECT "Id", "Login",  "Password"  FROM  "Users"  WHERE  "Login"  = $1;`
	conn, err := hook.pgp.NewConn()
	if err != nil {
		log.Error("selectOrdersHandler - c: ", err)
		return model.DataUser{}, err
	}
	err = conn.PostgresPool.QueryRow(context.Background(), query, k.Login).Scan(&buff.ID, &buff.Login, &buff.Password)
	if err != nil {
		log.Error("selectUserHandler", err)
		return model.DataUser{}, err
	}

	return buff, nil
}
