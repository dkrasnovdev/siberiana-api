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
	"github.com/dkrasnovdev/heritage-api/ent/predicate"
	"github.com/dkrasnovdev/heritage-api/ent/protectedarea"
	"github.com/dkrasnovdev/heritage-api/ent/protectedareacategory"
	"github.com/dkrasnovdev/heritage-api/ent/protectedareapicture"
)

// ProtectedAreaQuery is the builder for querying ProtectedArea entities.
type ProtectedAreaQuery struct {
	config
	ctx                            *QueryContext
	order                          []protectedarea.OrderOption
	inters                         []Interceptor
	predicates                     []predicate.ProtectedArea
	withProtectedAreaPictures      *ProtectedAreaPictureQuery
	withProtectedAreaCategory      *ProtectedAreaCategoryQuery
	withFKs                        bool
	modifiers                      []func(*sql.Selector)
	loadTotal                      []func(context.Context, []*ProtectedArea) error
	withNamedProtectedAreaPictures map[string]*ProtectedAreaPictureQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProtectedAreaQuery builder.
func (paq *ProtectedAreaQuery) Where(ps ...predicate.ProtectedArea) *ProtectedAreaQuery {
	paq.predicates = append(paq.predicates, ps...)
	return paq
}

// Limit the number of records to be returned by this query.
func (paq *ProtectedAreaQuery) Limit(limit int) *ProtectedAreaQuery {
	paq.ctx.Limit = &limit
	return paq
}

// Offset to start from.
func (paq *ProtectedAreaQuery) Offset(offset int) *ProtectedAreaQuery {
	paq.ctx.Offset = &offset
	return paq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (paq *ProtectedAreaQuery) Unique(unique bool) *ProtectedAreaQuery {
	paq.ctx.Unique = &unique
	return paq
}

// Order specifies how the records should be ordered.
func (paq *ProtectedAreaQuery) Order(o ...protectedarea.OrderOption) *ProtectedAreaQuery {
	paq.order = append(paq.order, o...)
	return paq
}

// QueryProtectedAreaPictures chains the current query on the "protected_area_pictures" edge.
func (paq *ProtectedAreaQuery) QueryProtectedAreaPictures() *ProtectedAreaPictureQuery {
	query := (&ProtectedAreaPictureClient{config: paq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := paq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := paq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(protectedarea.Table, protectedarea.FieldID, selector),
			sqlgraph.To(protectedareapicture.Table, protectedareapicture.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, protectedarea.ProtectedAreaPicturesTable, protectedarea.ProtectedAreaPicturesColumn),
		)
		fromU = sqlgraph.SetNeighbors(paq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryProtectedAreaCategory chains the current query on the "protected_area_category" edge.
func (paq *ProtectedAreaQuery) QueryProtectedAreaCategory() *ProtectedAreaCategoryQuery {
	query := (&ProtectedAreaCategoryClient{config: paq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := paq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := paq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(protectedarea.Table, protectedarea.FieldID, selector),
			sqlgraph.To(protectedareacategory.Table, protectedareacategory.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, protectedarea.ProtectedAreaCategoryTable, protectedarea.ProtectedAreaCategoryColumn),
		)
		fromU = sqlgraph.SetNeighbors(paq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ProtectedArea entity from the query.
// Returns a *NotFoundError when no ProtectedArea was found.
func (paq *ProtectedAreaQuery) First(ctx context.Context) (*ProtectedArea, error) {
	nodes, err := paq.Limit(1).All(setContextOp(ctx, paq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{protectedarea.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (paq *ProtectedAreaQuery) FirstX(ctx context.Context) *ProtectedArea {
	node, err := paq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ProtectedArea ID from the query.
// Returns a *NotFoundError when no ProtectedArea ID was found.
func (paq *ProtectedAreaQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = paq.Limit(1).IDs(setContextOp(ctx, paq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{protectedarea.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (paq *ProtectedAreaQuery) FirstIDX(ctx context.Context) int {
	id, err := paq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ProtectedArea entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ProtectedArea entity is found.
// Returns a *NotFoundError when no ProtectedArea entities are found.
func (paq *ProtectedAreaQuery) Only(ctx context.Context) (*ProtectedArea, error) {
	nodes, err := paq.Limit(2).All(setContextOp(ctx, paq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{protectedarea.Label}
	default:
		return nil, &NotSingularError{protectedarea.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (paq *ProtectedAreaQuery) OnlyX(ctx context.Context) *ProtectedArea {
	node, err := paq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ProtectedArea ID in the query.
// Returns a *NotSingularError when more than one ProtectedArea ID is found.
// Returns a *NotFoundError when no entities are found.
func (paq *ProtectedAreaQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = paq.Limit(2).IDs(setContextOp(ctx, paq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{protectedarea.Label}
	default:
		err = &NotSingularError{protectedarea.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (paq *ProtectedAreaQuery) OnlyIDX(ctx context.Context) int {
	id, err := paq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ProtectedAreas.
func (paq *ProtectedAreaQuery) All(ctx context.Context) ([]*ProtectedArea, error) {
	ctx = setContextOp(ctx, paq.ctx, "All")
	if err := paq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ProtectedArea, *ProtectedAreaQuery]()
	return withInterceptors[[]*ProtectedArea](ctx, paq, qr, paq.inters)
}

// AllX is like All, but panics if an error occurs.
func (paq *ProtectedAreaQuery) AllX(ctx context.Context) []*ProtectedArea {
	nodes, err := paq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ProtectedArea IDs.
func (paq *ProtectedAreaQuery) IDs(ctx context.Context) (ids []int, err error) {
	if paq.ctx.Unique == nil && paq.path != nil {
		paq.Unique(true)
	}
	ctx = setContextOp(ctx, paq.ctx, "IDs")
	if err = paq.Select(protectedarea.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (paq *ProtectedAreaQuery) IDsX(ctx context.Context) []int {
	ids, err := paq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (paq *ProtectedAreaQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, paq.ctx, "Count")
	if err := paq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, paq, querierCount[*ProtectedAreaQuery](), paq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (paq *ProtectedAreaQuery) CountX(ctx context.Context) int {
	count, err := paq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (paq *ProtectedAreaQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, paq.ctx, "Exist")
	switch _, err := paq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (paq *ProtectedAreaQuery) ExistX(ctx context.Context) bool {
	exist, err := paq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProtectedAreaQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (paq *ProtectedAreaQuery) Clone() *ProtectedAreaQuery {
	if paq == nil {
		return nil
	}
	return &ProtectedAreaQuery{
		config:                    paq.config,
		ctx:                       paq.ctx.Clone(),
		order:                     append([]protectedarea.OrderOption{}, paq.order...),
		inters:                    append([]Interceptor{}, paq.inters...),
		predicates:                append([]predicate.ProtectedArea{}, paq.predicates...),
		withProtectedAreaPictures: paq.withProtectedAreaPictures.Clone(),
		withProtectedAreaCategory: paq.withProtectedAreaCategory.Clone(),
		// clone intermediate query.
		sql:  paq.sql.Clone(),
		path: paq.path,
	}
}

// WithProtectedAreaPictures tells the query-builder to eager-load the nodes that are connected to
// the "protected_area_pictures" edge. The optional arguments are used to configure the query builder of the edge.
func (paq *ProtectedAreaQuery) WithProtectedAreaPictures(opts ...func(*ProtectedAreaPictureQuery)) *ProtectedAreaQuery {
	query := (&ProtectedAreaPictureClient{config: paq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	paq.withProtectedAreaPictures = query
	return paq
}

// WithProtectedAreaCategory tells the query-builder to eager-load the nodes that are connected to
// the "protected_area_category" edge. The optional arguments are used to configure the query builder of the edge.
func (paq *ProtectedAreaQuery) WithProtectedAreaCategory(opts ...func(*ProtectedAreaCategoryQuery)) *ProtectedAreaQuery {
	query := (&ProtectedAreaCategoryClient{config: paq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	paq.withProtectedAreaCategory = query
	return paq
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
//	client.ProtectedArea.Query().
//		GroupBy(protectedarea.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (paq *ProtectedAreaQuery) GroupBy(field string, fields ...string) *ProtectedAreaGroupBy {
	paq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ProtectedAreaGroupBy{build: paq}
	grbuild.flds = &paq.ctx.Fields
	grbuild.label = protectedarea.Label
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
//	client.ProtectedArea.Query().
//		Select(protectedarea.FieldCreatedAt).
//		Scan(ctx, &v)
func (paq *ProtectedAreaQuery) Select(fields ...string) *ProtectedAreaSelect {
	paq.ctx.Fields = append(paq.ctx.Fields, fields...)
	sbuild := &ProtectedAreaSelect{ProtectedAreaQuery: paq}
	sbuild.label = protectedarea.Label
	sbuild.flds, sbuild.scan = &paq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ProtectedAreaSelect configured with the given aggregations.
func (paq *ProtectedAreaQuery) Aggregate(fns ...AggregateFunc) *ProtectedAreaSelect {
	return paq.Select().Aggregate(fns...)
}

func (paq *ProtectedAreaQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range paq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, paq); err != nil {
				return err
			}
		}
	}
	for _, f := range paq.ctx.Fields {
		if !protectedarea.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if paq.path != nil {
		prev, err := paq.path(ctx)
		if err != nil {
			return err
		}
		paq.sql = prev
	}
	if protectedarea.Policy == nil {
		return errors.New("ent: uninitialized protectedarea.Policy (forgotten import ent/runtime?)")
	}
	if err := protectedarea.Policy.EvalQuery(ctx, paq); err != nil {
		return err
	}
	return nil
}

func (paq *ProtectedAreaQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ProtectedArea, error) {
	var (
		nodes       = []*ProtectedArea{}
		withFKs     = paq.withFKs
		_spec       = paq.querySpec()
		loadedTypes = [2]bool{
			paq.withProtectedAreaPictures != nil,
			paq.withProtectedAreaCategory != nil,
		}
	)
	if paq.withProtectedAreaCategory != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, protectedarea.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ProtectedArea).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ProtectedArea{config: paq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(paq.modifiers) > 0 {
		_spec.Modifiers = paq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, paq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := paq.withProtectedAreaPictures; query != nil {
		if err := paq.loadProtectedAreaPictures(ctx, query, nodes,
			func(n *ProtectedArea) { n.Edges.ProtectedAreaPictures = []*ProtectedAreaPicture{} },
			func(n *ProtectedArea, e *ProtectedAreaPicture) {
				n.Edges.ProtectedAreaPictures = append(n.Edges.ProtectedAreaPictures, e)
			}); err != nil {
			return nil, err
		}
	}
	if query := paq.withProtectedAreaCategory; query != nil {
		if err := paq.loadProtectedAreaCategory(ctx, query, nodes, nil,
			func(n *ProtectedArea, e *ProtectedAreaCategory) { n.Edges.ProtectedAreaCategory = e }); err != nil {
			return nil, err
		}
	}
	for name, query := range paq.withNamedProtectedAreaPictures {
		if err := paq.loadProtectedAreaPictures(ctx, query, nodes,
			func(n *ProtectedArea) { n.appendNamedProtectedAreaPictures(name) },
			func(n *ProtectedArea, e *ProtectedAreaPicture) { n.appendNamedProtectedAreaPictures(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range paq.loadTotal {
		if err := paq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (paq *ProtectedAreaQuery) loadProtectedAreaPictures(ctx context.Context, query *ProtectedAreaPictureQuery, nodes []*ProtectedArea, init func(*ProtectedArea), assign func(*ProtectedArea, *ProtectedAreaPicture)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*ProtectedArea)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.ProtectedAreaPicture(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(protectedarea.ProtectedAreaPicturesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.protected_area_protected_area_pictures
		if fk == nil {
			return fmt.Errorf(`foreign-key "protected_area_protected_area_pictures" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "protected_area_protected_area_pictures" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (paq *ProtectedAreaQuery) loadProtectedAreaCategory(ctx context.Context, query *ProtectedAreaCategoryQuery, nodes []*ProtectedArea, init func(*ProtectedArea), assign func(*ProtectedArea, *ProtectedAreaCategory)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*ProtectedArea)
	for i := range nodes {
		if nodes[i].protected_area_category_protected_areas == nil {
			continue
		}
		fk := *nodes[i].protected_area_category_protected_areas
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(protectedareacategory.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "protected_area_category_protected_areas" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (paq *ProtectedAreaQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := paq.querySpec()
	if len(paq.modifiers) > 0 {
		_spec.Modifiers = paq.modifiers
	}
	_spec.Node.Columns = paq.ctx.Fields
	if len(paq.ctx.Fields) > 0 {
		_spec.Unique = paq.ctx.Unique != nil && *paq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, paq.driver, _spec)
}

func (paq *ProtectedAreaQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(protectedarea.Table, protectedarea.Columns, sqlgraph.NewFieldSpec(protectedarea.FieldID, field.TypeInt))
	_spec.From = paq.sql
	if unique := paq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if paq.path != nil {
		_spec.Unique = true
	}
	if fields := paq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, protectedarea.FieldID)
		for i := range fields {
			if fields[i] != protectedarea.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := paq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := paq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := paq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := paq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (paq *ProtectedAreaQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(paq.driver.Dialect())
	t1 := builder.Table(protectedarea.Table)
	columns := paq.ctx.Fields
	if len(columns) == 0 {
		columns = protectedarea.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if paq.sql != nil {
		selector = paq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if paq.ctx.Unique != nil && *paq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range paq.predicates {
		p(selector)
	}
	for _, p := range paq.order {
		p(selector)
	}
	if offset := paq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := paq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedProtectedAreaPictures tells the query-builder to eager-load the nodes that are connected to the "protected_area_pictures"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (paq *ProtectedAreaQuery) WithNamedProtectedAreaPictures(name string, opts ...func(*ProtectedAreaPictureQuery)) *ProtectedAreaQuery {
	query := (&ProtectedAreaPictureClient{config: paq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if paq.withNamedProtectedAreaPictures == nil {
		paq.withNamedProtectedAreaPictures = make(map[string]*ProtectedAreaPictureQuery)
	}
	paq.withNamedProtectedAreaPictures[name] = query
	return paq
}

// ProtectedAreaGroupBy is the group-by builder for ProtectedArea entities.
type ProtectedAreaGroupBy struct {
	selector
	build *ProtectedAreaQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pagb *ProtectedAreaGroupBy) Aggregate(fns ...AggregateFunc) *ProtectedAreaGroupBy {
	pagb.fns = append(pagb.fns, fns...)
	return pagb
}

// Scan applies the selector query and scans the result into the given value.
func (pagb *ProtectedAreaGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pagb.build.ctx, "GroupBy")
	if err := pagb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProtectedAreaQuery, *ProtectedAreaGroupBy](ctx, pagb.build, pagb, pagb.build.inters, v)
}

func (pagb *ProtectedAreaGroupBy) sqlScan(ctx context.Context, root *ProtectedAreaQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pagb.fns))
	for _, fn := range pagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pagb.flds)+len(pagb.fns))
		for _, f := range *pagb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pagb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pagb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ProtectedAreaSelect is the builder for selecting fields of ProtectedArea entities.
type ProtectedAreaSelect struct {
	*ProtectedAreaQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pas *ProtectedAreaSelect) Aggregate(fns ...AggregateFunc) *ProtectedAreaSelect {
	pas.fns = append(pas.fns, fns...)
	return pas
}

// Scan applies the selector query and scans the result into the given value.
func (pas *ProtectedAreaSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pas.ctx, "Select")
	if err := pas.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProtectedAreaQuery, *ProtectedAreaSelect](ctx, pas.ProtectedAreaQuery, pas, pas.inters, v)
}

func (pas *ProtectedAreaSelect) sqlScan(ctx context.Context, root *ProtectedAreaQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pas.fns))
	for _, fn := range pas.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pas.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
