package migrate

import (
	"context"
	"database/sql"
	"fmt"

	"entgo.io/ent/dialect/sql/schema"
)

// EnablePostGISOption returns a schema.MigrateOption
// that will enable the Postgres Postgis extension if needed.
func EnablePostgisOption(db *sql.DB) schema.MigrateOption {
	return schema.WithHooks(func(next schema.Creator) schema.Creator {
		return schema.CreateFunc(func(ctx context.Context, tables ...*schema.Table) error {
			_, err := db.ExecContext(ctx, `CREATE EXTENSION IF NOT EXISTS postgis WITH SCHEMA public;`)
			if err != nil {
				return fmt.Errorf("could not enable postgis extension: %w", err)
			}

			return next.Create(ctx, tables...)
		})
	})
}
