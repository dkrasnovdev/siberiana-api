// Package ent provides functions to create a new ent client and apply schema migrations to the database.
package ent

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"ariga.io/entcache"
	"github.com/dkrasnovdev/heritage-api/config"
	"github.com/dkrasnovdev/heritage-api/ent"
	"github.com/dkrasnovdev/heritage-api/ent/migrate"
	mig "github.com/dkrasnovdev/heritage-api/internal/ent/migrate"
	"github.com/go-redis/redis/v8"

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

	// Create a context for the database operations.
	ctx := context.Background()

	// Create an ent driver using the SQL database connection.
	driver := entsql.OpenDB(dialect.Postgres, db)

	// Apply schema migrations to the database.
	if err := ent.NewClient(ent.Driver(driver)).Schema.Create(
		ctx,
		mig.EnablePostgisOption(db),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatal("failed to apply schema migrations", err)
	}

	// Create a debug driver to log database operations.
	drv := dialect.Debug(driver)

	// Create a Redis client to be used for caching.
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// Ping the Redis server to check the connection.
	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatal("failed pinging the redis derver", err)
	}

	// Create a caching driver for ent using Redis and LRU cache.
	drv = entcache.NewDriver(
		drv,
		entcache.TTL(time.Second*5),
		entcache.Levels(
			entcache.NewLRU(256),
			entcache.NewRedis(rdb),
		),
	)

	// Create and return the ent client using the caching driver.
	return ent.NewClient(ent.Driver(drv))
}
