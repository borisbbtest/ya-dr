package postgres

import (
	"context"

	"github.com/borisbbtest/ya-dr/internal/model"
)

const (
	keyPostgresSelectWithdrawals = "pgsql.select.tb.withdrawals"
)

func (p *Plugin) selectWithdrawalsHandler(conn *postgresConn, key string, params []interface{}) (interface{}, error) {

	var err error
	orders := []model.Wallet{}
	query := `SELECT "Order", "Sum", "Uploaded_at" FROM "Wallet" WHERE "Person" = $1;`

	rows, err := conn.postgresPool.Query(context.Background(), query, params...)

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
