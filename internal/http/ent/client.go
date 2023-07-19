package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/dkrasnovdev/heritage-api/config"
	"github.com/dkrasnovdev/heritage-api/ent"
	"github.com/dkrasnovdev/heritage-api/ent/migrate"

	_ "github.com/lib/pq"
)

func NewClient(env config.Config) *ent.Client {
	// Create the PostgreSQL connection string.
	psql := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		env.HOST, env.PORT, env.POSTGRES_USER, env.POSTGRES_PASSWORD, env.POSTGRES_DB, env.SSL_MODE)

	// Open a connection to the PostgreSQL database.
	client, err := ent.Open("postgres", psql)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	// Apply schema migrations to the database.
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatal("opening ent client", err)
	}

	return client
}