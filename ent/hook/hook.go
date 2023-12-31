// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/dkrasnovdev/siberiana-api/ent"
)

// The ArtFunc type is an adapter to allow the use of ordinary
// function as Art mutator.
type ArtFunc func(context.Context, *ent.ArtMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ArtFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ArtMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ArtMutation", m)
}

// The ArtGenreFunc type is an adapter to allow the use of ordinary
// function as ArtGenre mutator.
type ArtGenreFunc func(context.Context, *ent.ArtGenreMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ArtGenreFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ArtGenreMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ArtGenreMutation", m)
}

// The ArtStyleFunc type is an adapter to allow the use of ordinary
// function as ArtStyle mutator.
type ArtStyleFunc func(context.Context, *ent.ArtStyleMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ArtStyleFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ArtStyleMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ArtStyleMutation", m)
}

// The ArtifactFunc type is an adapter to allow the use of ordinary
// function as Artifact mutator.
type ArtifactFunc func(context.Context, *ent.ArtifactMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ArtifactFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ArtifactMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ArtifactMutation", m)
}

// The AuditLogFunc type is an adapter to allow the use of ordinary
// function as AuditLog mutator.
type AuditLogFunc func(context.Context, *ent.AuditLogMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AuditLogFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.AuditLogMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AuditLogMutation", m)
}

// The BookFunc type is an adapter to allow the use of ordinary
// function as Book mutator.
type BookFunc func(context.Context, *ent.BookMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f BookFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.BookMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.BookMutation", m)
}

// The BookGenreFunc type is an adapter to allow the use of ordinary
// function as BookGenre mutator.
type BookGenreFunc func(context.Context, *ent.BookGenreMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f BookGenreFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.BookGenreMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.BookGenreMutation", m)
}

// The CategoryFunc type is an adapter to allow the use of ordinary
// function as Category mutator.
type CategoryFunc func(context.Context, *ent.CategoryMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CategoryFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.CategoryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CategoryMutation", m)
}

// The CollectionFunc type is an adapter to allow the use of ordinary
// function as Collection mutator.
type CollectionFunc func(context.Context, *ent.CollectionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CollectionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.CollectionMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CollectionMutation", m)
}

// The CountryFunc type is an adapter to allow the use of ordinary
// function as Country mutator.
type CountryFunc func(context.Context, *ent.CountryMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CountryFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.CountryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CountryMutation", m)
}

// The CultureFunc type is an adapter to allow the use of ordinary
// function as Culture mutator.
type CultureFunc func(context.Context, *ent.CultureMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CultureFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.CultureMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CultureMutation", m)
}

// The DendrochronologicalAnalysisFunc type is an adapter to allow the use of ordinary
// function as DendrochronologicalAnalysis mutator.
type DendrochronologicalAnalysisFunc func(context.Context, *ent.DendrochronologicalAnalysisMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DendrochronologicalAnalysisFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.DendrochronologicalAnalysisMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DendrochronologicalAnalysisMutation", m)
}

// The DendrochronologyFunc type is an adapter to allow the use of ordinary
// function as Dendrochronology mutator.
type DendrochronologyFunc func(context.Context, *ent.DendrochronologyMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DendrochronologyFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.DendrochronologyMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DendrochronologyMutation", m)
}

// The DistrictFunc type is an adapter to allow the use of ordinary
// function as District mutator.
type DistrictFunc func(context.Context, *ent.DistrictMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DistrictFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.DistrictMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DistrictMutation", m)
}

// The EthnosFunc type is an adapter to allow the use of ordinary
// function as Ethnos mutator.
type EthnosFunc func(context.Context, *ent.EthnosMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f EthnosFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.EthnosMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.EthnosMutation", m)
}

// The FamiliaFunc type is an adapter to allow the use of ordinary
// function as Familia mutator.
type FamiliaFunc func(context.Context, *ent.FamiliaMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f FamiliaFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.FamiliaMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.FamiliaMutation", m)
}

// The FavouriteFunc type is an adapter to allow the use of ordinary
// function as Favourite mutator.
type FavouriteFunc func(context.Context, *ent.FavouriteMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f FavouriteFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.FavouriteMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.FavouriteMutation", m)
}

// The GenusFunc type is an adapter to allow the use of ordinary
// function as Genus mutator.
type GenusFunc func(context.Context, *ent.GenusMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GenusFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.GenusMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GenusMutation", m)
}

// The GroupFunc type is an adapter to allow the use of ordinary
// function as Group mutator.
type GroupFunc func(context.Context, *ent.GroupMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GroupFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.GroupMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GroupMutation", m)
}

// The HerbariumFunc type is an adapter to allow the use of ordinary
// function as Herbarium mutator.
type HerbariumFunc func(context.Context, *ent.HerbariumMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f HerbariumFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.HerbariumMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.HerbariumMutation", m)
}

// The InterviewFunc type is an adapter to allow the use of ordinary
// function as Interview mutator.
type InterviewFunc func(context.Context, *ent.InterviewMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f InterviewFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.InterviewMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.InterviewMutation", m)
}

// The KeywordFunc type is an adapter to allow the use of ordinary
// function as Keyword mutator.
type KeywordFunc func(context.Context, *ent.KeywordMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f KeywordFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.KeywordMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.KeywordMutation", m)
}

// The LicenseFunc type is an adapter to allow the use of ordinary
// function as License mutator.
type LicenseFunc func(context.Context, *ent.LicenseMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f LicenseFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.LicenseMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.LicenseMutation", m)
}

