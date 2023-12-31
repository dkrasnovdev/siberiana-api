package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.38

import (
	"context"
	"fmt"

	"github.com/dkrasnovdev/siberiana-api/ent"
	"github.com/dkrasnovdev/siberiana-api/ent/personalcollection"
	"github.com/dkrasnovdev/siberiana-api/internal/ent/privacy"
	"github.com/dkrasnovdev/siberiana-api/internal/gql"
)

// CreateArtifact is the resolver for the createArtifact field.
func (r *mutationResolver) CreateArtifact(ctx context.Context, input ent.CreateArtifactInput) (*ent.Artifact, error) {
	client := ent.FromContext(ctx)
	return client.Artifact.Create().SetInput(input).Save(ctx)
}

// UpdateArtifact is the resolver for the updateArtifact field.
func (r *mutationResolver) UpdateArtifact(ctx context.Context, id int, input ent.UpdateArtifactInput) (*ent.Artifact, error) {
	client := ent.FromContext(ctx)
	return client.Artifact.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeleteArtifact is the resolver for the deleteArtifact field.
func (r *mutationResolver) DeleteArtifact(ctx context.Context, id int) (string, error) {
	client := ent.FromContext(ctx)
	err := client.Artifact.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return "", err
	}
	return "Artifact has been deleted", nil
}

// CreateArt is the resolver for the createArt field.
func (r *mutationResolver) CreateArt(ctx context.Context, input ent.CreateArtInput) (*ent.Art, error) {
	client := ent.FromContext(ctx)
	return client.Art.Create().SetInput(input).Save(ctx)
}

// UpdateArt is the resolver for the updateArt field.
func (r *mutationResolver) UpdateArt(ctx context.Context, id int, input ent.UpdateArtInput) (*ent.Art, error) {
	client := ent.FromContext(ctx)
	return client.Art.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeleteArt is the resolver for the deleteArt field.
func (r *mutationResolver) DeleteArt(ctx context.Context, id int) (string, error) {
	client := ent.FromContext(ctx)
	err := client.Art.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return "", err
	}
	return "Art has been deleted", nil
}

// CreateArtGenre is the resolver for the createArtGenre field.
func (r *mutationResolver) CreateArtGenre(ctx context.Context, input ent.CreateArtGenreInput) (*ent.ArtGenre, error) {
	client := ent.FromContext(ctx)
	return client.ArtGenre.Create().SetInput(input).Save(ctx)
}

// UpdateArtGenre is the resolver for the updateArtGenre field.
func (r *mutationResolver) UpdateArtGenre(ctx context.Context, id int, input ent.UpdateArtGenreInput) (*ent.ArtGenre, error) {
	client := ent.FromContext(ctx)
	return client.ArtGenre.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeleteArtGenre is the resolver for the deleteArtGenre field.
func (r *mutationResolver) DeleteArtGenre(ctx context.Context, id int) (string, error) {
	client := ent.FromContext(ctx)
	err := client.ArtGenre.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return "", err
	}
	return "ArtGenre has been deleted", nil
}

// CreateArtStyle is the resolver for the createArtStyle field.
func (r *mutationResolver) CreateArtStyle(ctx context.Context, input ent.CreateArtStyleInput) (*ent.ArtStyle, error) {
	client := ent.FromContext(ctx)
	return client.ArtStyle.Create().SetInput(input).Save(ctx)
}

// UpdateArtStyle is the resolver for the updateArtStyle field.
func (r *mutationResolver) UpdateArtStyle(ctx context.Context, id int, input ent.UpdateArtStyleInput) (*ent.ArtStyle, error) {
	client := ent.FromContext(ctx)
	return client.ArtStyle.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeleteArtStyle is the resolver for the deleteArtStyle field.
func (r *mutationResolver) DeleteArtStyle(ctx context.Context, id int) (string, error) {
	client := ent.FromContext(ctx)
	err := client.ArtStyle.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return "", err
	}
	return "ArtStyle has been deleted", nil
}

// CreateBook is the resolver for the createBook field.
func (r *mutationResolver) CreateBook(ctx context.Context, input ent.CreateBookInput) (*ent.Book, error) {
	client := ent.FromContext(ctx)
	return client.Book.Create().SetInput(input).Save(ctx)
}

// UpdateBook is the resolver for the updateBook field.
func (r *mutationResolver) UpdateBook(ctx context.Context, id int, input ent.UpdateBookInput) (*ent.Book, error) {
	client := ent.FromContext(ctx)
	return client.Book.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeleteBook is the resolver for the deleteBook field.
func (r *mutationResolver) DeleteBook(ctx context.Context, id int) (string, error) {
	client := ent.FromContext(ctx)
	err := client.Book.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return "", err
	}
	return "Book has been deleted", nil
}

