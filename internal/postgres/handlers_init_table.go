package postgres

import (
	"context"
)

const (
	keyPostgresCreateDdLoyaltySystem = "pgsql.create.db.loyalty.system.url"
)

func (p *Plugin) CreateTableLoyaltySystemHandler(conn *postgresConn, key string, params []interface{}) (interface{}, error) {
	query := `
				    CREATE TABLE IF NOT EXISTS public."Users"
					(
						"Login" "text",
						"Password" "text" NOT NULL,
						"UserID" "text",
						CONSTRAINT "Login" PRIMARY KEY ("Login")
					)
					TABLESPACE pg_default;

					ALTER TABLE IF EXISTS public."Users" 	OWNER to postgres;

					COMMENT ON TABLE public."Users"  IS 'This table was created for storage data about persons users in within inside project';
			`

	if _, err := conn.postgresPool.Exec(context.Background(), query); err != nil {
		return "Table didn't create ", err
	}
	return "Table created ", nil
}
