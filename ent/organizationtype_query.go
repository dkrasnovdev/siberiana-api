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
	"github.com/dkrasnovdev/heritage-api/ent/organization"
	"github.com/dkrasnovdev/heritage-api/ent/organizationtype"
	"github.com/dkrasnovdev/heritage-api/ent/predicate"
)

// OrganizationTypeQuery is the builder for querying OrganizationType entities.
type OrganizationTypeQuery struct {
	config
	ctx                    *QueryContext
	order                  []organizationtype.OrderOption
	inters                 []Interceptor
	predicates             []predicate.OrganizationType
	withOrganizations      *OrganizationQuery
	modifiers              []func(*sql.Selector)
	loadTotal              []func(context.Context, []*OrganizationType) error
	withNamedOrganizations map[string]*OrganizationQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OrganizationTypeQuery builder.
func (otq *OrganizationTypeQuery) Where(ps ...predicate.OrganizationType) *OrganizationTypeQuery {
	otq.predicates = append(otq.predicates, ps...)
	return otq
}

// Limit the number of records to be returned by this query.
func (otq *OrganizationTypeQuery) Limit(limit int) *OrganizationTypeQuery {
	otq.ctx.Limit = &limit
	return otq
}

// Offset to start from.
func (otq *OrganizationTypeQuery) Offset(offset int) *OrganizationTypeQuery {
	otq.ctx.Offset = &offset
	return otq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (otq *OrganizationTypeQuery) Unique(unique bool) *OrganizationTypeQuery {
	otq.ctx.Unique = &unique
	return otq
}

// Order specifies how the records should be ordered.
func (otq *OrganizationTypeQuery) Order(o ...organizationtype.OrderOption) *OrganizationTypeQuery {
	otq.order = append(otq.order, o...)
	return otq
}

// QueryOrganizations chains the current query on the "organizations" edge.
func (otq *OrganizationTypeQuery) QueryOrganizations() *OrganizationQuery {
	query := (&OrganizationClient{config: otq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := otq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := otq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(organizationtype.Table, organizationtype.FieldID, selector),
			sqlgraph.To(organization.Table, organization.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, organizationtype.OrganizationsTable, organizationtype.OrganizationsColumn),
		)
		fromU = sqlgraph.SetNeighbors(otq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OrganizationType entity from the query.
// Returns a *NotFoundError when no OrganizationType was found.
func (otq *OrganizationTypeQuery) First(ctx context.Context) (*OrganizationType, error) {
	nodes, err := otq.Limit(1).All(setContextOp(ctx, otq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{organizationtype.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (otq *OrganizationTypeQuery) FirstX(ctx context.Context) *OrganizationType {
	node, err := otq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OrganizationType ID from the query.
// Returns a *NotFoundError when no OrganizationType ID was found.
func (otq *OrganizationTypeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = otq.Limit(1).IDs(setContextOp(ctx, otq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{organizationtype.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (otq *OrganizationTypeQuery) FirstIDX(ctx context.Context) int {
	id, err := otq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OrganizationType entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OrganizationType entity is found.
// Returns a *NotFoundError when no OrganizationType entities are found.
func (otq *OrganizationTypeQuery) Only(ctx context.Context) (*OrganizationType, error) {
	nodes, err := otq.Limit(2).All(setContextOp(ctx, otq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{organizationtype.Label}
	default:
		return nil, &NotSingularError{organizationtype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (otq *OrganizationTypeQuery) OnlyX(ctx context.Context) *OrganizationType {
	node, err := otq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OrganizationType ID in the query.
// Returns a *NotSingularError when more than one OrganizationType ID is found.
// Returns a *NotFoundError when no entities are found.
func (otq *OrganizationTypeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = otq.Limit(2).IDs(setContextOp(ctx, otq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{organizationtype.Label}
	default:
		err = &NotSingularError{organizationtype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (otq *OrganizationTypeQuery) OnlyIDX(ctx context.Context) int {
	id, err := otq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OrganizationTypes.
func (otq *OrganizationTypeQuery) All(ctx context.Context) ([]*OrganizationType, error) {
	ctx = setContextOp(ctx, otq.ctx, "All")
	if err := otq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*OrganizationType, *OrganizationTypeQuery]()
	return withInterceptors[[]*OrganizationType](ctx, otq, qr, otq.inters)
}

// AllX is like All, but panics if an error occurs.
func (otq *OrganizationTypeQuery) AllX(ctx context.Context) []*OrganizationType {
	nodes, err := otq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OrganizationType IDs.
func (otq *OrganizationTypeQuery) IDs(ctx context.Context) (ids []int, err error) {
	if otq.ctx.Unique == nil && otq.path != nil {
		otq.Unique(true)
	}
	ctx = setContextOp(ctx, otq.ctx, "IDs")
	if err = otq.Select(organizationtype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (otq *OrganizationTypeQuery) IDsX(ctx context.Context) []int {
	ids, err := otq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (otq *OrganizationTypeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, otq.ctx, "Count")
	if err := otq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, otq, querierCount[*OrganizationTypeQuery](), otq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (otq *OrganizationTypeQuery) CountX(ctx context.Context) int {
	count, err := otq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (otq *OrganizationTypeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, otq.ctx, "Exist")
	switch _, err := otq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (otq *OrganizationTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := otq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OrganizationTypeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (otq *OrganizationTypeQuery) Clone() *OrganizationTypeQuery {
	if otq == nil {
		return nil
	}
	return &OrganizationTypeQuery{
		config:            otq.config,
		ctx:               otq.ctx.Clone(),
		order:             append([]organizationtype.OrderOption{}, otq.order...),
		inters:            append([]Interceptor{}, otq.inters...),
		predicates:        append([]predicate.OrganizationType{}, otq.predicates...),
		withOrganizations: otq.withOrganizations.Clone(),
		// clone intermediate query.
		sql:  otq.sql.Clone(),
		path: otq.path,
	}
}

// WithOrganizations tells the query-builder to eager-load the nodes that are connected to
// the "organizations" edge. The optional arguments are used to configure the query builder of the edge.
func (otq *OrganizationTypeQuery) WithOrganizations(opts ...func(*OrganizationQuery)) *OrganizationTypeQuery {
	query := (&OrganizationClient{config: otq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	otq.withOrganizations = query
	return otq
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
//	client.OrganizationType.Query().
//		GroupBy(organizationtype.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (otq *OrganizationTypeQuery) GroupBy(field string, fields ...string) *OrganizationTypeGroupBy {
	otq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &OrganizationTypeGroupBy{build: otq}
	grbuild.flds = &otq.ctx.Fields
	grbuild.label = organizationtype.Label
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
//	client.OrganizationType.Query().
//		Select(organizationtype.FieldCreatedAt).
//		Scan(ctx, &v)
func (otq *OrganizationTypeQuery) Select(fields ...string) *OrganizationTypeSelect {
	otq.ctx.Fields = append(otq.ctx.Fields, fields...)
	sbuild := &OrganizationTypeSelect{OrganizationTypeQuery: otq}
	sbuild.label = organizationtype.Label
	sbuild.flds, sbuild.scan = &otq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a OrganizationTypeSelect configured with the given aggregations.
func (otq *OrganizationTypeQuery) Aggregate(fns ...AggregateFunc) *OrganizationTypeSelect {
	return otq.Select().Aggregate(fns...)
}

func (otq *OrganizationTypeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range otq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, otq); err != nil {
				return err
			}
		}
	}
	for _, f := range otq.ctx.Fields {
		if !organizationtype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if otq.path != nil {
		prev, err := otq.path(ctx)
		if err != nil {
			return err
		}
		otq.sql = prev
	}
	if organizationtype.Policy == nil {
		return errors.New("ent: uninitialized organizationtype.Policy (forgotten import ent/runtime?)")
	}
	if err := organizationtype.Policy.EvalQuery(ctx, otq); err != nil {
		return err
	}
	return nil
}

func (otq *OrganizationTypeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OrganizationType, error) {
	var (
		nodes       = []*OrganizationType{}
		_spec       = otq.querySpec()
		loadedTypes = [1]bool{
			otq.withOrganizations != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*OrganizationType).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &OrganizationType{config: otq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(otq.modifiers) > 0 {
		_spec.Modifiers = otq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, otq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := otq.withOrganizations; query != nil {
		if err := otq.loadOrganizations(ctx, query, nodes,
			func(n *OrganizationType) { n.Edges.Organizations = []*Organization{} },
			func(n *OrganizationType, e *Organization) { n.Edges.Organizations = append(n.Edges.Organizations, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range otq.withNamedOrganizations {
		if err := otq.loadOrganizations(ctx, query, nodes,
			func(n *OrganizationType) { n.appendNamedOrganizations(name) },
			func(n *OrganizationType, e *Organization) { n.appendNamedOrganizations(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range otq.loadTotal {
		if err := otq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (otq *OrganizationTypeQuery) loadOrganizations(ctx context.Context, query *OrganizationQuery, nodes []*OrganizationType, init func(*OrganizationType), assign func(*OrganizationType, *Organization)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*OrganizationType)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(organizationtype.OrganizationsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.organization_type_organizations
		if fk == nil {
			return fmt.Errorf(`foreign-key "organization_type_organizations" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "organization_type_organizations" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (otq *OrganizationTypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := otq.querySpec()
	if len(otq.modifiers) > 0 {
		_spec.Modifiers = otq.modifiers
	}
	_spec.Node.Columns = otq.ctx.Fields
	if len(otq.ctx.Fields) > 0 {
		_spec.Unique = otq.ctx.Unique != nil && *otq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, otq.driver, _spec)
}

func (otq *OrganizationTypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(organizationtype.Table, organizationtype.Columns, sqlgraph.NewFieldSpec(organizationtype.FieldID, field.TypeInt))
	_spec.From = otq.sql
	if unique := otq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if otq.path != nil {
		_spec.Unique = true
	}
	if fields := otq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, organizationtype.FieldID)
		for i := range fields {
			if fields[i] != organizationtype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := otq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := otq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := otq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := otq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (otq *OrganizationTypeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(otq.driver.Dialect())
	t1 := builder.Table(organizationtype.Table)
	columns := otq.ctx.Fields
	if len(columns) == 0 {
		columns = organizationtype.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if otq.sql != nil {
		selector = otq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if otq.ctx.Unique != nil && *otq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range otq.predicates {
		p(selector)
	}
	for _, p := range otq.order {
		p(selector)
	}
	if offset := otq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := otq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedOrganizations tells the query-builder to eager-load the nodes that are connected to the "organizations"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (otq *OrganizationTypeQuery) WithNamedOrganizations(name string, opts ...func(*OrganizationQuery)) *OrganizationTypeQuery {
	query := (&OrganizationClient{config: otq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if otq.withNamedOrganizations == nil {
		otq.withNamedOrganizations = make(map[string]*OrganizationQuery)
	}
	otq.withNamedOrganizations[name] = query
	return otq
}

// OrganizationTypeGroupBy is the group-by builder for OrganizationType entities.
type OrganizationTypeGroupBy struct {
	selector
	build *OrganizationTypeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (otgb *OrganizationTypeGroupBy) Aggregate(fns ...AggregateFunc) *OrganizationTypeGroupBy {
	otgb.fns = append(otgb.fns, fns...)
	return otgb
}

// Scan applies the selector query and scans the result into the given value.
func (otgb *OrganizationTypeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, otgb.build.ctx, "GroupBy")
	if err := otgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrganizationTypeQuery, *OrganizationTypeGroupBy](ctx, otgb.build, otgb, otgb.build.inters, v)
}

func (otgb *OrganizationTypeGroupBy) sqlScan(ctx context.Context, root *OrganizationTypeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(otgb.fns))
	for _, fn := range otgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*otgb.flds)+len(otgb.fns))
		for _, f := range *otgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*otgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := otgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// OrganizationTypeSelect is the builder for selecting fields of OrganizationType entities.
type OrganizationTypeSelect struct {
	*OrganizationTypeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ots *OrganizationTypeSelect) Aggregate(fns ...AggregateFunc) *OrganizationTypeSelect {
	ots.fns = append(ots.fns, fns...)
	return ots
}

// Scan applies the selector query and scans the result into the given value.
func (ots *OrganizationTypeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ots.ctx, "Select")
	if err := ots.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrganizationTypeQuery, *OrganizationTypeSelect](ctx, ots.OrganizationTypeQuery, ots, ots.inters, v)
}

func (ots *OrganizationTypeSelect) sqlScan(ctx context.Context, root *OrganizationTypeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ots.fns))
	for _, fn := range ots.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ots.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ots.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
