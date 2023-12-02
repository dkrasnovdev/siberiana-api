// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dkrasnovdev/siberiana-api/ent/favourite"
	"github.com/dkrasnovdev/siberiana-api/ent/predicate"
)

// FavouriteQuery is the builder for querying Favourite entities.
type FavouriteQuery struct {
	config
	ctx        *QueryContext
	order      []favourite.OrderOption
	inters     []Interceptor
	predicates []predicate.Favourite
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*Favourite) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FavouriteQuery builder.
func (fq *FavouriteQuery) Where(ps ...predicate.Favourite) *FavouriteQuery {
	fq.predicates = append(fq.predicates, ps...)
	return fq
}

// Limit the number of records to be returned by this query.
func (fq *FavouriteQuery) Limit(limit int) *FavouriteQuery {
	fq.ctx.Limit = &limit
	return fq
}

// Offset to start from.
func (fq *FavouriteQuery) Offset(offset int) *FavouriteQuery {
	fq.ctx.Offset = &offset
	return fq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fq *FavouriteQuery) Unique(unique bool) *FavouriteQuery {
	fq.ctx.Unique = &unique
	return fq
}

// Order specifies how the records should be ordered.
func (fq *FavouriteQuery) Order(o ...favourite.OrderOption) *FavouriteQuery {
	fq.order = append(fq.order, o...)
	return fq
}