// CreateBookGenre is the resolver for the createBookGenre field.
func (r *mutationResolver) CreateBookGenre(ctx context.Context, input ent.CreateBookGenreInput) (*ent.BookGenre, error) {
	client := ent.FromContext(ctx)
	return client.BookGenre.Create().SetInput(input).Save(ctx)
}

// UpdateBookGenre is the resolver for the updateBookGenre field.
func (r *mutationResolver) UpdateBookGenre(ctx context.Context, id int, input ent.UpdateBookGenreInput) (*ent.BookGenre, error) {
	client := ent.FromContext(ctx)
	return client.BookGenre.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeleteBookGenre is the resolver for the deleteBookGenre field.
func (r *mutationResolver) DeleteBookGenre(ctx context.Context, id int) (string, error) {
	client := ent.FromContext(ctx)
	err := client.BookGenre.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return "", err
	}
	return "BookGenre has been deleted", nil
}

// CreateCategory is the resolver for the createCategory field.
func (r *mutationResolver) CreateCategory(ctx context.Context, input ent.CreateCategoryInput) (*ent.Category, error) {
	client := ent.FromContext(ctx)
	return client.Category.Create().SetInput(input).Save(ctx)
}

// UpdateCategory is the resolver for the updateCategory field.
func (r *mutationResolver) UpdateCategory(ctx context.Context, id int, input ent.UpdateCategoryInput) (*ent.Category, error) {
	client := ent.FromContext(ctx)
	return client.Category.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeleteCategory is the resolver for the deleteCategory field.
func (r *mutationResolver) DeleteCategory(ctx context.Context, id int) (string, error) {
	client := ent.FromContext(ctx)
	err := client.Category.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return "", err
	}
	return "Category has been deleted", nil
}

// CreateCollection is the resolver for the createCollection field.
func (r *mutationResolver) CreateCollection(ctx context.Context, input ent.CreateCollectionInput) (*ent.Collection, error) {
	client := ent.FromContext(ctx)
	return client.Collection.Create().SetInput(input).Save(ctx)
}

// UpdateCollection is the resolver for the updateCollection field.
func (r *mutationResolver) UpdateCollection(ctx context.Context, id int, input ent.UpdateCollectionInput) (*ent.Collection, error) {
	client := ent.FromContext(ctx)
	return client.Collection.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeleteCollection is the resolver for the deleteCollection field.
func (r *mutationResolver) DeleteCollection(ctx context.Context, id int) (string, error) {
	client := ent.FromContext(ctx)
	err := client.Collection.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return "", err
	}
	return "Collection has been deleted", nil
}

// CreateCountry is the resolver for the createCountry field.
func (r *mutationResolver) CreateCountry(ctx context.Context, input ent.CreateCountryInput) (*ent.Country, error) {
	client := ent.FromContext(ctx)
	return client.Country.Create().SetInput(input).Save(ctx)
}

// UpdateCountry is the resolver for the updateCountry field.
func (r *mutationResolver) UpdateCountry(ctx context.Context, id int, input ent.UpdateCountryInput) (*ent.Country, error) {
	client := ent.FromContext(ctx)
	return client.Country.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeleteCountry is the resolver for the deleteCountry field.
func (r *mutationResolver) DeleteCountry(ctx context.Context, id int) (string, error) {
	client := ent.FromContext(ctx)
	err := client.Country.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return "", err
	}
	return "Country has been deleted", nil
}

// CreateCulture is the resolver for the createCulture field.
func (r *mutationResolver) CreateCulture(ctx context.Context, input ent.CreateCultureInput) (*ent.Culture, error) {
	client := ent.FromContext(ctx)
	return client.Culture.Create().SetInput(input).Save(ctx)
}

// UpdateCulture is the resolver for the updateCulture field.
func (r *mutationResolver) UpdateCulture(ctx context.Context, id int, input ent.UpdateCultureInput) (*ent.Culture, error) {
	client := ent.FromContext(ctx)
	return client.Culture.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeleteCulture is the resolver for the deleteCulture field.
func (r *mutationResolver) DeleteCulture(ctx context.Context, id int) (string, error) {
	client := ent.FromContext(ctx)
	err := client.Culture.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return "", err
	}
	return "Culture has been deleted", nil
}

// CreateFavourite is the resolver for the createFavourite field.
func (r *mutationResolver) CreateFavourite(ctx context.Context, input ent.CreateFavouriteInput) (*ent.Favourite, error) {
	client := ent.FromContext(ctx)
	return client.Favourite.Create().SetInput(input).Save(ctx)
}

