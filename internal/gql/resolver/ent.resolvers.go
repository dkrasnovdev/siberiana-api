package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"

	"entgo.io/contrib/entgql"
	"github.com/dkrasnovdev/siberiana-api/ent"
	"github.com/dkrasnovdev/siberiana-api/internal/gql"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	return r.client.Noder(ctx, id)
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	return r.client.Noders(ctx, ids)
}

// Arts is the resolver for the arts field.
func (r *queryResolver) Arts(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy []*ent.ArtOrder, where *ent.ArtWhereInput) (*ent.ArtConnection, error) {
	return r.client.Art.Query().Paginate(ctx, after, first, before, last, offset, ent.WithArtOrder(orderBy), ent.WithArtFilter(where.Filter))
}

// ArtGenres is the resolver for the artGenres field.
func (r *queryResolver) ArtGenres(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy []*ent.ArtGenreOrder, where *ent.ArtGenreWhereInput) (*ent.ArtGenreConnection, error) {
	return r.client.ArtGenre.Query().Paginate(ctx, after, first, before, last, offset, ent.WithArtGenreOrder(orderBy), ent.WithArtGenreFilter(where.Filter))
}

// ArtStyles is the resolver for the artStyles field.
func (r *queryResolver) ArtStyles(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy []*ent.ArtStyleOrder, where *ent.ArtStyleWhereInput) (*ent.ArtStyleConnection, error) {
	return r.client.ArtStyle.Query().Paginate(ctx, after, first, before, last, offset, ent.WithArtStyleOrder(orderBy), ent.WithArtStyleFilter(where.Filter))
}

// Artifacts is the resolver for the artifacts field.
func (r *queryResolver) Artifacts(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy []*ent.ArtifactOrder, where *ent.ArtifactWhereInput) (*ent.ArtifactConnection, error) {
	return r.client.Artifact.Query().Paginate(ctx, after, first, before, last, offset, ent.WithArtifactOrder(orderBy), ent.WithArtifactFilter(where.Filter))
}

// AuditLogs is the resolver for the auditLogs field.
func (r *queryResolver) AuditLogs(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy []*ent.AuditLogOrder, where *ent.AuditLogWhereInput) (*ent.AuditLogConnection, error) {
	return r.client.AuditLog.Query().Paginate(ctx, after, first, before, last, offset, ent.WithAuditLogOrder(orderBy), ent.WithAuditLogFilter(where.Filter))
}

// Books is the resolver for the books field.
func (r *queryResolver) Books(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy []*ent.BookOrder, where *ent.BookWhereInput) (*ent.BookConnection, error) {
	return r.client.Book.Query().Paginate(ctx, after, first, before, last, offset, ent.WithBookOrder(orderBy), ent.WithBookFilter(where.Filter))
}

// BookGenres is the resolver for the bookGenres field.
func (r *queryResolver) BookGenres(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy []*ent.BookGenreOrder, where *ent.BookGenreWhereInput) (*ent.BookGenreConnection, error) {
	return r.client.BookGenre.Query().Paginate(ctx, after, first, before, last, offset, ent.WithBookGenreOrder(orderBy), ent.WithBookGenreFilter(where.Filter))
}

// Categories is the resolver for the categories field.
func (r *queryResolver) Categories(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy []*ent.CategoryOrder, where *ent.CategoryWhereInput) (*ent.CategoryConnection, error) {
	return r.client.Category.Query().Paginate(ctx, after, first, before, last, offset, ent.WithCategoryOrder(orderBy), ent.WithCategoryFilter(where.Filter))
}

// Collections is the resolver for the collections field.
func (r *queryResolver) Collections(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy []*ent.CollectionOrder, where *ent.CollectionWhereInput) (*ent.CollectionConnection, error) {
	return r.client.Collection.Query().Paginate(ctx, after, first, before, last, offset, ent.WithCollectionOrder(orderBy), ent.WithCollectionFilter(where.Filter))
}

// Countries is the resolver for the countries field.
func (r *queryResolver) Countries(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy []*ent.CountryOrder, where *ent.CountryWhereInput) (*ent.CountryConnection, error) {
	return r.client.Country.Query().Paginate(ctx, after, first, before, last, offset, ent.WithCountryOrder(orderBy), ent.WithCountryFilter(where.Filter))
}

// Cultures is the resolver for the cultures field.
func (r *queryResolver) Cultures(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy []*ent.CultureOrder, where *ent.CultureWhereInput) (*ent.CultureConnection, error) {
	return r.client.Culture.Query().Paginate(ctx, after, first, before, last, offset, ent.WithCultureOrder(orderBy), ent.WithCultureFilter(where.Filter))
}

