package postgres

import (
	"context"

	"github.com/borisbbtest/ya-dr/internal/model"
)

const (
	keyPostgresSelectUser = "pgsql.select.tb.user"
)

func (p *Plugin) selectUserHandler(conn *postgresConn, key string, params []interface{}) (interface{}, error) {

	buff := model.DataUser{}
	query := `SELECT "Id", "Login",  "Password"  FROM  "Users"  WHERE  "Login"  = $1;`

	err := conn.PostgresPool.QueryRow(context.Background(), query, params[0]).Scan(&buff.ID, &buff.Login, &buff.Password)
	if err != nil {
		log.Error("selectUserHandler", err)
		return nil, err
	}

	return buff, nil
}
