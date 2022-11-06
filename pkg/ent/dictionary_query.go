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
	"github.com/suyuan32/simple-admin-core/pkg/ent/dictionary"
	"github.com/suyuan32/simple-admin-core/pkg/ent/dictionarydetail"
	"github.com/suyuan32/simple-admin-core/pkg/ent/predicate"
)

// DictionaryQuery is the builder for querying Dictionary entities.
type DictionaryQuery struct {
	config
	limit                 *int
	offset                *int
	unique                *bool
	order                 []OrderFunc
	fields                []string
	predicates            []predicate.Dictionary
	withDictionaryDetails *DictionaryDetailQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DictionaryQuery builder.
func (dq *DictionaryQuery) Where(ps ...predicate.Dictionary) *DictionaryQuery {
	dq.predicates = append(dq.predicates, ps...)
	return dq
}

// Limit adds a limit step to the query.
func (dq *DictionaryQuery) Limit(limit int) *DictionaryQuery {
	dq.limit = &limit
	return dq
}

// Offset adds an offset step to the query.
func (dq *DictionaryQuery) Offset(offset int) *DictionaryQuery {
	dq.offset = &offset
	return dq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dq *DictionaryQuery) Unique(unique bool) *DictionaryQuery {
	dq.unique = &unique
	return dq
}

// Order adds an order step to the query.
func (dq *DictionaryQuery) Order(o ...OrderFunc) *DictionaryQuery {
	dq.order = append(dq.order, o...)
	return dq
}

// QueryDictionaryDetails chains the current query on the "dictionary_details" edge.
func (dq *DictionaryQuery) QueryDictionaryDetails() *DictionaryDetailQuery {
	query := &DictionaryDetailQuery{config: dq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(dictionary.Table, dictionary.FieldID, selector),
			sqlgraph.To(dictionarydetail.Table, dictionarydetail.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, dictionary.DictionaryDetailsTable, dictionary.DictionaryDetailsColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Dictionary entity from the query.
// Returns a *NotFoundError when no Dictionary was found.
func (dq *DictionaryQuery) First(ctx context.Context) (*Dictionary, error) {
	nodes, err := dq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{dictionary.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dq *DictionaryQuery) FirstX(ctx context.Context) *Dictionary {
	node, err := dq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Dictionary ID from the query.
// Returns a *NotFoundError when no Dictionary ID was found.
func (dq *DictionaryQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = dq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{dictionary.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dq *DictionaryQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := dq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Dictionary entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Dictionary entity is found.
// Returns a *NotFoundError when no Dictionary entities are found.
func (dq *DictionaryQuery) Only(ctx context.Context) (*Dictionary, error) {
	nodes, err := dq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{dictionary.Label}
	default:
		return nil, &NotSingularError{dictionary.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dq *DictionaryQuery) OnlyX(ctx context.Context) *Dictionary {
	node, err := dq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Dictionary ID in the query.
// Returns a *NotSingularError when more than one Dictionary ID is found.
// Returns a *NotFoundError when no entities are found.
func (dq *DictionaryQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = dq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{dictionary.Label}
	default:
		err = &NotSingularError{dictionary.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dq *DictionaryQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := dq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Dictionaries.
func (dq *DictionaryQuery) All(ctx context.Context) ([]*Dictionary, error) {
	if err := dq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return dq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (dq *DictionaryQuery) AllX(ctx context.Context) []*Dictionary {
	nodes, err := dq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Dictionary IDs.
func (dq *DictionaryQuery) IDs(ctx context.Context) ([]uint64, error) {
	var ids []uint64
	if err := dq.Select(dictionary.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dq *DictionaryQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := dq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dq *DictionaryQuery) Count(ctx context.Context) (int, error) {
	if err := dq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return dq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (dq *DictionaryQuery) CountX(ctx context.Context) int {
	count, err := dq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dq *DictionaryQuery) Exist(ctx context.Context) (bool, error) {
	if err := dq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return dq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (dq *DictionaryQuery) ExistX(ctx context.Context) bool {
	exist, err := dq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DictionaryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dq *DictionaryQuery) Clone() *DictionaryQuery {
	if dq == nil {
		return nil
	}
	return &DictionaryQuery{
		config:                dq.config,
		limit:                 dq.limit,
		offset:                dq.offset,
		order:                 append([]OrderFunc{}, dq.order...),
		predicates:            append([]predicate.Dictionary{}, dq.predicates...),
		withDictionaryDetails: dq.withDictionaryDetails.Clone(),
		// clone intermediate query.
		sql:    dq.sql.Clone(),
		path:   dq.path,
		unique: dq.unique,
	}
}

// WithDictionaryDetails tells the query-builder to eager-load the nodes that are connected to
// the "dictionary_details" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DictionaryQuery) WithDictionaryDetails(opts ...func(*DictionaryDetailQuery)) *DictionaryQuery {
	query := &DictionaryDetailQuery{config: dq.config}
	for _, opt := range opts {
		opt(query)
	}
	dq.withDictionaryDetails = query
	return dq
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
//	client.Dictionary.Query().
//		GroupBy(dictionary.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dq *DictionaryQuery) GroupBy(field string, fields ...string) *DictionaryGroupBy {
	grbuild := &DictionaryGroupBy{config: dq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return dq.sqlQuery(ctx), nil
	}
	grbuild.label = dictionary.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
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
//	client.Dictionary.Query().
//		Select(dictionary.FieldCreatedAt).
//		Scan(ctx, &v)
func (dq *DictionaryQuery) Select(fields ...string) *DictionarySelect {
	dq.fields = append(dq.fields, fields...)
	selbuild := &DictionarySelect{DictionaryQuery: dq}
	selbuild.label = dictionary.Label
	selbuild.flds, selbuild.scan = &dq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a DictionarySelect configured with the given aggregations.
func (dq *DictionaryQuery) Aggregate(fns ...AggregateFunc) *DictionarySelect {
	return dq.Select().Aggregate(fns...)
}

func (dq *DictionaryQuery) prepareQuery(ctx context.Context) error {
	for _, f := range dq.fields {
		if !dictionary.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dq.path != nil {
		prev, err := dq.path(ctx)
		if err != nil {
			return err
		}
		dq.sql = prev
	}
	return nil
}

func (dq *DictionaryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Dictionary, error) {
	var (
		nodes       = []*Dictionary{}
		_spec       = dq.querySpec()
		loadedTypes = [1]bool{
			dq.withDictionaryDetails != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Dictionary).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Dictionary{config: dq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := dq.withDictionaryDetails; query != nil {
		if err := dq.loadDictionaryDetails(ctx, query, nodes,
			func(n *Dictionary) { n.Edges.DictionaryDetails = []*DictionaryDetail{} },
			func(n *Dictionary, e *DictionaryDetail) {
				n.Edges.DictionaryDetails = append(n.Edges.DictionaryDetails, e)
			}); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (dq *DictionaryQuery) loadDictionaryDetails(ctx context.Context, query *DictionaryDetailQuery, nodes []*Dictionary, init func(*Dictionary), assign func(*Dictionary, *DictionaryDetail)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uint64]*Dictionary)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.DictionaryDetail(func(s *sql.Selector) {
		s.Where(sql.InValues(dictionary.DictionaryDetailsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.dictionary_dictionary_details
		if fk == nil {
			return fmt.Errorf(`foreign-key "dictionary_dictionary_details" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "dictionary_dictionary_details" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (dq *DictionaryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dq.querySpec()
	_spec.Node.Columns = dq.fields
	if len(dq.fields) > 0 {
		_spec.Unique = dq.unique != nil && *dq.unique
	}
	return sqlgraph.CountNodes(ctx, dq.driver, _spec)
}

func (dq *DictionaryQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := dq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (dq *DictionaryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dictionary.Table,
			Columns: dictionary.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: dictionary.FieldID,
			},
		},
		From:   dq.sql,
		Unique: true,
	}
	if unique := dq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := dq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dictionary.FieldID)
		for i := range fields {
			if fields[i] != dictionary.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dq *DictionaryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dq.driver.Dialect())
	t1 := builder.Table(dictionary.Table)
	columns := dq.fields
	if len(columns) == 0 {
		columns = dictionary.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dq.sql != nil {
		selector = dq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dq.unique != nil && *dq.unique {
		selector.Distinct()
	}
	for _, p := range dq.predicates {
		p(selector)
	}
	for _, p := range dq.order {
		p(selector)
	}
	if offset := dq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DictionaryGroupBy is the group-by builder for Dictionary entities.
type DictionaryGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dgb *DictionaryGroupBy) Aggregate(fns ...AggregateFunc) *DictionaryGroupBy {
	dgb.fns = append(dgb.fns, fns...)
	return dgb
}

// Scan applies the group-by query and scans the result into the given value.
func (dgb *DictionaryGroupBy) Scan(ctx context.Context, v any) error {
	query, err := dgb.path(ctx)
	if err != nil {
		return err
	}
	dgb.sql = query
	return dgb.sqlScan(ctx, v)
}

func (dgb *DictionaryGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range dgb.fields {
		if !dictionary.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := dgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (dgb *DictionaryGroupBy) sqlQuery() *sql.Selector {
	selector := dgb.sql.Select()
	aggregation := make([]string, 0, len(dgb.fns))
	for _, fn := range dgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(dgb.fields)+len(dgb.fns))
		for _, f := range dgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(dgb.fields...)...)
}

// DictionarySelect is the builder for selecting fields of Dictionary entities.
type DictionarySelect struct {
	*DictionaryQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ds *DictionarySelect) Aggregate(fns ...AggregateFunc) *DictionarySelect {
	ds.fns = append(ds.fns, fns...)
	return ds
}

// Scan applies the selector query and scans the result into the given value.
func (ds *DictionarySelect) Scan(ctx context.Context, v any) error {
	if err := ds.prepareQuery(ctx); err != nil {
		return err
	}
	ds.sql = ds.DictionaryQuery.sqlQuery(ctx)
	return ds.sqlScan(ctx, v)
}

func (ds *DictionarySelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(ds.fns))
	for _, fn := range ds.fns {
		aggregation = append(aggregation, fn(ds.sql))
	}
	switch n := len(*ds.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		ds.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		ds.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := ds.sql.Query()
	if err := ds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
