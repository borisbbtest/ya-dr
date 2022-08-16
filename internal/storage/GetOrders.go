package storage

import (
	"context"

	"github.com/borisbbtest/ya-dr/internal/model"
)

func (hook *StoreDBinPostgreSQL) GetOrders(k int) ([]model.DataOrder, error) {

	var err error
	orders := []model.DataOrder{}
	query := `SELECT "Number","Status" ,"Person", "Accrual", "Uploaded_at"   FROM  "Orders"  WHERE  "Person"  = $1 ORDER BY "Uploaded_at";`
	conn, err := hook.pgp.NewConn()
	if err != nil {
		log.Error("selectOrdersHandler - c: ", err)
		return nil, err
	}
	rows, err := conn.PostgresPool.Query(context.Background(), query, k)

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
