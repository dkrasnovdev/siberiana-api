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
	"github.com/dkrasnovdev/siberiana-api/ent/mound"
	"github.com/dkrasnovdev/siberiana-api/ent/person"
	"github.com/dkrasnovdev/siberiana-api/ent/predicate"
	"github.com/dkrasnovdev/siberiana-api/ent/visit"
)

// VisitQuery is the builder for querying Visit entities.
type VisitQuery struct {
	config
	ctx               *QueryContext
	order             []visit.OrderOption
	inters            []Interceptor
	predicates        []predicate.Visit
	withMounds        *MoundQuery
	withVisitors      *PersonQuery
	modifiers         []func(*sql.Selector)
	loadTotal         []func(context.Context, []*Visit) error
	withNamedMounds   map[string]*MoundQuery
	withNamedVisitors map[string]*PersonQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the VisitQuery builder.
func (vq *VisitQuery) Where(ps ...predicate.Visit) *VisitQuery {
	vq.predicates = append(vq.predicates, ps...)
	return vq
}

// Limit the number of records to be returned by this query.
func (vq *VisitQuery) Limit(limit int) *VisitQuery {
	vq.ctx.Limit = &limit
	return vq
}

// Offset to start from.
func (vq *VisitQuery) Offset(offset int) *VisitQuery {
	vq.ctx.Offset = &offset
	return vq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (vq *VisitQuery) Unique(unique bool) *VisitQuery {
	vq.ctx.Unique = &unique
	return vq
}

// Order specifies how the records should be ordered.
func (vq *VisitQuery) Order(o ...visit.OrderOption) *VisitQuery {
	vq.order = append(vq.order, o...)
	return vq
}

// QueryMounds chains the current query on the "mounds" edge.
func (vq *VisitQuery) QueryMounds() *MoundQuery {
	query := (&MoundClient{config: vq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(visit.Table, visit.FieldID, selector),
			sqlgraph.To(mound.Table, mound.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, visit.MoundsTable, visit.MoundsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(vq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryVisitors chains the current query on the "visitors" edge.
func (vq *VisitQuery) QueryVisitors() *PersonQuery {
	query := (&PersonClient{config: vq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(visit.Table, visit.FieldID, selector),
			sqlgraph.To(person.Table, person.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, visit.VisitorsTable, visit.VisitorsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(vq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Visit entity from the query.
// Returns a *NotFoundError when no Visit was found.
func (vq *VisitQuery) First(ctx context.Context) (*Visit, error) {
	nodes, err := vq.Limit(1).All(setContextOp(ctx, vq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{visit.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (vq *VisitQuery) FirstX(ctx context.Context) *Visit {
	node, err := vq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Visit ID from the query.
// Returns a *NotFoundError when no Visit ID was found.
func (vq *VisitQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = vq.Limit(1).IDs(setContextOp(ctx, vq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{visit.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (vq *VisitQuery) FirstIDX(ctx context.Context) int {
	id, err := vq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Visit entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Visit entity is found.
// Returns a *NotFoundError when no Visit entities are found.
func (vq *VisitQuery) Only(ctx context.Context) (*Visit, error) {
	nodes, err := vq.Limit(2).All(setContextOp(ctx, vq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{visit.Label}
	default:
		return nil, &NotSingularError{visit.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (vq *VisitQuery) OnlyX(ctx context.Context) *Visit {
	node, err := vq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Visit ID in the query.
// Returns a *NotSingularError when more than one Visit ID is found.
// Returns a *NotFoundError when no entities are found.
func (vq *VisitQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = vq.Limit(2).IDs(setContextOp(ctx, vq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{visit.Label}
	default:
		err = &NotSingularError{visit.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (vq *VisitQuery) OnlyIDX(ctx context.Context) int {
	id, err := vq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Visits.
func (vq *VisitQuery) All(ctx context.Context) ([]*Visit, error) {
	ctx = setContextOp(ctx, vq.ctx, "All")
	if err := vq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Visit, *VisitQuery]()
	return withInterceptors[[]*Visit](ctx, vq, qr, vq.inters)
}

// AllX is like All, but panics if an error occurs.
func (vq *VisitQuery) AllX(ctx context.Context) []*Visit {
	nodes, err := vq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Visit IDs.
func (vq *VisitQuery) IDs(ctx context.Context) (ids []int, err error) {
	if vq.ctx.Unique == nil && vq.path != nil {
		vq.Unique(true)
	}
	ctx = setContextOp(ctx, vq.ctx, "IDs")
	if err = vq.Select(visit.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (vq *VisitQuery) IDsX(ctx context.Context) []int {
	ids, err := vq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (vq *VisitQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, vq.ctx, "Count")
	if err := vq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, vq, querierCount[*VisitQuery](), vq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (vq *VisitQuery) CountX(ctx context.Context) int {
	count, err := vq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (vq *VisitQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, vq.ctx, "Exist")
	switch _, err := vq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (vq *VisitQuery) ExistX(ctx context.Context) bool {
	exist, err := vq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the VisitQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (vq *VisitQuery) Clone() *VisitQuery {
	if vq == nil {
		return nil
	}
	return &VisitQuery{
		config:       vq.config,
		ctx:          vq.ctx.Clone(),
		order:        append([]visit.OrderOption{}, vq.order...),
		inters:       append([]Interceptor{}, vq.inters...),
		predicates:   append([]predicate.Visit{}, vq.predicates...),
		withMounds:   vq.withMounds.Clone(),
		withVisitors: vq.withVisitors.Clone(),
		// clone intermediate query.
		sql:  vq.sql.Clone(),
		path: vq.path,
	}
}

// WithMounds tells the query-builder to eager-load the nodes that are connected to
// the "mounds" edge. The optional arguments are used to configure the query builder of the edge.
func (vq *VisitQuery) WithMounds(opts ...func(*MoundQuery)) *VisitQuery {
	query := (&MoundClient{config: vq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vq.withMounds = query
	return vq
}

// WithVisitors tells the query-builder to eager-load the nodes that are connected to
// the "visitors" edge. The optional arguments are used to configure the query builder of the edge.
func (vq *VisitQuery) WithVisitors(opts ...func(*PersonQuery)) *VisitQuery {
	query := (&PersonClient{config: vq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vq.withVisitors = query
	return vq
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
//	client.Visit.Query().
//		GroupBy(visit.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (vq *VisitQuery) GroupBy(field string, fields ...string) *VisitGroupBy {
	vq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &VisitGroupBy{build: vq}
	grbuild.flds = &vq.ctx.Fields
	grbuild.label = visit.Label
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
//	client.Visit.Query().
//		Select(visit.FieldCreatedAt).
//		Scan(ctx, &v)
func (vq *VisitQuery) Select(fields ...string) *VisitSelect {
	vq.ctx.Fields = append(vq.ctx.Fields, fields...)
	sbuild := &VisitSelect{VisitQuery: vq}
	sbuild.label = visit.Label
	sbuild.flds, sbuild.scan = &vq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a VisitSelect configured with the given aggregations.
func (vq *VisitQuery) Aggregate(fns ...AggregateFunc) *VisitSelect {
	return vq.Select().Aggregate(fns...)
}

func (vq *VisitQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range vq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, vq); err != nil {
				return err
			}
		}
	}
	for _, f := range vq.ctx.Fields {
		if !visit.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if vq.path != nil {
		prev, err := vq.path(ctx)
		if err != nil {
			return err
		}
		vq.sql = prev
	}
	if visit.Policy == nil {
		return errors.New("ent: uninitialized visit.Policy (forgotten import ent/runtime?)")
	}
	if err := visit.Policy.EvalQuery(ctx, vq); err != nil {
		return err
	}
	return nil
}

func (vq *VisitQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Visit, error) {
	var (
		nodes       = []*Visit{}
		_spec       = vq.querySpec()
		loadedTypes = [2]bool{
			vq.withMounds != nil,
			vq.withVisitors != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Visit).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Visit{config: vq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(vq.modifiers) > 0 {
		_spec.Modifiers = vq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, vq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := vq.withMounds; query != nil {
		if err := vq.loadMounds(ctx, query, nodes,
			func(n *Visit) { n.Edges.Mounds = []*Mound{} },
			func(n *Visit, e *Mound) { n.Edges.Mounds = append(n.Edges.Mounds, e) }); err != nil {
			return nil, err
		}
	}
	if query := vq.withVisitors; query != nil {
		if err := vq.loadVisitors(ctx, query, nodes,
			func(n *Visit) { n.Edges.Visitors = []*Person{} },
			func(n *Visit, e *Person) { n.Edges.Visitors = append(n.Edges.Visitors, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range vq.withNamedMounds {
		if err := vq.loadMounds(ctx, query, nodes,
			func(n *Visit) { n.appendNamedMounds(name) },
			func(n *Visit, e *Mound) { n.appendNamedMounds(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range vq.withNamedVisitors {
		if err := vq.loadVisitors(ctx, query, nodes,
			func(n *Visit) { n.appendNamedVisitors(name) },
			func(n *Visit, e *Person) { n.appendNamedVisitors(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range vq.loadTotal {
		if err := vq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (vq *VisitQuery) loadMounds(ctx context.Context, query *MoundQuery, nodes []*Visit, init func(*Visit), assign func(*Visit, *Mound)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Visit)
	nids := make(map[int]map[*Visit]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(visit.MoundsTable)
		s.Join(joinT).On(s.C(mound.FieldID), joinT.C(visit.MoundsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(visit.MoundsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(visit.MoundsPrimaryKey[0]))
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
					nids[inValue] = map[*Visit]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Mound](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "mounds" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (vq *VisitQuery) loadVisitors(ctx context.Context, query *PersonQuery, nodes []*Visit, init func(*Visit), assign func(*Visit, *Person)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Visit)
	nids := make(map[int]map[*Visit]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(visit.VisitorsTable)
		s.Join(joinT).On(s.C(person.FieldID), joinT.C(visit.VisitorsPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(visit.VisitorsPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(visit.VisitorsPrimaryKey[1]))
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
					nids[inValue] = map[*Visit]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Person](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "visitors" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (vq *VisitQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := vq.querySpec()
	if len(vq.modifiers) > 0 {
		_spec.Modifiers = vq.modifiers
	}
	_spec.Node.Columns = vq.ctx.Fields
	if len(vq.ctx.Fields) > 0 {
		_spec.Unique = vq.ctx.Unique != nil && *vq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, vq.driver, _spec)
}

func (vq *VisitQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(visit.Table, visit.Columns, sqlgraph.NewFieldSpec(visit.FieldID, field.TypeInt))
	_spec.From = vq.sql
	if unique := vq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if vq.path != nil {
		_spec.Unique = true
	}
	if fields := vq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, visit.FieldID)
		for i := range fields {
			if fields[i] != visit.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := vq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := vq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := vq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := vq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (vq *VisitQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(vq.driver.Dialect())
	t1 := builder.Table(visit.Table)
	columns := vq.ctx.Fields
	if len(columns) == 0 {
		columns = visit.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if vq.sql != nil {
		selector = vq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if vq.ctx.Unique != nil && *vq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range vq.predicates {
		p(selector)
	}
	for _, p := range vq.order {
		p(selector)
	}
	if offset := vq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := vq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedMounds tells the query-builder to eager-load the nodes that are connected to the "mounds"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (vq *VisitQuery) WithNamedMounds(name string, opts ...func(*MoundQuery)) *VisitQuery {
	query := (&MoundClient{config: vq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if vq.withNamedMounds == nil {
		vq.withNamedMounds = make(map[string]*MoundQuery)
	}
	vq.withNamedMounds[name] = query
	return vq
}

// WithNamedVisitors tells the query-builder to eager-load the nodes that are connected to the "visitors"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (vq *VisitQuery) WithNamedVisitors(name string, opts ...func(*PersonQuery)) *VisitQuery {
	query := (&PersonClient{config: vq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if vq.withNamedVisitors == nil {
		vq.withNamedVisitors = make(map[string]*PersonQuery)
	}
	vq.withNamedVisitors[name] = query
	return vq
}

// VisitGroupBy is the group-by builder for Visit entities.
type VisitGroupBy struct {
	selector
	build *VisitQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (vgb *VisitGroupBy) Aggregate(fns ...AggregateFunc) *VisitGroupBy {
	vgb.fns = append(vgb.fns, fns...)
	return vgb
}

// Scan applies the selector query and scans the result into the given value.
func (vgb *VisitGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vgb.build.ctx, "GroupBy")
	if err := vgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VisitQuery, *VisitGroupBy](ctx, vgb.build, vgb, vgb.build.inters, v)
}

func (vgb *VisitGroupBy) sqlScan(ctx context.Context, root *VisitQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(vgb.fns))
	for _, fn := range vgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*vgb.flds)+len(vgb.fns))
		for _, f := range *vgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*vgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// VisitSelect is the builder for selecting fields of Visit entities.
type VisitSelect struct {
	*VisitQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (vs *VisitSelect) Aggregate(fns ...AggregateFunc) *VisitSelect {
	vs.fns = append(vs.fns, fns...)
	return vs
}

// Scan applies the selector query and scans the result into the given value.
func (vs *VisitSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vs.ctx, "Select")
	if err := vs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VisitQuery, *VisitSelect](ctx, vs.VisitQuery, vs, vs.inters, v)
}

func (vs *VisitSelect) sqlScan(ctx context.Context, root *VisitQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(vs.fns))
	for _, fn := range vs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*vs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
