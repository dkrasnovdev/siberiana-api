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
	"github.com/dkrasnovdev/siberiana-api/ent/book"
	"github.com/dkrasnovdev/siberiana-api/ent/organization"
	"github.com/dkrasnovdev/siberiana-api/ent/person"
	"github.com/dkrasnovdev/siberiana-api/ent/predicate"
)

// OrganizationQuery is the builder for querying Organization entities.
type OrganizationQuery struct {
	config
	ctx                *QueryContext
	order              []organization.OrderOption
	inters             []Interceptor
	predicates         []predicate.Organization
	withArtifacts      *ArtifactQuery
	withBooks          *BookQuery
	withPeople         *PersonQuery
	modifiers          []func(*sql.Selector)
	loadTotal          []func(context.Context, []*Organization) error
	withNamedArtifacts map[string]*ArtifactQuery
	withNamedBooks     map[string]*BookQuery
	withNamedPeople    map[string]*PersonQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OrganizationQuery builder.
func (oq *OrganizationQuery) Where(ps ...predicate.Organization) *OrganizationQuery {
	oq.predicates = append(oq.predicates, ps...)
	return oq
}

// Limit the number of records to be returned by this query.
func (oq *OrganizationQuery) Limit(limit int) *OrganizationQuery {
	oq.ctx.Limit = &limit
	return oq
}

// Offset to start from.
func (oq *OrganizationQuery) Offset(offset int) *OrganizationQuery {
	oq.ctx.Offset = &offset
	return oq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (oq *OrganizationQuery) Unique(unique bool) *OrganizationQuery {
	oq.ctx.Unique = &unique
	return oq
}

// Order specifies how the records should be ordered.
func (oq *OrganizationQuery) Order(o ...organization.OrderOption) *OrganizationQuery {
	oq.order = append(oq.order, o...)
	return oq
}

// QueryArtifacts chains the current query on the "artifacts" edge.
func (oq *OrganizationQuery) QueryArtifacts() *ArtifactQuery {
	query := (&ArtifactClient{config: oq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(organization.Table, organization.FieldID, selector),
			sqlgraph.To(artifact.Table, artifact.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, organization.ArtifactsTable, organization.ArtifactsColumn),
		)
		fromU = sqlgraph.SetNeighbors(oq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBooks chains the current query on the "books" edge.
func (oq *OrganizationQuery) QueryBooks() *BookQuery {
	query := (&BookClient{config: oq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(organization.Table, organization.FieldID, selector),
			sqlgraph.To(book.Table, book.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, organization.BooksTable, organization.BooksColumn),
		)
		fromU = sqlgraph.SetNeighbors(oq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPeople chains the current query on the "people" edge.
func (oq *OrganizationQuery) QueryPeople() *PersonQuery {
	query := (&PersonClient{config: oq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(organization.Table, organization.FieldID, selector),
			sqlgraph.To(person.Table, person.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, organization.PeopleTable, organization.PeopleColumn),
		)
		fromU = sqlgraph.SetNeighbors(oq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Organization entity from the query.
// Returns a *NotFoundError when no Organization was found.
func (oq *OrganizationQuery) First(ctx context.Context) (*Organization, error) {
	nodes, err := oq.Limit(1).All(setContextOp(ctx, oq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{organization.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (oq *OrganizationQuery) FirstX(ctx context.Context) *Organization {
	node, err := oq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Organization ID from the query.
// Returns a *NotFoundError when no Organization ID was found.
func (oq *OrganizationQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oq.Limit(1).IDs(setContextOp(ctx, oq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{organization.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (oq *OrganizationQuery) FirstIDX(ctx context.Context) int {
	id, err := oq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Organization entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Organization entity is found.
// Returns a *NotFoundError when no Organization entities are found.
func (oq *OrganizationQuery) Only(ctx context.Context) (*Organization, error) {
	nodes, err := oq.Limit(2).All(setContextOp(ctx, oq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{organization.Label}
	default:
		return nil, &NotSingularError{organization.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (oq *OrganizationQuery) OnlyX(ctx context.Context) *Organization {
	node, err := oq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Organization ID in the query.
// Returns a *NotSingularError when more than one Organization ID is found.
// Returns a *NotFoundError when no entities are found.
func (oq *OrganizationQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oq.Limit(2).IDs(setContextOp(ctx, oq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{organization.Label}
	default:
		err = &NotSingularError{organization.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (oq *OrganizationQuery) OnlyIDX(ctx context.Context) int {
	id, err := oq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Organizations.
func (oq *OrganizationQuery) All(ctx context.Context) ([]*Organization, error) {
	ctx = setContextOp(ctx, oq.ctx, "All")
	if err := oq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Organization, *OrganizationQuery]()
	return withInterceptors[[]*Organization](ctx, oq, qr, oq.inters)
}

// AllX is like All, but panics if an error occurs.
func (oq *OrganizationQuery) AllX(ctx context.Context) []*Organization {
	nodes, err := oq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Organization IDs.
func (oq *OrganizationQuery) IDs(ctx context.Context) (ids []int, err error) {
	if oq.ctx.Unique == nil && oq.path != nil {
		oq.Unique(true)
	}
	ctx = setContextOp(ctx, oq.ctx, "IDs")
	if err = oq.Select(organization.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (oq *OrganizationQuery) IDsX(ctx context.Context) []int {
	ids, err := oq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (oq *OrganizationQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, oq.ctx, "Count")
	if err := oq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, oq, querierCount[*OrganizationQuery](), oq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (oq *OrganizationQuery) CountX(ctx context.Context) int {
	count, err := oq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (oq *OrganizationQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, oq.ctx, "Exist")
	switch _, err := oq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (oq *OrganizationQuery) ExistX(ctx context.Context) bool {
	exist, err := oq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OrganizationQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (oq *OrganizationQuery) Clone() *OrganizationQuery {
	if oq == nil {
		return nil
	}
	return &OrganizationQuery{
		config:        oq.config,
		ctx:           oq.ctx.Clone(),
		order:         append([]organization.OrderOption{}, oq.order...),
		inters:        append([]Interceptor{}, oq.inters...),
		predicates:    append([]predicate.Organization{}, oq.predicates...),
		withArtifacts: oq.withArtifacts.Clone(),
		withBooks:     oq.withBooks.Clone(),
		withPeople:    oq.withPeople.Clone(),
		// clone intermediate query.
		sql:  oq.sql.Clone(),
		path: oq.path,
	}
}

// WithArtifacts tells the query-builder to eager-load the nodes that are connected to
// the "artifacts" edge. The optional arguments are used to configure the query builder of the edge.
func (oq *OrganizationQuery) WithArtifacts(opts ...func(*ArtifactQuery)) *OrganizationQuery {
	query := (&ArtifactClient{config: oq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	oq.withArtifacts = query
	return oq
}

// WithBooks tells the query-builder to eager-load the nodes that are connected to
// the "books" edge. The optional arguments are used to configure the query builder of the edge.
func (oq *OrganizationQuery) WithBooks(opts ...func(*BookQuery)) *OrganizationQuery {
	query := (&BookClient{config: oq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	oq.withBooks = query
	return oq
}

// WithPeople tells the query-builder to eager-load the nodes that are connected to
// the "people" edge. The optional arguments are used to configure the query builder of the edge.
func (oq *OrganizationQuery) WithPeople(opts ...func(*PersonQuery)) *OrganizationQuery {
	query := (&PersonClient{config: oq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	oq.withPeople = query
	return oq
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
//	client.Organization.Query().
//		GroupBy(organization.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (oq *OrganizationQuery) GroupBy(field string, fields ...string) *OrganizationGroupBy {
	oq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &OrganizationGroupBy{build: oq}
	grbuild.flds = &oq.ctx.Fields
	grbuild.label = organization.Label
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
//	client.Organization.Query().
//		Select(organization.FieldCreatedAt).
//		Scan(ctx, &v)
func (oq *OrganizationQuery) Select(fields ...string) *OrganizationSelect {
	oq.ctx.Fields = append(oq.ctx.Fields, fields...)
	sbuild := &OrganizationSelect{OrganizationQuery: oq}
	sbuild.label = organization.Label
	sbuild.flds, sbuild.scan = &oq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a OrganizationSelect configured with the given aggregations.
func (oq *OrganizationQuery) Aggregate(fns ...AggregateFunc) *OrganizationSelect {
	return oq.Select().Aggregate(fns...)
}

func (oq *OrganizationQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range oq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, oq); err != nil {
				return err
			}
		}
	}
	for _, f := range oq.ctx.Fields {
		if !organization.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if oq.path != nil {
		prev, err := oq.path(ctx)
		if err != nil {
			return err
		}
		oq.sql = prev
	}
	if organization.Policy == nil {
		return errors.New("ent: uninitialized organization.Policy (forgotten import ent/runtime?)")
	}
	if err := organization.Policy.EvalQuery(ctx, oq); err != nil {
		return err
	}
	return nil
}

func (oq *OrganizationQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Organization, error) {
	var (
		nodes       = []*Organization{}
		_spec       = oq.querySpec()
		loadedTypes = [3]bool{
			oq.withArtifacts != nil,
			oq.withBooks != nil,
			oq.withPeople != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Organization).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Organization{config: oq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(oq.modifiers) > 0 {
		_spec.Modifiers = oq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, oq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := oq.withArtifacts; query != nil {
		if err := oq.loadArtifacts(ctx, query, nodes,
			func(n *Organization) { n.Edges.Artifacts = []*Artifact{} },
			func(n *Organization, e *Artifact) { n.Edges.Artifacts = append(n.Edges.Artifacts, e) }); err != nil {
			return nil, err
		}
	}
	if query := oq.withBooks; query != nil {
		if err := oq.loadBooks(ctx, query, nodes,
			func(n *Organization) { n.Edges.Books = []*Book{} },
			func(n *Organization, e *Book) { n.Edges.Books = append(n.Edges.Books, e) }); err != nil {
			return nil, err
		}
	}
	if query := oq.withPeople; query != nil {
		if err := oq.loadPeople(ctx, query, nodes,
			func(n *Organization) { n.Edges.People = []*Person{} },
			func(n *Organization, e *Person) { n.Edges.People = append(n.Edges.People, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range oq.withNamedArtifacts {
		if err := oq.loadArtifacts(ctx, query, nodes,
			func(n *Organization) { n.appendNamedArtifacts(name) },
			func(n *Organization, e *Artifact) { n.appendNamedArtifacts(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range oq.withNamedBooks {
		if err := oq.loadBooks(ctx, query, nodes,
			func(n *Organization) { n.appendNamedBooks(name) },
			func(n *Organization, e *Book) { n.appendNamedBooks(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range oq.withNamedPeople {
		if err := oq.loadPeople(ctx, query, nodes,
			func(n *Organization) { n.appendNamedPeople(name) },
			func(n *Organization, e *Person) { n.appendNamedPeople(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range oq.loadTotal {
		if err := oq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (oq *OrganizationQuery) loadArtifacts(ctx context.Context, query *ArtifactQuery, nodes []*Organization, init func(*Organization), assign func(*Organization, *Artifact)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Organization)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Artifact(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(organization.ArtifactsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.organization_artifacts
		if fk == nil {
			return fmt.Errorf(`foreign-key "organization_artifacts" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "organization_artifacts" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (oq *OrganizationQuery) loadBooks(ctx context.Context, query *BookQuery, nodes []*Organization, init func(*Organization), assign func(*Organization, *Book)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Organization)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Book(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(organization.BooksColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.organization_books
		if fk == nil {
			return fmt.Errorf(`foreign-key "organization_books" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "organization_books" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (oq *OrganizationQuery) loadPeople(ctx context.Context, query *PersonQuery, nodes []*Organization, init func(*Organization), assign func(*Organization, *Person)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Organization)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Person(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(organization.PeopleColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.organization_people
		if fk == nil {
			return fmt.Errorf(`foreign-key "organization_people" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "organization_people" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (oq *OrganizationQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := oq.querySpec()
	if len(oq.modifiers) > 0 {
		_spec.Modifiers = oq.modifiers
	}
	_spec.Node.Columns = oq.ctx.Fields
	if len(oq.ctx.Fields) > 0 {
		_spec.Unique = oq.ctx.Unique != nil && *oq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, oq.driver, _spec)
}

func (oq *OrganizationQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(organization.Table, organization.Columns, sqlgraph.NewFieldSpec(organization.FieldID, field.TypeInt))
	_spec.From = oq.sql
	if unique := oq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if oq.path != nil {
		_spec.Unique = true
	}
	if fields := oq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, organization.FieldID)
		for i := range fields {
			if fields[i] != organization.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := oq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := oq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := oq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := oq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (oq *OrganizationQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(oq.driver.Dialect())
	t1 := builder.Table(organization.Table)
	columns := oq.ctx.Fields
	if len(columns) == 0 {
		columns = organization.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if oq.sql != nil {
		selector = oq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if oq.ctx.Unique != nil && *oq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range oq.predicates {
		p(selector)
	}
	for _, p := range oq.order {
		p(selector)
	}
	if offset := oq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := oq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedArtifacts tells the query-builder to eager-load the nodes that are connected to the "artifacts"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (oq *OrganizationQuery) WithNamedArtifacts(name string, opts ...func(*ArtifactQuery)) *OrganizationQuery {
	query := (&ArtifactClient{config: oq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if oq.withNamedArtifacts == nil {
		oq.withNamedArtifacts = make(map[string]*ArtifactQuery)
	}
	oq.withNamedArtifacts[name] = query
	return oq
}

// WithNamedBooks tells the query-builder to eager-load the nodes that are connected to the "books"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (oq *OrganizationQuery) WithNamedBooks(name string, opts ...func(*BookQuery)) *OrganizationQuery {
	query := (&BookClient{config: oq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if oq.withNamedBooks == nil {
		oq.withNamedBooks = make(map[string]*BookQuery)
	}
	oq.withNamedBooks[name] = query
	return oq
}

// WithNamedPeople tells the query-builder to eager-load the nodes that are connected to the "people"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (oq *OrganizationQuery) WithNamedPeople(name string, opts ...func(*PersonQuery)) *OrganizationQuery {
	query := (&PersonClient{config: oq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if oq.withNamedPeople == nil {
		oq.withNamedPeople = make(map[string]*PersonQuery)
	}
	oq.withNamedPeople[name] = query
	return oq
}

// OrganizationGroupBy is the group-by builder for Organization entities.
type OrganizationGroupBy struct {
	selector
	build *OrganizationQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ogb *OrganizationGroupBy) Aggregate(fns ...AggregateFunc) *OrganizationGroupBy {
	ogb.fns = append(ogb.fns, fns...)
	return ogb
}

// Scan applies the selector query and scans the result into the given value.
func (ogb *OrganizationGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ogb.build.ctx, "GroupBy")
	if err := ogb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrganizationQuery, *OrganizationGroupBy](ctx, ogb.build, ogb, ogb.build.inters, v)
}

func (ogb *OrganizationGroupBy) sqlScan(ctx context.Context, root *OrganizationQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ogb.fns))
	for _, fn := range ogb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ogb.flds)+len(ogb.fns))
		for _, f := range *ogb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ogb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ogb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// OrganizationSelect is the builder for selecting fields of Organization entities.
type OrganizationSelect struct {
	*OrganizationQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (os *OrganizationSelect) Aggregate(fns ...AggregateFunc) *OrganizationSelect {
	os.fns = append(os.fns, fns...)
	return os
}

// Scan applies the selector query and scans the result into the given value.
func (os *OrganizationSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, os.ctx, "Select")
	if err := os.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrganizationQuery, *OrganizationSelect](ctx, os.OrganizationQuery, os, os.inters, v)
}

func (os *OrganizationSelect) sqlScan(ctx context.Context, root *OrganizationQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(os.fns))
	for _, fn := range os.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*os.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := os.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
