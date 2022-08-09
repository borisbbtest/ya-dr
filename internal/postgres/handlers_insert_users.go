package postgres

import (
	"context"
)

const (
	keyPostgresInsertUser = "pgsql.insert.tb.users"
)

func (p *Plugin) insertUserHandler(conn *postgresConn, key string, params []interface{}) (interface{}, error) {

	var err error
	var users string
	query := `
	WITH cte AS (
		INSERT INTO public."Users"(
		"Login", "Password" )
		VALUES ($1, $2)
		ON CONFLICT ("Login") DO NOTHING
		RETURNING "Login"
	)
	SELECT NULL AS result
	WHERE EXISTS (SELECT 1 FROM cte)
	UNION ALL
    SELECT "Login"  FROM  "Users"  WHERE  "Login"  = $1;`

	err = conn.postgresPool.QueryRow(context.Background(), query, params...).Scan(&users)

	if err != nil {
		log.Info("Custom  --  ", err)
		return nil, err
	}

	return users, nil
}
