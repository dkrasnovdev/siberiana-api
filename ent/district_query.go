// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dkrasnovdev/siberiana-api/ent/district"
	"github.com/dkrasnovdev/siberiana-api/ent/location"
	"github.com/dkrasnovdev/siberiana-api/ent/predicate"
)

// DistrictQuery is the builder for querying District entities.
type DistrictQuery struct {
	config
	ctx                *QueryContext
	order              []district.OrderOption
	inters             []Interceptor
	predicates         []predicate.District
	withLocations      *LocationQuery
	modifiers          []func(*sql.Selector)
	loadTotal          []func(context.Context, []*District) error
	withNamedLocations map[string]*LocationQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DistrictQuery builder.
func (dq *DistrictQuery) Where(ps ...predicate.District) *DistrictQuery {
	dq.predicates = append(dq.predicates, ps...)
	return dq
}

// Limit the number of records to be returned by this query.
func (dq *DistrictQuery) Limit(limit int) *DistrictQuery {
	dq.ctx.Limit = &limit
	return dq
}

// Offset to start from.
func (dq *DistrictQuery) Offset(offset int) *DistrictQuery {
	dq.ctx.Offset = &offset
	return dq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dq *DistrictQuery) Unique(unique bool) *DistrictQuery {
	dq.ctx.Unique = &unique
	return dq
}

// Order specifies how the records should be ordered.
func (dq *DistrictQuery) Order(o ...district.OrderOption) *DistrictQuery {
	dq.order = append(dq.order, o...)
	return dq
}

// QueryLocations chains the current query on the "locations" edge.
func (dq *DistrictQuery) QueryLocations() *LocationQuery {
	query := (&LocationClient{config: dq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(district.Table, district.FieldID, selector),
			sqlgraph.To(location.Table, location.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, district.LocationsTable, district.LocationsColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first District entity from the query.
// Returns a *NotFoundError when no District was found.
func (dq *DistrictQuery) First(ctx context.Context) (*District, error) {
	nodes, err := dq.Limit(1).All(setContextOp(ctx, dq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{district.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dq *DistrictQuery) FirstX(ctx context.Context) *District {
	node, err := dq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first District ID from the query.
// Returns a *NotFoundError when no District ID was found.
func (dq *DistrictQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dq.Limit(1).IDs(setContextOp(ctx, dq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{district.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dq *DistrictQuery) FirstIDX(ctx context.Context) int {
	id, err := dq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single District entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one District entity is found.
// Returns a *NotFoundError when no District entities are found.
func (dq *DistrictQuery) Only(ctx context.Context) (*District, error) {
	nodes, err := dq.Limit(2).All(setContextOp(ctx, dq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{district.Label}
	default:
		return nil, &NotSingularError{district.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dq *DistrictQuery) OnlyX(ctx context.Context) *District {
	node, err := dq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only District ID in the query.
// Returns a *NotSingularError when more than one District ID is found.
// Returns a *NotFoundError when no entities are found.
func (dq *DistrictQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dq.Limit(2).IDs(setContextOp(ctx, dq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{district.Label}
	default:
		err = &NotSingularError{district.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dq *DistrictQuery) OnlyIDX(ctx context.Context) int {
	id, err := dq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Districts.
func (dq *DistrictQuery) All(ctx context.Context) ([]*District, error) {
	ctx = setContextOp(ctx, dq.ctx, "All")
	if err := dq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*District, *DistrictQuery]()
	return withInterceptors[[]*District](ctx, dq, qr, dq.inters)
}

// AllX is like All, but panics if an error occurs.
func (dq *DistrictQuery) AllX(ctx context.Context) []*District {
	nodes, err := dq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of District IDs.
func (dq *DistrictQuery) IDs(ctx context.Context) (ids []int, err error) {
	if dq.ctx.Unique == nil && dq.path != nil {
		dq.Unique(true)
	}
	ctx = setContextOp(ctx, dq.ctx, "IDs")
	if err = dq.Select(district.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dq *DistrictQuery) IDsX(ctx context.Context) []int {
	ids, err := dq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dq *DistrictQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, dq.ctx, "Count")
	if err := dq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, dq, querierCount[*DistrictQuery](), dq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (dq *DistrictQuery) CountX(ctx context.Context) int {
	count, err := dq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dq *DistrictQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, dq.ctx, "Exist")
	switch _, err := dq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (dq *DistrictQuery) ExistX(ctx context.Context) bool {
	exist, err := dq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DistrictQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dq *DistrictQuery) Clone() *DistrictQuery {
	if dq == nil {
		return nil
	}
	return &DistrictQuery{
		config:        dq.config,
		ctx:           dq.ctx.Clone(),
		order:         append([]district.OrderOption{}, dq.order...),
		inters:        append([]Interceptor{}, dq.inters...),
		predicates:    append([]predicate.District{}, dq.predicates...),
		withLocations: dq.withLocations.Clone(),
		// clone intermediate query.
		sql:  dq.sql.Clone(),
		path: dq.path,
	}
}

// WithLocations tells the query-builder to eager-load the nodes that are connected to
// the "locations" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DistrictQuery) WithLocations(opts ...func(*LocationQuery)) *DistrictQuery {
	query := (&LocationClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dq.withLocations = query
	return dq
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
//	client.District.Query().
//		GroupBy(district.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dq *DistrictQuery) GroupBy(field string, fields ...string) *DistrictGroupBy {
	dq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &DistrictGroupBy{build: dq}
	grbuild.flds = &dq.ctx.Fields
	grbuild.label = district.Label
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
//	client.District.Query().
//		Select(district.FieldCreatedAt).
//		Scan(ctx, &v)
func (dq *DistrictQuery) Select(fields ...string) *DistrictSelect {
	dq.ctx.Fields = append(dq.ctx.Fields, fields...)
	sbuild := &DistrictSelect{DistrictQuery: dq}
	sbuild.label = district.Label
	sbuild.flds, sbuild.scan = &dq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DistrictSelect configured with the given aggregations.
func (dq *DistrictQuery) Aggregate(fns ...AggregateFunc) *DistrictSelect {
	return dq.Select().Aggregate(fns...)
}

func (dq *DistrictQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range dq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, dq); err != nil {
				return err
			}
		}
	}
	for _, f := range dq.ctx.Fields {
		if !district.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dq.path != nil {
		prev, err := dq.path(ctx)
		if err != nil {
			return err
		}
		dq.sql = prev
	}
	if district.Policy == nil {
		return errors.New("ent: uninitialized district.Policy (forgotten import ent/runtime?)")
	}
	if err := district.Policy.EvalQuery(ctx, dq); err != nil {
		return err
	}
	return nil
}

func (dq *DistrictQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*District, error) {
	var (
		nodes       = []*District{}
		_spec       = dq.querySpec()
		loadedTypes = [1]bool{
			dq.withLocations != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*District).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &District{config: dq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(dq.modifiers) > 0 {
		_spec.Modifiers = dq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := dq.withLocations; query != nil {
		if err := dq.loadLocations(ctx, query, nodes,
			func(n *District) { n.Edges.Locations = []*Location{} },
			func(n *District, e *Location) { n.Edges.Locations = append(n.Edges.Locations, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range dq.withNamedLocations {
		if err := dq.loadLocations(ctx, query, nodes,
			func(n *District) { n.appendNamedLocations(name) },
			func(n *District, e *Location) { n.appendNamedLocations(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range dq.loadTotal {
		if err := dq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (dq *DistrictQuery) loadLocations(ctx context.Context, query *LocationQuery, nodes []*District, init func(*District), assign func(*District, *Location)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*District)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Location(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(district.LocationsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.location_district
		if fk == nil {
			return fmt.Errorf(`foreign-key "location_district" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "location_district" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (dq *DistrictQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dq.querySpec()
	if len(dq.modifiers) > 0 {
		_spec.Modifiers = dq.modifiers
	}
	_spec.Node.Columns = dq.ctx.Fields
	if len(dq.ctx.Fields) > 0 {
		_spec.Unique = dq.ctx.Unique != nil && *dq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, dq.driver, _spec)
}

func (dq *DistrictQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(district.Table, district.Columns, sqlgraph.NewFieldSpec(district.FieldID, field.TypeInt))
	_spec.From = dq.sql
	if unique := dq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if dq.path != nil {
		_spec.Unique = true
	}
	if fields := dq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, district.FieldID)
		for i := range fields {
			if fields[i] != district.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dq *DistrictQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dq.driver.Dialect())
	t1 := builder.Table(district.Table)
	columns := dq.ctx.Fields
	if len(columns) == 0 {
		columns = district.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dq.sql != nil {
		selector = dq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dq.ctx.Unique != nil && *dq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range dq.predicates {
		p(selector)
	}
	for _, p := range dq.order {
		p(selector)
	}
	if offset := dq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedLocations tells the query-builder to eager-load the nodes that are connected to the "locations"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (dq *DistrictQuery) WithNamedLocations(name string, opts ...func(*LocationQuery)) *DistrictQuery {
	query := (&LocationClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if dq.withNamedLocations == nil {
		dq.withNamedLocations = make(map[string]*LocationQuery)
	}
	dq.withNamedLocations[name] = query
	return dq
}

// DistrictGroupBy is the group-by builder for District entities.
type DistrictGroupBy struct {
	selector
	build *DistrictQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dgb *DistrictGroupBy) Aggregate(fns ...AggregateFunc) *DistrictGroupBy {
	dgb.fns = append(dgb.fns, fns...)
	return dgb
}

// Scan applies the selector query and scans the result into the given value.
func (dgb *DistrictGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dgb.build.ctx, "GroupBy")
	if err := dgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DistrictQuery, *DistrictGroupBy](ctx, dgb.build, dgb, dgb.build.inters, v)
}

func (dgb *DistrictGroupBy) sqlScan(ctx context.Context, root *DistrictQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(dgb.fns))
	for _, fn := range dgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*dgb.flds)+len(dgb.fns))
		for _, f := range *dgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*dgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// DistrictSelect is the builder for selecting fields of District entities.
type DistrictSelect struct {
	*DistrictQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ds *DistrictSelect) Aggregate(fns ...AggregateFunc) *DistrictSelect {
	ds.fns = append(ds.fns, fns...)
	return ds
}

// Scan applies the selector query and scans the result into the given value.
func (ds *DistrictSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ds.ctx, "Select")
	if err := ds.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DistrictQuery, *DistrictSelect](ctx, ds.DistrictQuery, ds, ds.inters, v)
}

func (ds *DistrictSelect) sqlScan(ctx context.Context, root *DistrictQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ds.fns))
	for _, fn := range ds.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ds.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
