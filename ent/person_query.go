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
	"github.com/dkrasnovdev/heritage-api/ent/holder"
	"github.com/dkrasnovdev/heritage-api/ent/person"
	"github.com/dkrasnovdev/heritage-api/ent/predicate"
	"github.com/dkrasnovdev/heritage-api/ent/project"
	"github.com/dkrasnovdev/heritage-api/ent/publication"
)

// PersonQuery is the builder for querying Person entities.
type PersonQuery struct {
	config
	ctx                   *QueryContext
	order                 []person.OrderOption
	inters                []Interceptor
	predicates            []predicate.Person
	withArtifacts         *ArtifactQuery
	withProjects          *ProjectQuery
	withPublications      *PublicationQuery
	withHolder            *HolderQuery
	withFKs               bool
	modifiers             []func(*sql.Selector)
	loadTotal             []func(context.Context, []*Person) error
	withNamedArtifacts    map[string]*ArtifactQuery
	withNamedProjects     map[string]*ProjectQuery
	withNamedPublications map[string]*PublicationQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PersonQuery builder.
func (pq *PersonQuery) Where(ps ...predicate.Person) *PersonQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit the number of records to be returned by this query.
func (pq *PersonQuery) Limit(limit int) *PersonQuery {
	pq.ctx.Limit = &limit
	return pq
}

// Offset to start from.
func (pq *PersonQuery) Offset(offset int) *PersonQuery {
	pq.ctx.Offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *PersonQuery) Unique(unique bool) *PersonQuery {
	pq.ctx.Unique = &unique
	return pq
}

// Order specifies how the records should be ordered.
func (pq *PersonQuery) Order(o ...person.OrderOption) *PersonQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryArtifacts chains the current query on the "artifacts" edge.
func (pq *PersonQuery) QueryArtifacts() *ArtifactQuery {
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
			sqlgraph.From(person.Table, person.FieldID, selector),
			sqlgraph.To(artifact.Table, artifact.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, person.ArtifactsTable, person.ArtifactsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryProjects chains the current query on the "projects" edge.
func (pq *PersonQuery) QueryProjects() *ProjectQuery {
	query := (&ProjectClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(person.Table, person.FieldID, selector),
			sqlgraph.To(project.Table, project.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, person.ProjectsTable, person.ProjectsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPublications chains the current query on the "publications" edge.
func (pq *PersonQuery) QueryPublications() *PublicationQuery {
	query := (&PublicationClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(person.Table, person.FieldID, selector),
			sqlgraph.To(publication.Table, publication.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, person.PublicationsTable, person.PublicationsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryHolder chains the current query on the "holder" edge.
func (pq *PersonQuery) QueryHolder() *HolderQuery {
	query := (&HolderClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(person.Table, person.FieldID, selector),
			sqlgraph.To(holder.Table, holder.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, person.HolderTable, person.HolderColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Person entity from the query.
// Returns a *NotFoundError when no Person was found.
func (pq *PersonQuery) First(ctx context.Context) (*Person, error) {
	nodes, err := pq.Limit(1).All(setContextOp(ctx, pq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{person.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PersonQuery) FirstX(ctx context.Context) *Person {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Person ID from the query.
// Returns a *NotFoundError when no Person ID was found.
func (pq *PersonQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(1).IDs(setContextOp(ctx, pq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{person.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *PersonQuery) FirstIDX(ctx context.Context) int {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Person entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Person entity is found.
// Returns a *NotFoundError when no Person entities are found.
func (pq *PersonQuery) Only(ctx context.Context) (*Person, error) {
	nodes, err := pq.Limit(2).All(setContextOp(ctx, pq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{person.Label}
	default:
		return nil, &NotSingularError{person.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PersonQuery) OnlyX(ctx context.Context) *Person {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Person ID in the query.
// Returns a *NotSingularError when more than one Person ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *PersonQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(2).IDs(setContextOp(ctx, pq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{person.Label}
	default:
		err = &NotSingularError{person.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *PersonQuery) OnlyIDX(ctx context.Context) int {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Persons.
func (pq *PersonQuery) All(ctx context.Context) ([]*Person, error) {
	ctx = setContextOp(ctx, pq.ctx, "All")
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Person, *PersonQuery]()
	return withInterceptors[[]*Person](ctx, pq, qr, pq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pq *PersonQuery) AllX(ctx context.Context) []*Person {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Person IDs.
func (pq *PersonQuery) IDs(ctx context.Context) (ids []int, err error) {
	if pq.ctx.Unique == nil && pq.path != nil {
		pq.Unique(true)
	}
	ctx = setContextOp(ctx, pq.ctx, "IDs")
	if err = pq.Select(person.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PersonQuery) IDsX(ctx context.Context) []int {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PersonQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pq.ctx, "Count")
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pq, querierCount[*PersonQuery](), pq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PersonQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PersonQuery) Exist(ctx context.Context) (bool, error) {
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
func (pq *PersonQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PersonQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PersonQuery) Clone() *PersonQuery {
	if pq == nil {
		return nil
	}
	return &PersonQuery{
		config:           pq.config,
		ctx:              pq.ctx.Clone(),
		order:            append([]person.OrderOption{}, pq.order...),
		inters:           append([]Interceptor{}, pq.inters...),
		predicates:       append([]predicate.Person{}, pq.predicates...),
		withArtifacts:    pq.withArtifacts.Clone(),
		withProjects:     pq.withProjects.Clone(),
		withPublications: pq.withPublications.Clone(),
		withHolder:       pq.withHolder.Clone(),
		// clone intermediate query.
		sql:  pq.sql.Clone(),
		path: pq.path,
	}
}

// WithArtifacts tells the query-builder to eager-load the nodes that are connected to
// the "artifacts" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PersonQuery) WithArtifacts(opts ...func(*ArtifactQuery)) *PersonQuery {
	query := (&ArtifactClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withArtifacts = query
	return pq
}

// WithProjects tells the query-builder to eager-load the nodes that are connected to
// the "projects" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PersonQuery) WithProjects(opts ...func(*ProjectQuery)) *PersonQuery {
	query := (&ProjectClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withProjects = query
	return pq
}

// WithPublications tells the query-builder to eager-load the nodes that are connected to
// the "publications" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PersonQuery) WithPublications(opts ...func(*PublicationQuery)) *PersonQuery {
	query := (&PublicationClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withPublications = query
	return pq
}

// WithHolder tells the query-builder to eager-load the nodes that are connected to
// the "holder" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PersonQuery) WithHolder(opts ...func(*HolderQuery)) *PersonQuery {
	query := (&HolderClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withHolder = query
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
//	client.Person.Query().
//		GroupBy(person.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pq *PersonQuery) GroupBy(field string, fields ...string) *PersonGroupBy {
	pq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PersonGroupBy{build: pq}
	grbuild.flds = &pq.ctx.Fields
	grbuild.label = person.Label
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
//	client.Person.Query().
//		Select(person.FieldCreatedAt).
//		Scan(ctx, &v)
func (pq *PersonQuery) Select(fields ...string) *PersonSelect {
	pq.ctx.Fields = append(pq.ctx.Fields, fields...)
	sbuild := &PersonSelect{PersonQuery: pq}
	sbuild.label = person.Label
	sbuild.flds, sbuild.scan = &pq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PersonSelect configured with the given aggregations.
func (pq *PersonQuery) Aggregate(fns ...AggregateFunc) *PersonSelect {
	return pq.Select().Aggregate(fns...)
}

func (pq *PersonQuery) prepareQuery(ctx context.Context) error {
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
		if !person.ValidColumn(f) {
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
	if person.Policy == nil {
		return errors.New("ent: uninitialized person.Policy (forgotten import ent/runtime?)")
	}
	if err := person.Policy.EvalQuery(ctx, pq); err != nil {
		return err
	}
	return nil
}

func (pq *PersonQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Person, error) {
	var (
		nodes       = []*Person{}
		withFKs     = pq.withFKs
		_spec       = pq.querySpec()
		loadedTypes = [4]bool{
			pq.withArtifacts != nil,
			pq.withProjects != nil,
			pq.withPublications != nil,
			pq.withHolder != nil,
		}
	)
	if pq.withHolder != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, person.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Person).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Person{config: pq.config}
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
			func(n *Person) { n.Edges.Artifacts = []*Artifact{} },
			func(n *Person, e *Artifact) { n.Edges.Artifacts = append(n.Edges.Artifacts, e) }); err != nil {
			return nil, err
		}
	}
	if query := pq.withProjects; query != nil {
		if err := pq.loadProjects(ctx, query, nodes,
			func(n *Person) { n.Edges.Projects = []*Project{} },
			func(n *Person, e *Project) { n.Edges.Projects = append(n.Edges.Projects, e) }); err != nil {
			return nil, err
		}
	}
	if query := pq.withPublications; query != nil {
		if err := pq.loadPublications(ctx, query, nodes,
			func(n *Person) { n.Edges.Publications = []*Publication{} },
			func(n *Person, e *Publication) { n.Edges.Publications = append(n.Edges.Publications, e) }); err != nil {
			return nil, err
		}
	}
	if query := pq.withHolder; query != nil {
		if err := pq.loadHolder(ctx, query, nodes, nil,
			func(n *Person, e *Holder) { n.Edges.Holder = e }); err != nil {
			return nil, err
		}
	}
	for name, query := range pq.withNamedArtifacts {
		if err := pq.loadArtifacts(ctx, query, nodes,
			func(n *Person) { n.appendNamedArtifacts(name) },
			func(n *Person, e *Artifact) { n.appendNamedArtifacts(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range pq.withNamedProjects {
		if err := pq.loadProjects(ctx, query, nodes,
			func(n *Person) { n.appendNamedProjects(name) },
			func(n *Person, e *Project) { n.appendNamedProjects(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range pq.withNamedPublications {
		if err := pq.loadPublications(ctx, query, nodes,
			func(n *Person) { n.appendNamedPublications(name) },
			func(n *Person, e *Publication) { n.appendNamedPublications(name, e) }); err != nil {
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

func (pq *PersonQuery) loadArtifacts(ctx context.Context, query *ArtifactQuery, nodes []*Person, init func(*Person), assign func(*Person, *Artifact)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Person)
	nids := make(map[int]map[*Person]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(person.ArtifactsTable)
		s.Join(joinT).On(s.C(artifact.FieldID), joinT.C(person.ArtifactsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(person.ArtifactsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(person.ArtifactsPrimaryKey[0]))
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
					nids[inValue] = map[*Person]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Artifact](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "artifacts" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (pq *PersonQuery) loadProjects(ctx context.Context, query *ProjectQuery, nodes []*Person, init func(*Person), assign func(*Person, *Project)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Person)
	nids := make(map[int]map[*Person]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(person.ProjectsTable)
		s.Join(joinT).On(s.C(project.FieldID), joinT.C(person.ProjectsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(person.ProjectsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(person.ProjectsPrimaryKey[0]))
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
					nids[inValue] = map[*Person]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Project](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "projects" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (pq *PersonQuery) loadPublications(ctx context.Context, query *PublicationQuery, nodes []*Person, init func(*Person), assign func(*Person, *Publication)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Person)
	nids := make(map[int]map[*Person]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(person.PublicationsTable)
		s.Join(joinT).On(s.C(publication.FieldID), joinT.C(person.PublicationsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(person.PublicationsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(person.PublicationsPrimaryKey[0]))
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
					nids[inValue] = map[*Person]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Publication](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "publications" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (pq *PersonQuery) loadHolder(ctx context.Context, query *HolderQuery, nodes []*Person, init func(*Person), assign func(*Person, *Holder)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Person)
	for i := range nodes {
		if nodes[i].holder_person == nil {
			continue
		}
		fk := *nodes[i].holder_person
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(holder.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "holder_person" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (pq *PersonQuery) sqlCount(ctx context.Context) (int, error) {
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

func (pq *PersonQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(person.Table, person.Columns, sqlgraph.NewFieldSpec(person.FieldID, field.TypeInt))
	_spec.From = pq.sql
	if unique := pq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pq.path != nil {
		_spec.Unique = true
	}
	if fields := pq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, person.FieldID)
		for i := range fields {
			if fields[i] != person.FieldID {
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

func (pq *PersonQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(person.Table)
	columns := pq.ctx.Fields
	if len(columns) == 0 {
		columns = person.Columns
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
func (pq *PersonQuery) WithNamedArtifacts(name string, opts ...func(*ArtifactQuery)) *PersonQuery {
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

// WithNamedProjects tells the query-builder to eager-load the nodes that are connected to the "projects"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (pq *PersonQuery) WithNamedProjects(name string, opts ...func(*ProjectQuery)) *PersonQuery {
	query := (&ProjectClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if pq.withNamedProjects == nil {
		pq.withNamedProjects = make(map[string]*ProjectQuery)
	}
	pq.withNamedProjects[name] = query
	return pq
}

// WithNamedPublications tells the query-builder to eager-load the nodes that are connected to the "publications"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (pq *PersonQuery) WithNamedPublications(name string, opts ...func(*PublicationQuery)) *PersonQuery {
	query := (&PublicationClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if pq.withNamedPublications == nil {
		pq.withNamedPublications = make(map[string]*PublicationQuery)
	}
	pq.withNamedPublications[name] = query
	return pq
}

// PersonGroupBy is the group-by builder for Person entities.
type PersonGroupBy struct {
	selector
	build *PersonQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PersonGroupBy) Aggregate(fns ...AggregateFunc) *PersonGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the selector query and scans the result into the given value.
func (pgb *PersonGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pgb.build.ctx, "GroupBy")
	if err := pgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PersonQuery, *PersonGroupBy](ctx, pgb.build, pgb, pgb.build.inters, v)
}

func (pgb *PersonGroupBy) sqlScan(ctx context.Context, root *PersonQuery, v any) error {
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

// PersonSelect is the builder for selecting fields of Person entities.
type PersonSelect struct {
	*PersonQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ps *PersonSelect) Aggregate(fns ...AggregateFunc) *PersonSelect {
	ps.fns = append(ps.fns, fns...)
	return ps
}

// Scan applies the selector query and scans the result into the given value.
func (ps *PersonSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ps.ctx, "Select")
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PersonQuery, *PersonSelect](ctx, ps.PersonQuery, ps, ps.inters, v)
}

func (ps *PersonSelect) sqlScan(ctx context.Context, root *PersonQuery, v any) error {
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