// UpdateFavourite is the resolver for the updateFavourite field.
func (r *mutationResolver) UpdateFavourite(ctx context.Context, id int, input ent.UpdateFavouriteInput) (*ent.Favourite, error) {
	client := ent.FromContext(ctx)
	return client.Favourite.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeleteFavourite is the resolver for the deleteFavourite field.
func (r *mutationResolver) DeleteFavourite(ctx context.Context, id int) (string, error) {
	client := ent.FromContext(ctx)
	err := client.Favourite.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return "", err
	}
	return "Favourite has been deleted", nil
}

// CreateDendrochronology is the resolver for the createDendrochronology field.
func (r *mutationResolver) CreateDendrochronology(ctx context.Context, input ent.CreateDendrochronologyInput) (*ent.Dendrochronology, error) {
	client := ent.FromContext(ctx)
	return client.Dendrochronology.Create().SetInput(input).Save(ctx)
}

// UpdateDendrochronology is the resolver for the updateDendrochronology field.
func (r *mutationResolver) UpdateDendrochronology(ctx context.Context, id int, input ent.UpdateDendrochronologyInput) (*ent.Dendrochronology, error) {
	client := ent.FromContext(ctx)
	return client.Dendrochronology.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeleteDendrochronology is the resolver for the deleteDendrochronology field.
func (r *mutationResolver) DeleteDendrochronology(ctx context.Context, id int) (string, error) {
	client := ent.FromContext(ctx)
	err := client.Dendrochronology.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return "", err
	}
	return "Dendrochronology has been deleted", nil
}

// CreateDendrochronologicalAnalysis is the resolver for the createDendrochronologicalAnalysis field.
func (r *mutationResolver) CreateDendrochronologicalAnalysis(ctx context.Context, input ent.CreateDendrochronologicalAnalysisInput) (*ent.DendrochronologicalAnalysis, error) {
	client := ent.FromContext(ctx)
	return client.DendrochronologicalAnalysis.Create().SetInput(input).Save(ctx)
}

// UpdateDendrochronologicalAnalysis is the resolver for the updateDendrochronologicalAnalysis field.
func (r *mutationResolver) UpdateDendrochronologicalAnalysis(ctx context.Context, id int, input ent.UpdateDendrochronologicalAnalysisInput) (*ent.DendrochronologicalAnalysis, error) {
	client := ent.FromContext(ctx)
	return client.DendrochronologicalAnalysis.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeleteDendrochronologicalAnalysis is the resolver for the deleteDendrochronologicalAnalysis field.
func (r *mutationResolver) DeleteDendrochronologicalAnalysis(ctx context.Context, id int) (string, error) {
	client := ent.FromContext(ctx)
	err := client.DendrochronologicalAnalysis.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return "", err
	}
	return "DendrochronologicalAnalysis has been deleted", nil
}

// CreateDistrict is the resolver for the createDistrict field.
func (r *mutationResolver) CreateDistrict(ctx context.Context, input ent.CreateDistrictInput) (*ent.District, error) {
	client := ent.FromContext(ctx)
	return client.District.Create().SetInput(input).Save(ctx)
}

// UpdateDistrict is the resolver for the updateDistrict field.
func (r *mutationResolver) UpdateDistrict(ctx context.Context, id int, input ent.UpdateDistrictInput) (*ent.District, error) {
	client := ent.FromContext(ctx)
	return client.District.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeleteDistrict is the resolver for the deleteDistrict field.
func (r *mutationResolver) DeleteDistrict(ctx context.Context, id int) (string, error) {
	client := ent.FromContext(ctx)
	err := client.District.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return "", err
	}
	return "District has been deleted", nil
}

// CreateEthnos is the resolver for the createEthnos field.
func (r *mutationResolver) CreateEthnos(ctx context.Context, input ent.CreateEthnosInput) (*ent.Ethnos, error) {
	client := ent.FromContext(ctx)
	return client.Ethnos.Create().SetInput(input).Save(ctx)
}

// UpdateEthnos is the resolver for the updateEthnos field.
func (r *mutationResolver) UpdateEthnos(ctx context.Context, id int, input ent.UpdateEthnosInput) (*ent.Ethnos, error) {
	client := ent.FromContext(ctx)
	return client.Ethnos.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeleteEthnos is the resolver for the deleteEthnos field.
func (r *mutationResolver) DeleteEthnos(ctx context.Context, id int) (string, error) {
	client := ent.FromContext(ctx)
	err := client.Ethnos.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return "", err
	}
	return "Ethnos has been deleted", nil
}

