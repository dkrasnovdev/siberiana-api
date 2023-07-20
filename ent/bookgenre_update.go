// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dkrasnovdev/heritage-api/ent/bookgenre"
	"github.com/dkrasnovdev/heritage-api/ent/predicate"
)

// BookGenreUpdate is the builder for updating BookGenre entities.
type BookGenreUpdate struct {
	config
	hooks    []Hook
	mutation *BookGenreMutation
}

// Where appends a list predicates to the BookGenreUpdate builder.
func (bgu *BookGenreUpdate) Where(ps ...predicate.BookGenre) *BookGenreUpdate {
	bgu.mutation.Where(ps...)
	return bgu
}

// Mutation returns the BookGenreMutation object of the builder.
func (bgu *BookGenreUpdate) Mutation() *BookGenreMutation {
	return bgu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bgu *BookGenreUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, bgu.sqlSave, bgu.mutation, bgu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bgu *BookGenreUpdate) SaveX(ctx context.Context) int {
	affected, err := bgu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bgu *BookGenreUpdate) Exec(ctx context.Context) error {
	_, err := bgu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bgu *BookGenreUpdate) ExecX(ctx context.Context) {
	if err := bgu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bgu *BookGenreUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(bookgenre.Table, bookgenre.Columns, sqlgraph.NewFieldSpec(bookgenre.FieldID, field.TypeInt))
	if ps := bgu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bgu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bookgenre.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bgu.mutation.done = true
	return n, nil
}

// BookGenreUpdateOne is the builder for updating a single BookGenre entity.
type BookGenreUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BookGenreMutation
}

// Mutation returns the BookGenreMutation object of the builder.
func (bguo *BookGenreUpdateOne) Mutation() *BookGenreMutation {
	return bguo.mutation
}

// Where appends a list predicates to the BookGenreUpdate builder.
func (bguo *BookGenreUpdateOne) Where(ps ...predicate.BookGenre) *BookGenreUpdateOne {
	bguo.mutation.Where(ps...)
	return bguo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (bguo *BookGenreUpdateOne) Select(field string, fields ...string) *BookGenreUpdateOne {
	bguo.fields = append([]string{field}, fields...)
	return bguo
}

// Save executes the query and returns the updated BookGenre entity.
func (bguo *BookGenreUpdateOne) Save(ctx context.Context) (*BookGenre, error) {
	return withHooks(ctx, bguo.sqlSave, bguo.mutation, bguo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bguo *BookGenreUpdateOne) SaveX(ctx context.Context) *BookGenre {
	node, err := bguo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (bguo *BookGenreUpdateOne) Exec(ctx context.Context) error {
	_, err := bguo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bguo *BookGenreUpdateOne) ExecX(ctx context.Context) {
	if err := bguo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bguo *BookGenreUpdateOne) sqlSave(ctx context.Context) (_node *BookGenre, err error) {
	_spec := sqlgraph.NewUpdateSpec(bookgenre.Table, bookgenre.Columns, sqlgraph.NewFieldSpec(bookgenre.FieldID, field.TypeInt))
	id, ok := bguo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "BookGenre.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := bguo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, bookgenre.FieldID)
		for _, f := range fields {
			if !bookgenre.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != bookgenre.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := bguo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	_node = &BookGenre{config: bguo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, bguo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bookgenre.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	bguo.mutation.done = true
	return _node, nil
}
