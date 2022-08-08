package postgres

import (
	"context"
	"fmt"
)

const (
	keyPostgresPing     = "pgsql.ping"
	postgresPingUnknown = -1
	postgresPingFailed  = 0
	postgresPingOk      = 1
)

// pingHandler executes 'SELECT 1 as pingOk' commands and returns pingOK if a connection is alive or postgresPingFailed otherwise.
func (p *Plugin) pingHandler(conn *postgresConn, key string, params []interface{}) (interface{}, error) {
	var pingOK int64 = postgresPingUnknown

	_ = conn.postgresPool.QueryRow(context.Background(), fmt.Sprintf("SELECT %d as pingOk", postgresPingOk)).Scan(&pingOK)
	if pingOK != postgresPingOk {
		log.Error(errorPostgresPing)
		return postgresPingFailed, nil
	}

	return pingOK, nil
}
