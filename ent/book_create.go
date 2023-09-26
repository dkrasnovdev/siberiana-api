// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dkrasnovdev/siberiana-api/ent/book"
	"github.com/dkrasnovdev/siberiana-api/ent/bookgenre"
	"github.com/dkrasnovdev/siberiana-api/ent/collection"
	"github.com/dkrasnovdev/siberiana-api/ent/holder"
	"github.com/dkrasnovdev/siberiana-api/ent/license"
	"github.com/dkrasnovdev/siberiana-api/ent/location"
	"github.com/dkrasnovdev/siberiana-api/ent/person"
	"github.com/dkrasnovdev/siberiana-api/ent/publisher"
)

// BookCreate is the builder for creating a Book entity.
type BookCreate struct {
	config
	mutation *BookMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (bc *BookCreate) SetCreatedAt(t time.Time) *BookCreate {
	bc.mutation.SetCreatedAt(t)
	return bc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (bc *BookCreate) SetNillableCreatedAt(t *time.Time) *BookCreate {
	if t != nil {
		bc.SetCreatedAt(*t)
	}
	return bc
}

// SetCreatedBy sets the "created_by" field.
func (bc *BookCreate) SetCreatedBy(s string) *BookCreate {
	bc.mutation.SetCreatedBy(s)
	return bc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (bc *BookCreate) SetNillableCreatedBy(s *string) *BookCreate {
	if s != nil {
		bc.SetCreatedBy(*s)
	}
	return bc
}

// SetUpdatedAt sets the "updated_at" field.
func (bc *BookCreate) SetUpdatedAt(t time.Time) *BookCreate {
	bc.mutation.SetUpdatedAt(t)
	return bc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (bc *BookCreate) SetNillableUpdatedAt(t *time.Time) *BookCreate {
	if t != nil {
		bc.SetUpdatedAt(*t)
	}
	return bc
}

// SetUpdatedBy sets the "updated_by" field.
func (bc *BookCreate) SetUpdatedBy(s string) *BookCreate {
	bc.mutation.SetUpdatedBy(s)
	return bc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (bc *BookCreate) SetNillableUpdatedBy(s *string) *BookCreate {
	if s != nil {
		bc.SetUpdatedBy(*s)
	}
	return bc
}

// SetDisplayName sets the "display_name" field.
func (bc *BookCreate) SetDisplayName(s string) *BookCreate {
	bc.mutation.SetDisplayName(s)
	return bc
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (bc *BookCreate) SetNillableDisplayName(s *string) *BookCreate {
	if s != nil {
		bc.SetDisplayName(*s)
	}
	return bc
}

// SetAbbreviation sets the "abbreviation" field.
func (bc *BookCreate) SetAbbreviation(s string) *BookCreate {
	bc.mutation.SetAbbreviation(s)
	return bc
}

// SetNillableAbbreviation sets the "abbreviation" field if the given value is not nil.
func (bc *BookCreate) SetNillableAbbreviation(s *string) *BookCreate {
	if s != nil {
		bc.SetAbbreviation(*s)
	}
	return bc
}

// SetDescription sets the "description" field.
func (bc *BookCreate) SetDescription(s string) *BookCreate {
	bc.mutation.SetDescription(s)
	return bc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (bc *BookCreate) SetNillableDescription(s *string) *BookCreate {
	if s != nil {
		bc.SetDescription(*s)
	}
	return bc
}

// SetExternalLink sets the "external_link" field.
func (bc *BookCreate) SetExternalLink(s string) *BookCreate {
	bc.mutation.SetExternalLink(s)
	return bc
}

// SetNillableExternalLink sets the "external_link" field if the given value is not nil.
func (bc *BookCreate) SetNillableExternalLink(s *string) *BookCreate {
	if s != nil {
		bc.SetExternalLink(*s)
	}
	return bc
}

// SetType sets the "type" field.
func (bc *BookCreate) SetType(b book.Type) *BookCreate {
	bc.mutation.SetType(b)
	return bc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (bc *BookCreate) SetNillableType(b *book.Type) *BookCreate {
	if b != nil {
		bc.SetType(*b)
	}
	return bc
}

// SetPrimaryImageURL sets the "primary_image_url" field.
func (bc *BookCreate) SetPrimaryImageURL(s string) *BookCreate {
	bc.mutation.SetPrimaryImageURL(s)
	return bc
}

// SetNillablePrimaryImageURL sets the "primary_image_url" field if the given value is not nil.
func (bc *BookCreate) SetNillablePrimaryImageURL(s *string) *BookCreate {
	if s != nil {
		bc.SetPrimaryImageURL(*s)
	}
	return bc
}

// SetAdditionalImagesUrls sets the "additional_images_urls" field.
func (bc *BookCreate) SetAdditionalImagesUrls(s []string) *BookCreate {
	bc.mutation.SetAdditionalImagesUrls(s)
	return bc
}

// SetFiles sets the "files" field.
func (bc *BookCreate) SetFiles(s []string) *BookCreate {
	bc.mutation.SetFiles(s)
	return bc
}

// SetYear sets the "year" field.
func (bc *BookCreate) SetYear(i int) *BookCreate {
	bc.mutation.SetYear(i)
	return bc
}

// SetNillableYear sets the "year" field if the given value is not nil.
func (bc *BookCreate) SetNillableYear(i *int) *BookCreate {
	if i != nil {
		bc.SetYear(*i)
	}
	return bc
}

// AddAuthorIDs adds the "authors" edge to the Person entity by IDs.
func (bc *BookCreate) AddAuthorIDs(ids ...int) *BookCreate {
	bc.mutation.AddAuthorIDs(ids...)
	return bc
}

// AddAuthors adds the "authors" edges to the Person entity.
func (bc *BookCreate) AddAuthors(p ...*Person) *BookCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return bc.AddAuthorIDs(ids...)
}

// AddBookGenreIDs adds the "book_genres" edge to the BookGenre entity by IDs.
func (bc *BookCreate) AddBookGenreIDs(ids ...int) *BookCreate {
	bc.mutation.AddBookGenreIDs(ids...)
	return bc
}

// AddBookGenres adds the "book_genres" edges to the BookGenre entity.
func (bc *BookCreate) AddBookGenres(b ...*BookGenre) *BookCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bc.AddBookGenreIDs(ids...)
}

// SetCollectionID sets the "collection" edge to the Collection entity by ID.
func (bc *BookCreate) SetCollectionID(id int) *BookCreate {
	bc.mutation.SetCollectionID(id)
	return bc
}

// SetCollection sets the "collection" edge to the Collection entity.
func (bc *BookCreate) SetCollection(c *Collection) *BookCreate {
	return bc.SetCollectionID(c.ID)
}

// AddHolderIDs adds the "holders" edge to the Holder entity by IDs.
func (bc *BookCreate) AddHolderIDs(ids ...int) *BookCreate {
	bc.mutation.AddHolderIDs(ids...)
	return bc
}

// AddHolders adds the "holders" edges to the Holder entity.
func (bc *BookCreate) AddHolders(h ...*Holder) *BookCreate {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return bc.AddHolderIDs(ids...)
}

// SetPublisherID sets the "publisher" edge to the Publisher entity by ID.
func (bc *BookCreate) SetPublisherID(id int) *BookCreate {
	bc.mutation.SetPublisherID(id)
	return bc
}

// SetNillablePublisherID sets the "publisher" edge to the Publisher entity by ID if the given value is not nil.
func (bc *BookCreate) SetNillablePublisherID(id *int) *BookCreate {
	if id != nil {
		bc = bc.SetPublisherID(*id)
	}
	return bc
}

// SetPublisher sets the "publisher" edge to the Publisher entity.
func (bc *BookCreate) SetPublisher(p *Publisher) *BookCreate {
	return bc.SetPublisherID(p.ID)
}

// SetLicenseID sets the "license" edge to the License entity by ID.
func (bc *BookCreate) SetLicenseID(id int) *BookCreate {
	bc.mutation.SetLicenseID(id)
	return bc
}

// SetNillableLicenseID sets the "license" edge to the License entity by ID if the given value is not nil.
func (bc *BookCreate) SetNillableLicenseID(id *int) *BookCreate {
	if id != nil {
		bc = bc.SetLicenseID(*id)
	}
	return bc
}

// SetLicense sets the "license" edge to the License entity.
func (bc *BookCreate) SetLicense(l *License) *BookCreate {
	return bc.SetLicenseID(l.ID)
}

// SetLocationID sets the "location" edge to the Location entity by ID.
func (bc *BookCreate) SetLocationID(id int) *BookCreate {
	bc.mutation.SetLocationID(id)
	return bc
}

// SetNillableLocationID sets the "location" edge to the Location entity by ID if the given value is not nil.
func (bc *BookCreate) SetNillableLocationID(id *int) *BookCreate {
	if id != nil {
		bc = bc.SetLocationID(*id)
	}
	return bc
}

// SetLocation sets the "location" edge to the Location entity.
func (bc *BookCreate) SetLocation(l *Location) *BookCreate {
	return bc.SetLocationID(l.ID)
}

// Mutation returns the BookMutation object of the builder.
func (bc *BookCreate) Mutation() *BookMutation {
	return bc.mutation
}

// Save creates the Book in the database.
func (bc *BookCreate) Save(ctx context.Context) (*Book, error) {
	if err := bc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, bc.sqlSave, bc.mutation, bc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BookCreate) SaveX(ctx context.Context) *Book {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bc *BookCreate) Exec(ctx context.Context) error {
	_, err := bc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bc *BookCreate) ExecX(ctx context.Context) {
	if err := bc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bc *BookCreate) defaults() error {
	if _, ok := bc.mutation.CreatedAt(); !ok {
		if book.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized book.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := book.DefaultCreatedAt()
		bc.mutation.SetCreatedAt(v)
	}
	if _, ok := bc.mutation.UpdatedAt(); !ok {
		if book.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized book.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := book.DefaultUpdatedAt()
		bc.mutation.SetUpdatedAt(v)
	}
	if _, ok := bc.mutation.GetType(); !ok {
		v := book.DefaultType
		bc.mutation.SetType(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (bc *BookCreate) check() error {
	if _, ok := bc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Book.created_at"`)}
	}
	if _, ok := bc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Book.updated_at"`)}
	}
	if v, ok := bc.mutation.GetType(); ok {
		if err := book.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Book.type": %w`, err)}
		}
	}
	if v, ok := bc.mutation.Year(); ok {
		if err := book.YearValidator(v); err != nil {
			return &ValidationError{Name: "year", err: fmt.Errorf(`ent: validator failed for field "Book.year": %w`, err)}
		}
	}
	if _, ok := bc.mutation.CollectionID(); !ok {
		return &ValidationError{Name: "collection", err: errors.New(`ent: missing required edge "Book.collection"`)}
	}
	return nil
}

func (bc *BookCreate) sqlSave(ctx context.Context) (*Book, error) {
	if err := bc.check(); err != nil {
		return nil, err
	}
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	bc.mutation.id = &_node.ID
	bc.mutation.done = true
	return _node, nil
}

func (bc *BookCreate) createSpec() (*Book, *sqlgraph.CreateSpec) {
	var (
		_node = &Book{config: bc.config}
		_spec = sqlgraph.NewCreateSpec(book.Table, sqlgraph.NewFieldSpec(book.FieldID, field.TypeInt))
	)
	if value, ok := bc.mutation.CreatedAt(); ok {
		_spec.SetField(book.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := bc.mutation.CreatedBy(); ok {
		_spec.SetField(book.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := bc.mutation.UpdatedAt(); ok {
		_spec.SetField(book.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := bc.mutation.UpdatedBy(); ok {
		_spec.SetField(book.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := bc.mutation.DisplayName(); ok {
		_spec.SetField(book.FieldDisplayName, field.TypeString, value)
		_node.DisplayName = value
	}
	if value, ok := bc.mutation.Abbreviation(); ok {
		_spec.SetField(book.FieldAbbreviation, field.TypeString, value)
		_node.Abbreviation = value
	}
	if value, ok := bc.mutation.Description(); ok {
		_spec.SetField(book.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := bc.mutation.ExternalLink(); ok {
		_spec.SetField(book.FieldExternalLink, field.TypeString, value)
		_node.ExternalLink = value
	}
	if value, ok := bc.mutation.GetType(); ok {
		_spec.SetField(book.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if value, ok := bc.mutation.PrimaryImageURL(); ok {
		_spec.SetField(book.FieldPrimaryImageURL, field.TypeString, value)
		_node.PrimaryImageURL = value
	}
	if value, ok := bc.mutation.AdditionalImagesUrls(); ok {
		_spec.SetField(book.FieldAdditionalImagesUrls, field.TypeJSON, value)
		_node.AdditionalImagesUrls = value
	}
	if value, ok := bc.mutation.Files(); ok {
		_spec.SetField(book.FieldFiles, field.TypeJSON, value)
		_node.Files = value
	}
	if value, ok := bc.mutation.Year(); ok {
		_spec.SetField(book.FieldYear, field.TypeInt, value)
		_node.Year = value
	}
	if nodes := bc.mutation.AuthorsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   book.AuthorsTable,
			Columns: book.AuthorsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(person.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.BookGenresIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   book.BookGenresTable,
			Columns: book.BookGenresPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bookgenre.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.CollectionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   book.CollectionTable,
			Columns: []string{book.CollectionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(collection.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.collection_books = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.HoldersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   book.HoldersTable,
			Columns: book.HoldersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(holder.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.PublisherIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   book.PublisherTable,
			Columns: []string{book.PublisherColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(publisher.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.publisher_books = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.LicenseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   book.LicenseTable,
			Columns: []string{book.LicenseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(license.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.license_books = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.LocationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   book.LocationTable,
			Columns: []string{book.LocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(location.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.location_books = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// BookCreateBulk is the builder for creating many Book entities in bulk.
type BookCreateBulk struct {
	config
	err      error
	builders []*BookCreate
}

// Save creates the Book entities in the database.
func (bcb *BookCreateBulk) Save(ctx context.Context) ([]*Book, error) {
	if bcb.err != nil {
		return nil, bcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Book, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BookMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BookCreateBulk) SaveX(ctx context.Context) []*Book {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcb *BookCreateBulk) Exec(ctx context.Context) error {
	_, err := bcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcb *BookCreateBulk) ExecX(ctx context.Context) {
	if err := bcb.Exec(ctx); err != nil {
		panic(err)
	}
}
