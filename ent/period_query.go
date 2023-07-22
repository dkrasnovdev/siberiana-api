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
	"github.com/dkrasnovdev/heritage-api/ent/artifact"
	"github.com/dkrasnovdev/heritage-api/ent/period"
	"github.com/dkrasnovdev/heritage-api/ent/predicate"
)

// PeriodQuery is the builder for querying Period entities.
type PeriodQuery struct {
	config
	ctx                *QueryContext
	order              []period.OrderOption
	inters             []Interceptor
	predicates         []predicate.Period
	withArtifacts      *ArtifactQuery
	modifiers          []func(*sql.Selector)
	loadTotal          []func(context.Context, []*Period) error
	withNamedArtifacts map[string]*ArtifactQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PeriodQuery builder.
func (pq *PeriodQuery) Where(ps ...predicate.Period) *PeriodQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit the number of records to be returned by this query.
func (pq *PeriodQuery) Limit(limit int) *PeriodQuery {
	pq.ctx.Limit = &limit
	return pq
}

// Offset to start from.
func (pq *PeriodQuery) Offset(offset int) *PeriodQuery {
	pq.ctx.Offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *PeriodQuery) Unique(unique bool) *PeriodQuery {
	pq.ctx.Unique = &unique
	return pq
}

// Order specifies how the records should be ordered.
func (pq *PeriodQuery) Order(o ...period.OrderOption) *PeriodQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryArtifacts chains the current query on the "artifacts" edge.
func (pq *PeriodQuery) QueryArtifacts() *ArtifactQuery {
	query := (&ArtifactClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(period.Table, period.FieldID, selector),
			sqlgraph.To(artifact.Table, artifact.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, period.ArtifactsTable, period.ArtifactsColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Period entity from the query.
// Returns a *NotFoundError when no Period was found.
func (pq *PeriodQuery) First(ctx context.Context) (*Period, error) {
	nodes, err := pq.Limit(1).All(setContextOp(ctx, pq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{period.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PeriodQuery) FirstX(ctx context.Context) *Period {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Period ID from the query.
// Returns a *NotFoundError when no Period ID was found.
func (pq *PeriodQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(1).IDs(setContextOp(ctx, pq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{period.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *PeriodQuery) FirstIDX(ctx context.Context) int {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Period entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Period entity is found.
// Returns a *NotFoundError when no Period entities are found.
func (pq *PeriodQuery) Only(ctx context.Context) (*Period, error) {
	nodes, err := pq.Limit(2).All(setContextOp(ctx, pq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{period.Label}
	default:
		return nil, &NotSingularError{period.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PeriodQuery) OnlyX(ctx context.Context) *Period {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Period ID in the query.
// Returns a *NotSingularError when more than one Period ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *PeriodQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(2).IDs(setContextOp(ctx, pq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{period.Label}
	default:
		err = &NotSingularError{period.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *PeriodQuery) OnlyIDX(ctx context.Context) int {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Periods.
func (pq *PeriodQuery) All(ctx context.Context) ([]*Period, error) {
	ctx = setContextOp(ctx, pq.ctx, "All")
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Period, *PeriodQuery]()
	return withInterceptors[[]*Period](ctx, pq, qr, pq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pq *PeriodQuery) AllX(ctx context.Context) []*Period {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Period IDs.
func (pq *PeriodQuery) IDs(ctx context.Context) (ids []int, err error) {
	if pq.ctx.Unique == nil && pq.path != nil {
		pq.Unique(true)
	}
	ctx = setContextOp(ctx, pq.ctx, "IDs")
	if err = pq.Select(period.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PeriodQuery) IDsX(ctx context.Context) []int {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PeriodQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pq.ctx, "Count")
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pq, querierCount[*PeriodQuery](), pq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PeriodQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PeriodQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pq.ctx, "Exist")
	switch _, err := pq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *PeriodQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PeriodQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PeriodQuery) Clone() *PeriodQuery {
	if pq == nil {
		return nil
	}
	return &PeriodQuery{
		config:        pq.config,
		ctx:           pq.ctx.Clone(),
		order:         append([]period.OrderOption{}, pq.order...),
		inters:        append([]Interceptor{}, pq.inters...),
		predicates:    append([]predicate.Period{}, pq.predicates...),
		withArtifacts: pq.withArtifacts.Clone(),
		// clone intermediate query.
		sql:  pq.sql.Clone(),
		path: pq.path,
	}
}

// WithArtifacts tells the query-builder to eager-load the nodes that are connected to
// the "artifacts" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PeriodQuery) WithArtifacts(opts ...func(*ArtifactQuery)) *PeriodQuery {
	query := (&ArtifactClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withArtifacts = query
	return pq
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
//	client.Period.Query().
//		GroupBy(period.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pq *PeriodQuery) GroupBy(field string, fields ...string) *PeriodGroupBy {
	pq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PeriodGroupBy{build: pq}
	grbuild.flds = &pq.ctx.Fields
	grbuild.label = period.Label
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
//	client.Period.Query().
//		Select(period.FieldCreatedAt).
//		Scan(ctx, &v)
func (pq *PeriodQuery) Select(fields ...string) *PeriodSelect {
	pq.ctx.Fields = append(pq.ctx.Fields, fields...)
	sbuild := &PeriodSelect{PeriodQuery: pq}
	sbuild.label = period.Label
	sbuild.flds, sbuild.scan = &pq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PeriodSelect configured with the given aggregations.
func (pq *PeriodQuery) Aggregate(fns ...AggregateFunc) *PeriodSelect {
	return pq.Select().Aggregate(fns...)
}

func (pq *PeriodQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pq); err != nil {
				return err
			}
		}
	}
	for _, f := range pq.ctx.Fields {
		if !period.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	if period.Policy == nil {
		return errors.New("ent: uninitialized period.Policy (forgotten import ent/runtime?)")
	}
	if err := period.Policy.EvalQuery(ctx, pq); err != nil {
		return err
	}
	return nil
}

func (pq *PeriodQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Period, error) {
	var (
		nodes       = []*Period{}
		_spec       = pq.querySpec()
		loadedTypes = [1]bool{
			pq.withArtifacts != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Period).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Period{config: pq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(pq.modifiers) > 0 {
		_spec.Modifiers = pq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pq.withArtifacts; query != nil {
		if err := pq.loadArtifacts(ctx, query, nodes,
			func(n *Period) { n.Edges.Artifacts = []*Artifact{} },
			func(n *Period, e *Artifact) { n.Edges.Artifacts = append(n.Edges.Artifacts, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range pq.withNamedArtifacts {
		if err := pq.loadArtifacts(ctx, query, nodes,
			func(n *Period) { n.appendNamedArtifacts(name) },
			func(n *Period, e *Artifact) { n.appendNamedArtifacts(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range pq.loadTotal {
		if err := pq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pq *PeriodQuery) loadArtifacts(ctx context.Context, query *ArtifactQuery, nodes []*Period, init func(*Period), assign func(*Period, *Artifact)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Period)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(period.ArtifactsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.period_artifacts
		if fk == nil {
			return fmt.Errorf(`foreign-key "period_artifacts" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "period_artifacts" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (pq *PeriodQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	if len(pq.modifiers) > 0 {
		_spec.Modifiers = pq.modifiers
	}
	_spec.Node.Columns = pq.ctx.Fields
	if len(pq.ctx.Fields) > 0 {
		_spec.Unique = pq.ctx.Unique != nil && *pq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *PeriodQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(period.Table, period.Columns, sqlgraph.NewFieldSpec(period.FieldID, field.TypeInt))
	_spec.From = pq.sql
	if unique := pq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pq.path != nil {
		_spec.Unique = true
	}
	if fields := pq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, period.FieldID)
		for i := range fields {
			if fields[i] != period.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *PeriodQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(period.Table)
	columns := pq.ctx.Fields
	if len(columns) == 0 {
		columns = period.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.ctx.Unique != nil && *pq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedArtifacts tells the query-builder to eager-load the nodes that are connected to the "artifacts"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (pq *PeriodQuery) WithNamedArtifacts(name string, opts ...func(*ArtifactQuery)) *PeriodQuery {
	query := (&ArtifactClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if pq.withNamedArtifacts == nil {
		pq.withNamedArtifacts = make(map[string]*ArtifactQuery)
	}
	pq.withNamedArtifacts[name] = query
	return pq
}

// PeriodGroupBy is the group-by builder for Period entities.
type PeriodGroupBy struct {
	selector
	build *PeriodQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PeriodGroupBy) Aggregate(fns ...AggregateFunc) *PeriodGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the selector query and scans the result into the given value.
func (pgb *PeriodGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pgb.build.ctx, "GroupBy")
	if err := pgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PeriodQuery, *PeriodGroupBy](ctx, pgb.build, pgb, pgb.build.inters, v)
}

func (pgb *PeriodGroupBy) sqlScan(ctx context.Context, root *PeriodQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pgb.flds)+len(pgb.fns))
		for _, f := range *pgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PeriodSelect is the builder for selecting fields of Period entities.
type PeriodSelect struct {
	*PeriodQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ps *PeriodSelect) Aggregate(fns ...AggregateFunc) *PeriodSelect {
	ps.fns = append(ps.fns, fns...)
	return ps
}

// Scan applies the selector query and scans the result into the given value.
func (ps *PeriodSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ps.ctx, "Select")
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PeriodQuery, *PeriodSelect](ctx, ps.PeriodQuery, ps, ps.inters, v)
}

func (ps *PeriodSelect) sqlScan(ctx context.Context, root *PeriodQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ps.fns))
	for _, fn := range ps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
