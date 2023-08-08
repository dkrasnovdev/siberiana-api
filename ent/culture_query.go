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
	"github.com/dkrasnovdev/siberiana-api/ent/artifact"
	"github.com/dkrasnovdev/siberiana-api/ent/culture"
	"github.com/dkrasnovdev/siberiana-api/ent/predicate"
)

// CultureQuery is the builder for querying Culture entities.
type CultureQuery struct {
	config
	ctx                *QueryContext
	order              []culture.OrderOption
	inters             []Interceptor
	predicates         []predicate.Culture
	withArtifacts      *ArtifactQuery
	modifiers          []func(*sql.Selector)
	loadTotal          []func(context.Context, []*Culture) error
	withNamedArtifacts map[string]*ArtifactQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CultureQuery builder.
func (cq *CultureQuery) Where(ps ...predicate.Culture) *CultureQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit the number of records to be returned by this query.
func (cq *CultureQuery) Limit(limit int) *CultureQuery {
	cq.ctx.Limit = &limit
	return cq
}

// Offset to start from.
func (cq *CultureQuery) Offset(offset int) *CultureQuery {
	cq.ctx.Offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *CultureQuery) Unique(unique bool) *CultureQuery {
	cq.ctx.Unique = &unique
	return cq
}

// Order specifies how the records should be ordered.
func (cq *CultureQuery) Order(o ...culture.OrderOption) *CultureQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryArtifacts chains the current query on the "artifacts" edge.
func (cq *CultureQuery) QueryArtifacts() *ArtifactQuery {
	query := (&ArtifactClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(culture.Table, culture.FieldID, selector),
			sqlgraph.To(artifact.Table, artifact.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, culture.ArtifactsTable, culture.ArtifactsColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Culture entity from the query.
// Returns a *NotFoundError when no Culture was found.
func (cq *CultureQuery) First(ctx context.Context) (*Culture, error) {
	nodes, err := cq.Limit(1).All(setContextOp(ctx, cq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{culture.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *CultureQuery) FirstX(ctx context.Context) *Culture {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Culture ID from the query.
// Returns a *NotFoundError when no Culture ID was found.
func (cq *CultureQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(1).IDs(setContextOp(ctx, cq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{culture.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *CultureQuery) FirstIDX(ctx context.Context) int {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Culture entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Culture entity is found.
// Returns a *NotFoundError when no Culture entities are found.
func (cq *CultureQuery) Only(ctx context.Context) (*Culture, error) {
	nodes, err := cq.Limit(2).All(setContextOp(ctx, cq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{culture.Label}
	default:
		return nil, &NotSingularError{culture.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *CultureQuery) OnlyX(ctx context.Context) *Culture {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Culture ID in the query.
// Returns a *NotSingularError when more than one Culture ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *CultureQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(2).IDs(setContextOp(ctx, cq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{culture.Label}
	default:
		err = &NotSingularError{culture.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *CultureQuery) OnlyIDX(ctx context.Context) int {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Cultures.
func (cq *CultureQuery) All(ctx context.Context) ([]*Culture, error) {
	ctx = setContextOp(ctx, cq.ctx, "All")
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Culture, *CultureQuery]()
	return withInterceptors[[]*Culture](ctx, cq, qr, cq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cq *CultureQuery) AllX(ctx context.Context) []*Culture {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Culture IDs.
func (cq *CultureQuery) IDs(ctx context.Context) (ids []int, err error) {
	if cq.ctx.Unique == nil && cq.path != nil {
		cq.Unique(true)
	}
	ctx = setContextOp(ctx, cq.ctx, "IDs")
	if err = cq.Select(culture.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *CultureQuery) IDsX(ctx context.Context) []int {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *CultureQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cq.ctx, "Count")
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cq, querierCount[*CultureQuery](), cq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cq *CultureQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *CultureQuery) Exist(ctx context.Context) (bool, error) {
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
func (cq *CultureQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CultureQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *CultureQuery) Clone() *CultureQuery {
	if cq == nil {
		return nil
	}
	return &CultureQuery{
		config:        cq.config,
		ctx:           cq.ctx.Clone(),
		order:         append([]culture.OrderOption{}, cq.order...),
		inters:        append([]Interceptor{}, cq.inters...),
		predicates:    append([]predicate.Culture{}, cq.predicates...),
		withArtifacts: cq.withArtifacts.Clone(),
		// clone intermediate query.
		sql:  cq.sql.Clone(),
		path: cq.path,
	}
}

// WithArtifacts tells the query-builder to eager-load the nodes that are connected to
// the "artifacts" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CultureQuery) WithArtifacts(opts ...func(*ArtifactQuery)) *CultureQuery {
	query := (&ArtifactClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withArtifacts = query
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
//	client.Culture.Query().
//		GroupBy(culture.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cq *CultureQuery) GroupBy(field string, fields ...string) *CultureGroupBy {
	cq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CultureGroupBy{build: cq}
	grbuild.flds = &cq.ctx.Fields
	grbuild.label = culture.Label
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
//	client.Culture.Query().
//		Select(culture.FieldCreatedAt).
//		Scan(ctx, &v)
func (cq *CultureQuery) Select(fields ...string) *CultureSelect {
	cq.ctx.Fields = append(cq.ctx.Fields, fields...)
	sbuild := &CultureSelect{CultureQuery: cq}
	sbuild.label = culture.Label
	sbuild.flds, sbuild.scan = &cq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CultureSelect configured with the given aggregations.
func (cq *CultureQuery) Aggregate(fns ...AggregateFunc) *CultureSelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *CultureQuery) prepareQuery(ctx context.Context) error {
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
		if !culture.ValidColumn(f) {
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
	if culture.Policy == nil {
		return errors.New("ent: uninitialized culture.Policy (forgotten import ent/runtime?)")
	}
	if err := culture.Policy.EvalQuery(ctx, cq); err != nil {
		return err
	}
	return nil
}

func (cq *CultureQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Culture, error) {
	var (
		nodes       = []*Culture{}
		_spec       = cq.querySpec()
		loadedTypes = [1]bool{
			cq.withArtifacts != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Culture).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Culture{config: cq.config}
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
	if query := cq.withArtifacts; query != nil {
		if err := cq.loadArtifacts(ctx, query, nodes,
			func(n *Culture) { n.Edges.Artifacts = []*Artifact{} },
			func(n *Culture, e *Artifact) { n.Edges.Artifacts = append(n.Edges.Artifacts, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range cq.withNamedArtifacts {
		if err := cq.loadArtifacts(ctx, query, nodes,
			func(n *Culture) { n.appendNamedArtifacts(name) },
			func(n *Culture, e *Artifact) { n.appendNamedArtifacts(name, e) }); err != nil {
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

func (cq *CultureQuery) loadArtifacts(ctx context.Context, query *ArtifactQuery, nodes []*Culture, init func(*Culture), assign func(*Culture, *Artifact)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Culture)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(culture.ArtifactsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.culture_artifacts
		if fk == nil {
			return fmt.Errorf(`foreign-key "culture_artifacts" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "culture_artifacts" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (cq *CultureQuery) sqlCount(ctx context.Context) (int, error) {
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

func (cq *CultureQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(culture.Table, culture.Columns, sqlgraph.NewFieldSpec(culture.FieldID, field.TypeInt))
	_spec.From = cq.sql
	if unique := cq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cq.path != nil {
		_spec.Unique = true
	}
	if fields := cq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, culture.FieldID)
		for i := range fields {
			if fields[i] != culture.FieldID {
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

func (cq *CultureQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(culture.Table)
	columns := cq.ctx.Fields
	if len(columns) == 0 {
		columns = culture.Columns
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

// WithNamedArtifacts tells the query-builder to eager-load the nodes that are connected to the "artifacts"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (cq *CultureQuery) WithNamedArtifacts(name string, opts ...func(*ArtifactQuery)) *CultureQuery {
	query := (&ArtifactClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if cq.withNamedArtifacts == nil {
		cq.withNamedArtifacts = make(map[string]*ArtifactQuery)
	}
	cq.withNamedArtifacts[name] = query
	return cq
}

// CultureGroupBy is the group-by builder for Culture entities.
type CultureGroupBy struct {
	selector
	build *CultureQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *CultureGroupBy) Aggregate(fns ...AggregateFunc) *CultureGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *CultureGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cgb.build.ctx, "GroupBy")
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CultureQuery, *CultureGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (cgb *CultureGroupBy) sqlScan(ctx context.Context, root *CultureQuery, v any) error {
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

// CultureSelect is the builder for selecting fields of Culture entities.
type CultureSelect struct {
	*CultureQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *CultureSelect) Aggregate(fns ...AggregateFunc) *CultureSelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *CultureSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cs.ctx, "Select")
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CultureQuery, *CultureSelect](ctx, cs.CultureQuery, cs, cs.inters, v)
}

func (cs *CultureSelect) sqlScan(ctx context.Context, root *CultureQuery, v any) error {
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