// The LocationFunc type is an adapter to allow the use of ordinary
// function as Location mutator.
type LocationFunc func(context.Context, *ent.LocationMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f LocationFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.LocationMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.LocationMutation", m)
}

// The MediumFunc type is an adapter to allow the use of ordinary
// function as Medium mutator.
type MediumFunc func(context.Context, *ent.MediumMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f MediumFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.MediumMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.MediumMutation", m)
}

// The ModelFunc type is an adapter to allow the use of ordinary
// function as Model mutator.
type ModelFunc func(context.Context, *ent.ModelMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ModelFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ModelMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ModelMutation", m)
}

// The MonumentFunc type is an adapter to allow the use of ordinary
// function as Monument mutator.
type MonumentFunc func(context.Context, *ent.MonumentMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f MonumentFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.MonumentMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.MonumentMutation", m)
}

// The MoundFunc type is an adapter to allow the use of ordinary
// function as Mound mutator.
type MoundFunc func(context.Context, *ent.MoundMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f MoundFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.MoundMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.MoundMutation", m)
}

// The OrganizationFunc type is an adapter to allow the use of ordinary
// function as Organization mutator.
type OrganizationFunc func(context.Context, *ent.OrganizationMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f OrganizationFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.OrganizationMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.OrganizationMutation", m)
}

// The PeriodicalFunc type is an adapter to allow the use of ordinary
// function as Periodical mutator.
type PeriodicalFunc func(context.Context, *ent.PeriodicalMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PeriodicalFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.PeriodicalMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PeriodicalMutation", m)
}

// The PersonFunc type is an adapter to allow the use of ordinary
// function as Person mutator.
type PersonFunc func(context.Context, *ent.PersonMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PersonFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.PersonMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PersonMutation", m)
}

// The PersonalCollectionFunc type is an adapter to allow the use of ordinary
// function as PersonalCollection mutator.
type PersonalCollectionFunc func(context.Context, *ent.PersonalCollectionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PersonalCollectionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.PersonalCollectionMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PersonalCollectionMutation", m)
}

// The PetroglyphFunc type is an adapter to allow the use of ordinary
// function as Petroglyph mutator.
type PetroglyphFunc func(context.Context, *ent.PetroglyphMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PetroglyphFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.PetroglyphMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PetroglyphMutation", m)
}

// The ProjectFunc type is an adapter to allow the use of ordinary
// function as Project mutator.
type ProjectFunc func(context.Context, *ent.ProjectMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ProjectFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ProjectMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ProjectMutation", m)
}

// The ProtectedAreaFunc type is an adapter to allow the use of ordinary
// function as ProtectedArea mutator.
type ProtectedAreaFunc func(context.Context, *ent.ProtectedAreaMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ProtectedAreaFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ProtectedAreaMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ProtectedAreaMutation", m)
}

// The ProtectedAreaCategoryFunc type is an adapter to allow the use of ordinary
// function as ProtectedAreaCategory mutator.
type ProtectedAreaCategoryFunc func(context.Context, *ent.ProtectedAreaCategoryMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ProtectedAreaCategoryFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ProtectedAreaCategoryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ProtectedAreaCategoryMutation", m)
}

// The ProtectedAreaPictureFunc type is an adapter to allow the use of ordinary
// function as ProtectedAreaPicture mutator.
type ProtectedAreaPictureFunc func(context.Context, *ent.ProtectedAreaPictureMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ProtectedAreaPictureFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ProtectedAreaPictureMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ProtectedAreaPictureMutation", m)
}

// The PublicationFunc type is an adapter to allow the use of ordinary
// function as Publication mutator.
type PublicationFunc func(context.Context, *ent.PublicationMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PublicationFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.PublicationMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PublicationMutation", m)
}

// The PublisherFunc type is an adapter to allow the use of ordinary
// function as Publisher mutator.
type PublisherFunc func(context.Context, *ent.PublisherMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PublisherFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.PublisherMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PublisherMutation", m)
}

// The RegionFunc type is an adapter to allow the use of ordinary
// function as Region mutator.
type RegionFunc func(context.Context, *ent.RegionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f RegionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.RegionMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.RegionMutation", m)
}

// The SetFunc type is an adapter to allow the use of ordinary
// function as Set mutator.
type SetFunc func(context.Context, *ent.SetMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SetFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.SetMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SetMutation", m)
}

// The SettlementFunc type is an adapter to allow the use of ordinary
// function as Settlement mutator.
type SettlementFunc func(context.Context, *ent.SettlementMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SettlementFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.SettlementMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SettlementMutation", m)
}

// The SpeciesFunc type is an adapter to allow the use of ordinary
// function as Species mutator.
type SpeciesFunc func(context.Context, *ent.SpeciesMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SpeciesFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.SpeciesMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SpeciesMutation", m)
}

// The TechniqueFunc type is an adapter to allow the use of ordinary
// function as Technique mutator.
type TechniqueFunc func(context.Context, *ent.TechniqueMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TechniqueFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.TechniqueMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TechniqueMutation", m)
}

// The VisitFunc type is an adapter to allow the use of ordinary
// function as Visit mutator.
type VisitFunc func(context.Context, *ent.VisitMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f VisitFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.VisitMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.VisitMutation", m)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) ent.Hook {
	return func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(context.Context, ent.Mutation) (ent.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
func Reject(op ent.Op) ent.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