// CreateFamilia is the resolver for the createFamilia field.
func (r *mutationResolver) CreateFamilia(ctx context.Context, input ent.CreateFamiliaInput) (*ent.Familia, error) {
	client := ent.FromContext(ctx)
	return client.Familia.Create().SetInput(input).Save(ctx)
}

// UpdateFamilia is the resolver for the updateFamilia field.
func (r *mutationResolver) UpdateFamilia(ctx context.Context, id int, input ent.UpdateFamiliaInput) (*ent.Familia, error) {
	client := ent.FromContext(ctx)
	return client.Familia.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeleteFamilia is the resolver for the deleteFamilia field.
func (r *mutationResolver) DeleteFamilia(ctx context.Context, id int) (string, error) {
	client := ent.FromContext(ctx)
	err := client.Familia.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return "", err
	}
	return "Familia has been deleted", nil
}

// CreateGenus is the resolver for the createGenus field.
func (r *mutationResolver) CreateGenus(ctx context.Context, input ent.CreateGenusInput) (*ent.Genus, error) {
	client := ent.FromContext(ctx)
	return client.Genus.Create().SetInput(input).Save(ctx)
}

// UpdateGenus is the resolver for the updateGenus field.
func (r *mutationResolver) UpdateGenus(ctx context.Context, id int, input ent.UpdateGenusInput) (*ent.Genus, error) {
	client := ent.FromContext(ctx)
	return client.Genus.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeleteGenus is the resolver for the deleteGenus field.
func (r *mutationResolver) DeleteGenus(ctx context.Context, id int) (string, error) {
	client := ent.FromContext(ctx)
	err := client.Genus.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return "", err
	}
	return "Genus has been deleted", nil
}

// CreateGroup is the resolver for the createGroup field.
func (r *mutationResolver) CreateGroup(ctx context.Context, input ent.CreateGroupInput) (*ent.Group, error) {
	client := ent.FromContext(ctx)
	return client.Group.Create().SetInput(input).Save(ctx)
}

// UpdateGroup is the resolver for the updateGroup field.
func (r *mutationResolver) UpdateGroup(ctx context.Context, id int, input ent.UpdateGroupInput) (*ent.Group, error) {
	client := ent.FromContext(ctx)
	return client.Group.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeleteGroup is the resolver for the deleteGroup field.
func (r *mutationResolver) DeleteGroup(ctx context.Context, id int) (string, error) {
	client := ent.FromContext(ctx)
	err := client.Group.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return "", err
	}
	return "Group has been deleted", nil
}

// CreateHerbarium is the resolver for the createHerbarium field.
func (r *mutationResolver) CreateHerbarium(ctx context.Context, input ent.CreateHerbariumInput) (*ent.Herbarium, error) {
	client := ent.FromContext(ctx)
	return client.Herbarium.Create().SetInput(input).Save(ctx)
}

// UpdateHerbarium is the resolver for the updateHerbarium field.
func (r *mutationResolver) UpdateHerbarium(ctx context.Context, id int, input ent.UpdateHerbariumInput) (*ent.Herbarium, error) {
	client := ent.FromContext(ctx)
	return client.Herbarium.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeleteHerbarium is the resolver for the deleteHerbarium field.
func (r *mutationResolver) DeleteHerbarium(ctx context.Context, id int) (string, error) {
	client := ent.FromContext(ctx)
	err := client.Herbarium.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return "", err
	}
	return "Herbarium has been deleted", nil
}

// CreateLicense is the resolver for the createLicense field.
func (r *mutationResolver) CreateLicense(ctx context.Context, input ent.CreateLicenseInput) (*ent.License, error) {
	client := ent.FromContext(ctx)
	return client.License.Create().SetInput(input).Save(ctx)
}

// UpdateLicense is the resolver for the updateLicense field.
func (r *mutationResolver) UpdateLicense(ctx context.Context, id int, input ent.UpdateLicenseInput) (*ent.License, error) {
	client := ent.FromContext(ctx)
	return client.License.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeleteLicense is the resolver for the deleteLicense field.
func (r *mutationResolver) DeleteLicense(ctx context.Context, id int) (string, error) {
	client := ent.FromContext(ctx)
	err := client.License.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return "", err
	}
	return "License has been deleted", nil
}

// CreateLocation is the resolver for the createLocation field.
func (r *mutationResolver) CreateLocation(ctx context.Context, input ent.CreateLocationInput) (*ent.Location, error) {
	client := ent.FromContext(ctx)
	return client.Location.Create().SetInput(input).Save(ctx)
}

