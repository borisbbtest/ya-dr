package postgres

import (
	"context"
)

const (
	keyPostgresUpdateOrder = "pgsql.update.tb.order"
)

func (p *Plugin) updateOrderHandler(conn *postgresConn, key string, params []interface{}) (interface{}, error) {

	var err error
	var person string
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

	err = conn.postgresPool.QueryRow(context.Background(), query, params...).Scan(&person)

	if err != nil {
		log.Info("Custom  --  ", err)
		return nil, err
	}

	return person, nil
}
