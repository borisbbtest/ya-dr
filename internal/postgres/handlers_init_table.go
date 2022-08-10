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
						CONSTRAINT "Login" PRIMARY KEY ("Login")
					)
					TABLESPACE pg_default;

					ALTER TABLE IF EXISTS public."Users" 	OWNER to postgres;

					COMMENT ON TABLE public."Users"  IS 'This table was created for storage data about persons users in within inside project';


					CREATE TABLE IF NOT EXISTS public."Orders"
					(
						"Number" "text" NOT NULL,
						"Status" "text" NOT NULL,
						"Person" "text" NOT NULL,
						"Accrual" "numeric",
						"Uploaded_at" "timestamptz" NOT NULL,
						CONSTRAINT "Number" PRIMARY KEY ("Number")
					)
					TABLESPACE pg_default;

					ALTER TABLE IF EXISTS public."Orders" 	OWNER to postgres;

					COMMENT ON TABLE public."Orders"  IS 'This table was created for storage data about orders';

			`

	if _, err := conn.postgresPool.Exec(context.Background(), query); err != nil {
		return "Table didn't create ", err
	}
	return "Table created ", nil
}