// UpdateLocation is the resolver for the updateLocation field.
func (r *mutationResolver) UpdateLocation(ctx context.Context, id int, input ent.UpdateLocationInput) (*ent.Location, error) {
	client := ent.FromContext(ctx)
	return client.Location.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeleteLocation is the resolver for the deleteLocation field.
func (r *mutationResolver) DeleteLocation(ctx context.Context, id int) (string, error) {
	client := ent.FromContext(ctx)
	err := client.Location.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return "", err
	}
	return "Location has been deleted", nil
}

// CreateMedium is the resolver for the createMedium field.
func (r *mutationResolver) CreateMedium(ctx context.Context, input ent.CreateMediumInput) (*ent.Medium, error) {
	client := ent.FromContext(ctx)
	return client.Medium.Create().SetInput(input).Save(ctx)
}

// UpdateMedium is the resolver for the updateMedium field.
func (r *mutationResolver) UpdateMedium(ctx context.Context, id int, input ent.UpdateMediumInput) (*ent.Medium, error) {
	client := ent.FromContext(ctx)
	return client.Medium.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeleteMedium is the resolver for the deleteMedium field.
func (r *mutationResolver) DeleteMedium(ctx context.Context, id int) (string, error) {
	client := ent.FromContext(ctx)
	err := client.Medium.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return "", err
	}
	return "Medium has been deleted", nil
}

// CreateMound is the resolver for the createMound field.
func (r *mutationResolver) CreateMound(ctx context.Context, input ent.CreateMoundInput) (*ent.Mound, error) {
	client := ent.FromContext(ctx)
	return client.Mound.Create().SetInput(input).Save(ctx)
}

// UpdateMound is the resolver for the updateMound field.
func (r *mutationResolver) UpdateMound(ctx context.Context, id int, input ent.UpdateMoundInput) (*ent.Mound, error) {
	client := ent.FromContext(ctx)
	return client.Mound.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeleteMound is the resolver for the deleteMound field.
func (r *mutationResolver) DeleteMound(ctx context.Context, id int) (string, error) {
	client := ent.FromContext(ctx)
	err := client.Mound.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return "", err
	}
	return "Mound has been deleted", nil
}

// CreateModel is the resolver for the createModel field.
func (r *mutationResolver) CreateModel(ctx context.Context, input ent.CreateModelInput) (*ent.Model, error) {
	client := ent.FromContext(ctx)
	return client.Model.Create().SetInput(input).Save(ctx)
}

// UpdateModel is the resolver for the updateModel field.
func (r *mutationResolver) UpdateModel(ctx context.Context, id int, input ent.UpdateModelInput) (*ent.Model, error) {
	client := ent.FromContext(ctx)
	return client.Model.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeleteModel is the resolver for the deleteModel field.
func (r *mutationResolver) DeleteModel(ctx context.Context, id int) (string, error) {
	client := ent.FromContext(ctx)
	err := client.Model.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return "", err
	}
	return "Model has been deleted", nil
}

// CreateMonument is the resolver for the createMonument field.
func (r *mutationResolver) CreateMonument(ctx context.Context, input ent.CreateMonumentInput) (*ent.Monument, error) {
	client := ent.FromContext(ctx)
	return client.Monument.Create().SetInput(input).Save(ctx)
}

// UpdateMonument is the resolver for the updateMonument field.
func (r *mutationResolver) UpdateMonument(ctx context.Context, id int, input ent.UpdateMonumentInput) (*ent.Monument, error) {
	client := ent.FromContext(ctx)
	return client.Monument.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeleteMonument is the resolver for the deleteMonument field.
func (r *mutationResolver) DeleteMonument(ctx context.Context, id int) (string, error) {
	client := ent.FromContext(ctx)
	err := client.Monument.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return "", err
	}
	return "Monument has been deleted", nil
}

// CreateOrganization is the resolver for the createOrganization field.
func (r *mutationResolver) CreateOrganization(ctx context.Context, input ent.CreateOrganizationInput) (*ent.Organization, error) {
	client := ent.FromContext(ctx)
	return client.Organization.Create().SetInput(input).Save(ctx)
}

// UpdateOrganization is the resolver for the updateOrganization field.
func (r *mutationResolver) UpdateOrganization(ctx context.Context, id int, input ent.UpdateOrganizationInput) (*ent.Organization, error) {
	client := ent.FromContext(ctx)
	return client.Organization.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeleteOrganization is the resolver for the deleteOrganization field.
func (r *mutationResolver) DeleteOrganization(ctx context.Context, id int) (string, error) {
	client := ent.FromContext(ctx)
	err := client.Organization.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return "", err
	}
	return "Organization has been deleted", nil
}

// CreatePeriodical is the resolver for the createPeriodical field.
func (r *mutationResolver) CreatePeriodical(ctx context.Context, input ent.CreatePeriodicalInput) (*ent.Periodical, error) {
	client := ent.FromContext(ctx)
	return client.Periodical.Create().SetInput(input).Save(ctx)
}

