// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/dkrasnovdev/siberiana-api/ent/category"
	"github.com/dkrasnovdev/siberiana-api/ent/collection"
)

// Collection is the model entity for the Collection schema.
type Collection struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy string `json:"created_by,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy string `json:"updated_by,omitempty"`
	// DisplayName holds the value of the "display_name" field.
	DisplayName string `json:"display_name,omitempty"`
	// Abbreviation holds the value of the "abbreviation" field.
	Abbreviation string `json:"abbreviation,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// ExternalLink holds the value of the "external_link" field.
	ExternalLink string `json:"external_link,omitempty"`
	// PrimaryImageURL holds the value of the "primary_image_url" field.
	PrimaryImageURL string `json:"primary_image_url,omitempty"`
	// AdditionalImagesUrls holds the value of the "additional_images_urls" field.
	AdditionalImagesUrls []string `json:"additional_images_urls,omitempty"`
	// Slug holds the value of the "slug" field.
	Slug string `json:"slug,omitempty"`
	// Type holds the value of the "type" field.
	Type collection.Type `json:"type,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CollectionQuery when eager-loading is set.
	Edges                CollectionEdges `json:"edges"`
	category_collections *int
	selectValues         sql.SelectValues
}

// CollectionEdges holds the relations/edges for other nodes in the graph.
type CollectionEdges struct {
	// Art holds the value of the art edge.
	Art []*Art `json:"art,omitempty"`
	// Artifacts holds the value of the artifacts edge.
	Artifacts []*Artifact `json:"artifacts,omitempty"`
	// Books holds the value of the books edge.
	Books []*Book `json:"books,omitempty"`
	// ProtectedAreaPictures holds the value of the protected_area_pictures edge.
	ProtectedAreaPictures []*ProtectedAreaPicture `json:"protected_area_pictures,omitempty"`
	// Category holds the value of the category edge.
	Category *Category `json:"category,omitempty"`
	// Authors holds the value of the authors edge.
	Authors []*Person `json:"authors,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
	// totalCount holds the count of the edges above.
	totalCount [6]map[string]int

	namedArt                   map[string][]*Art
	namedArtifacts             map[string][]*Artifact
	namedBooks                 map[string][]*Book
	namedProtectedAreaPictures map[string][]*ProtectedAreaPicture
	namedAuthors               map[string][]*Person
}

// ArtOrErr returns the Art value or an error if the edge
// was not loaded in eager-loading.
func (e CollectionEdges) ArtOrErr() ([]*Art, error) {
	if e.loadedTypes[0] {
		return e.Art, nil
	}
	return nil, &NotLoadedError{edge: "art"}
}

// ArtifactsOrErr returns the Artifacts value or an error if the edge
// was not loaded in eager-loading.
func (e CollectionEdges) ArtifactsOrErr() ([]*Artifact, error) {
	if e.loadedTypes[1] {
		return e.Artifacts, nil
	}
	return nil, &NotLoadedError{edge: "artifacts"}
}

// BooksOrErr returns the Books value or an error if the edge
// was not loaded in eager-loading.
func (e CollectionEdges) BooksOrErr() ([]*Book, error) {
	if e.loadedTypes[2] {
		return e.Books, nil
	}
	return nil, &NotLoadedError{edge: "books"}
}

// ProtectedAreaPicturesOrErr returns the ProtectedAreaPictures value or an error if the edge
// was not loaded in eager-loading.
func (e CollectionEdges) ProtectedAreaPicturesOrErr() ([]*ProtectedAreaPicture, error) {
	if e.loadedTypes[3] {
		return e.ProtectedAreaPictures, nil
	}
	return nil, &NotLoadedError{edge: "protected_area_pictures"}
}

// CategoryOrErr returns the Category value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CollectionEdges) CategoryOrErr() (*Category, error) {
	if e.loadedTypes[4] {
		if e.Category == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: category.Label}
		}
		return e.Category, nil
	}
	return nil, &NotLoadedError{edge: "category"}
}

