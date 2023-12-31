package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.38

import (
	"context"
	"fmt"

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

// DendrochronologicalAnalyses is the resolver for the dendrochronologicalAnalyses field.
func (r *queryResolver) DendrochronologicalAnalyses(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy []*ent.DendrochronologicalAnalysisOrder, where *ent.DendrochronologicalAnalysisWhereInput) (*ent.DendrochronologicalAnalysisConnection, error) {
	return r.client.DendrochronologicalAnalysis.Query().Paginate(ctx, after, first, before, last, offset, ent.WithDendrochronologicalAnalysisOrder(orderBy), ent.WithDendrochronologicalAnalysisFilter(where.Filter))
}

// Dendrochronologies is the resolver for the dendrochronologies field.
func (r *queryResolver) Dendrochronologies(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy []*ent.DendrochronologyOrder, where *ent.DendrochronologyWhereInput) (*ent.DendrochronologyConnection, error) {
	return r.client.Dendrochronology.Query().Paginate(ctx, after, first, before, last, offset, ent.WithDendrochronologyOrder(orderBy), ent.WithDendrochronologyFilter(where.Filter))
}

// Districts is the resolver for the districts field.
func (r *queryResolver) Districts(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy []*ent.DistrictOrder, where *ent.DistrictWhereInput) (*ent.DistrictConnection, error) {
	return r.client.District.Query().Paginate(ctx, after, first, before, last, offset, ent.WithDistrictOrder(orderBy), ent.WithDistrictFilter(where.Filter))
}

// EthnosSlice is the resolver for the ethnosSlice field.
func (r *queryResolver) EthnosSlice(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy []*ent.EthnosOrder, where *ent.EthnosWhereInput) (*ent.EthnosConnection, error) {
	return r.client.Ethnos.Query().Paginate(ctx, after, first, before, last, offset, ent.WithEthnosOrder(orderBy), ent.WithEthnosFilter(where.Filter))
}

// FamiliaSlice is the resolver for the familiaSlice field.
func (r *queryResolver) FamiliaSlice(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy []*ent.FamiliaOrder, where *ent.FamiliaWhereInput) (*ent.FamiliaConnection, error) {
	return r.client.Familia.Query().Paginate(ctx, after, first, before, last, offset, ent.WithFamiliaOrder(orderBy), ent.WithFamiliaFilter(where.Filter))
}

// Favourites is the resolver for the favourites field.
func (r *queryResolver) Favourites(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy []*ent.FavouriteOrder, where *ent.FavouriteWhereInput) (*ent.FavouriteConnection, error) {
	return r.client.Favourite.Query().Paginate(ctx, after, first, before, last, offset, ent.WithFavouriteOrder(orderBy), ent.WithFavouriteFilter(where.Filter))
}

// GenusSlice is the resolver for the genusSlice field.
func (r *queryResolver) GenusSlice(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy []*ent.GenusOrder, where *ent.GenusWhereInput) (*ent.GenusConnection, error) {
	return r.client.Genus.Query().Paginate(ctx, after, first, before, last, offset, ent.WithGenusOrder(orderBy), ent.WithGenusFilter(where.Filter))
}

// Groups is the resolver for the groups field.
func (r *queryResolver) Groups(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy []*ent.GroupOrder, where *ent.GroupWhereInput) (*ent.GroupConnection, error) {
	return r.client.Group.Query().Paginate(ctx, after, first, before, last, offset, ent.WithGroupOrder(orderBy), ent.WithGroupFilter(where.Filter))
}

// Herbaria is the resolver for the herbaria field.
func (r *queryResolver) Herbaria(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy []*ent.HerbariumOrder, where *ent.HerbariumWhereInput) (*ent.HerbariumConnection, error) {
	return r.client.Herbarium.Query().Paginate(ctx, after, first, before, last, offset, ent.WithHerbariumOrder(orderBy), ent.WithHerbariumFilter(where.Filter))
}

// Interviews is the resolver for the interviews field.
func (r *queryResolver) Interviews(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy []*ent.InterviewOrder, where *ent.InterviewWhereInput) (*ent.InterviewConnection, error) {
	panic(fmt.Errorf("not implemented: Interviews - interviews"))
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

// Mounds is the resolver for the mounds field.
func (r *queryResolver) Mounds(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy []*ent.MoundOrder, where *ent.MoundWhereInput) (*ent.MoundConnection, error) {
	return r.client.Mound.Query().Paginate(ctx, after, first, before, last, offset, ent.WithMoundOrder(orderBy), ent.WithMoundFilter(where.Filter))
}

// Organizations is the resolver for the organizations field.
func (r *queryResolver) Organizations(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy *ent.OrganizationOrder, where *ent.OrganizationWhereInput) (*ent.OrganizationConnection, error) {
	return r.client.Organization.Query().Paginate(ctx, after, first, before, last, offset, ent.WithOrganizationOrder(orderBy), ent.WithOrganizationFilter(where.Filter))
}

// Periodicals is the resolver for the periodicals field.
func (r *queryResolver) Periodicals(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy []*ent.PeriodicalOrder, where *ent.PeriodicalWhereInput) (*ent.PeriodicalConnection, error) {
	return r.client.Periodical.Query().Paginate(ctx, after, first, before, last, offset, ent.WithPeriodicalOrder(orderBy), ent.WithPeriodicalFilter(where.Filter))
}

// Persons is the resolver for the persons field.
func (r *queryResolver) Persons(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy []*ent.PersonOrder, where *ent.PersonWhereInput) (*ent.PersonConnection, error) {
	return r.client.Person.Query().Paginate(ctx, after, first, before, last, offset, ent.WithPersonOrder(orderBy), ent.WithPersonFilter(where.Filter))
}

// PersonalCollections is the resolver for the personalCollections field.
func (r *queryResolver) PersonalCollections(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy []*ent.PersonalCollectionOrder, where *ent.PersonalCollectionWhereInput) (*ent.PersonalCollectionConnection, error) {
	return r.client.PersonalCollection.Query().Paginate(ctx, after, first, before, last, offset, ent.WithPersonalCollectionOrder(orderBy), ent.WithPersonalCollectionFilter(where.Filter))
}

// Petroglyphs is the resolver for the petroglyphs field.
func (r *queryResolver) Petroglyphs(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy []*ent.PetroglyphOrder, where *ent.PetroglyphWhereInput) (*ent.PetroglyphConnection, error) {
	return r.client.Petroglyph.Query().Paginate(ctx, after, first, before, last, offset, ent.WithPetroglyphOrder(orderBy), ent.WithPetroglyphFilter(where.Filter))
}

// Projects is the resolver for the projects field.
func (r *queryResolver) Projects(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy []*ent.ProjectOrder, where *ent.ProjectWhereInput) (*ent.ProjectConnection, error) {
	return r.client.Project.Query().Paginate(ctx, after, first, before, last, offset, ent.WithProjectOrder(orderBy), ent.WithProjectFilter(where.Filter))
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

// SpeciesSlice is the resolver for the speciesSlice field.
func (r *queryResolver) SpeciesSlice(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy []*ent.SpeciesOrder, where *ent.SpeciesWhereInput) (*ent.SpeciesConnection, error) {
	return r.client.Species.Query().Paginate(ctx, after, first, before, last, offset, ent.WithSpeciesOrder(orderBy), ent.WithSpeciesFilter(where.Filter))
}

// Techniques is the resolver for the techniques field.
func (r *queryResolver) Techniques(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy []*ent.TechniqueOrder, where *ent.TechniqueWhereInput) (*ent.TechniqueConnection, error) {
	return r.client.Technique.Query().Paginate(ctx, after, first, before, last, offset, ent.WithTechniqueOrder(orderBy), ent.WithTechniqueFilter(where.Filter))
}

// Visits is the resolver for the visits field.
func (r *queryResolver) Visits(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, offset *int, orderBy []*ent.VisitOrder, where *ent.VisitWhereInput) (*ent.VisitConnection, error) {
	panic(fmt.Errorf("not implemented: Visits - visits"))
}

// Query returns gql.QueryResolver implementation.
func (r *Resolver) Query() gql.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