// UpdatePeriodical is the resolver for the updatePeriodical field.
func (r *mutationResolver) UpdatePeriodical(ctx context.Context, id int, input ent.UpdatePeriodicalInput) (*ent.Periodical, error) {
	client := ent.FromContext(ctx)
	return client.Periodical.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeletePeriodical is the resolver for the deletePeriodical field.
func (r *mutationResolver) DeletePeriodical(ctx context.Context, id int) (string, error) {
	client := ent.FromContext(ctx)
	err := client.Periodical.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return "", err
	}
	return "Periodical has been deleted", nil
}

// CreatePetroglyph is the resolver for the createPetroglyph field.
func (r *mutationResolver) CreatePetroglyph(ctx context.Context, input ent.CreatePetroglyphInput) (*ent.Petroglyph, error) {
	client := ent.FromContext(ctx)
	return client.Petroglyph.Create().SetInput(input).Save(ctx)
}

// UpdatePetroglyph is the resolver for the updatePetroglyph field.
func (r *mutationResolver) UpdatePetroglyph(ctx context.Context, id int, input ent.UpdatePetroglyphInput) (*ent.Petroglyph, error) {
	client := ent.FromContext(ctx)
	return client.Petroglyph.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeletePetroglyph is the resolver for the deletePetroglyph field.
func (r *mutationResolver) DeletePetroglyph(ctx context.Context, id int) (string, error) {
	client := ent.FromContext(ctx)
	err := client.Petroglyph.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return "", err
	}
	return "Petroglyph has been deleted", nil
}

// CreatePerson is the resolver for the createPerson field.
func (r *mutationResolver) CreatePerson(ctx context.Context, input ent.CreatePersonInput) (*ent.Person, error) {
	client := ent.FromContext(ctx)
	return client.Person.Create().SetInput(input).Save(ctx)
}

// UpdatePerson is the resolver for the updatePerson field.
func (r *mutationResolver) UpdatePerson(ctx context.Context, id int, input ent.UpdatePersonInput) (*ent.Person, error) {
	client := ent.FromContext(ctx)
	return client.Person.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeletePerson is the resolver for the deletePerson field.
func (r *mutationResolver) DeletePerson(ctx context.Context, id int) (string, error) {
	client := ent.FromContext(ctx)
	err := client.Person.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return "", err
	}
	return "Person has been deleted", nil
}

// CreatePersonalCollection is the resolver for the createPersonalCollection field.
func (r *mutationResolver) CreatePersonalCollection(ctx context.Context, input ent.CreatePersonalCollectionInput) (*ent.PersonalCollection, error) {
	client := ent.FromContext(ctx)
	return client.PersonalCollection.Create().SetInput(input).Save(ctx)
}

// UpdatePersonalCollection is the resolver for the updatePersonalCollection field.
func (r *mutationResolver) UpdatePersonalCollection(ctx context.Context, id int, input ent.UpdatePersonalCollectionInput) (*ent.PersonalCollection, error) {
	client := ent.FromContext(ctx)
	return client.PersonalCollection.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeletePersonalCollection is the resolver for the deletePersonalCollection field.
func (r *mutationResolver) DeletePersonalCollection(ctx context.Context, id int) (string, error) {
	client := ent.FromContext(ctx)
	viewer := privacy.FromContext(ctx)
	if viewer == nil {
		return "", fmt.Errorf("Only logged in users can delete personal collections")
	}
	user := viewer.GetPreferredUsername()
	_, err := client.PersonalCollection.Query().
		Where(personalcollection.And(personalcollection.IDEQ(id), personalcollection.CreatedByEQ(user))).
		Only(ctx)
	if err != nil {
		return "", fmt.Errorf("Operation was not successful")
	}
	err = client.PersonalCollection.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return "", err
	}
	return "Personal Collection has been deleted", nil
}

// CreateProject is the resolver for the createProject field.
func (r *mutationResolver) CreateProject(ctx context.Context, input ent.CreateProjectInput) (*ent.Project, error) {
	client := ent.FromContext(ctx)
	return client.Project.Create().SetInput(input).Save(ctx)
}

// UpdateProject is the resolver for the updateProject field.
func (r *mutationResolver) UpdateProject(ctx context.Context, id int, input ent.UpdateProjectInput) (*ent.Project, error) {
	client := ent.FromContext(ctx)
	return client.Project.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeleteProject is the resolver for the deleteProject field.
func (r *mutationResolver) DeleteProject(ctx context.Context, id int) (string, error) {
	client := ent.FromContext(ctx)
	err := client.Project.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return "", err
	}
	return "Project has been deleted", nil
}