// AuthorsOrErr returns the Authors value or an error if the edge
// was not loaded in eager-loading.
func (e CollectionEdges) AuthorsOrErr() ([]*Person, error) {
	if e.loadedTypes[5] {
		return e.Authors, nil
	}
	return nil, &NotLoadedError{edge: "authors"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Collection) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case collection.FieldAdditionalImagesUrls:
			values[i] = new([]byte)
		case collection.FieldID:
			values[i] = new(sql.NullInt64)
		case collection.FieldCreatedBy, collection.FieldUpdatedBy, collection.FieldDisplayName, collection.FieldAbbreviation, collection.FieldDescription, collection.FieldExternalLink, collection.FieldPrimaryImageURL, collection.FieldSlug, collection.FieldType:
			values[i] = new(sql.NullString)
		case collection.FieldCreatedAt, collection.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case collection.ForeignKeys[0]: // category_collections
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Collection fields.
func (c *Collection) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case collection.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case collection.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case collection.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				c.CreatedBy = value.String
			}
		case collection.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case collection.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				c.UpdatedBy = value.String
			}
		case collection.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				c.DisplayName = value.String
			}
		case collection.FieldAbbreviation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field abbreviation", values[i])
			} else if value.Valid {
				c.Abbreviation = value.String
			}
		case collection.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				c.Description = value.String
			}
		case collection.FieldExternalLink:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field external_link", values[i])
			} else if value.Valid {
				c.ExternalLink = value.String
			}
		case collection.FieldPrimaryImageURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field primary_image_url", values[i])
			} else if value.Valid {
				c.PrimaryImageURL = value.String
			}
		case collection.FieldAdditionalImagesUrls:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field additional_images_urls", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &c.AdditionalImagesUrls); err != nil {
					return fmt.Errorf("unmarshal field additional_images_urls: %w", err)
				}
			}
		case collection.FieldSlug:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field slug", values[i])
			} else if value.Valid {
				c.Slug = value.String
			}
		case collection.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				c.Type = collection.Type(value.String)
			}
		case collection.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field category_collections", value)
			} else if value.Valid {
				c.category_collections = new(int)
				*c.category_collections = int(value.Int64)
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Collection.
// This includes values selected through modifiers, order, etc.
func (c *Collection) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryArt queries the "art" edge of the Collection entity.
func (c *Collection) QueryArt() *ArtQuery {
	return NewCollectionClient(c.config).QueryArt(c)
}

// QueryArtifacts queries the "artifacts" edge of the Collection entity.
func (c *Collection) QueryArtifacts() *ArtifactQuery {
	return NewCollectionClient(c.config).QueryArtifacts(c)
}

// QueryBooks queries the "books" edge of the Collection entity.
func (c *Collection) QueryBooks() *BookQuery {
	return NewCollectionClient(c.config).QueryBooks(c)
}

// QueryProtectedAreaPictures queries the "protected_area_pictures" edge of the Collection entity.
func (c *Collection) QueryProtectedAreaPictures() *ProtectedAreaPictureQuery {
	return NewCollectionClient(c.config).QueryProtectedAreaPictures(c)
}

// QueryCategory queries the "category" edge of the Collection entity.
func (c *Collection) QueryCategory() *CategoryQuery {
	return NewCollectionClient(c.config).QueryCategory(c)
}

// QueryAuthors queries the "authors" edge of the Collection entity.
func (c *Collection) QueryAuthors() *PersonQuery {
	return NewCollectionClient(c.config).QueryAuthors(c)
}

// Update returns a builder for updating this Collection.
// Note that you need to call Collection.Unwrap() before calling this method if this Collection
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Collection) Update() *CollectionUpdateOne {
	return NewCollectionClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Collection entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Collection) Unwrap() *Collection {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Collection is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Collection) String() string {
	var builder strings.Builder
	builder.WriteString("Collection(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(c.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(c.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("display_name=")
	builder.WriteString(c.DisplayName)
	builder.WriteString(", ")
	builder.WriteString("abbreviation=")
	builder.WriteString(c.Abbreviation)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(c.Description)
	builder.WriteString(", ")
	builder.WriteString("external_link=")
	builder.WriteString(c.ExternalLink)
	builder.WriteString(", ")
	builder.WriteString("primary_image_url=")
	builder.WriteString(c.PrimaryImageURL)
	builder.WriteString(", ")
	builder.WriteString("additional_images_urls=")
	builder.WriteString(fmt.Sprintf("%v", c.AdditionalImagesUrls))
	builder.WriteString(", ")
	builder.WriteString("slug=")
	builder.WriteString(c.Slug)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", c.Type))
	builder.WriteByte(')')
	return builder.String()
}

// NamedArt returns the Art named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Collection) NamedArt(name string) ([]*Art, error) {
	if c.Edges.namedArt == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedArt[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Collection) appendNamedArt(name string, edges ...*Art) {
	if c.Edges.namedArt == nil {
		c.Edges.namedArt = make(map[string][]*Art)
	}
	if len(edges) == 0 {
		c.Edges.namedArt[name] = []*Art{}
	} else {
		c.Edges.namedArt[name] = append(c.Edges.namedArt[name], edges...)
	}
}

// NamedArtifacts returns the Artifacts named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Collection) NamedArtifacts(name string) ([]*Artifact, error) {
	if c.Edges.namedArtifacts == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedArtifacts[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Collection) appendNamedArtifacts(name string, edges ...*Artifact) {
	if c.Edges.namedArtifacts == nil {
		c.Edges.namedArtifacts = make(map[string][]*Artifact)
	}
	if len(edges) == 0 {
		c.Edges.namedArtifacts[name] = []*Artifact{}
	} else {
		c.Edges.namedArtifacts[name] = append(c.Edges.namedArtifacts[name], edges...)
	}
}

// NamedBooks returns the Books named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Collection) NamedBooks(name string) ([]*Book, error) {
	if c.Edges.namedBooks == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedBooks[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Collection) appendNamedBooks(name string, edges ...*Book) {
	if c.Edges.namedBooks == nil {
		c.Edges.namedBooks = make(map[string][]*Book)
	}
	if len(edges) == 0 {
		c.Edges.namedBooks[name] = []*Book{}
	} else {
		c.Edges.namedBooks[name] = append(c.Edges.namedBooks[name], edges...)
	}
}

// NamedProtectedAreaPictures returns the ProtectedAreaPictures named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Collection) NamedProtectedAreaPictures(name string) ([]*ProtectedAreaPicture, error) {
	if c.Edges.namedProtectedAreaPictures == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedProtectedAreaPictures[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Collection) appendNamedProtectedAreaPictures(name string, edges ...*ProtectedAreaPicture) {
	if c.Edges.namedProtectedAreaPictures == nil {
		c.Edges.namedProtectedAreaPictures = make(map[string][]*ProtectedAreaPicture)
	}
	if len(edges) == 0 {
		c.Edges.namedProtectedAreaPictures[name] = []*ProtectedAreaPicture{}
	} else {
		c.Edges.namedProtectedAreaPictures[name] = append(c.Edges.namedProtectedAreaPictures[name], edges...)
	}
}

// NamedAuthors returns the Authors named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Collection) NamedAuthors(name string) ([]*Person, error) {
	if c.Edges.namedAuthors == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedAuthors[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Collection) appendNamedAuthors(name string, edges ...*Person) {
	if c.Edges.namedAuthors == nil {
		c.Edges.namedAuthors = make(map[string][]*Person)
	}
	if len(edges) == 0 {
		c.Edges.namedAuthors[name] = []*Person{}
	} else {
		c.Edges.namedAuthors[name] = append(c.Edges.namedAuthors[name], edges...)
	}
}

// Collections is a parsable slice of Collection.
type Collections []*Collection