// Districts is the resolver for the districts field.
func (r *queryResolver) Districts(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy []*ent.DistrictOrder, where *ent.DistrictWhereInput) (*ent.DistrictConnection, error) {
	return r.client.District.Query().Paginate(ctx, after, first, before, last, offset, ent.WithDistrictOrder(orderBy), ent.WithDistrictFilter(where.Filter))
}

// Holders is the resolver for the holders field.
func (r *queryResolver) Holders(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy []*ent.HolderOrder, where *ent.HolderWhereInput) (*ent.HolderConnection, error) {
	return r.client.Holder.Query().Paginate(ctx, after, first, before, last, offset, ent.WithHolderOrder(orderBy), ent.WithHolderFilter(where.Filter))
}

// HolderResponsibilities is the resolver for the holderResponsibilities field.
func (r *queryResolver) HolderResponsibilities(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy []*ent.HolderResponsibilityOrder, where *ent.HolderResponsibilityWhereInput) (*ent.HolderResponsibilityConnection, error) {
	return r.client.HolderResponsibility.Query().Paginate(ctx, after, first, before, last, offset, ent.WithHolderResponsibilityOrder(orderBy), ent.WithHolderResponsibilityFilter(where.Filter))
}

// Licenses is the resolver for the licenses field.
func (r *queryResolver) Licenses(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy []*ent.LicenseOrder, where *ent.LicenseWhereInput) (*ent.LicenseConnection, error) {
	return r.client.License.Query().Paginate(ctx, after, first, before, last, offset, ent.WithLicenseOrder(orderBy), ent.WithLicenseFilter(where.Filter))
}

// Locations is the resolver for the locations field.
func (r *queryResolver) Locations(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy []*ent.LocationOrder, where *ent.LocationWhereInput) (*ent.LocationConnection, error) {
	return r.client.Location.Query().Paginate(ctx, after, first, before, last, offset, ent.WithLocationOrder(orderBy), ent.WithLocationFilter(where.Filter))
}

// Media is the resolver for the media field.
func (r *queryResolver) Media(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy []*ent.MediumOrder, where *ent.MediumWhereInput) (*ent.MediumConnection, error) {
	return r.client.Medium.Query().Paginate(ctx, after, first, before, last, offset, ent.WithMediumOrder(orderBy), ent.WithMediumFilter(where.Filter))
}

// Models is the resolver for the models field.
func (r *queryResolver) Models(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy []*ent.ModelOrder, where *ent.ModelWhereInput) (*ent.ModelConnection, error) {
	return r.client.Model.Query().Paginate(ctx, after, first, before, last, offset, ent.WithModelOrder(orderBy), ent.WithModelFilter(where.Filter))
}

// Monuments is the resolver for the monuments field.
func (r *queryResolver) Monuments(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy []*ent.MonumentOrder, where *ent.MonumentWhereInput) (*ent.MonumentConnection, error) {
	return r.client.Monument.Query().Paginate(ctx, after, first, before, last, offset, ent.WithMonumentOrder(orderBy), ent.WithMonumentFilter(where.Filter))
}

// Organizations is the resolver for the organizations field.
func (r *queryResolver) Organizations(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy *ent.OrganizationOrder, where *ent.OrganizationWhereInput) (*ent.OrganizationConnection, error) {
	return r.client.Organization.Query().Paginate(ctx, after, first, before, last, offset, ent.WithOrganizationOrder(orderBy), ent.WithOrganizationFilter(where.Filter))
}

// OrganizationTypes is the resolver for the organizationTypes field.
func (r *queryResolver) OrganizationTypes(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy []*ent.OrganizationTypeOrder, where *ent.OrganizationTypeWhereInput) (*ent.OrganizationTypeConnection, error) {
	return r.client.OrganizationType.Query().Paginate(ctx, after, first, before, last, offset, ent.WithOrganizationTypeOrder(orderBy), ent.WithOrganizationTypeFilter(where.Filter))
}

// Periods is the resolver for the periods field.
func (r *queryResolver) Periods(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy []*ent.PeriodOrder, where *ent.PeriodWhereInput) (*ent.PeriodConnection, error) {
	return r.client.Period.Query().Paginate(ctx, after, first, before, last, offset, ent.WithPeriodOrder(orderBy), ent.WithPeriodFilter(where.Filter))
}

// Persons is the resolver for the persons field.
func (r *queryResolver) Persons(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy []*ent.PersonOrder, where *ent.PersonWhereInput) (*ent.PersonConnection, error) {
	return r.client.Person.Query().Paginate(ctx, after, first, before, last, offset, ent.WithPersonOrder(orderBy), ent.WithPersonFilter(where.Filter))
}

// PersonRoles is the resolver for the personRoles field.
func (r *queryResolver) PersonRoles(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy []*ent.PersonRoleOrder, where *ent.PersonRoleWhereInput) (*ent.PersonRoleConnection, error) {
	return r.client.PersonRole.Query().Paginate(ctx, after, first, before, last, offset, ent.WithPersonRoleOrder(orderBy), ent.WithPersonRoleFilter(where.Filter))
}

