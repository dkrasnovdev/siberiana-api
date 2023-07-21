// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dkrasnovdev/heritage-api/ent/holder"
	"github.com/dkrasnovdev/heritage-api/ent/holderresponsibility"
	"github.com/dkrasnovdev/heritage-api/ent/predicate"
)

// HolderResponsibilityQuery is the builder for querying HolderResponsibility entities.
type HolderResponsibilityQuery struct {
	config
	ctx             *QueryContext
	order           []holderresponsibility.OrderOption
	inters          []Interceptor
	predicates      []predicate.HolderResponsibility
	withHolder      *HolderQuery
	modifiers       []func(*sql.Selector)
	loadTotal       []func(context.Context, []*HolderResponsibility) error
	withNamedHolder map[string]*HolderQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the HolderResponsibilityQuery builder.
func (hrq *HolderResponsibilityQuery) Where(ps ...predicate.HolderResponsibility) *HolderResponsibilityQuery {
	hrq.predicates = append(hrq.predicates, ps...)
	return hrq
}

// Limit the number of records to be returned by this query.
func (hrq *HolderResponsibilityQuery) Limit(limit int) *HolderResponsibilityQuery {
	hrq.ctx.Limit = &limit
	return hrq
}

// Offset to start from.
func (hrq *HolderResponsibilityQuery) Offset(offset int) *HolderResponsibilityQuery {
	hrq.ctx.Offset = &offset
	return hrq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (hrq *HolderResponsibilityQuery) Unique(unique bool) *HolderResponsibilityQuery {
	hrq.ctx.Unique = &unique
	return hrq
}

// Order specifies how the records should be ordered.
func (hrq *HolderResponsibilityQuery) Order(o ...holderresponsibility.OrderOption) *HolderResponsibilityQuery {
	hrq.order = append(hrq.order, o...)
	return hrq
}

// QueryHolder chains the current query on the "holder" edge.
func (hrq *HolderResponsibilityQuery) QueryHolder() *HolderQuery {
	query := (&HolderClient{config: hrq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := hrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := hrq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(holderresponsibility.Table, holderresponsibility.FieldID, selector),
			sqlgraph.To(holder.Table, holder.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, holderresponsibility.HolderTable, holderresponsibility.HolderPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(hrq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first HolderResponsibility entity from the query.
// Returns a *NotFoundError when no HolderResponsibility was found.
func (hrq *HolderResponsibilityQuery) First(ctx context.Context) (*HolderResponsibility, error) {
	nodes, err := hrq.Limit(1).All(setContextOp(ctx, hrq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{holderresponsibility.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (hrq *HolderResponsibilityQuery) FirstX(ctx context.Context) *HolderResponsibility {
	node, err := hrq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first HolderResponsibility ID from the query.
// Returns a *NotFoundError when no HolderResponsibility ID was found.
func (hrq *HolderResponsibilityQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = hrq.Limit(1).IDs(setContextOp(ctx, hrq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{holderresponsibility.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (hrq *HolderResponsibilityQuery) FirstIDX(ctx context.Context) int {
	id, err := hrq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single HolderResponsibility entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one HolderResponsibility entity is found.
// Returns a *NotFoundError when no HolderResponsibility entities are found.
func (hrq *HolderResponsibilityQuery) Only(ctx context.Context) (*HolderResponsibility, error) {
	nodes, err := hrq.Limit(2).All(setContextOp(ctx, hrq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{holderresponsibility.Label}
	default:
		return nil, &NotSingularError{holderresponsibility.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (hrq *HolderResponsibilityQuery) OnlyX(ctx context.Context) *HolderResponsibility {
	node, err := hrq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only HolderResponsibility ID in the query.
// Returns a *NotSingularError when more than one HolderResponsibility ID is found.
// Returns a *NotFoundError when no entities are found.
func (hrq *HolderResponsibilityQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = hrq.Limit(2).IDs(setContextOp(ctx, hrq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{holderresponsibility.Label}
	default:
		err = &NotSingularError{holderresponsibility.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (hrq *HolderResponsibilityQuery) OnlyIDX(ctx context.Context) int {
	id, err := hrq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of HolderResponsibilities.
func (hrq *HolderResponsibilityQuery) All(ctx context.Context) ([]*HolderResponsibility, error) {
	ctx = setContextOp(ctx, hrq.ctx, "All")
	if err := hrq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*HolderResponsibility, *HolderResponsibilityQuery]()
	return withInterceptors[[]*HolderResponsibility](ctx, hrq, qr, hrq.inters)
}

// AllX is like All, but panics if an error occurs.
func (hrq *HolderResponsibilityQuery) AllX(ctx context.Context) []*HolderResponsibility {
	nodes, err := hrq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of HolderResponsibility IDs.
func (hrq *HolderResponsibilityQuery) IDs(ctx context.Context) (ids []int, err error) {
	if hrq.ctx.Unique == nil && hrq.path != nil {
		hrq.Unique(true)
	}
	ctx = setContextOp(ctx, hrq.ctx, "IDs")
	if err = hrq.Select(holderresponsibility.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (hrq *HolderResponsibilityQuery) IDsX(ctx context.Context) []int {
	ids, err := hrq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (hrq *HolderResponsibilityQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, hrq.ctx, "Count")
	if err := hrq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, hrq, querierCount[*HolderResponsibilityQuery](), hrq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (hrq *HolderResponsibilityQuery) CountX(ctx context.Context) int {
	count, err := hrq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (hrq *HolderResponsibilityQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, hrq.ctx, "Exist")
	switch _, err := hrq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (hrq *HolderResponsibilityQuery) ExistX(ctx context.Context) bool {
	exist, err := hrq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the HolderResponsibilityQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (hrq *HolderResponsibilityQuery) Clone() *HolderResponsibilityQuery {
	if hrq == nil {
		return nil
	}
	return &HolderResponsibilityQuery{
		config:     hrq.config,
		ctx:        hrq.ctx.Clone(),
		order:      append([]holderresponsibility.OrderOption{}, hrq.order...),
		inters:     append([]Interceptor{}, hrq.inters...),
		predicates: append([]predicate.HolderResponsibility{}, hrq.predicates...),
		withHolder: hrq.withHolder.Clone(),
		// clone intermediate query.
		sql:  hrq.sql.Clone(),
		path: hrq.path,
	}
}

// WithHolder tells the query-builder to eager-load the nodes that are connected to
// the "holder" edge. The optional arguments are used to configure the query builder of the edge.
func (hrq *HolderResponsibilityQuery) WithHolder(opts ...func(*HolderQuery)) *HolderResponsibilityQuery {
	query := (&HolderClient{config: hrq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	hrq.withHolder = query
	return hrq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (hrq *HolderResponsibilityQuery) GroupBy(field string, fields ...string) *HolderResponsibilityGroupBy {
	hrq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &HolderResponsibilityGroupBy{build: hrq}
	grbuild.flds = &hrq.ctx.Fields
	grbuild.label = holderresponsibility.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (hrq *HolderResponsibilityQuery) Select(fields ...string) *HolderResponsibilitySelect {
	hrq.ctx.Fields = append(hrq.ctx.Fields, fields...)
	sbuild := &HolderResponsibilitySelect{HolderResponsibilityQuery: hrq}
	sbuild.label = holderresponsibility.Label
	sbuild.flds, sbuild.scan = &hrq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a HolderResponsibilitySelect configured with the given aggregations.
func (hrq *HolderResponsibilityQuery) Aggregate(fns ...AggregateFunc) *HolderResponsibilitySelect {
	return hrq.Select().Aggregate(fns...)
}

func (hrq *HolderResponsibilityQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range hrq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, hrq); err != nil {
				return err
			}
		}
	}
	for _, f := range hrq.ctx.Fields {
		if !holderresponsibility.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if hrq.path != nil {
		prev, err := hrq.path(ctx)
		if err != nil {
			return err
		}
		hrq.sql = prev
	}
	return nil
}

func (hrq *HolderResponsibilityQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*HolderResponsibility, error) {
	var (
		nodes       = []*HolderResponsibility{}
		_spec       = hrq.querySpec()
		loadedTypes = [1]bool{
			hrq.withHolder != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*HolderResponsibility).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &HolderResponsibility{config: hrq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(hrq.modifiers) > 0 {
		_spec.Modifiers = hrq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, hrq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := hrq.withHolder; query != nil {
		if err := hrq.loadHolder(ctx, query, nodes,
			func(n *HolderResponsibility) { n.Edges.Holder = []*Holder{} },
			func(n *HolderResponsibility, e *Holder) { n.Edges.Holder = append(n.Edges.Holder, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range hrq.withNamedHolder {
		if err := hrq.loadHolder(ctx, query, nodes,
			func(n *HolderResponsibility) { n.appendNamedHolder(name) },
			func(n *HolderResponsibility, e *Holder) { n.appendNamedHolder(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range hrq.loadTotal {
		if err := hrq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (hrq *HolderResponsibilityQuery) loadHolder(ctx context.Context, query *HolderQuery, nodes []*HolderResponsibility, init func(*HolderResponsibility), assign func(*HolderResponsibility, *Holder)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*HolderResponsibility)
	nids := make(map[int]map[*HolderResponsibility]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(holderresponsibility.HolderTable)
		s.Join(joinT).On(s.C(holder.FieldID), joinT.C(holderresponsibility.HolderPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(holderresponsibility.HolderPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(holderresponsibility.HolderPrimaryKey[1]))
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
					nids[inValue] = map[*HolderResponsibility]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Holder](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "holder" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (hrq *HolderResponsibilityQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := hrq.querySpec()
	if len(hrq.modifiers) > 0 {
		_spec.Modifiers = hrq.modifiers
	}
	_spec.Node.Columns = hrq.ctx.Fields
	if len(hrq.ctx.Fields) > 0 {
		_spec.Unique = hrq.ctx.Unique != nil && *hrq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, hrq.driver, _spec)
}

func (hrq *HolderResponsibilityQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(holderresponsibility.Table, holderresponsibility.Columns, sqlgraph.NewFieldSpec(holderresponsibility.FieldID, field.TypeInt))
	_spec.From = hrq.sql
	if unique := hrq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if hrq.path != nil {
		_spec.Unique = true
	}
	if fields := hrq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, holderresponsibility.FieldID)
		for i := range fields {
			if fields[i] != holderresponsibility.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := hrq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := hrq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := hrq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := hrq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (hrq *HolderResponsibilityQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(hrq.driver.Dialect())
	t1 := builder.Table(holderresponsibility.Table)
	columns := hrq.ctx.Fields
	if len(columns) == 0 {
		columns = holderresponsibility.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if hrq.sql != nil {
		selector = hrq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if hrq.ctx.Unique != nil && *hrq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range hrq.predicates {
		p(selector)
	}
	for _, p := range hrq.order {
		p(selector)
	}
	if offset := hrq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := hrq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedHolder tells the query-builder to eager-load the nodes that are connected to the "holder"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (hrq *HolderResponsibilityQuery) WithNamedHolder(name string, opts ...func(*HolderQuery)) *HolderResponsibilityQuery {
	query := (&HolderClient{config: hrq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if hrq.withNamedHolder == nil {
		hrq.withNamedHolder = make(map[string]*HolderQuery)
	}
	hrq.withNamedHolder[name] = query
	return hrq
}

// HolderResponsibilityGroupBy is the group-by builder for HolderResponsibility entities.
type HolderResponsibilityGroupBy struct {
	selector
	build *HolderResponsibilityQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (hrgb *HolderResponsibilityGroupBy) Aggregate(fns ...AggregateFunc) *HolderResponsibilityGroupBy {
	hrgb.fns = append(hrgb.fns, fns...)
	return hrgb
}

// Scan applies the selector query and scans the result into the given value.
func (hrgb *HolderResponsibilityGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, hrgb.build.ctx, "GroupBy")
	if err := hrgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*HolderResponsibilityQuery, *HolderResponsibilityGroupBy](ctx, hrgb.build, hrgb, hrgb.build.inters, v)
}

func (hrgb *HolderResponsibilityGroupBy) sqlScan(ctx context.Context, root *HolderResponsibilityQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(hrgb.fns))
	for _, fn := range hrgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*hrgb.flds)+len(hrgb.fns))
		for _, f := range *hrgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*hrgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := hrgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// HolderResponsibilitySelect is the builder for selecting fields of HolderResponsibility entities.
type HolderResponsibilitySelect struct {
	*HolderResponsibilityQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (hrs *HolderResponsibilitySelect) Aggregate(fns ...AggregateFunc) *HolderResponsibilitySelect {
	hrs.fns = append(hrs.fns, fns...)
	return hrs
}

// Scan applies the selector query and scans the result into the given value.
func (hrs *HolderResponsibilitySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, hrs.ctx, "Select")
	if err := hrs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*HolderResponsibilityQuery, *HolderResponsibilitySelect](ctx, hrs.HolderResponsibilityQuery, hrs, hrs.inters, v)
}

func (hrs *HolderResponsibilitySelect) sqlScan(ctx context.Context, root *HolderResponsibilityQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(hrs.fns))
	for _, fn := range hrs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*hrs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := hrs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
