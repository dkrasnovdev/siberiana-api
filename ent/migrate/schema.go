// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ArtifactsColumns holds the columns for the "artifacts" table.
	ArtifactsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "display_name", Type: field.TypeString, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "primary_image_url", Type: field.TypeString, Nullable: true},
		{Name: "additional_image_urls", Type: field.TypeJSON, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_by", Type: field.TypeString, Nullable: true},
		{Name: "collection_artifacts", Type: field.TypeInt, Nullable: true},
		{Name: "culture_artifacts", Type: field.TypeInt, Nullable: true},
		{Name: "license_artifacts", Type: field.TypeInt, Nullable: true},
		{Name: "location_artifacts", Type: field.TypeInt, Nullable: true},
		{Name: "model_artifacts", Type: field.TypeInt, Nullable: true},
		{Name: "monument_artifacts", Type: field.TypeInt, Nullable: true},
		{Name: "set_artifacts", Type: field.TypeInt, Nullable: true},
	}
	// ArtifactsTable holds the schema information for the "artifacts" table.
	ArtifactsTable = &schema.Table{
		Name:       "artifacts",
		Columns:    ArtifactsColumns,
		PrimaryKey: []*schema.Column{ArtifactsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "artifacts_collections_artifacts",
				Columns:    []*schema.Column{ArtifactsColumns[11]},
				RefColumns: []*schema.Column{CollectionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "artifacts_cultures_artifacts",
				Columns:    []*schema.Column{ArtifactsColumns[12]},
				RefColumns: []*schema.Column{CulturesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "artifacts_licenses_artifacts",
				Columns:    []*schema.Column{ArtifactsColumns[13]},
				RefColumns: []*schema.Column{LicensesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "artifacts_locations_artifacts",
				Columns:    []*schema.Column{ArtifactsColumns[14]},
				RefColumns: []*schema.Column{LocationsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "artifacts_models_artifacts",
				Columns:    []*schema.Column{ArtifactsColumns[15]},
				RefColumns: []*schema.Column{ModelsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "artifacts_monuments_artifacts",
				Columns:    []*schema.Column{ArtifactsColumns[16]},
				RefColumns: []*schema.Column{MonumentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "artifacts_sets_artifacts",
				Columns:    []*schema.Column{ArtifactsColumns[17]},
				RefColumns: []*schema.Column{SetsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// AuditLogsColumns holds the columns for the "audit_logs" table.
	AuditLogsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "table", Type: field.TypeString, Nullable: true},
		{Name: "ref_id", Type: field.TypeInt, Nullable: true},
		{Name: "operation", Type: field.TypeString, Nullable: true},
		{Name: "changes", Type: field.TypeJSON, Nullable: true},
		{Name: "added_ids", Type: field.TypeJSON, Nullable: true},
		{Name: "removed_ids", Type: field.TypeJSON, Nullable: true},
		{Name: "cleared_edges", Type: field.TypeJSON, Nullable: true},
		{Name: "blame", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
	}
	// AuditLogsTable holds the schema information for the "audit_logs" table.
	AuditLogsTable = &schema.Table{
		Name:       "audit_logs",
		Columns:    AuditLogsColumns,
		PrimaryKey: []*schema.Column{AuditLogsColumns[0]},
	}
	// CategoriesColumns holds the columns for the "categories" table.
	CategoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "display_name", Type: field.TypeString, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
	}
	// CategoriesTable holds the schema information for the "categories" table.
	CategoriesTable = &schema.Table{
		Name:       "categories",
		Columns:    CategoriesColumns,
		PrimaryKey: []*schema.Column{CategoriesColumns[0]},
	}
	// CollectionsColumns holds the columns for the "collections" table.
	CollectionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "display_name", Type: field.TypeString, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "category_collections", Type: field.TypeInt, Nullable: true},
	}
	// CollectionsTable holds the schema information for the "collections" table.
	CollectionsTable = &schema.Table{
		Name:       "collections",
		Columns:    CollectionsColumns,
		PrimaryKey: []*schema.Column{CollectionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "collections_categories_collections",
				Columns:    []*schema.Column{CollectionsColumns[7]},
				RefColumns: []*schema.Column{CategoriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CulturesColumns holds the columns for the "cultures" table.
	CulturesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "display_name", Type: field.TypeString, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
	}
	// CulturesTable holds the schema information for the "cultures" table.
	CulturesTable = &schema.Table{
		Name:       "cultures",
		Columns:    CulturesColumns,
		PrimaryKey: []*schema.Column{CulturesColumns[0]},
	}
	// DistrictsColumns holds the columns for the "districts" table.
	DistrictsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "display_name", Type: field.TypeString, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "location_district", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// DistrictsTable holds the schema information for the "districts" table.
	DistrictsTable = &schema.Table{
		Name:       "districts",
		Columns:    DistrictsColumns,
		PrimaryKey: []*schema.Column{DistrictsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "districts_locations_district",
				Columns:    []*schema.Column{DistrictsColumns[7]},
				RefColumns: []*schema.Column{LocationsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// HoldersColumns holds the columns for the "holders" table.
	HoldersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "display_name", Type: field.TypeString, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
	}
	// HoldersTable holds the schema information for the "holders" table.
	HoldersTable = &schema.Table{
		Name:       "holders",
		Columns:    HoldersColumns,
		PrimaryKey: []*schema.Column{HoldersColumns[0]},
	}
	// LicensesColumns holds the columns for the "licenses" table.
	LicensesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "display_name", Type: field.TypeString, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
	}
	// LicensesTable holds the schema information for the "licenses" table.
	LicensesTable = &schema.Table{
		Name:       "licenses",
		Columns:    LicensesColumns,
		PrimaryKey: []*schema.Column{LicensesColumns[0]},
	}
	// LocationsColumns holds the columns for the "locations" table.
	LocationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "display_name", Type: field.TypeString, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
	}
	// LocationsTable holds the schema information for the "locations" table.
	LocationsTable = &schema.Table{
		Name:       "locations",
		Columns:    LocationsColumns,
		PrimaryKey: []*schema.Column{LocationsColumns[0]},
	}
	// MediaColumns holds the columns for the "media" table.
	MediaColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "display_name", Type: field.TypeString, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
	}
	// MediaTable holds the schema information for the "media" table.
	MediaTable = &schema.Table{
		Name:       "media",
		Columns:    MediaColumns,
		PrimaryKey: []*schema.Column{MediaColumns[0]},
	}
	// ModelsColumns holds the columns for the "models" table.
	ModelsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "display_name", Type: field.TypeString, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
	}
	// ModelsTable holds the schema information for the "models" table.
	ModelsTable = &schema.Table{
		Name:       "models",
		Columns:    ModelsColumns,
		PrimaryKey: []*schema.Column{ModelsColumns[0]},
	}
	// MonumentsColumns holds the columns for the "monuments" table.
	MonumentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "display_name", Type: field.TypeString, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
	}
	// MonumentsTable holds the schema information for the "monuments" table.
	MonumentsTable = &schema.Table{
		Name:       "monuments",
		Columns:    MonumentsColumns,
		PrimaryKey: []*schema.Column{MonumentsColumns[0]},
	}
	// OrganizationsColumns holds the columns for the "organizations" table.
	OrganizationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "display_name", Type: field.TypeString, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "holder_organization", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// OrganizationsTable holds the schema information for the "organizations" table.
	OrganizationsTable = &schema.Table{
		Name:       "organizations",
		Columns:    OrganizationsColumns,
		PrimaryKey: []*schema.Column{OrganizationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "organizations_holders_organization",
				Columns:    []*schema.Column{OrganizationsColumns[7]},
				RefColumns: []*schema.Column{HoldersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PersonsColumns holds the columns for the "persons" table.
	PersonsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "display_name", Type: field.TypeString, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "holder_person", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// PersonsTable holds the schema information for the "persons" table.
	PersonsTable = &schema.Table{
		Name:       "persons",
		Columns:    PersonsColumns,
		PrimaryKey: []*schema.Column{PersonsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "persons_holders_person",
				Columns:    []*schema.Column{PersonsColumns[7]},
				RefColumns: []*schema.Column{HoldersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ProjectsColumns holds the columns for the "projects" table.
	ProjectsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "display_name", Type: field.TypeString, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
	}
	// ProjectsTable holds the schema information for the "projects" table.
	ProjectsTable = &schema.Table{
		Name:       "projects",
		Columns:    ProjectsColumns,
		PrimaryKey: []*schema.Column{ProjectsColumns[0]},
	}
	// PublicationsColumns holds the columns for the "publications" table.
	PublicationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "display_name", Type: field.TypeString, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
	}
	// PublicationsTable holds the schema information for the "publications" table.
	PublicationsTable = &schema.Table{
		Name:       "publications",
		Columns:    PublicationsColumns,
		PrimaryKey: []*schema.Column{PublicationsColumns[0]},
	}
	// RegionsColumns holds the columns for the "regions" table.
	RegionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "display_name", Type: field.TypeString, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "location_region", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// RegionsTable holds the schema information for the "regions" table.
	RegionsTable = &schema.Table{
		Name:       "regions",
		Columns:    RegionsColumns,
		PrimaryKey: []*schema.Column{RegionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "regions_locations_region",
				Columns:    []*schema.Column{RegionsColumns[7]},
				RefColumns: []*schema.Column{LocationsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SetsColumns holds the columns for the "sets" table.
	SetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "display_name", Type: field.TypeString, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
	}
	// SetsTable holds the schema information for the "sets" table.
	SetsTable = &schema.Table{
		Name:       "sets",
		Columns:    SetsColumns,
		PrimaryKey: []*schema.Column{SetsColumns[0]},
	}
	// SettlementsColumns holds the columns for the "settlements" table.
	SettlementsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "display_name", Type: field.TypeString, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "location_settlement", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// SettlementsTable holds the schema information for the "settlements" table.
	SettlementsTable = &schema.Table{
		Name:       "settlements",
		Columns:    SettlementsColumns,
		PrimaryKey: []*schema.Column{SettlementsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "settlements_locations_settlement",
				Columns:    []*schema.Column{SettlementsColumns[7]},
				RefColumns: []*schema.Column{LocationsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TechniquesColumns holds the columns for the "techniques" table.
	TechniquesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "display_name", Type: field.TypeString, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
	}
	// TechniquesTable holds the schema information for the "techniques" table.
	TechniquesTable = &schema.Table{
		Name:       "techniques",
		Columns:    TechniquesColumns,
		PrimaryKey: []*schema.Column{TechniquesColumns[0]},
	}
	// HolderArtifactsColumns holds the columns for the "holder_artifacts" table.
	HolderArtifactsColumns = []*schema.Column{
		{Name: "holder_id", Type: field.TypeInt},
		{Name: "artifact_id", Type: field.TypeInt},
	}
	// HolderArtifactsTable holds the schema information for the "holder_artifacts" table.
	HolderArtifactsTable = &schema.Table{
		Name:       "holder_artifacts",
		Columns:    HolderArtifactsColumns,
		PrimaryKey: []*schema.Column{HolderArtifactsColumns[0], HolderArtifactsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "holder_artifacts_holder_id",
				Columns:    []*schema.Column{HolderArtifactsColumns[0]},
				RefColumns: []*schema.Column{HoldersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "holder_artifacts_artifact_id",
				Columns:    []*schema.Column{HolderArtifactsColumns[1]},
				RefColumns: []*schema.Column{ArtifactsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// MediumArtifactsColumns holds the columns for the "medium_artifacts" table.
	MediumArtifactsColumns = []*schema.Column{
		{Name: "medium_id", Type: field.TypeInt},
		{Name: "artifact_id", Type: field.TypeInt},
	}
	// MediumArtifactsTable holds the schema information for the "medium_artifacts" table.
	MediumArtifactsTable = &schema.Table{
		Name:       "medium_artifacts",
		Columns:    MediumArtifactsColumns,
		PrimaryKey: []*schema.Column{MediumArtifactsColumns[0], MediumArtifactsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "medium_artifacts_medium_id",
				Columns:    []*schema.Column{MediumArtifactsColumns[0]},
				RefColumns: []*schema.Column{MediaColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "medium_artifacts_artifact_id",
				Columns:    []*schema.Column{MediumArtifactsColumns[1]},
				RefColumns: []*schema.Column{ArtifactsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// PersonArtifactsColumns holds the columns for the "person_artifacts" table.
	PersonArtifactsColumns = []*schema.Column{
		{Name: "person_id", Type: field.TypeInt},
		{Name: "artifact_id", Type: field.TypeInt},
	}
	// PersonArtifactsTable holds the schema information for the "person_artifacts" table.
	PersonArtifactsTable = &schema.Table{
		Name:       "person_artifacts",
		Columns:    PersonArtifactsColumns,
		PrimaryKey: []*schema.Column{PersonArtifactsColumns[0], PersonArtifactsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "person_artifacts_person_id",
				Columns:    []*schema.Column{PersonArtifactsColumns[0]},
				RefColumns: []*schema.Column{PersonsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "person_artifacts_artifact_id",
				Columns:    []*schema.Column{PersonArtifactsColumns[1]},
				RefColumns: []*schema.Column{ArtifactsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// PersonProjectsColumns holds the columns for the "person_projects" table.
	PersonProjectsColumns = []*schema.Column{
		{Name: "person_id", Type: field.TypeInt},
		{Name: "project_id", Type: field.TypeInt},
	}
	// PersonProjectsTable holds the schema information for the "person_projects" table.
	PersonProjectsTable = &schema.Table{
		Name:       "person_projects",
		Columns:    PersonProjectsColumns,
		PrimaryKey: []*schema.Column{PersonProjectsColumns[0], PersonProjectsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "person_projects_person_id",
				Columns:    []*schema.Column{PersonProjectsColumns[0]},
				RefColumns: []*schema.Column{PersonsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "person_projects_project_id",
				Columns:    []*schema.Column{PersonProjectsColumns[1]},
				RefColumns: []*schema.Column{ProjectsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// PersonPublicationsColumns holds the columns for the "person_publications" table.
	PersonPublicationsColumns = []*schema.Column{
		{Name: "person_id", Type: field.TypeInt},
		{Name: "publication_id", Type: field.TypeInt},
	}
	// PersonPublicationsTable holds the schema information for the "person_publications" table.
	PersonPublicationsTable = &schema.Table{
		Name:       "person_publications",
		Columns:    PersonPublicationsColumns,
		PrimaryKey: []*schema.Column{PersonPublicationsColumns[0], PersonPublicationsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "person_publications_person_id",
				Columns:    []*schema.Column{PersonPublicationsColumns[0]},
				RefColumns: []*schema.Column{PersonsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "person_publications_publication_id",
				Columns:    []*schema.Column{PersonPublicationsColumns[1]},
				RefColumns: []*schema.Column{PublicationsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// ProjectArtifactsColumns holds the columns for the "project_artifacts" table.
	ProjectArtifactsColumns = []*schema.Column{
		{Name: "project_id", Type: field.TypeInt},
		{Name: "artifact_id", Type: field.TypeInt},
	}
	// ProjectArtifactsTable holds the schema information for the "project_artifacts" table.
	ProjectArtifactsTable = &schema.Table{
		Name:       "project_artifacts",
		Columns:    ProjectArtifactsColumns,
		PrimaryKey: []*schema.Column{ProjectArtifactsColumns[0], ProjectArtifactsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "project_artifacts_project_id",
				Columns:    []*schema.Column{ProjectArtifactsColumns[0]},
				RefColumns: []*schema.Column{ProjectsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "project_artifacts_artifact_id",
				Columns:    []*schema.Column{ProjectArtifactsColumns[1]},
				RefColumns: []*schema.Column{ArtifactsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// PublicationArtifactsColumns holds the columns for the "publication_artifacts" table.
	PublicationArtifactsColumns = []*schema.Column{
		{Name: "publication_id", Type: field.TypeInt},
		{Name: "artifact_id", Type: field.TypeInt},
	}
	// PublicationArtifactsTable holds the schema information for the "publication_artifacts" table.
	PublicationArtifactsTable = &schema.Table{
		Name:       "publication_artifacts",
		Columns:    PublicationArtifactsColumns,
		PrimaryKey: []*schema.Column{PublicationArtifactsColumns[0], PublicationArtifactsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "publication_artifacts_publication_id",
				Columns:    []*schema.Column{PublicationArtifactsColumns[0]},
				RefColumns: []*schema.Column{PublicationsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "publication_artifacts_artifact_id",
				Columns:    []*schema.Column{PublicationArtifactsColumns[1]},
				RefColumns: []*schema.Column{ArtifactsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// TechniqueArtifactsColumns holds the columns for the "technique_artifacts" table.
	TechniqueArtifactsColumns = []*schema.Column{
		{Name: "technique_id", Type: field.TypeInt},
		{Name: "artifact_id", Type: field.TypeInt},
	}
	// TechniqueArtifactsTable holds the schema information for the "technique_artifacts" table.
	TechniqueArtifactsTable = &schema.Table{
		Name:       "technique_artifacts",
		Columns:    TechniqueArtifactsColumns,
		PrimaryKey: []*schema.Column{TechniqueArtifactsColumns[0], TechniqueArtifactsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "technique_artifacts_technique_id",
				Columns:    []*schema.Column{TechniqueArtifactsColumns[0]},
				RefColumns: []*schema.Column{TechniquesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "technique_artifacts_artifact_id",
				Columns:    []*schema.Column{TechniqueArtifactsColumns[1]},
				RefColumns: []*schema.Column{ArtifactsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ArtifactsTable,
		AuditLogsTable,
		CategoriesTable,
		CollectionsTable,
		CulturesTable,
		DistrictsTable,
		HoldersTable,
		LicensesTable,
		LocationsTable,
		MediaTable,
		ModelsTable,
		MonumentsTable,
		OrganizationsTable,
		PersonsTable,
		ProjectsTable,
		PublicationsTable,
		RegionsTable,
		SetsTable,
		SettlementsTable,
		TechniquesTable,
		HolderArtifactsTable,
		MediumArtifactsTable,
		PersonArtifactsTable,
		PersonProjectsTable,
		PersonPublicationsTable,
		ProjectArtifactsTable,
		PublicationArtifactsTable,
		TechniqueArtifactsTable,
	}
)

func init() {
	ArtifactsTable.ForeignKeys[0].RefTable = CollectionsTable
	ArtifactsTable.ForeignKeys[1].RefTable = CulturesTable
	ArtifactsTable.ForeignKeys[2].RefTable = LicensesTable
	ArtifactsTable.ForeignKeys[3].RefTable = LocationsTable
	ArtifactsTable.ForeignKeys[4].RefTable = ModelsTable
	ArtifactsTable.ForeignKeys[5].RefTable = MonumentsTable
	ArtifactsTable.ForeignKeys[6].RefTable = SetsTable
	CollectionsTable.ForeignKeys[0].RefTable = CategoriesTable
	DistrictsTable.ForeignKeys[0].RefTable = LocationsTable
	OrganizationsTable.ForeignKeys[0].RefTable = HoldersTable
	PersonsTable.ForeignKeys[0].RefTable = HoldersTable
	RegionsTable.ForeignKeys[0].RefTable = LocationsTable
	SettlementsTable.ForeignKeys[0].RefTable = LocationsTable
	HolderArtifactsTable.ForeignKeys[0].RefTable = HoldersTable
	HolderArtifactsTable.ForeignKeys[1].RefTable = ArtifactsTable
	MediumArtifactsTable.ForeignKeys[0].RefTable = MediaTable
	MediumArtifactsTable.ForeignKeys[1].RefTable = ArtifactsTable
	PersonArtifactsTable.ForeignKeys[0].RefTable = PersonsTable
	PersonArtifactsTable.ForeignKeys[1].RefTable = ArtifactsTable
	PersonProjectsTable.ForeignKeys[0].RefTable = PersonsTable
	PersonProjectsTable.ForeignKeys[1].RefTable = ProjectsTable
	PersonPublicationsTable.ForeignKeys[0].RefTable = PersonsTable
	PersonPublicationsTable.ForeignKeys[1].RefTable = PublicationsTable
	ProjectArtifactsTable.ForeignKeys[0].RefTable = ProjectsTable
	ProjectArtifactsTable.ForeignKeys[1].RefTable = ArtifactsTable
	PublicationArtifactsTable.ForeignKeys[0].RefTable = PublicationsTable
	PublicationArtifactsTable.ForeignKeys[1].RefTable = ArtifactsTable
	TechniqueArtifactsTable.ForeignKeys[0].RefTable = TechniquesTable
	TechniqueArtifactsTable.ForeignKeys[1].RefTable = ArtifactsTable
}
