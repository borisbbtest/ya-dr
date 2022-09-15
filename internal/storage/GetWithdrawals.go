package storage

import (
	"context"

	"github.com/borisbbtest/ya-dr/internal/model"
)

func (hook *StoreDBinPostgreSQL) GetWithdrawals(k int) ([]model.Wallet, error) {
	var err error
	orders := []model.Wallet{}
	query := `SELECT "Order", "Sum", "Uploaded_at" FROM "Wallet" WHERE "Person" = $1 ORDER BY  "Uploaded_at";`

	conn, err := hook.pgp.NewConn()
	if err != nil {
		log.Error("selectOrdersHandler - c: ", err)
		return nil, err
	}
	rows, err := conn.PostgresPool.Query(context.Background(), query, k)

	for rows.Next() {
		m := model.Wallet{}
		err = rows.Scan(&m.Order, &m.Sum, &m.ProcessedAt)
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
