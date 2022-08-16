package storage

import (
	"context"

	"github.com/borisbbtest/ya-dr/internal/model"
)

func (hook *StoreDBinPostgreSQL) PutOrder(v model.DataOrder) (int, error) {
	var err error
	var person int
	query := `
	WITH cte AS (
		INSERT INTO public."Orders" ("Number","Status", "Person","Uploaded_at")
		VALUES ($1, $2, $3, NOW())
		ON CONFLICT ("Number") DO NOTHING
		RETURNING "Number"
	)
	SELECT NULL AS result
	WHERE EXISTS (SELECT 1 FROM cte)
	UNION ALL
    SELECT  "Person"  FROM  "Orders"  WHERE  "Number"  = $1;`
	conn, err := hook.pgp.NewConn()
	if err != nil {
		log.Error("selectOrdersHandler - c: ", err)
		return -1, err
	}
	err = conn.PostgresPool.QueryRow(context.Background(), query, v.Number, v.Status, v.Person).Scan(&person)

	if err != nil {
		log.Info("insertOrderHandler  --  ", err)
		return -1, err
	}

	return person, nil
}
