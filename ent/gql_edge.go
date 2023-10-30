// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (a *Art) Author(ctx context.Context) (*Person, error) {
	result, err := a.Edges.AuthorOrErr()
	if IsNotLoaded(err) {
		result, err = a.QueryAuthor().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (a *Art) ArtGenre(ctx context.Context) (result []*ArtGenre, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = a.NamedArtGenre(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = a.Edges.ArtGenreOrErr()
	}
	if IsNotLoaded(err) {
		result, err = a.QueryArtGenre().All(ctx)
	}
	return result, err
}

func (a *Art) ArtStyle(ctx context.Context) (result []*ArtStyle, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = a.NamedArtStyle(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = a.Edges.ArtStyleOrErr()
	}
	if IsNotLoaded(err) {
		result, err = a.QueryArtStyle().All(ctx)
	}
	return result, err
}

func (a *Art) Mediums(ctx context.Context) (result []*Medium, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = a.NamedMediums(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = a.Edges.MediumsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = a.QueryMediums().All(ctx)
	}
	return result, err
}

func (a *Art) Collection(ctx context.Context) (*Collection, error) {
	result, err := a.Edges.CollectionOrErr()
	if IsNotLoaded(err) {
		result, err = a.QueryCollection().Only(ctx)
	}
	return result, err
}

func (ag *ArtGenre) Art(ctx context.Context) (result []*Art, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = ag.NamedArt(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = ag.Edges.ArtOrErr()
	}
	if IsNotLoaded(err) {
		result, err = ag.QueryArt().All(ctx)
	}
	return result, err
}

func (as *ArtStyle) Art(ctx context.Context) (result []*Art, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = as.NamedArt(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = as.Edges.ArtOrErr()
	}
	if IsNotLoaded(err) {
		result, err = as.QueryArt().All(ctx)
	}
	return result, err
}

func (a *Artifact) Authors(ctx context.Context) (result []*Person, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = a.NamedAuthors(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = a.Edges.AuthorsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = a.QueryAuthors().All(ctx)
	}
	return result, err
}

func (a *Artifact) Mediums(ctx context.Context) (result []*Medium, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = a.NamedMediums(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = a.Edges.MediumsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = a.QueryMediums().All(ctx)
	}
	return result, err
}

func (a *Artifact) Techniques(ctx context.Context) (result []*Technique, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = a.NamedTechniques(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = a.Edges.TechniquesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = a.QueryTechniques().All(ctx)
	}
	return result, err
}

func (a *Artifact) Projects(ctx context.Context) (result []*Project, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = a.NamedProjects(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = a.Edges.ProjectsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = a.QueryProjects().All(ctx)
	}
	return result, err
}

func (a *Artifact) Publications(ctx context.Context) (result []*Publication, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = a.NamedPublications(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = a.Edges.PublicationsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = a.QueryPublications().All(ctx)
	}
	return result, err
}

func (a *Artifact) CulturalAffiliation(ctx context.Context) (*Culture, error) {
	result, err := a.Edges.CulturalAffiliationOrErr()
	if IsNotLoaded(err) {
		result, err = a.QueryCulturalAffiliation().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (a *Artifact) Monument(ctx context.Context) (*Monument, error) {
	result, err := a.Edges.MonumentOrErr()
	if IsNotLoaded(err) {
		result, err = a.QueryMonument().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (a *Artifact) Model(ctx context.Context) (*Model, error) {
	result, err := a.Edges.ModelOrErr()
	if IsNotLoaded(err) {
		result, err = a.QueryModel().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (a *Artifact) Set(ctx context.Context) (*Set, error) {
	result, err := a.Edges.SetOrErr()
	if IsNotLoaded(err) {
		result, err = a.QuerySet().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (a *Artifact) Location(ctx context.Context) (*Location, error) {
	result, err := a.Edges.LocationOrErr()
	if IsNotLoaded(err) {
		result, err = a.QueryLocation().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (a *Artifact) Collection(ctx context.Context) (*Collection, error) {
	result, err := a.Edges.CollectionOrErr()
	if IsNotLoaded(err) {
		result, err = a.QueryCollection().Only(ctx)
	}
	return result, err
}

func (a *Artifact) License(ctx context.Context) (*License, error) {
	result, err := a.Edges.LicenseOrErr()
	if IsNotLoaded(err) {
		result, err = a.QueryLicense().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (b *Book) Authors(ctx context.Context) (result []*Person, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = b.NamedAuthors(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = b.Edges.AuthorsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = b.QueryAuthors().All(ctx)
	}
	return result, err
}

func (b *Book) BookGenres(ctx context.Context) (result []*BookGenre, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = b.NamedBookGenres(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = b.Edges.BookGenresOrErr()
	}
	if IsNotLoaded(err) {
		result, err = b.QueryBookGenres().All(ctx)
	}
	return result, err
}

func (b *Book) Collection(ctx context.Context) (*Collection, error) {
	result, err := b.Edges.CollectionOrErr()
	if IsNotLoaded(err) {
		result, err = b.QueryCollection().Only(ctx)
	}
	return result, err
}

func (b *Book) Periodical(ctx context.Context) (*Periodical, error) {
	result, err := b.Edges.PeriodicalOrErr()
	if IsNotLoaded(err) {
		result, err = b.QueryPeriodical().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (b *Book) Publisher(ctx context.Context) (*Publisher, error) {
	result, err := b.Edges.PublisherOrErr()
	if IsNotLoaded(err) {
		result, err = b.QueryPublisher().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (b *Book) License(ctx context.Context) (*License, error) {
	result, err := b.Edges.LicenseOrErr()
	if IsNotLoaded(err) {
		result, err = b.QueryLicense().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (b *Book) Location(ctx context.Context) (*Location, error) {
	result, err := b.Edges.LocationOrErr()
	if IsNotLoaded(err) {
		result, err = b.QueryLocation().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (b *Book) PlaceOfPublication(ctx context.Context) (*Settlement, error) {
	result, err := b.Edges.PlaceOfPublicationOrErr()
	if IsNotLoaded(err) {
		result, err = b.QueryPlaceOfPublication().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (b *Book) Library(ctx context.Context) (*Organization, error) {
	result, err := b.Edges.LibraryOrErr()
	if IsNotLoaded(err) {
		result, err = b.QueryLibrary().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (bg *BookGenre) Books(ctx context.Context) (result []*Book, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = bg.NamedBooks(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = bg.Edges.BooksOrErr()
	}
	if IsNotLoaded(err) {
		result, err = bg.QueryBooks().All(ctx)
	}
	return result, err
}

func (c *Category) Collections(ctx context.Context) (result []*Collection, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = c.NamedCollections(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = c.Edges.CollectionsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = c.QueryCollections().All(ctx)
	}
	return result, err
}

func (c *Collection) Arts(ctx context.Context) (result []*Art, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = c.NamedArts(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = c.Edges.ArtsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = c.QueryArts().All(ctx)
	}
	return result, err
}

func (c *Collection) Artifacts(ctx context.Context) (result []*Artifact, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = c.NamedArtifacts(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = c.Edges.ArtifactsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = c.QueryArtifacts().All(ctx)
	}
	return result, err
}

func (c *Collection) Books(ctx context.Context) (result []*Book, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = c.NamedBooks(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = c.Edges.BooksOrErr()
	}
	if IsNotLoaded(err) {
		result, err = c.QueryBooks().All(ctx)
	}
	return result, err
}

func (c *Collection) ProtectedAreaPictures(ctx context.Context) (result []*ProtectedAreaPicture, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = c.NamedProtectedAreaPictures(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = c.Edges.ProtectedAreaPicturesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = c.QueryProtectedAreaPictures().All(ctx)
	}
	return result, err
}

func (c *Collection) Category(ctx context.Context) (*Category, error) {
	result, err := c.Edges.CategoryOrErr()
	if IsNotLoaded(err) {
		result, err = c.QueryCategory().Only(ctx)
	}
	return result, err
}

func (c *Collection) Authors(ctx context.Context) (result []*Person, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = c.NamedAuthors(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = c.Edges.AuthorsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = c.QueryAuthors().All(ctx)
	}
	return result, err
}

func (c *Country) Locations(ctx context.Context) (result []*Location, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = c.NamedLocations(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = c.Edges.LocationsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = c.QueryLocations().All(ctx)
	}
	return result, err
}

func (c *Culture) Artifacts(ctx context.Context) (result []*Artifact, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = c.NamedArtifacts(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = c.Edges.ArtifactsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = c.QueryArtifacts().All(ctx)
	}
	return result, err
}

func (d *District) Locations(ctx context.Context) (result []*Location, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = d.NamedLocations(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = d.Edges.LocationsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = d.QueryLocations().All(ctx)
	}
	return result, err
}

func (f *Favourite) Proxies(ctx context.Context) (result []*Proxy, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = f.NamedProxies(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = f.Edges.ProxiesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = f.QueryProxies().All(ctx)
	}
	return result, err
}

func (l *License) Artifacts(ctx context.Context) (result []*Artifact, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = l.NamedArtifacts(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = l.Edges.ArtifactsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = l.QueryArtifacts().All(ctx)
	}
	return result, err
}

func (l *License) Books(ctx context.Context) (result []*Book, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = l.NamedBooks(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = l.Edges.BooksOrErr()
	}
	if IsNotLoaded(err) {
		result, err = l.QueryBooks().All(ctx)
	}
	return result, err
}

func (l *License) ProtectedAreaPictures(ctx context.Context) (result []*ProtectedAreaPicture, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = l.NamedProtectedAreaPictures(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = l.Edges.ProtectedAreaPicturesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = l.QueryProtectedAreaPictures().All(ctx)
	}
	return result, err
}

func (l *Location) Artifacts(ctx context.Context) (result []*Artifact, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = l.NamedArtifacts(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = l.Edges.ArtifactsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = l.QueryArtifacts().All(ctx)
	}
	return result, err
}

func (l *Location) Books(ctx context.Context) (result []*Book, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = l.NamedBooks(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = l.Edges.BooksOrErr()
	}
	if IsNotLoaded(err) {
		result, err = l.QueryBooks().All(ctx)
	}
	return result, err
}

func (l *Location) ProtectedAreaPictures(ctx context.Context) (result []*ProtectedAreaPicture, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = l.NamedProtectedAreaPictures(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = l.Edges.ProtectedAreaPicturesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = l.QueryProtectedAreaPictures().All(ctx)
	}
	return result, err
}

func (l *Location) Country(ctx context.Context) (*Country, error) {
	result, err := l.Edges.CountryOrErr()
	if IsNotLoaded(err) {
		result, err = l.QueryCountry().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (l *Location) District(ctx context.Context) (*District, error) {
	result, err := l.Edges.DistrictOrErr()
	if IsNotLoaded(err) {
		result, err = l.QueryDistrict().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (l *Location) Settlement(ctx context.Context) (*Settlement, error) {
	result, err := l.Edges.SettlementOrErr()
	if IsNotLoaded(err) {
		result, err = l.QuerySettlement().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (l *Location) Region(ctx context.Context) (*Region, error) {
	result, err := l.Edges.RegionOrErr()
	if IsNotLoaded(err) {
		result, err = l.QueryRegion().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (m *Medium) Arts(ctx context.Context) (result []*Art, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = m.NamedArts(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = m.Edges.ArtsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = m.QueryArts().All(ctx)
	}
	return result, err
}

func (m *Medium) Artifacts(ctx context.Context) (result []*Artifact, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = m.NamedArtifacts(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = m.Edges.ArtifactsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = m.QueryArtifacts().All(ctx)
	}
	return result, err
}

func (m *Model) Artifacts(ctx context.Context) (result []*Artifact, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = m.NamedArtifacts(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = m.Edges.ArtifactsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = m.QueryArtifacts().All(ctx)
	}
	return result, err
}

func (m *Monument) Artifacts(ctx context.Context) (result []*Artifact, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = m.NamedArtifacts(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = m.Edges.ArtifactsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = m.QueryArtifacts().All(ctx)
	}
	return result, err
}

func (m *Monument) Sets(ctx context.Context) (result []*Set, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = m.NamedSets(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = m.Edges.SetsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = m.QuerySets().All(ctx)
	}
	return result, err
}

func (o *Organization) Books(ctx context.Context) (result []*Book, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = o.NamedBooks(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = o.Edges.BooksOrErr()
	}
	if IsNotLoaded(err) {
		result, err = o.QueryBooks().All(ctx)
	}
	return result, err
}

func (o *Organization) People(ctx context.Context) (result []*Person, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = o.NamedPeople(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = o.Edges.PeopleOrErr()
	}
	if IsNotLoaded(err) {
		result, err = o.QueryPeople().All(ctx)
	}
	return result, err
}

func (pe *Periodical) Books(ctx context.Context) (result []*Book, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = pe.NamedBooks(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = pe.Edges.BooksOrErr()
	}
	if IsNotLoaded(err) {
		result, err = pe.QueryBooks().All(ctx)
	}
	return result, err
}

func (pe *Person) Collections(ctx context.Context) (result []*Collection, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = pe.NamedCollections(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = pe.Edges.CollectionsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = pe.QueryCollections().All(ctx)
	}
	return result, err
}

func (pe *Person) Arts(ctx context.Context) (result []*Art, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = pe.NamedArts(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = pe.Edges.ArtsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = pe.QueryArts().All(ctx)
	}
	return result, err
}

func (pe *Person) Artifacts(ctx context.Context) (result []*Artifact, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = pe.NamedArtifacts(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = pe.Edges.ArtifactsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = pe.QueryArtifacts().All(ctx)
	}
	return result, err
}

func (pe *Person) Books(ctx context.Context) (result []*Book, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = pe.NamedBooks(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = pe.Edges.BooksOrErr()
	}
	if IsNotLoaded(err) {
		result, err = pe.QueryBooks().All(ctx)
	}
	return result, err
}

func (pe *Person) Projects(ctx context.Context) (result []*Project, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = pe.NamedProjects(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = pe.Edges.ProjectsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = pe.QueryProjects().All(ctx)
	}
	return result, err
}

func (pe *Person) Publications(ctx context.Context) (result []*Publication, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = pe.NamedPublications(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = pe.Edges.PublicationsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = pe.QueryPublications().All(ctx)
	}
	return result, err
}

func (pe *Person) Affiliation(ctx context.Context) (*Organization, error) {
	result, err := pe.Edges.AffiliationOrErr()
	if IsNotLoaded(err) {
		result, err = pe.QueryAffiliation().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (pe *Personal) Proxies(ctx context.Context) (result []*Proxy, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = pe.NamedProxies(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = pe.Edges.ProxiesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = pe.QueryProxies().All(ctx)
	}
	return result, err
}

func (pr *Project) Artifacts(ctx context.Context) (result []*Artifact, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = pr.NamedArtifacts(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = pr.Edges.ArtifactsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = pr.QueryArtifacts().All(ctx)
	}
	return result, err
}

func (pr *Project) Team(ctx context.Context) (result []*Person, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = pr.NamedTeam(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = pr.Edges.TeamOrErr()
	}
	if IsNotLoaded(err) {
		result, err = pr.QueryTeam().All(ctx)
	}
	return result, err
}

func (pa *ProtectedArea) ProtectedAreaPictures(ctx context.Context) (result []*ProtectedAreaPicture, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = pa.NamedProtectedAreaPictures(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = pa.Edges.ProtectedAreaPicturesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = pa.QueryProtectedAreaPictures().All(ctx)
	}
	return result, err
}

func (pa *ProtectedArea) ProtectedAreaCategory(ctx context.Context) (*ProtectedAreaCategory, error) {
	result, err := pa.Edges.ProtectedAreaCategoryOrErr()
	if IsNotLoaded(err) {
		result, err = pa.QueryProtectedAreaCategory().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (pac *ProtectedAreaCategory) ProtectedAreas(ctx context.Context) (result []*ProtectedArea, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = pac.NamedProtectedAreas(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = pac.Edges.ProtectedAreasOrErr()
	}
	if IsNotLoaded(err) {
		result, err = pac.QueryProtectedAreas().All(ctx)
	}
	return result, err
}

func (pap *ProtectedAreaPicture) Collection(ctx context.Context) (*Collection, error) {
	result, err := pap.Edges.CollectionOrErr()
	if IsNotLoaded(err) {
		result, err = pap.QueryCollection().Only(ctx)
	}
	return result, err
}

func (pap *ProtectedAreaPicture) ProtectedArea(ctx context.Context) (*ProtectedArea, error) {
	result, err := pap.Edges.ProtectedAreaOrErr()
	if IsNotLoaded(err) {
		result, err = pap.QueryProtectedArea().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (pap *ProtectedAreaPicture) Location(ctx context.Context) (*Location, error) {
	result, err := pap.Edges.LocationOrErr()
	if IsNotLoaded(err) {
		result, err = pap.QueryLocation().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (pap *ProtectedAreaPicture) License(ctx context.Context) (*License, error) {
	result, err := pap.Edges.LicenseOrErr()
	if IsNotLoaded(err) {
		result, err = pap.QueryLicense().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (pr *Proxy) Favourite(ctx context.Context) (*Favourite, error) {
	result, err := pr.Edges.FavouriteOrErr()
	if IsNotLoaded(err) {
		result, err = pr.QueryFavourite().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (pr *Proxy) Personal(ctx context.Context) (*Personal, error) {
	result, err := pr.Edges.PersonalOrErr()
	if IsNotLoaded(err) {
		result, err = pr.QueryPersonal().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (pu *Publication) Artifacts(ctx context.Context) (result []*Artifact, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = pu.NamedArtifacts(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = pu.Edges.ArtifactsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = pu.QueryArtifacts().All(ctx)
	}
	return result, err
}

func (pu *Publication) Authors(ctx context.Context) (result []*Person, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = pu.NamedAuthors(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = pu.Edges.AuthorsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = pu.QueryAuthors().All(ctx)
	}
	return result, err
}

func (pu *Publisher) Books(ctx context.Context) (result []*Book, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = pu.NamedBooks(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = pu.Edges.BooksOrErr()
	}
	if IsNotLoaded(err) {
		result, err = pu.QueryBooks().All(ctx)
	}
	return result, err
}

func (r *Region) Locations(ctx context.Context) (result []*Location, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = r.NamedLocations(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = r.Edges.LocationsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = r.QueryLocations().All(ctx)
	}
	return result, err
}

func (s *Set) Artifacts(ctx context.Context) (result []*Artifact, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = s.NamedArtifacts(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = s.Edges.ArtifactsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = s.QueryArtifacts().All(ctx)
	}
	return result, err
}

func (s *Set) Monuments(ctx context.Context) (result []*Monument, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = s.NamedMonuments(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = s.Edges.MonumentsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = s.QueryMonuments().All(ctx)
	}
	return result, err
}

func (s *Settlement) Books(ctx context.Context) (result []*Book, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = s.NamedBooks(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = s.Edges.BooksOrErr()
	}
	if IsNotLoaded(err) {
		result, err = s.QueryBooks().All(ctx)
	}
	return result, err
}

func (s *Settlement) Locations(ctx context.Context) (result []*Location, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = s.NamedLocations(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = s.Edges.LocationsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = s.QueryLocations().All(ctx)
	}
	return result, err
}

func (t *Technique) Artifacts(ctx context.Context) (result []*Artifact, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = t.NamedArtifacts(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = t.Edges.ArtifactsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = t.QueryArtifacts().All(ctx)
	}
	return result, err
}