// CreateProtectedArea is the resolver for the createProtectedArea field.
func (r *mutationResolver) CreateProtectedArea(ctx context.Context, input ent.CreateProtectedAreaInput) (*ent.ProtectedArea, error) {
	client := ent.FromContext(ctx)
	return client.ProtectedArea.Create().SetInput(input).Save(ctx)
}

// UpdateProtectedArea is the resolver for the updateProtectedArea field.
func (r *mutationResolver) UpdateProtectedArea(ctx context.Context, id int, input ent.UpdateProtectedAreaInput) (*ent.ProtectedArea, error) {
	client := ent.FromContext(ctx)
	return client.ProtectedArea.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeleteProtectedArea is the resolver for the deleteProtectedArea field.
func (r *mutationResolver) DeleteProtectedArea(ctx context.Context, id int) (string, error) {
	client := ent.FromContext(ctx)
	err := client.ProtectedArea.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return "", err
	}
	return "ProtectedArea has been deleted", nil
}

// CreateProtectedAreaCategory is the resolver for the createProtectedAreaCategory field.
func (r *mutationResolver) CreateProtectedAreaCategory(ctx context.Context, input ent.CreateProtectedAreaCategoryInput) (*ent.ProtectedAreaCategory, error) {
	client := ent.FromContext(ctx)
	return client.ProtectedAreaCategory.Create().SetInput(input).Save(ctx)
}

// UpdateProtectedAreaCategory is the resolver for the updateProtectedAreaCategory field.
func (r *mutationResolver) UpdateProtectedAreaCategory(ctx context.Context, id int, input ent.UpdateProtectedAreaCategoryInput) (*ent.ProtectedAreaCategory, error) {
	client := ent.FromContext(ctx)
	return client.ProtectedAreaCategory.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeleteProtectedAreaCategory is the resolver for the deleteProtectedAreaCategory field.
func (r *mutationResolver) DeleteProtectedAreaCategory(ctx context.Context, id int) (string, error) {
	client := ent.FromContext(ctx)
	err := client.ProtectedAreaCategory.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return "", err
	}
	return "ProtectedAreaCategory has been deleted", nil
}

// CreateProtectedAreaPicture is the resolver for the createProtectedAreaPicture field.
func (r *mutationResolver) CreateProtectedAreaPicture(ctx context.Context, input ent.CreateProtectedAreaPictureInput) (*ent.ProtectedAreaPicture, error) {
	client := ent.FromContext(ctx)
	return client.ProtectedAreaPicture.Create().SetInput(input).Save(ctx)
}

// UpdateProtectedAreaPicture is the resolver for the updateProtectedAreaPicture field.
func (r *mutationResolver) UpdateProtectedAreaPicture(ctx context.Context, id int, input ent.UpdateProtectedAreaPictureInput) (*ent.ProtectedAreaPicture, error) {
	client := ent.FromContext(ctx)
	return client.ProtectedAreaPicture.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeleteProtectedAreaPicture is the resolver for the deleteProtectedAreaPicture field.
func (r *mutationResolver) DeleteProtectedAreaPicture(ctx context.Context, id int) (string, error) {
	client := ent.FromContext(ctx)
	err := client.ProtectedAreaPicture.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return "", err
	}
	return "ProtectedAreaPicture has been deleted", nil
}

// CreatePublication is the resolver for the createPublication field.
func (r *mutationResolver) CreatePublication(ctx context.Context, input ent.CreatePublicationInput) (*ent.Publication, error) {
	client := ent.FromContext(ctx)
	return client.Publication.Create().SetInput(input).Save(ctx)
}

// UpdatePublication is the resolver for the updatePublication field.
func (r *mutationResolver) UpdatePublication(ctx context.Context, id int, input ent.UpdatePublicationInput) (*ent.Publication, error) {
	client := ent.FromContext(ctx)
	return client.Publication.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeletePublication is the resolver for the deletePublication field.
func (r *mutationResolver) DeletePublication(ctx context.Context, id int) (string, error) {
	client := ent.FromContext(ctx)
	err := client.Publication.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return "", err
	}
	return "Publication has been deleted", nil
}

// CreatePublisher is the resolver for the createPublisher field.
func (r *mutationResolver) CreatePublisher(ctx context.Context, input ent.CreatePublisherInput) (*ent.Publisher, error) {
	client := ent.FromContext(ctx)
	return client.Publisher.Create().SetInput(input).Save(ctx)
}

