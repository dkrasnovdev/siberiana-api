//go:build ignore

package main

import (
	"log"
	"os"
	"path/filepath"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/dkrasnovdev/siberiana-api/config"
	"github.com/dkrasnovdev/siberiana-api/internal/minio"
)

func main() {
	// Get the present working directory.
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("getting present working directory: %v", err)
	}

	// Load the configuration from the .env file.
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
		return
	}

	// Create a new MinIO client using the loaded configuration.
	minioClient := minio.NewClient(config)

	// Create a new entgql extension.
	gql, err := entgql.NewExtension(
		entgql.WithWhereInputs(true),                                           // Enable where inputs in GraphQL schema.
		entgql.WithConfigPath(filepath.Join(pwd, "/internal/gql/gqlgen.yml")),  // Set the path to the gqlgen configuration file.
		entgql.WithSchemaGenerator(),                                           // Enable schema generation.
		entgql.WithSchemaPath(filepath.Join(pwd, "/internal/gql/ent.graphql")), // Set the path to the ent GraphQL schema file.
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}

	// Set options for entc code generation.
	opts := []entc.Option{
		entc.Extensions(gql), // Add the entgql extension.
		entc.FeatureNames("intercept", "privacy", "schema/snapshot"), // Enable specific ent features.
		entc.Dependency(
			entc.DependencyName("Minio"),
			entc.DependencyType(minioClient),
		),
	}

	// Run the ent code generation.
	err = entc.Generate(filepath.Join(pwd, "/ent/schema"), &gen.Config{}, opts...)
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
