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
	"github.com/dkrasnovdev/heritage-api/ent/book"
	"github.com/dkrasnovdev/heritage-api/ent/bookgenre"
	"github.com/dkrasnovdev/heritage-api/ent/predicate"
)

// BookGenreQuery is the builder for querying BookGenre entities.
type BookGenreQuery struct {
	config
	ctx            *QueryContext
	order          []bookgenre.OrderOption
	inters         []Interceptor
	predicates     []predicate.BookGenre
	withBooks      *BookQuery
	modifiers      []func(*sql.Selector)
	loadTotal      []func(context.Context, []*BookGenre) error
	withNamedBooks map[string]*BookQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BookGenreQuery builder.
func (bgq *BookGenreQuery) Where(ps ...predicate.BookGenre) *BookGenreQuery {
	bgq.predicates = append(bgq.predicates, ps...)
	return bgq
}

// Limit the number of records to be returned by this query.
func (bgq *BookGenreQuery) Limit(limit int) *BookGenreQuery {
	bgq.ctx.Limit = &limit
	return bgq
}

// Offset to start from.
func (bgq *BookGenreQuery) Offset(offset int) *BookGenreQuery {
	bgq.ctx.Offset = &offset
	return bgq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bgq *BookGenreQuery) Unique(unique bool) *BookGenreQuery {
	bgq.ctx.Unique = &unique
	return bgq
}

// Order specifies how the records should be ordered.
func (bgq *BookGenreQuery) Order(o ...bookgenre.OrderOption) *BookGenreQuery {
	bgq.order = append(bgq.order, o...)
	return bgq
}

// QueryBooks chains the current query on the "books" edge.
func (bgq *BookGenreQuery) QueryBooks() *BookQuery {
	query := (&BookClient{config: bgq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bgq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bgq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(bookgenre.Table, bookgenre.FieldID, selector),
			sqlgraph.To(book.Table, book.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, bookgenre.BooksTable, bookgenre.BooksPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(bgq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first BookGenre entity from the query.
// Returns a *NotFoundError when no BookGenre was found.
func (bgq *BookGenreQuery) First(ctx context.Context) (*BookGenre, error) {
	nodes, err := bgq.Limit(1).All(setContextOp(ctx, bgq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{bookgenre.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bgq *BookGenreQuery) FirstX(ctx context.Context) *BookGenre {
	node, err := bgq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BookGenre ID from the query.
// Returns a *NotFoundError when no BookGenre ID was found.
func (bgq *BookGenreQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bgq.Limit(1).IDs(setContextOp(ctx, bgq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{bookgenre.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bgq *BookGenreQuery) FirstIDX(ctx context.Context) int {
	id, err := bgq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BookGenre entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one BookGenre entity is found.
// Returns a *NotFoundError when no BookGenre entities are found.
func (bgq *BookGenreQuery) Only(ctx context.Context) (*BookGenre, error) {
	nodes, err := bgq.Limit(2).All(setContextOp(ctx, bgq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{bookgenre.Label}
	default:
		return nil, &NotSingularError{bookgenre.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bgq *BookGenreQuery) OnlyX(ctx context.Context) *BookGenre {
	node, err := bgq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BookGenre ID in the query.
// Returns a *NotSingularError when more than one BookGenre ID is found.
// Returns a *NotFoundError when no entities are found.
func (bgq *BookGenreQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bgq.Limit(2).IDs(setContextOp(ctx, bgq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{bookgenre.Label}
	default:
		err = &NotSingularError{bookgenre.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bgq *BookGenreQuery) OnlyIDX(ctx context.Context) int {
	id, err := bgq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BookGenres.
func (bgq *BookGenreQuery) All(ctx context.Context) ([]*BookGenre, error) {
	ctx = setContextOp(ctx, bgq.ctx, "All")
	if err := bgq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*BookGenre, *BookGenreQuery]()
	return withInterceptors[[]*BookGenre](ctx, bgq, qr, bgq.inters)
}

// AllX is like All, but panics if an error occurs.
func (bgq *BookGenreQuery) AllX(ctx context.Context) []*BookGenre {
	nodes, err := bgq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BookGenre IDs.
func (bgq *BookGenreQuery) IDs(ctx context.Context) (ids []int, err error) {
	if bgq.ctx.Unique == nil && bgq.path != nil {
		bgq.Unique(true)
	}
	ctx = setContextOp(ctx, bgq.ctx, "IDs")
	if err = bgq.Select(bookgenre.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bgq *BookGenreQuery) IDsX(ctx context.Context) []int {
	ids, err := bgq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bgq *BookGenreQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, bgq.ctx, "Count")
	if err := bgq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, bgq, querierCount[*BookGenreQuery](), bgq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (bgq *BookGenreQuery) CountX(ctx context.Context) int {
	count, err := bgq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bgq *BookGenreQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, bgq.ctx, "Exist")
	switch _, err := bgq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (bgq *BookGenreQuery) ExistX(ctx context.Context) bool {
	exist, err := bgq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BookGenreQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bgq *BookGenreQuery) Clone() *BookGenreQuery {
	if bgq == nil {
		return nil
	}
	return &BookGenreQuery{
		config:     bgq.config,
		ctx:        bgq.ctx.Clone(),
		order:      append([]bookgenre.OrderOption{}, bgq.order...),
		inters:     append([]Interceptor{}, bgq.inters...),
		predicates: append([]predicate.BookGenre{}, bgq.predicates...),
		withBooks:  bgq.withBooks.Clone(),
		// clone intermediate query.
		sql:  bgq.sql.Clone(),
		path: bgq.path,
	}
}

// WithBooks tells the query-builder to eager-load the nodes that are connected to
// the "books" edge. The optional arguments are used to configure the query builder of the edge.
func (bgq *BookGenreQuery) WithBooks(opts ...func(*BookQuery)) *BookGenreQuery {
	query := (&BookClient{config: bgq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	bgq.withBooks = query
	return bgq
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
//	client.BookGenre.Query().
//		GroupBy(bookgenre.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (bgq *BookGenreQuery) GroupBy(field string, fields ...string) *BookGenreGroupBy {
	bgq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &BookGenreGroupBy{build: bgq}
	grbuild.flds = &bgq.ctx.Fields
	grbuild.label = bookgenre.Label
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
//	client.BookGenre.Query().
//		Select(bookgenre.FieldCreatedAt).
//		Scan(ctx, &v)
func (bgq *BookGenreQuery) Select(fields ...string) *BookGenreSelect {
	bgq.ctx.Fields = append(bgq.ctx.Fields, fields...)
	sbuild := &BookGenreSelect{BookGenreQuery: bgq}
	sbuild.label = bookgenre.Label
	sbuild.flds, sbuild.scan = &bgq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a BookGenreSelect configured with the given aggregations.
func (bgq *BookGenreQuery) Aggregate(fns ...AggregateFunc) *BookGenreSelect {
	return bgq.Select().Aggregate(fns...)
}

func (bgq *BookGenreQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range bgq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, bgq); err != nil {
				return err
			}
		}
	}
	for _, f := range bgq.ctx.Fields {
		if !bookgenre.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if bgq.path != nil {
		prev, err := bgq.path(ctx)
		if err != nil {
			return err
		}
		bgq.sql = prev
	}
	if bookgenre.Policy == nil {
		return errors.New("ent: uninitialized bookgenre.Policy (forgotten import ent/runtime?)")
	}
	if err := bookgenre.Policy.EvalQuery(ctx, bgq); err != nil {
		return err
	}
	return nil
}

func (bgq *BookGenreQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*BookGenre, error) {
	var (
		nodes       = []*BookGenre{}
		_spec       = bgq.querySpec()
		loadedTypes = [1]bool{
			bgq.withBooks != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*BookGenre).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &BookGenre{config: bgq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(bgq.modifiers) > 0 {
		_spec.Modifiers = bgq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bgq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := bgq.withBooks; query != nil {
		if err := bgq.loadBooks(ctx, query, nodes,
			func(n *BookGenre) { n.Edges.Books = []*Book{} },
			func(n *BookGenre, e *Book) { n.Edges.Books = append(n.Edges.Books, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range bgq.withNamedBooks {
		if err := bgq.loadBooks(ctx, query, nodes,
			func(n *BookGenre) { n.appendNamedBooks(name) },
			func(n *BookGenre, e *Book) { n.appendNamedBooks(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range bgq.loadTotal {
		if err := bgq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (bgq *BookGenreQuery) loadBooks(ctx context.Context, query *BookQuery, nodes []*BookGenre, init func(*BookGenre), assign func(*BookGenre, *Book)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*BookGenre)
	nids := make(map[int]map[*BookGenre]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(bookgenre.BooksTable)
		s.Join(joinT).On(s.C(book.FieldID), joinT.C(bookgenre.BooksPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(bookgenre.BooksPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(bookgenre.BooksPrimaryKey[0]))
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
					nids[inValue] = map[*BookGenre]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Book](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "books" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (bgq *BookGenreQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bgq.querySpec()
	if len(bgq.modifiers) > 0 {
		_spec.Modifiers = bgq.modifiers
	}
	_spec.Node.Columns = bgq.ctx.Fields
	if len(bgq.ctx.Fields) > 0 {
		_spec.Unique = bgq.ctx.Unique != nil && *bgq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, bgq.driver, _spec)
}

func (bgq *BookGenreQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(bookgenre.Table, bookgenre.Columns, sqlgraph.NewFieldSpec(bookgenre.FieldID, field.TypeInt))
	_spec.From = bgq.sql
	if unique := bgq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if bgq.path != nil {
		_spec.Unique = true
	}
	if fields := bgq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, bookgenre.FieldID)
		for i := range fields {
			if fields[i] != bookgenre.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bgq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bgq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bgq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bgq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bgq *BookGenreQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bgq.driver.Dialect())
	t1 := builder.Table(bookgenre.Table)
	columns := bgq.ctx.Fields
	if len(columns) == 0 {
		columns = bookgenre.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bgq.sql != nil {
		selector = bgq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bgq.ctx.Unique != nil && *bgq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range bgq.predicates {
		p(selector)
	}
	for _, p := range bgq.order {
		p(selector)
	}
	if offset := bgq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bgq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedBooks tells the query-builder to eager-load the nodes that are connected to the "books"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (bgq *BookGenreQuery) WithNamedBooks(name string, opts ...func(*BookQuery)) *BookGenreQuery {
	query := (&BookClient{config: bgq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if bgq.withNamedBooks == nil {
		bgq.withNamedBooks = make(map[string]*BookQuery)
	}
	bgq.withNamedBooks[name] = query
	return bgq
}

// BookGenreGroupBy is the group-by builder for BookGenre entities.
type BookGenreGroupBy struct {
	selector
	build *BookGenreQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bggb *BookGenreGroupBy) Aggregate(fns ...AggregateFunc) *BookGenreGroupBy {
	bggb.fns = append(bggb.fns, fns...)
	return bggb
}

// Scan applies the selector query and scans the result into the given value.
func (bggb *BookGenreGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bggb.build.ctx, "GroupBy")
	if err := bggb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BookGenreQuery, *BookGenreGroupBy](ctx, bggb.build, bggb, bggb.build.inters, v)
}

func (bggb *BookGenreGroupBy) sqlScan(ctx context.Context, root *BookGenreQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(bggb.fns))
	for _, fn := range bggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*bggb.flds)+len(bggb.fns))
		for _, f := range *bggb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*bggb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bggb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// BookGenreSelect is the builder for selecting fields of BookGenre entities.
type BookGenreSelect struct {
	*BookGenreQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (bgs *BookGenreSelect) Aggregate(fns ...AggregateFunc) *BookGenreSelect {
	bgs.fns = append(bgs.fns, fns...)
	return bgs
}

// Scan applies the selector query and scans the result into the given value.
func (bgs *BookGenreSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bgs.ctx, "Select")
	if err := bgs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BookGenreQuery, *BookGenreSelect](ctx, bgs.BookGenreQuery, bgs, bgs.inters, v)
}

func (bgs *BookGenreSelect) sqlScan(ctx context.Context, root *BookGenreQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(bgs.fns))
	for _, fn := range bgs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*bgs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bgs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