// UpdatePublisher is the resolver for the updatePublisher field.
func (r *mutationResolver) UpdatePublisher(ctx context.Context, id int, input ent.UpdatePublisherInput) (*ent.Publisher, error) {
	client := ent.FromContext(ctx)
	return client.Publisher.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeletePublisher is the resolver for the deletePublisher field.
func (r *mutationResolver) DeletePublisher(ctx context.Context, id int) (string, error) {
	client := ent.FromContext(ctx)
	err := client.Publisher.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return "", err
	}
	return "Publisher has been deleted", nil
}

// CreateRegion is the resolver for the createRegion field.
func (r *mutationResolver) CreateRegion(ctx context.Context, input ent.CreateRegionInput) (*ent.Region, error) {
	client := ent.FromContext(ctx)
	return client.Region.Create().SetInput(input).Save(ctx)
}

// UpdateRegion is the resolver for the updateRegion field.
func (r *mutationResolver) UpdateRegion(ctx context.Context, id int, input ent.UpdateRegionInput) (*ent.Region, error) {
	client := ent.FromContext(ctx)
	return client.Region.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeleteRegion is the resolver for the deleteRegion field.
func (r *mutationResolver) DeleteRegion(ctx context.Context, id int) (string, error) {
	client := ent.FromContext(ctx)
	err := client.Region.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return "", err
	}
	return "Region has been deleted", nil
}

// CreateSet is the resolver for the createSet field.
func (r *mutationResolver) CreateSet(ctx context.Context, input ent.CreateSetInput) (*ent.Set, error) {
	client := ent.FromContext(ctx)
	return client.Set.Create().SetInput(input).Save(ctx)
}

// UpdateSet is the resolver for the updateSet field.
func (r *mutationResolver) UpdateSet(ctx context.Context, id int, input ent.UpdateSetInput) (*ent.Set, error) {
	client := ent.FromContext(ctx)
	return client.Set.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeleteSet is the resolver for the deleteSet field.
func (r *mutationResolver) DeleteSet(ctx context.Context, id int) (string, error) {
	client := ent.FromContext(ctx)
	err := client.Set.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return "", err
	}
	return "Set has been deleted", nil
}

// CreateSettlement is the resolver for the createSettlement field.
func (r *mutationResolver) CreateSettlement(ctx context.Context, input ent.CreateSettlementInput) (*ent.Settlement, error) {
	client := ent.FromContext(ctx)
	return client.Settlement.Create().SetInput(input).Save(ctx)
}

// UpdateSettlement is the resolver for the updateSettlement field.
func (r *mutationResolver) UpdateSettlement(ctx context.Context, id int, input ent.UpdateSettlementInput) (*ent.Settlement, error) {
	client := ent.FromContext(ctx)
	return client.Settlement.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeleteSettlement is the resolver for the deleteSettlement field.
func (r *mutationResolver) DeleteSettlement(ctx context.Context, id int) (string, error) {
	client := ent.FromContext(ctx)
	err := client.Settlement.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return "", err
	}
	return "Settlement has been deleted", nil
}

// CreateSpecies is the resolver for the createSpecies field.
func (r *mutationResolver) CreateSpecies(ctx context.Context, input ent.CreateSpeciesInput) (*ent.Species, error) {
	client := ent.FromContext(ctx)
	return client.Species.Create().SetInput(input).Save(ctx)
}

// UpdateSpecies is the resolver for the updateSpecies field.
func (r *mutationResolver) UpdateSpecies(ctx context.Context, id int, input ent.UpdateSpeciesInput) (*ent.Species, error) {
	client := ent.FromContext(ctx)
	return client.Species.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeleteSpecies is the resolver for the deleteSpecies field.
func (r *mutationResolver) DeleteSpecies(ctx context.Context, id int) (string, error) {
	client := ent.FromContext(ctx)
	err := client.Species.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return "", err
	}
	return "Species has been deleted", nil
}

// CreateTechnique is the resolver for the createTechnique field.
func (r *mutationResolver) CreateTechnique(ctx context.Context, input ent.CreateTechniqueInput) (*ent.Technique, error) {
	client := ent.FromContext(ctx)
	return client.Technique.Create().SetInput(input).Save(ctx)
}

// UpdateTechnique is the resolver for the updateTechnique field.
func (r *mutationResolver) UpdateTechnique(ctx context.Context, id int, input ent.UpdateTechniqueInput) (*ent.Technique, error) {
	client := ent.FromContext(ctx)
	return client.Technique.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeleteTechnique is the resolver for the deleteTechnique field.
func (r *mutationResolver) DeleteTechnique(ctx context.Context, id int) (string, error) {
	client := ent.FromContext(ctx)
	err := client.Technique.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return "", err
	}
	return "Technique has been deleted", nil
}

// Mutation returns gql.MutationResolver implementation.
func (r *Resolver) Mutation() gql.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
