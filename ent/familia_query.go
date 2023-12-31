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
	"github.com/dkrasnovdev/siberiana-api/ent/familia"
	"github.com/dkrasnovdev/siberiana-api/ent/herbarium"
	"github.com/dkrasnovdev/siberiana-api/ent/predicate"
)

// FamiliaQuery is the builder for querying Familia entities.
type FamiliaQuery struct {
	config
	ctx               *QueryContext
	order             []familia.OrderOption
	inters            []Interceptor
	predicates        []predicate.Familia
	withHerbaria      *HerbariumQuery
	modifiers         []func(*sql.Selector)
	loadTotal         []func(context.Context, []*Familia) error
	withNamedHerbaria map[string]*HerbariumQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FamiliaQuery builder.
func (fq *FamiliaQuery) Where(ps ...predicate.Familia) *FamiliaQuery {
	fq.predicates = append(fq.predicates, ps...)
	return fq
}

// Limit the number of records to be returned by this query.
func (fq *FamiliaQuery) Limit(limit int) *FamiliaQuery {
	fq.ctx.Limit = &limit
	return fq
}

// Offset to start from.
func (fq *FamiliaQuery) Offset(offset int) *FamiliaQuery {
	fq.ctx.Offset = &offset
	return fq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fq *FamiliaQuery) Unique(unique bool) *FamiliaQuery {
	fq.ctx.Unique = &unique
	return fq
}

// Order specifies how the records should be ordered.
func (fq *FamiliaQuery) Order(o ...familia.OrderOption) *FamiliaQuery {
	fq.order = append(fq.order, o...)
	return fq
}

// QueryHerbaria chains the current query on the "herbaria" edge.
func (fq *FamiliaQuery) QueryHerbaria() *HerbariumQuery {
	query := (&HerbariumClient{config: fq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(familia.Table, familia.FieldID, selector),
			sqlgraph.To(herbarium.Table, herbarium.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, familia.HerbariaTable, familia.HerbariaColumn),
		)
		fromU = sqlgraph.SetNeighbors(fq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Familia entity from the query.
// Returns a *NotFoundError when no Familia was found.
func (fq *FamiliaQuery) First(ctx context.Context) (*Familia, error) {
	nodes, err := fq.Limit(1).All(setContextOp(ctx, fq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{familia.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fq *FamiliaQuery) FirstX(ctx context.Context) *Familia {
	node, err := fq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Familia ID from the query.
// Returns a *NotFoundError when no Familia ID was found.
func (fq *FamiliaQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fq.Limit(1).IDs(setContextOp(ctx, fq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{familia.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fq *FamiliaQuery) FirstIDX(ctx context.Context) int {
	id, err := fq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Familia entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Familia entity is found.
// Returns a *NotFoundError when no Familia entities are found.
func (fq *FamiliaQuery) Only(ctx context.Context) (*Familia, error) {
	nodes, err := fq.Limit(2).All(setContextOp(ctx, fq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{familia.Label}
	default:
		return nil, &NotSingularError{familia.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fq *FamiliaQuery) OnlyX(ctx context.Context) *Familia {
	node, err := fq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Familia ID in the query.
// Returns a *NotSingularError when more than one Familia ID is found.
// Returns a *NotFoundError when no entities are found.
func (fq *FamiliaQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fq.Limit(2).IDs(setContextOp(ctx, fq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{familia.Label}
	default:
		err = &NotSingularError{familia.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fq *FamiliaQuery) OnlyIDX(ctx context.Context) int {
	id, err := fq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FamiliaSlice.
func (fq *FamiliaQuery) All(ctx context.Context) ([]*Familia, error) {
	ctx = setContextOp(ctx, fq.ctx, "All")
	if err := fq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Familia, *FamiliaQuery]()
	return withInterceptors[[]*Familia](ctx, fq, qr, fq.inters)
}

// AllX is like All, but panics if an error occurs.
func (fq *FamiliaQuery) AllX(ctx context.Context) []*Familia {
	nodes, err := fq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Familia IDs.
func (fq *FamiliaQuery) IDs(ctx context.Context) (ids []int, err error) {
	if fq.ctx.Unique == nil && fq.path != nil {
		fq.Unique(true)
	}
	ctx = setContextOp(ctx, fq.ctx, "IDs")
	if err = fq.Select(familia.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fq *FamiliaQuery) IDsX(ctx context.Context) []int {
	ids, err := fq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fq *FamiliaQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, fq.ctx, "Count")
	if err := fq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, fq, querierCount[*FamiliaQuery](), fq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (fq *FamiliaQuery) CountX(ctx context.Context) int {
	count, err := fq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fq *FamiliaQuery) Exist(ctx context.Context) (bool, error) {
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
func (fq *FamiliaQuery) ExistX(ctx context.Context) bool {
	exist, err := fq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FamiliaQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fq *FamiliaQuery) Clone() *FamiliaQuery {
	if fq == nil {
		return nil
	}
	return &FamiliaQuery{
		config:       fq.config,
		ctx:          fq.ctx.Clone(),
		order:        append([]familia.OrderOption{}, fq.order...),
		inters:       append([]Interceptor{}, fq.inters...),
		predicates:   append([]predicate.Familia{}, fq.predicates...),
		withHerbaria: fq.withHerbaria.Clone(),
		// clone intermediate query.
		sql:  fq.sql.Clone(),
		path: fq.path,
	}
}

// WithHerbaria tells the query-builder to eager-load the nodes that are connected to
// the "herbaria" edge. The optional arguments are used to configure the query builder of the edge.
func (fq *FamiliaQuery) WithHerbaria(opts ...func(*HerbariumQuery)) *FamiliaQuery {
	query := (&HerbariumClient{config: fq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	fq.withHerbaria = query
	return fq
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
//	client.Familia.Query().
//		GroupBy(familia.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (fq *FamiliaQuery) GroupBy(field string, fields ...string) *FamiliaGroupBy {
	fq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &FamiliaGroupBy{build: fq}
	grbuild.flds = &fq.ctx.Fields
	grbuild.label = familia.Label
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
//	client.Familia.Query().
//		Select(familia.FieldCreatedAt).
//		Scan(ctx, &v)
func (fq *FamiliaQuery) Select(fields ...string) *FamiliaSelect {
	fq.ctx.Fields = append(fq.ctx.Fields, fields...)
	sbuild := &FamiliaSelect{FamiliaQuery: fq}
	sbuild.label = familia.Label
	sbuild.flds, sbuild.scan = &fq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a FamiliaSelect configured with the given aggregations.
func (fq *FamiliaQuery) Aggregate(fns ...AggregateFunc) *FamiliaSelect {
	return fq.Select().Aggregate(fns...)
}

func (fq *FamiliaQuery) prepareQuery(ctx context.Context) error {
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
		if !familia.ValidColumn(f) {
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
	if familia.Policy == nil {
		return errors.New("ent: uninitialized familia.Policy (forgotten import ent/runtime?)")
	}
	if err := familia.Policy.EvalQuery(ctx, fq); err != nil {
		return err
	}
	return nil
}

func (fq *FamiliaQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Familia, error) {
	var (
		nodes       = []*Familia{}
		_spec       = fq.querySpec()
		loadedTypes = [1]bool{
			fq.withHerbaria != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Familia).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Familia{config: fq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
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
	if query := fq.withHerbaria; query != nil {
		if err := fq.loadHerbaria(ctx, query, nodes,
			func(n *Familia) { n.Edges.Herbaria = []*Herbarium{} },
			func(n *Familia, e *Herbarium) { n.Edges.Herbaria = append(n.Edges.Herbaria, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range fq.withNamedHerbaria {
		if err := fq.loadHerbaria(ctx, query, nodes,
			func(n *Familia) { n.appendNamedHerbaria(name) },
			func(n *Familia, e *Herbarium) { n.appendNamedHerbaria(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range fq.loadTotal {
		if err := fq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (fq *FamiliaQuery) loadHerbaria(ctx context.Context, query *HerbariumQuery, nodes []*Familia, init func(*Familia), assign func(*Familia, *Herbarium)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Familia)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Herbarium(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(familia.HerbariaColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.familia_herbaria
		if fk == nil {
			return fmt.Errorf(`foreign-key "familia_herbaria" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "familia_herbaria" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (fq *FamiliaQuery) sqlCount(ctx context.Context) (int, error) {
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

func (fq *FamiliaQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(familia.Table, familia.Columns, sqlgraph.NewFieldSpec(familia.FieldID, field.TypeInt))
	_spec.From = fq.sql
	if unique := fq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if fq.path != nil {
		_spec.Unique = true
	}
	if fields := fq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, familia.FieldID)
		for i := range fields {
			if fields[i] != familia.FieldID {
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

func (fq *FamiliaQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fq.driver.Dialect())
	t1 := builder.Table(familia.Table)
	columns := fq.ctx.Fields
	if len(columns) == 0 {
		columns = familia.Columns
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

// WithNamedHerbaria tells the query-builder to eager-load the nodes that are connected to the "herbaria"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (fq *FamiliaQuery) WithNamedHerbaria(name string, opts ...func(*HerbariumQuery)) *FamiliaQuery {
	query := (&HerbariumClient{config: fq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if fq.withNamedHerbaria == nil {
		fq.withNamedHerbaria = make(map[string]*HerbariumQuery)
	}
	fq.withNamedHerbaria[name] = query
	return fq
}

// FamiliaGroupBy is the group-by builder for Familia entities.
type FamiliaGroupBy struct {
	selector
	build *FamiliaQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fgb *FamiliaGroupBy) Aggregate(fns ...AggregateFunc) *FamiliaGroupBy {
	fgb.fns = append(fgb.fns, fns...)
	return fgb
}

// Scan applies the selector query and scans the result into the given value.
func (fgb *FamiliaGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fgb.build.ctx, "GroupBy")
	if err := fgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FamiliaQuery, *FamiliaGroupBy](ctx, fgb.build, fgb, fgb.build.inters, v)
}

func (fgb *FamiliaGroupBy) sqlScan(ctx context.Context, root *FamiliaQuery, v any) error {
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

// FamiliaSelect is the builder for selecting fields of Familia entities.
type FamiliaSelect struct {
	*FamiliaQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (fs *FamiliaSelect) Aggregate(fns ...AggregateFunc) *FamiliaSelect {
	fs.fns = append(fs.fns, fns...)
	return fs
}

// Scan applies the selector query and scans the result into the given value.
func (fs *FamiliaSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fs.ctx, "Select")
	if err := fs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FamiliaQuery, *FamiliaSelect](ctx, fs.FamiliaQuery, fs, fs.inters, v)
}

func (fs *FamiliaSelect) sqlScan(ctx context.Context, root *FamiliaQuery, v any) error {
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
