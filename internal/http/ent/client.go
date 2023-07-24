// Package ent provides functions to create a new ent client and apply schema migrations to the database.

package ent

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/dkrasnovdev/heritage-api/config"
	"github.com/dkrasnovdev/heritage-api/ent"
	"github.com/dkrasnovdev/heritage-api/ent/migrate"
	mig "github.com/dkrasnovdev/heritage-api/internal/ent/migrate"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/dkrasnovdev/heritage-api/ent/runtime"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// NewClient creates a new ent client and applies schema migrations to the database.
// It sets up hooks for various entity types to log changes to the respective entities.
func NewClient(env config.Config) *ent.Client {
	// Create the PostgreSQL connection string.
	psql := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		env.HOST, env.PORT, env.POSTGRES_USER, env.POSTGRES_PASSWORD, env.POSTGRES_DB, env.SSL_MODE)

	// Open a connection to the PostgreSQL database.
	db, err := sql.Open("pgx", psql)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	// Create an ent driver using the SQL database connection.
	driver := entsql.OpenDB(dialect.Postgres, db)
	client := ent.NewClient(ent.Driver(driver))

	// Apply schema migrations to the database.
	if err := client.Schema.Create(
		context.Background(),
		mig.EnablePostgisOption(db),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatal("failed to apply schema migrations", err)
	}

	return client
}