// Projects is the resolver for the projects field.
func (r *queryResolver) Projects(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy []*ent.ProjectOrder, where *ent.ProjectWhereInput) (*ent.ProjectConnection, error) {
	return r.client.Project.Query().Paginate(ctx, after, first, before, last, offset, ent.WithProjectOrder(orderBy), ent.WithProjectFilter(where.Filter))
}

// ProjectTypes is the resolver for the projectTypes field.
func (r *queryResolver) ProjectTypes(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy []*ent.ProjectTypeOrder, where *ent.ProjectTypeWhereInput) (*ent.ProjectTypeConnection, error) {
	return r.client.ProjectType.Query().Paginate(ctx, after, first, before, last, offset, ent.WithProjectTypeOrder(orderBy), ent.WithProjectTypeFilter(where.Filter))
}

// ProtectedAreas is the resolver for the protectedAreas field.
func (r *queryResolver) ProtectedAreas(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy []*ent.ProtectedAreaOrder, where *ent.ProtectedAreaWhereInput) (*ent.ProtectedAreaConnection, error) {
	return r.client.ProtectedArea.Query().Paginate(ctx, after, first, before, last, offset, ent.WithProtectedAreaOrder(orderBy), ent.WithProtectedAreaFilter(where.Filter))
}

// ProtectedAreaCategories is the resolver for the protectedAreaCategories field.
func (r *queryResolver) ProtectedAreaCategories(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy []*ent.ProtectedAreaCategoryOrder, where *ent.ProtectedAreaCategoryWhereInput) (*ent.ProtectedAreaCategoryConnection, error) {
	return r.client.ProtectedAreaCategory.Query().Paginate(ctx, after, first, before, last, offset, ent.WithProtectedAreaCategoryOrder(orderBy), ent.WithProtectedAreaCategoryFilter(where.Filter))
}

// ProtectedAreaPictures is the resolver for the protectedAreaPictures field.
func (r *queryResolver) ProtectedAreaPictures(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy []*ent.ProtectedAreaPictureOrder, where *ent.ProtectedAreaPictureWhereInput) (*ent.ProtectedAreaPictureConnection, error) {
	return r.client.ProtectedAreaPicture.Query().Paginate(ctx, after, first, before, last, offset, ent.WithProtectedAreaPictureOrder(orderBy), ent.WithProtectedAreaPictureFilter(where.Filter))
}

// Publications is the resolver for the publications field.
func (r *queryResolver) Publications(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy []*ent.PublicationOrder, where *ent.PublicationWhereInput) (*ent.PublicationConnection, error) {
	return r.client.Publication.Query().Paginate(ctx, after, first, before, last, offset, ent.WithPublicationOrder(orderBy), ent.WithPublicationFilter(where.Filter))
}

// Publishers is the resolver for the publishers field.
func (r *queryResolver) Publishers(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy []*ent.PublisherOrder, where *ent.PublisherWhereInput) (*ent.PublisherConnection, error) {
	return r.client.Publisher.Query().Paginate(ctx, after, first, before, last, offset, ent.WithPublisherOrder(orderBy), ent.WithPublisherFilter(where.Filter))
}

// Regions is the resolver for the regions field.
func (r *queryResolver) Regions(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy []*ent.RegionOrder, where *ent.RegionWhereInput) (*ent.RegionConnection, error) {
	return r.client.Region.Query().Paginate(ctx, after, first, before, last, offset, ent.WithRegionOrder(orderBy), ent.WithRegionFilter(where.Filter))
}

// Sets is the resolver for the sets field.
func (r *queryResolver) Sets(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy []*ent.SetOrder, where *ent.SetWhereInput) (*ent.SetConnection, error) {
	return r.client.Set.Query().Paginate(ctx, after, first, before, last, offset, ent.WithSetOrder(orderBy), ent.WithSetFilter(where.Filter))
}

// Settlements is the resolver for the settlements field.
func (r *queryResolver) Settlements(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy []*ent.SettlementOrder, where *ent.SettlementWhereInput) (*ent.SettlementConnection, error) {
	return r.client.Settlement.Query().Paginate(ctx, after, first, before, last, offset, ent.WithSettlementOrder(orderBy), ent.WithSettlementFilter(where.Filter))
}

// Techniques is the resolver for the techniques field.
func (r *queryResolver) Techniques(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy []*ent.TechniqueOrder, where *ent.TechniqueWhereInput) (*ent.TechniqueConnection, error) {
	return r.client.Technique.Query().Paginate(ctx, after, first, before, last, offset, ent.WithTechniqueOrder(orderBy), ent.WithTechniqueFilter(where.Filter))
}

// Query returns gql.QueryResolver implementation.
func (r *Resolver) Query() gql.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