// First returns the first Favourite entity from the query.
// Returns a *NotFoundError when no Favourite was found.
func (fq *FavouriteQuery) First(ctx context.Context) (*Favourite, error) {
	nodes, err := fq.Limit(1).All(setContextOp(ctx, fq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{favourite.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fq *FavouriteQuery) FirstX(ctx context.Context) *Favourite {
	node, err := fq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Favourite ID from the query.
// Returns a *NotFoundError when no Favourite ID was found.
func (fq *FavouriteQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fq.Limit(1).IDs(setContextOp(ctx, fq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{favourite.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fq *FavouriteQuery) FirstIDX(ctx context.Context) int {
	id, err := fq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Favourite entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Favourite entity is found.
// Returns a *NotFoundError when no Favourite entities are found.
func (fq *FavouriteQuery) Only(ctx context.Context) (*Favourite, error) {
	nodes, err := fq.Limit(2).All(setContextOp(ctx, fq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{favourite.Label}
	default:
		return nil, &NotSingularError{favourite.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fq *FavouriteQuery) OnlyX(ctx context.Context) *Favourite {
	node, err := fq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Favourite ID in the query.
// Returns a *NotSingularError when more than one Favourite ID is found.
// Returns a *NotFoundError when no entities are found.
func (fq *FavouriteQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fq.Limit(2).IDs(setContextOp(ctx, fq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{favourite.Label}
	default:
		err = &NotSingularError{favourite.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fq *FavouriteQuery) OnlyIDX(ctx context.Context) int {
	id, err := fq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Favourites.
func (fq *FavouriteQuery) All(ctx context.Context) ([]*Favourite, error) {
	ctx = setContextOp(ctx, fq.ctx, "All")
	if err := fq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Favourite, *FavouriteQuery]()
	return withInterceptors[[]*Favourite](ctx, fq, qr, fq.inters)
}

// AllX is like All, but panics if an error occurs.
func (fq *FavouriteQuery) AllX(ctx context.Context) []*Favourite {
	nodes, err := fq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Favourite IDs.
func (fq *FavouriteQuery) IDs(ctx context.Context) (ids []int, err error) {
	if fq.ctx.Unique == nil && fq.path != nil {
		fq.Unique(true)
	}
	ctx = setContextOp(ctx, fq.ctx, "IDs")
	if err = fq.Select(favourite.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fq *FavouriteQuery) IDsX(ctx context.Context) []int {
	ids, err := fq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fq *FavouriteQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, fq.ctx, "Count")
	if err := fq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, fq, querierCount[*FavouriteQuery](), fq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (fq *FavouriteQuery) CountX(ctx context.Context) int {
	count, err := fq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fq *FavouriteQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, fq.ctx, "Exist")
	switch _, err := fq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (fq *FavouriteQuery) ExistX(ctx context.Context) bool {
	exist, err := fq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FavouriteQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fq *FavouriteQuery) Clone() *FavouriteQuery {
	if fq == nil {
		return nil
	}
	return &FavouriteQuery{
		config:     fq.config,
		ctx:        fq.ctx.Clone(),
		order:      append([]favourite.OrderOption{}, fq.order...),
		inters:     append([]Interceptor{}, fq.inters...),
		predicates: append([]predicate.Favourite{}, fq.predicates...),
		// clone intermediate query.
		sql:  fq.sql.Clone(),
		path: fq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Favourite.Query().
//		GroupBy(favourite.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (fq *FavouriteQuery) GroupBy(field string, fields ...string) *FavouriteGroupBy {
	fq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &FavouriteGroupBy{build: fq}
	grbuild.flds = &fq.ctx.Fields
	grbuild.label = favourite.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.Favourite.Query().
//		Select(favourite.FieldCreatedAt).
//		Scan(ctx, &v)
func (fq *FavouriteQuery) Select(fields ...string) *FavouriteSelect {
	fq.ctx.Fields = append(fq.ctx.Fields, fields...)
	sbuild := &FavouriteSelect{FavouriteQuery: fq}
	sbuild.label = favourite.Label
	sbuild.flds, sbuild.scan = &fq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a FavouriteSelect configured with the given aggregations.
func (fq *FavouriteQuery) Aggregate(fns ...AggregateFunc) *FavouriteSelect {
	return fq.Select().Aggregate(fns...)
}

func (fq *FavouriteQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range fq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, fq); err != nil {
				return err
			}
		}
	}
	for _, f := range fq.ctx.Fields {
		if !favourite.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fq.path != nil {
		prev, err := fq.path(ctx)
		if err != nil {
			return err
		}
		fq.sql = prev
	}
	if favourite.Policy == nil {
		return errors.New("ent: uninitialized favourite.Policy (forgotten import ent/runtime?)")
	}
	if err := favourite.Policy.EvalQuery(ctx, fq); err != nil {
		return err
	}
	return nil
}

func (fq *FavouriteQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Favourite, error) {
	var (
		nodes = []*Favourite{}
		_spec = fq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Favourite).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Favourite{config: fq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(fq.modifiers) > 0 {
		_spec.Modifiers = fq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, fq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	for i := range fq.loadTotal {
		if err := fq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (fq *FavouriteQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fq.querySpec()
	if len(fq.modifiers) > 0 {
		_spec.Modifiers = fq.modifiers
	}
	_spec.Node.Columns = fq.ctx.Fields
	if len(fq.ctx.Fields) > 0 {
		_spec.Unique = fq.ctx.Unique != nil && *fq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, fq.driver, _spec)
}

func (fq *FavouriteQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(favourite.Table, favourite.Columns, sqlgraph.NewFieldSpec(favourite.FieldID, field.TypeInt))
	_spec.From = fq.sql
	if unique := fq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if fq.path != nil {
		_spec.Unique = true
	}
	if fields := fq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, favourite.FieldID)
		for i := range fields {
			if fields[i] != favourite.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := fq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fq *FavouriteQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fq.driver.Dialect())
	t1 := builder.Table(favourite.Table)
	columns := fq.ctx.Fields
	if len(columns) == 0 {
		columns = favourite.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fq.sql != nil {
		selector = fq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fq.ctx.Unique != nil && *fq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range fq.predicates {
		p(selector)
	}
	for _, p := range fq.order {
		p(selector)
	}
	if offset := fq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FavouriteGroupBy is the group-by builder for Favourite entities.
type FavouriteGroupBy struct {
	selector
	build *FavouriteQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fgb *FavouriteGroupBy) Aggregate(fns ...AggregateFunc) *FavouriteGroupBy {
	fgb.fns = append(fgb.fns, fns...)
	return fgb
}

// Scan applies the selector query and scans the result into the given value.
func (fgb *FavouriteGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fgb.build.ctx, "GroupBy")
	if err := fgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FavouriteQuery, *FavouriteGroupBy](ctx, fgb.build, fgb, fgb.build.inters, v)
}

func (fgb *FavouriteGroupBy) sqlScan(ctx context.Context, root *FavouriteQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(fgb.fns))
	for _, fn := range fgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*fgb.flds)+len(fgb.fns))
		for _, f := range *fgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*fgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// FavouriteSelect is the builder for selecting fields of Favourite entities.
type FavouriteSelect struct {
	*FavouriteQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (fs *FavouriteSelect) Aggregate(fns ...AggregateFunc) *FavouriteSelect {
	fs.fns = append(fs.fns, fns...)
	return fs
}

// Scan applies the selector query and scans the result into the given value.
func (fs *FavouriteSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fs.ctx, "Select")
	if err := fs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FavouriteQuery, *FavouriteSelect](ctx, fs.FavouriteQuery, fs, fs.inters, v)
}

func (fs *FavouriteSelect) sqlScan(ctx context.Context, root *FavouriteQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(fs.fns))
	for _, fn := range fs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*fs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
