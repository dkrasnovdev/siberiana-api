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
	"github.com/dkrasnovdev/heritage-api/ent/art"
	"github.com/dkrasnovdev/heritage-api/ent/artgenre"
	"github.com/dkrasnovdev/heritage-api/ent/predicate"
)

// ArtGenreQuery is the builder for querying ArtGenre entities.
type ArtGenreQuery struct {
	config
	ctx          *QueryContext
	order        []artgenre.OrderOption
	inters       []Interceptor
	predicates   []predicate.ArtGenre
	withArt      *ArtQuery
	modifiers    []func(*sql.Selector)
	loadTotal    []func(context.Context, []*ArtGenre) error
	withNamedArt map[string]*ArtQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ArtGenreQuery builder.
func (agq *ArtGenreQuery) Where(ps ...predicate.ArtGenre) *ArtGenreQuery {
	agq.predicates = append(agq.predicates, ps...)
	return agq
}

// Limit the number of records to be returned by this query.
func (agq *ArtGenreQuery) Limit(limit int) *ArtGenreQuery {
	agq.ctx.Limit = &limit
	return agq
}

// Offset to start from.
func (agq *ArtGenreQuery) Offset(offset int) *ArtGenreQuery {
	agq.ctx.Offset = &offset
	return agq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (agq *ArtGenreQuery) Unique(unique bool) *ArtGenreQuery {
	agq.ctx.Unique = &unique
	return agq
}

// Order specifies how the records should be ordered.
func (agq *ArtGenreQuery) Order(o ...artgenre.OrderOption) *ArtGenreQuery {
	agq.order = append(agq.order, o...)
	return agq
}

// QueryArt chains the current query on the "art" edge.
func (agq *ArtGenreQuery) QueryArt() *ArtQuery {
	query := (&ArtClient{config: agq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := agq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := agq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(artgenre.Table, artgenre.FieldID, selector),
			sqlgraph.To(art.Table, art.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, artgenre.ArtTable, artgenre.ArtPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(agq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ArtGenre entity from the query.
// Returns a *NotFoundError when no ArtGenre was found.
func (agq *ArtGenreQuery) First(ctx context.Context) (*ArtGenre, error) {
	nodes, err := agq.Limit(1).All(setContextOp(ctx, agq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{artgenre.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (agq *ArtGenreQuery) FirstX(ctx context.Context) *ArtGenre {
	node, err := agq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ArtGenre ID from the query.
// Returns a *NotFoundError when no ArtGenre ID was found.
func (agq *ArtGenreQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = agq.Limit(1).IDs(setContextOp(ctx, agq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{artgenre.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (agq *ArtGenreQuery) FirstIDX(ctx context.Context) int {
	id, err := agq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ArtGenre entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ArtGenre entity is found.
// Returns a *NotFoundError when no ArtGenre entities are found.
func (agq *ArtGenreQuery) Only(ctx context.Context) (*ArtGenre, error) {
	nodes, err := agq.Limit(2).All(setContextOp(ctx, agq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{artgenre.Label}
	default:
		return nil, &NotSingularError{artgenre.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (agq *ArtGenreQuery) OnlyX(ctx context.Context) *ArtGenre {
	node, err := agq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ArtGenre ID in the query.
// Returns a *NotSingularError when more than one ArtGenre ID is found.
// Returns a *NotFoundError when no entities are found.
func (agq *ArtGenreQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = agq.Limit(2).IDs(setContextOp(ctx, agq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{artgenre.Label}
	default:
		err = &NotSingularError{artgenre.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (agq *ArtGenreQuery) OnlyIDX(ctx context.Context) int {
	id, err := agq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ArtGenres.
func (agq *ArtGenreQuery) All(ctx context.Context) ([]*ArtGenre, error) {
	ctx = setContextOp(ctx, agq.ctx, "All")
	if err := agq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ArtGenre, *ArtGenreQuery]()
	return withInterceptors[[]*ArtGenre](ctx, agq, qr, agq.inters)
}

// AllX is like All, but panics if an error occurs.
func (agq *ArtGenreQuery) AllX(ctx context.Context) []*ArtGenre {
	nodes, err := agq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ArtGenre IDs.
func (agq *ArtGenreQuery) IDs(ctx context.Context) (ids []int, err error) {
	if agq.ctx.Unique == nil && agq.path != nil {
		agq.Unique(true)
	}
	ctx = setContextOp(ctx, agq.ctx, "IDs")
	if err = agq.Select(artgenre.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (agq *ArtGenreQuery) IDsX(ctx context.Context) []int {
	ids, err := agq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (agq *ArtGenreQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, agq.ctx, "Count")
	if err := agq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, agq, querierCount[*ArtGenreQuery](), agq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (agq *ArtGenreQuery) CountX(ctx context.Context) int {
	count, err := agq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (agq *ArtGenreQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, agq.ctx, "Exist")
	switch _, err := agq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (agq *ArtGenreQuery) ExistX(ctx context.Context) bool {
	exist, err := agq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ArtGenreQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (agq *ArtGenreQuery) Clone() *ArtGenreQuery {
	if agq == nil {
		return nil
	}
	return &ArtGenreQuery{
		config:     agq.config,
		ctx:        agq.ctx.Clone(),
		order:      append([]artgenre.OrderOption{}, agq.order...),
		inters:     append([]Interceptor{}, agq.inters...),
		predicates: append([]predicate.ArtGenre{}, agq.predicates...),
		withArt:    agq.withArt.Clone(),
		// clone intermediate query.
		sql:  agq.sql.Clone(),
		path: agq.path,
	}
}

// WithArt tells the query-builder to eager-load the nodes that are connected to
// the "art" edge. The optional arguments are used to configure the query builder of the edge.
func (agq *ArtGenreQuery) WithArt(opts ...func(*ArtQuery)) *ArtGenreQuery {
	query := (&ArtClient{config: agq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	agq.withArt = query
	return agq
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
//	client.ArtGenre.Query().
//		GroupBy(artgenre.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (agq *ArtGenreQuery) GroupBy(field string, fields ...string) *ArtGenreGroupBy {
	agq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ArtGenreGroupBy{build: agq}
	grbuild.flds = &agq.ctx.Fields
	grbuild.label = artgenre.Label
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
//	client.ArtGenre.Query().
//		Select(artgenre.FieldCreatedAt).
//		Scan(ctx, &v)
func (agq *ArtGenreQuery) Select(fields ...string) *ArtGenreSelect {
	agq.ctx.Fields = append(agq.ctx.Fields, fields...)
	sbuild := &ArtGenreSelect{ArtGenreQuery: agq}
	sbuild.label = artgenre.Label
	sbuild.flds, sbuild.scan = &agq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ArtGenreSelect configured with the given aggregations.
func (agq *ArtGenreQuery) Aggregate(fns ...AggregateFunc) *ArtGenreSelect {
	return agq.Select().Aggregate(fns...)
}

func (agq *ArtGenreQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range agq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, agq); err != nil {
				return err
			}
		}
	}
	for _, f := range agq.ctx.Fields {
		if !artgenre.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if agq.path != nil {
		prev, err := agq.path(ctx)
		if err != nil {
			return err
		}
		agq.sql = prev
	}
	if artgenre.Policy == nil {
		return errors.New("ent: uninitialized artgenre.Policy (forgotten import ent/runtime?)")
	}
	if err := artgenre.Policy.EvalQuery(ctx, agq); err != nil {
		return err
	}
	return nil
}

func (agq *ArtGenreQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ArtGenre, error) {
	var (
		nodes       = []*ArtGenre{}
		_spec       = agq.querySpec()
		loadedTypes = [1]bool{
			agq.withArt != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ArtGenre).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ArtGenre{config: agq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(agq.modifiers) > 0 {
		_spec.Modifiers = agq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, agq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := agq.withArt; query != nil {
		if err := agq.loadArt(ctx, query, nodes,
			func(n *ArtGenre) { n.Edges.Art = []*Art{} },
			func(n *ArtGenre, e *Art) { n.Edges.Art = append(n.Edges.Art, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range agq.withNamedArt {
		if err := agq.loadArt(ctx, query, nodes,
			func(n *ArtGenre) { n.appendNamedArt(name) },
			func(n *ArtGenre, e *Art) { n.appendNamedArt(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range agq.loadTotal {
		if err := agq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (agq *ArtGenreQuery) loadArt(ctx context.Context, query *ArtQuery, nodes []*ArtGenre, init func(*ArtGenre), assign func(*ArtGenre, *Art)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*ArtGenre)
	nids := make(map[int]map[*ArtGenre]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(artgenre.ArtTable)
		s.Join(joinT).On(s.C(art.FieldID), joinT.C(artgenre.ArtPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(artgenre.ArtPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(artgenre.ArtPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*ArtGenre]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Art](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "art" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (agq *ArtGenreQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := agq.querySpec()
	if len(agq.modifiers) > 0 {
		_spec.Modifiers = agq.modifiers
	}
	_spec.Node.Columns = agq.ctx.Fields
	if len(agq.ctx.Fields) > 0 {
		_spec.Unique = agq.ctx.Unique != nil && *agq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, agq.driver, _spec)
}

func (agq *ArtGenreQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(artgenre.Table, artgenre.Columns, sqlgraph.NewFieldSpec(artgenre.FieldID, field.TypeInt))
	_spec.From = agq.sql
	if unique := agq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if agq.path != nil {
		_spec.Unique = true
	}
	if fields := agq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, artgenre.FieldID)
		for i := range fields {
			if fields[i] != artgenre.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := agq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := agq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := agq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := agq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (agq *ArtGenreQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(agq.driver.Dialect())
	t1 := builder.Table(artgenre.Table)
	columns := agq.ctx.Fields
	if len(columns) == 0 {
		columns = artgenre.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if agq.sql != nil {
		selector = agq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if agq.ctx.Unique != nil && *agq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range agq.predicates {
		p(selector)
	}
	for _, p := range agq.order {
		p(selector)
	}
	if offset := agq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := agq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedArt tells the query-builder to eager-load the nodes that are connected to the "art"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (agq *ArtGenreQuery) WithNamedArt(name string, opts ...func(*ArtQuery)) *ArtGenreQuery {
	query := (&ArtClient{config: agq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if agq.withNamedArt == nil {
		agq.withNamedArt = make(map[string]*ArtQuery)
	}
	agq.withNamedArt[name] = query
	return agq
}

// ArtGenreGroupBy is the group-by builder for ArtGenre entities.
type ArtGenreGroupBy struct {
	selector
	build *ArtGenreQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (aggb *ArtGenreGroupBy) Aggregate(fns ...AggregateFunc) *ArtGenreGroupBy {
	aggb.fns = append(aggb.fns, fns...)
	return aggb
}

// Scan applies the selector query and scans the result into the given value.
func (aggb *ArtGenreGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, aggb.build.ctx, "GroupBy")
	if err := aggb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ArtGenreQuery, *ArtGenreGroupBy](ctx, aggb.build, aggb, aggb.build.inters, v)
}

func (aggb *ArtGenreGroupBy) sqlScan(ctx context.Context, root *ArtGenreQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(aggb.fns))
	for _, fn := range aggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*aggb.flds)+len(aggb.fns))
		for _, f := range *aggb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*aggb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := aggb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ArtGenreSelect is the builder for selecting fields of ArtGenre entities.
type ArtGenreSelect struct {
	*ArtGenreQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ags *ArtGenreSelect) Aggregate(fns ...AggregateFunc) *ArtGenreSelect {
	ags.fns = append(ags.fns, fns...)
	return ags
}

// Scan applies the selector query and scans the result into the given value.
func (ags *ArtGenreSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ags.ctx, "Select")
	if err := ags.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ArtGenreQuery, *ArtGenreSelect](ctx, ags.ArtGenreQuery, ags, ags.inters, v)
}

func (ags *ArtGenreSelect) sqlScan(ctx context.Context, root *ArtGenreQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ags.fns))
	for _, fn := range ags.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ags.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ags.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
