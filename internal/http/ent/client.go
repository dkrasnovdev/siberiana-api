package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/dkrasnovdev/heritage-api/config"
	"github.com/dkrasnovdev/heritage-api/ent"
	"github.com/dkrasnovdev/heritage-api/ent/migrate"
	"github.com/dkrasnovdev/heritage-api/internal/ent/hook"

	_ "github.com/dkrasnovdev/heritage-api/ent/runtime"
	_ "github.com/lib/pq"
)

// NewClient creates a new ent client and applies schema migrations to the database.
// It sets up hooks for various entity types to log changes to the respective entities.
func NewClient(env config.Config) *ent.Client {
	// Create the PostgreSQL connection string.
	psql := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		env.HOST, env.PORT, env.POSTGRES_USER, env.POSTGRES_PASSWORD, env.POSTGRES_DB, env.SSL_MODE)

	// Open a connection to the PostgreSQL database.
	client, err := ent.Open("postgres", psql)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	// Apply schema migrations to the database.
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatal("opening ent client", err)
	}

	// Attach hooks for various entity types to log changes.
	client.Artifact.Use(hook.ArtifactLogger(client))
	client.Category.Use(hook.CategoryLogger(client))
	client.Collection.Use(hook.CollectionLogger(client))
	client.Culture.Use(hook.CultureLogger(client))
	client.District.Use(hook.DistrictLogger(client))
	client.Holder.Use(hook.HolderLogger(client))
	client.License.Use(hook.LicenseLogger(client))
	client.Location.Use(hook.LocationLogger(client))
	client.Medium.Use(hook.MediumLogger(client))
	client.Model.Use(hook.ModelLogger(client))
	client.Monument.Use(hook.MonumentLogger(client))
	client.Organization.Use(hook.OrganizationLogger(client))
	client.Person.Use(hook.PersonLogger(client))
	client.Project.Use(hook.ProjectLogger(client))
	client.Publication.Use(hook.PublicationLogger(client))
	client.Region.Use(hook.RegionLogger(client))
	client.Set.Use(hook.SetLogger(client))
	client.Settlement.Use(hook.SettlementLogger(client))
	client.Technique.Use(hook.TechniqueLogger(client))

	return client
}
