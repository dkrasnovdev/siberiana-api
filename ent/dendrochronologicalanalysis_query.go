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
	"github.com/dkrasnovdev/siberiana-api/ent/dendrochronologicalanalysis"
	"github.com/dkrasnovdev/siberiana-api/ent/dendrochronology"
	"github.com/dkrasnovdev/siberiana-api/ent/predicate"
)

// DendrochronologicalAnalysisQuery is the builder for querying DendrochronologicalAnalysis entities.
type DendrochronologicalAnalysisQuery struct {
	config
	ctx                  *QueryContext
	order                []dendrochronologicalanalysis.OrderOption
	inters               []Interceptor
	predicates           []predicate.DendrochronologicalAnalysis
	withDendrochronology *DendrochronologyQuery
	withFKs              bool
	modifiers            []func(*sql.Selector)
	loadTotal            []func(context.Context, []*DendrochronologicalAnalysis) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DendrochronologicalAnalysisQuery builder.
func (daq *DendrochronologicalAnalysisQuery) Where(ps ...predicate.DendrochronologicalAnalysis) *DendrochronologicalAnalysisQuery {
	daq.predicates = append(daq.predicates, ps...)
	return daq
}

// Limit the number of records to be returned by this query.
func (daq *DendrochronologicalAnalysisQuery) Limit(limit int) *DendrochronologicalAnalysisQuery {
	daq.ctx.Limit = &limit
	return daq
}

// Offset to start from.
func (daq *DendrochronologicalAnalysisQuery) Offset(offset int) *DendrochronologicalAnalysisQuery {
	daq.ctx.Offset = &offset
	return daq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (daq *DendrochronologicalAnalysisQuery) Unique(unique bool) *DendrochronologicalAnalysisQuery {
	daq.ctx.Unique = &unique
	return daq
}

// Order specifies how the records should be ordered.
func (daq *DendrochronologicalAnalysisQuery) Order(o ...dendrochronologicalanalysis.OrderOption) *DendrochronologicalAnalysisQuery {
	daq.order = append(daq.order, o...)
	return daq
}

// QueryDendrochronology chains the current query on the "dendrochronology" edge.
func (daq *DendrochronologicalAnalysisQuery) QueryDendrochronology() *DendrochronologyQuery {
	query := (&DendrochronologyClient{config: daq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := daq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := daq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(dendrochronologicalanalysis.Table, dendrochronologicalanalysis.FieldID, selector),
			sqlgraph.To(dendrochronology.Table, dendrochronology.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, dendrochronologicalanalysis.DendrochronologyTable, dendrochronologicalanalysis.DendrochronologyColumn),
		)
		fromU = sqlgraph.SetNeighbors(daq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first DendrochronologicalAnalysis entity from the query.
// Returns a *NotFoundError when no DendrochronologicalAnalysis was found.
func (daq *DendrochronologicalAnalysisQuery) First(ctx context.Context) (*DendrochronologicalAnalysis, error) {
	nodes, err := daq.Limit(1).All(setContextOp(ctx, daq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{dendrochronologicalanalysis.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (daq *DendrochronologicalAnalysisQuery) FirstX(ctx context.Context) *DendrochronologicalAnalysis {
	node, err := daq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DendrochronologicalAnalysis ID from the query.
// Returns a *NotFoundError when no DendrochronologicalAnalysis ID was found.
func (daq *DendrochronologicalAnalysisQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = daq.Limit(1).IDs(setContextOp(ctx, daq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{dendrochronologicalanalysis.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (daq *DendrochronologicalAnalysisQuery) FirstIDX(ctx context.Context) int {
	id, err := daq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DendrochronologicalAnalysis entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one DendrochronologicalAnalysis entity is found.
// Returns a *NotFoundError when no DendrochronologicalAnalysis entities are found.
func (daq *DendrochronologicalAnalysisQuery) Only(ctx context.Context) (*DendrochronologicalAnalysis, error) {
	nodes, err := daq.Limit(2).All(setContextOp(ctx, daq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{dendrochronologicalanalysis.Label}
	default:
		return nil, &NotSingularError{dendrochronologicalanalysis.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (daq *DendrochronologicalAnalysisQuery) OnlyX(ctx context.Context) *DendrochronologicalAnalysis {
	node, err := daq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DendrochronologicalAnalysis ID in the query.
// Returns a *NotSingularError when more than one DendrochronologicalAnalysis ID is found.
// Returns a *NotFoundError when no entities are found.
func (daq *DendrochronologicalAnalysisQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = daq.Limit(2).IDs(setContextOp(ctx, daq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{dendrochronologicalanalysis.Label}
	default:
		err = &NotSingularError{dendrochronologicalanalysis.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (daq *DendrochronologicalAnalysisQuery) OnlyIDX(ctx context.Context) int {
	id, err := daq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DendrochronologicalAnalyses.
func (daq *DendrochronologicalAnalysisQuery) All(ctx context.Context) ([]*DendrochronologicalAnalysis, error) {
	ctx = setContextOp(ctx, daq.ctx, "All")
	if err := daq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*DendrochronologicalAnalysis, *DendrochronologicalAnalysisQuery]()
	return withInterceptors[[]*DendrochronologicalAnalysis](ctx, daq, qr, daq.inters)
}

// AllX is like All, but panics if an error occurs.
func (daq *DendrochronologicalAnalysisQuery) AllX(ctx context.Context) []*DendrochronologicalAnalysis {
	nodes, err := daq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DendrochronologicalAnalysis IDs.
func (daq *DendrochronologicalAnalysisQuery) IDs(ctx context.Context) (ids []int, err error) {
	if daq.ctx.Unique == nil && daq.path != nil {
		daq.Unique(true)
	}
	ctx = setContextOp(ctx, daq.ctx, "IDs")
	if err = daq.Select(dendrochronologicalanalysis.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (daq *DendrochronologicalAnalysisQuery) IDsX(ctx context.Context) []int {
	ids, err := daq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (daq *DendrochronologicalAnalysisQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, daq.ctx, "Count")
	if err := daq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, daq, querierCount[*DendrochronologicalAnalysisQuery](), daq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (daq *DendrochronologicalAnalysisQuery) CountX(ctx context.Context) int {
	count, err := daq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (daq *DendrochronologicalAnalysisQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, daq.ctx, "Exist")
	switch _, err := daq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (daq *DendrochronologicalAnalysisQuery) ExistX(ctx context.Context) bool {
	exist, err := daq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DendrochronologicalAnalysisQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (daq *DendrochronologicalAnalysisQuery) Clone() *DendrochronologicalAnalysisQuery {
	if daq == nil {
		return nil
	}
	return &DendrochronologicalAnalysisQuery{
		config:               daq.config,
		ctx:                  daq.ctx.Clone(),
		order:                append([]dendrochronologicalanalysis.OrderOption{}, daq.order...),
		inters:               append([]Interceptor{}, daq.inters...),
		predicates:           append([]predicate.DendrochronologicalAnalysis{}, daq.predicates...),
		withDendrochronology: daq.withDendrochronology.Clone(),
		// clone intermediate query.
		sql:  daq.sql.Clone(),
		path: daq.path,
	}
}

// WithDendrochronology tells the query-builder to eager-load the nodes that are connected to
// the "dendrochronology" edge. The optional arguments are used to configure the query builder of the edge.
func (daq *DendrochronologicalAnalysisQuery) WithDendrochronology(opts ...func(*DendrochronologyQuery)) *DendrochronologicalAnalysisQuery {
	query := (&DendrochronologyClient{config: daq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	daq.withDendrochronology = query
	return daq
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
//	client.DendrochronologicalAnalysis.Query().
//		GroupBy(dendrochronologicalanalysis.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (daq *DendrochronologicalAnalysisQuery) GroupBy(field string, fields ...string) *DendrochronologicalAnalysisGroupBy {
	daq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &DendrochronologicalAnalysisGroupBy{build: daq}
	grbuild.flds = &daq.ctx.Fields
	grbuild.label = dendrochronologicalanalysis.Label
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
//	client.DendrochronologicalAnalysis.Query().
//		Select(dendrochronologicalanalysis.FieldCreatedAt).
//		Scan(ctx, &v)
func (daq *DendrochronologicalAnalysisQuery) Select(fields ...string) *DendrochronologicalAnalysisSelect {
	daq.ctx.Fields = append(daq.ctx.Fields, fields...)
	sbuild := &DendrochronologicalAnalysisSelect{DendrochronologicalAnalysisQuery: daq}
	sbuild.label = dendrochronologicalanalysis.Label
	sbuild.flds, sbuild.scan = &daq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DendrochronologicalAnalysisSelect configured with the given aggregations.
func (daq *DendrochronologicalAnalysisQuery) Aggregate(fns ...AggregateFunc) *DendrochronologicalAnalysisSelect {
	return daq.Select().Aggregate(fns...)
}

func (daq *DendrochronologicalAnalysisQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range daq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, daq); err != nil {
				return err
			}
		}
	}
	for _, f := range daq.ctx.Fields {
		if !dendrochronologicalanalysis.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if daq.path != nil {
		prev, err := daq.path(ctx)
		if err != nil {
			return err
		}
		daq.sql = prev
	}
	if dendrochronologicalanalysis.Policy == nil {
		return errors.New("ent: uninitialized dendrochronologicalanalysis.Policy (forgotten import ent/runtime?)")
	}
	if err := dendrochronologicalanalysis.Policy.EvalQuery(ctx, daq); err != nil {
		return err
	}
	return nil
}

func (daq *DendrochronologicalAnalysisQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*DendrochronologicalAnalysis, error) {
	var (
		nodes       = []*DendrochronologicalAnalysis{}
		withFKs     = daq.withFKs
		_spec       = daq.querySpec()
		loadedTypes = [1]bool{
			daq.withDendrochronology != nil,
		}
	)
	if daq.withDendrochronology != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, dendrochronologicalanalysis.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*DendrochronologicalAnalysis).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &DendrochronologicalAnalysis{config: daq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(daq.modifiers) > 0 {
		_spec.Modifiers = daq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, daq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := daq.withDendrochronology; query != nil {
		if err := daq.loadDendrochronology(ctx, query, nodes, nil,
			func(n *DendrochronologicalAnalysis, e *Dendrochronology) { n.Edges.Dendrochronology = e }); err != nil {
			return nil, err
		}
	}
	for i := range daq.loadTotal {
		if err := daq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (daq *DendrochronologicalAnalysisQuery) loadDendrochronology(ctx context.Context, query *DendrochronologyQuery, nodes []*DendrochronologicalAnalysis, init func(*DendrochronologicalAnalysis), assign func(*DendrochronologicalAnalysis, *Dendrochronology)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*DendrochronologicalAnalysis)
	for i := range nodes {
		if nodes[i].dendrochronological_analysis_dendrochronology == nil {
			continue
		}
		fk := *nodes[i].dendrochronological_analysis_dendrochronology
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(dendrochronology.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "dendrochronological_analysis_dendrochronology" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (daq *DendrochronologicalAnalysisQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := daq.querySpec()
	if len(daq.modifiers) > 0 {
		_spec.Modifiers = daq.modifiers
	}
	_spec.Node.Columns = daq.ctx.Fields
	if len(daq.ctx.Fields) > 0 {
		_spec.Unique = daq.ctx.Unique != nil && *daq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, daq.driver, _spec)
}

func (daq *DendrochronologicalAnalysisQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(dendrochronologicalanalysis.Table, dendrochronologicalanalysis.Columns, sqlgraph.NewFieldSpec(dendrochronologicalanalysis.FieldID, field.TypeInt))
	_spec.From = daq.sql
	if unique := daq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if daq.path != nil {
		_spec.Unique = true
	}
	if fields := daq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dendrochronologicalanalysis.FieldID)
		for i := range fields {
			if fields[i] != dendrochronologicalanalysis.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := daq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := daq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := daq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := daq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (daq *DendrochronologicalAnalysisQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(daq.driver.Dialect())
	t1 := builder.Table(dendrochronologicalanalysis.Table)
	columns := daq.ctx.Fields
	if len(columns) == 0 {
		columns = dendrochronologicalanalysis.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if daq.sql != nil {
		selector = daq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if daq.ctx.Unique != nil && *daq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range daq.predicates {
		p(selector)
	}
	for _, p := range daq.order {
		p(selector)
	}
	if offset := daq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := daq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DendrochronologicalAnalysisGroupBy is the group-by builder for DendrochronologicalAnalysis entities.
type DendrochronologicalAnalysisGroupBy struct {
	selector
	build *DendrochronologicalAnalysisQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dagb *DendrochronologicalAnalysisGroupBy) Aggregate(fns ...AggregateFunc) *DendrochronologicalAnalysisGroupBy {
	dagb.fns = append(dagb.fns, fns...)
	return dagb
}

// Scan applies the selector query and scans the result into the given value.
func (dagb *DendrochronologicalAnalysisGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dagb.build.ctx, "GroupBy")
	if err := dagb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DendrochronologicalAnalysisQuery, *DendrochronologicalAnalysisGroupBy](ctx, dagb.build, dagb, dagb.build.inters, v)
}

func (dagb *DendrochronologicalAnalysisGroupBy) sqlScan(ctx context.Context, root *DendrochronologicalAnalysisQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(dagb.fns))
	for _, fn := range dagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*dagb.flds)+len(dagb.fns))
		for _, f := range *dagb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*dagb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dagb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// DendrochronologicalAnalysisSelect is the builder for selecting fields of DendrochronologicalAnalysis entities.
type DendrochronologicalAnalysisSelect struct {
	*DendrochronologicalAnalysisQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (das *DendrochronologicalAnalysisSelect) Aggregate(fns ...AggregateFunc) *DendrochronologicalAnalysisSelect {
	das.fns = append(das.fns, fns...)
	return das
}

// Scan applies the selector query and scans the result into the given value.
func (das *DendrochronologicalAnalysisSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, das.ctx, "Select")
	if err := das.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DendrochronologicalAnalysisQuery, *DendrochronologicalAnalysisSelect](ctx, das.DendrochronologicalAnalysisQuery, das, das.inters, v)
}

func (das *DendrochronologicalAnalysisSelect) sqlScan(ctx context.Context, root *DendrochronologicalAnalysisQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(das.fns))
	for _, fn := range das.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*das.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := das.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}