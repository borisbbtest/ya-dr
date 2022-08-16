package storage

import (
	"github.com/borisbbtest/ya-dr/internal/model"
)

func (hook *StoreDBinPostgreSQL) PutUser(v model.DataUser) (string, error) {
	buff := []interface{}{v.Login, v.Password}
	res, err := hook.pgp.NewDBConn("pgsql.insert.tb.user", []string{}, hook.connStr, buff)
	if err != nil {
		log.Info(res)
		hook.pgp.NewDBConn("pgsql.insert.tb.balance", []string{}, hook.connStr, buff)
		return "", err
	}

	return res.(string), err

	// var err error
	// var users string

	// conn, err := hook.pgp.NewConn()
	// if err != nil {
	// 	log.Error("selectOrdersHandler - c: ", err)
	// 	return "", err
	// }

	// query := `
	// WITH cte AS (
	// 	INSERT INTO public."Users"(
	// 	"Login", "Password" )
	// 	VALUES ($1, $2)
	// 	ON CONFLICT ("Login") DO NOTHING
	// 	RETURNING "Login"
	// )
	// SELECT NULL AS result
	// WHERE EXISTS (SELECT 1 FROM cte)
	// UNION ALL
	// SELECT "Login"  FROM  "Users"  WHERE  "Login"  = $1;`

	// err = conn.PostgresPool.QueryRow(context.Background(), query, v.Login, v.Password).Scan(&users)

	// if err != nil {
	// 	log.Info("insertUserHandler  --  ", err)
	// 	return "", err
	// }

	// return users, nil

	// query = `INSERT INTO "Balance" ("Person", "Withdrawn" , "CurrentAccrual" )
	// SELECT "Id", 0 ,0  From "Users" WHERE  "Login"  = $1 ;`

	// if _, err := conn.PostgresPool.Exec(context.Background(), query, v.Login); err != nil {
	// 	log.Info("insertBalanceHandler --- ", err)
	// }

	// return "", nil
}
