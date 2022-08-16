package postgres

import (
	"time"
)

type Plugin struct {
	connMgr *connManager
	dsn     string
}

type requestHandler func(conn *postgresConn, key string, params []interface{}) (res interface{}, err error)

func (p *Plugin) Start(dsn string) {
	p.dsn = dsn
	p.connMgr = p.NewConnManager(
		time.Duration(20000)*time.Second,
		time.Duration(20000)*time.Second,
	)
}

func (p *Plugin) Stop() {
	p.connMgr.stop()
	p.connMgr = nil
}

func (p *Plugin) NewConn() (conn *postgresConn, err error) {

	conn, err = p.connMgr.GetPostgresConnection(p.dsn)
	if err != nil {
		log.Errorf("connection error: %s", err)
		return nil, err
	}
	return
}
