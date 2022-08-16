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
						"Id" "serial",
						"Login" "text" UNIQUE NOT NULL,
						"Password" "text" NOT NULL,
						CONSTRAINT "Id" PRIMARY KEY ("Id")
					)
					TABLESPACE pg_default;

					ALTER TABLE IF EXISTS public."Users" 	OWNER to postgres;

					COMMENT ON TABLE public."Users"  IS 'This table was created for storage data about persons users in within inside project';

					CREATE TABLE IF NOT EXISTS public."Orders"
					(
						"Number" "text" NOT NULL,
						"Status" "text" NOT NULL,
						"Person" "numeric" NOT NULL,
						"Accrual" "numeric",
						"Uploaded_at" "timestamptz" NOT NULL,
						CONSTRAINT "Number" PRIMARY KEY ("Number")
					)
					TABLESPACE pg_default;

					ALTER TABLE IF EXISTS public."Orders" 	OWNER to postgres;

					COMMENT ON TABLE public."Orders"  IS 'This table was created for storage data about orders';

					CREATE TABLE IF NOT EXISTS public."Wallet"
					(
						"Order" "text" NOT NULL,
						"Person" "numeric" NOT NULL,
						"Sum" "numeric",
						"Uploaded_at" "timestamptz" NOT NULL
					)
					TABLESPACE pg_default;

					ALTER TABLE IF EXISTS public."Wallet" 	OWNER to postgres;

					COMMENT ON TABLE public."Wallet"  IS 'This table was created for storage data about Wallet';

					CREATE TABLE IF NOT EXISTS public."Balance"
					(
						"Person"           "numeric",
						"Withdrawn"        "numeric",
						"CurrentAccrual"   "numeric",
					    CONSTRAINT "Person" PRIMARY KEY ("Person")
					)
					TABLESPACE pg_default;

					ALTER TABLE IF EXISTS public."Balance" 	OWNER to postgres;

					COMMENT ON TABLE public."Balance"  IS 'This table was created for storage data about Balance';

			`

	if _, err := conn.PostgresPool.Exec(context.Background(), query); err != nil {
		return "Table didn't create ", err
	}
	return "Table created ", nil
}
