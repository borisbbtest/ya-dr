package postgres

import (
	"context"

	"github.com/borisbbtest/ya-dr/internal/model"
)

const (
	keyPostgresSelectOrders = "pgsql.select.tb.orders"
)

func (p *Plugin) selectOrdersHandler(conn *postgresConn, key string, params []interface{}) (interface{}, error) {

	var err error
	orders := []model.DataOrder{}
	query := `SELECT "Number","Status" ,"Person", "Accrual", "Uploaded_at"   FROM  "Orders"  WHERE  "Person"  = $1 ORDER BY "Uploaded_at";`

	rows, err := conn.postgresPool.Query(context.Background(), query, params...)

	for rows.Next() {
		m := model.DataOrder{}
		err = rows.Scan(&m.Number, &m.Status, &m.Person, &m.Accrual, &m.UploadedAt)
		if err != nil {
			log.Info("selectOrdersHandler - c: ", err)
			return nil, err
		}
		orders = append(orders, m)
	}
	if err != nil {
		log.Info("selectOrdersHandler  --  ", err)
		return nil, err
	}

	return orders, nil
}
