package postgres

import (
	"context"

	"github.com/borisbbtest/ya-dr/internal/model"
)

const (
	keyPostgresSelectUser = "pgsql.select.tb.user"
)

func (p *Plugin) selectUserHandler(conn *postgresConn, key string, params []interface{}) (interface{}, error) {

	buff := model.DataUsers{}
	query := `SELECT "Login", "Password", "UserID"  FROM  "Users"  WHERE  "Login"  = $1;`

	err := conn.postgresPool.QueryRow(context.Background(), query, params[0]).Scan(&buff.Login, &buff.Password)
	if err != nil {
		log.Error(err)
		return postgresPingFailed, err
	}

	return buff, nil
}
