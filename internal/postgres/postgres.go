package postgres

import "time"

type Plugin struct {
	connMgr *connManager
}

type requestHandler func(conn *postgresConn, key string, params []interface{}) (res interface{}, err error)

func (p *Plugin) Start() {
	p.connMgr = p.NewConnManager(
		time.Duration(20000)*time.Second,
		time.Duration(20000)*time.Second,
	)
}

func (p *Plugin) Stop() {
	p.connMgr.stop()
	p.connMgr = nil
}

// whereToConnect builds a session based on key's parameters and a configuration file.
func whereToConnect(params []string) (u *URI, err error) {
	var uri string
	user := ""
	if len(params) > 1 {
		user = params[1]
	}

	password := ""
	if len(params) > 2 {
		password = params[2]
	}

	database := ""
	if len(params) > 3 && len(params) < 5 {
		database = params[3]
	}

	// The first param can be either a URI or a session identifier
	if len(params) > 0 && len(params[0]) > 0 {
		if isLooksLikeURI(params[0]) {
			// Use the URI defined as key's parameter
			uri = params[0]
		}
	}

	if len(user) > 0 || len(password) > 0 || len(database) > 0 {
		return newURIWithCreds(uri, user, password, database)
	}

	return parseURI(uri)
}

// Export implements the Exporter interface.
func (p *Plugin) NewDBConn(key string, params []string, dsnString string, handlerParams []interface{}) (result interface{}, err error) {
	var (
		handler requestHandler
	)
	var connString string
	u, err := whereToConnect(params)
	if err != nil {
		connString = dsnString
	} else {
		// get connection string for PostgreSQL
		connString = u.URI()
	}
	switch key {
	case keyPostgresPing:
		handler = p.pingHandler // postgres.ping[[connString]]
	case keyPostgresCreateDdLoyaltySystem:
		handler = p.CreateTableLoyaltySystemHandler
	case keyPostgresInsertUser:
		handler = p.insertUserHandler
	case keyPostgresSelectUser:
		handler = p.selectUserHandler
	case keyPostgresInsertOrder:
		handler = p.insertOrderHandler
	case keyPostgresUpdateOrder:
		handler = p.updateOrderHandler
	case keyPostgresSelectOrders:
		handler = p.selectOrdersHandler
	case keyPostgresInsertWithdraw:
		handler = p.insertWithdrawHandler
	case keyPostgresSelectBalance:
		handler = p.selectBalanceHandler
	case keyPostgresSelectWithdrawals:
		handler = p.selectWithdrawalsHandler
	case keyPostgresInsertBalance:
		handler = p.insertBalanceHandler
	case keyPostgresUpdateBalance:
		handler = p.updateBalanceHandler
	default:
		return nil, errorUnsupportedQuery
	}

	conn, err := p.connMgr.GetPostgresConnection(connString)
	if err != nil {
		// Here is another logic of processing connection errors if postgres.ping is requested
		if key == keyPostgresPing {
			return postgresPingFailed, nil
		}
		log.Errorf("connection error: %s", err)
		log.Debugf("parameters: %+v", params)
		return nil, err
	}

	return handler(conn, key, handlerParams)
}
