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
	"github.com/dkrasnovdev/siberiana-api/ent/country"
	"github.com/dkrasnovdev/siberiana-api/ent/location"
	"github.com/dkrasnovdev/siberiana-api/ent/predicate"
)

// CountryQuery is the builder for querying Country entities.
type CountryQuery struct {
	config
	ctx               *QueryContext
	order             []country.OrderOption
	inters            []Interceptor
	predicates        []predicate.Country
	withLocation      *LocationQuery
	modifiers         []func(*sql.Selector)
	loadTotal         []func(context.Context, []*Country) error
	withNamedLocation map[string]*LocationQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CountryQuery builder.
func (cq *CountryQuery) Where(ps ...predicate.Country) *CountryQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit the number of records to be returned by this query.
func (cq *CountryQuery) Limit(limit int) *CountryQuery {
	cq.ctx.Limit = &limit
	return cq
}

// Offset to start from.
func (cq *CountryQuery) Offset(offset int) *CountryQuery {
	cq.ctx.Offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *CountryQuery) Unique(unique bool) *CountryQuery {
	cq.ctx.Unique = &unique
	return cq
}

// Order specifies how the records should be ordered.
func (cq *CountryQuery) Order(o ...country.OrderOption) *CountryQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryLocation chains the current query on the "location" edge.
func (cq *CountryQuery) QueryLocation() *LocationQuery {
	query := (&LocationClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(country.Table, country.FieldID, selector),
			sqlgraph.To(location.Table, location.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, country.LocationTable, country.LocationColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Country entity from the query.
// Returns a *NotFoundError when no Country was found.
func (cq *CountryQuery) First(ctx context.Context) (*Country, error) {
	nodes, err := cq.Limit(1).All(setContextOp(ctx, cq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{country.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *CountryQuery) FirstX(ctx context.Context) *Country {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Country ID from the query.
// Returns a *NotFoundError when no Country ID was found.
func (cq *CountryQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(1).IDs(setContextOp(ctx, cq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{country.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *CountryQuery) FirstIDX(ctx context.Context) int {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Country entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Country entity is found.
// Returns a *NotFoundError when no Country entities are found.
func (cq *CountryQuery) Only(ctx context.Context) (*Country, error) {
	nodes, err := cq.Limit(2).All(setContextOp(ctx, cq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{country.Label}
	default:
		return nil, &NotSingularError{country.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *CountryQuery) OnlyX(ctx context.Context) *Country {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Country ID in the query.
// Returns a *NotSingularError when more than one Country ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *CountryQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(2).IDs(setContextOp(ctx, cq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{country.Label}
	default:
		err = &NotSingularError{country.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *CountryQuery) OnlyIDX(ctx context.Context) int {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Countries.
func (cq *CountryQuery) All(ctx context.Context) ([]*Country, error) {
	ctx = setContextOp(ctx, cq.ctx, "All")
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Country, *CountryQuery]()
	return withInterceptors[[]*Country](ctx, cq, qr, cq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cq *CountryQuery) AllX(ctx context.Context) []*Country {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Country IDs.
func (cq *CountryQuery) IDs(ctx context.Context) (ids []int, err error) {
	if cq.ctx.Unique == nil && cq.path != nil {
		cq.Unique(true)
	}
	ctx = setContextOp(ctx, cq.ctx, "IDs")
	if err = cq.Select(country.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *CountryQuery) IDsX(ctx context.Context) []int {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *CountryQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cq.ctx, "Count")
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cq, querierCount[*CountryQuery](), cq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cq *CountryQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *CountryQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cq.ctx, "Exist")
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *CountryQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CountryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *CountryQuery) Clone() *CountryQuery {
	if cq == nil {
		return nil
	}
	return &CountryQuery{
		config:       cq.config,
		ctx:          cq.ctx.Clone(),
		order:        append([]country.OrderOption{}, cq.order...),
		inters:       append([]Interceptor{}, cq.inters...),
		predicates:   append([]predicate.Country{}, cq.predicates...),
		withLocation: cq.withLocation.Clone(),
		// clone intermediate query.
		sql:  cq.sql.Clone(),
		path: cq.path,
	}
}

// WithLocation tells the query-builder to eager-load the nodes that are connected to
// the "location" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CountryQuery) WithLocation(opts ...func(*LocationQuery)) *CountryQuery {
	query := (&LocationClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withLocation = query
	return cq
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
//	client.Country.Query().
//		GroupBy(country.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cq *CountryQuery) GroupBy(field string, fields ...string) *CountryGroupBy {
	cq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CountryGroupBy{build: cq}
	grbuild.flds = &cq.ctx.Fields
	grbuild.label = country.Label
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
//	client.Country.Query().
//		Select(country.FieldCreatedAt).
//		Scan(ctx, &v)
func (cq *CountryQuery) Select(fields ...string) *CountrySelect {
	cq.ctx.Fields = append(cq.ctx.Fields, fields...)
	sbuild := &CountrySelect{CountryQuery: cq}
	sbuild.label = country.Label
	sbuild.flds, sbuild.scan = &cq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CountrySelect configured with the given aggregations.
func (cq *CountryQuery) Aggregate(fns ...AggregateFunc) *CountrySelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *CountryQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cq); err != nil {
				return err
			}
		}
	}
	for _, f := range cq.ctx.Fields {
		if !country.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	if country.Policy == nil {
		return errors.New("ent: uninitialized country.Policy (forgotten import ent/runtime?)")
	}
	if err := country.Policy.EvalQuery(ctx, cq); err != nil {
		return err
	}
	return nil
}

func (cq *CountryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Country, error) {
	var (
		nodes       = []*Country{}
		_spec       = cq.querySpec()
		loadedTypes = [1]bool{
			cq.withLocation != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Country).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Country{config: cq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(cq.modifiers) > 0 {
		_spec.Modifiers = cq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cq.withLocation; query != nil {
		if err := cq.loadLocation(ctx, query, nodes,
			func(n *Country) { n.Edges.Location = []*Location{} },
			func(n *Country, e *Location) { n.Edges.Location = append(n.Edges.Location, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range cq.withNamedLocation {
		if err := cq.loadLocation(ctx, query, nodes,
			func(n *Country) { n.appendNamedLocation(name) },
			func(n *Country, e *Location) { n.appendNamedLocation(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range cq.loadTotal {
		if err := cq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cq *CountryQuery) loadLocation(ctx context.Context, query *LocationQuery, nodes []*Country, init func(*Country), assign func(*Country, *Location)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Country)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Location(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(country.LocationColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.location_country
		if fk == nil {
			return fmt.Errorf(`foreign-key "location_country" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "location_country" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (cq *CountryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	if len(cq.modifiers) > 0 {
		_spec.Modifiers = cq.modifiers
	}
	_spec.Node.Columns = cq.ctx.Fields
	if len(cq.ctx.Fields) > 0 {
		_spec.Unique = cq.ctx.Unique != nil && *cq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *CountryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(country.Table, country.Columns, sqlgraph.NewFieldSpec(country.FieldID, field.TypeInt))
	_spec.From = cq.sql
	if unique := cq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cq.path != nil {
		_spec.Unique = true
	}
	if fields := cq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, country.FieldID)
		for i := range fields {
			if fields[i] != country.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *CountryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(country.Table)
	columns := cq.ctx.Fields
	if len(columns) == 0 {
		columns = country.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.ctx.Unique != nil && *cq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedLocation tells the query-builder to eager-load the nodes that are connected to the "location"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (cq *CountryQuery) WithNamedLocation(name string, opts ...func(*LocationQuery)) *CountryQuery {
	query := (&LocationClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if cq.withNamedLocation == nil {
		cq.withNamedLocation = make(map[string]*LocationQuery)
	}
	cq.withNamedLocation[name] = query
	return cq
}

// CountryGroupBy is the group-by builder for Country entities.
type CountryGroupBy struct {
	selector
	build *CountryQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *CountryGroupBy) Aggregate(fns ...AggregateFunc) *CountryGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *CountryGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cgb.build.ctx, "GroupBy")
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CountryQuery, *CountryGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (cgb *CountryGroupBy) sqlScan(ctx context.Context, root *CountryQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cgb.flds)+len(cgb.fns))
		for _, f := range *cgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CountrySelect is the builder for selecting fields of Country entities.
type CountrySelect struct {
	*CountryQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *CountrySelect) Aggregate(fns ...AggregateFunc) *CountrySelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *CountrySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cs.ctx, "Select")
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CountryQuery, *CountrySelect](ctx, cs.CountryQuery, cs, cs.inters, v)
}

func (cs *CountrySelect) sqlScan(ctx context.Context, root *CountryQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cs.fns))
	for _, fn := range cs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
