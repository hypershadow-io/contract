package qb

import "github.com/hypershadow-io/contract/db"

type (
	// QueryBuilder defines the interface for building SQL queries.
	// It is expected to produce an object implementing the Query interface,
	// typically by assembling SQL strings and associated arguments dynamically.
	QueryBuilder interface {
		// Select returns a new SelectQuery
		Select() SelectQuery

		// Insert returns a new InsertQuery with the given table name
		Insert(table string) InsertQuery

		// Replace returns a new InsertQuery with the statement keyword set to
		// "REPLACE" and with the given table name
		Replace(table string) InsertQuery

		// Update returns a new UpdateQuery with the given table name
		Update(table string) UpdateQuery

		// Delete returns a new DeleteQuery with the given table name
		Delete(table string) DeleteQuery

		// Placeholders returns a string with count ? placeholders joined with commas
		Placeholders(count int) string

		// Sql builds an expression from a SQL fragment and arguments
		//
		// Example:
		//  Expr("FROM_UNIXTIME(?)", t)
		Sql(sql string, args ...any) db.Query

		// Case returns a new CaseQuery "what" represents case value
		Case(what ...any) CaseQuery

		// Concat builds an expression by concatenating strings and other expressions
		//
		// Example:
		//
		//  name_expr := qb.Sql("CONCAT(?, ' ', ?)", firstName, lastName)
		//  Concat("COALESCE(full_name,", name_expr, ")")
		Concat(parts ...any) db.Query

		// Alias allows to define alias for column in SelectQuery. Useful when column is
		// defined as complex expression like IF or CASE
		//
		// Example:
		// .Column(qb.Alias(caseStmt, "case_column"))
		Alias(expr db.Query, alias string) db.Query

		// Eq is syntactic sugar for use with Where/Having/Set methods
		//
		// Example:
		//  .Where(qb.Eq(map[string]any{"id": 1})) => "id == 1"
		Eq(v map[string]any) db.Query

		// NotEq is syntactic sugar for use with Where/Having/Set methods
		//
		// Example:
		//  .Where(qb.NotEq(map[string]any{"id": 1})) => "id != 1"
		NotEq(v map[string]any) db.Query

		// Like is syntactic sugar for use with LIKE conditions
		//
		// Example:
		//  .Where(qb.Like(map[string]any{"name": "%irrel"})) => "name LIKE '%irrel'"
		Like(v map[string]any) db.Query

		// NotLike is syntactic sugar for use with NOT LIKE conditions
		//
		// Example:
		//  .Where(qb.NotLike(map[string]any{"name": "%irrel"})) => "name NOT LIKE '%irrel'"
		NotLike(v map[string]any) db.Query

		// ILike is syntactic sugar for use with ILIKE conditions
		//
		// Example:
		//  .Where(qb.ILike(map[string]any{"name": "%irrel"})) => "name ILIKE '%irrel'"
		ILike(v map[string]any) db.Query

		// NotILike is syntactic sugar for use with NOT ILIKE conditions
		//
		// Example:
		//  .Where(qb.NotILike(map[string]any{"name": "%irrel"})) => "name NOT ILIKE '%irrel'"
		NotILike(v map[string]any) db.Query

		// Lt is syntactic sugar for use with Where/Having/Set methods
		//
		// Example:
		//  .Where(qb.Lt(map[string]any{"id": 1})) => "id < 1"
		Lt(v map[string]any) db.Query

		// LtOrEq is syntactic sugar for use with Where/Having/Set methods
		//
		// Example:
		//  .Where(qb.LtOrEq(map[string]any{"id": 1})) => "id <= 1"
		LtOrEq(v map[string]any) db.Query

		// Gt is syntactic sugar for use with Where/Having/Set methods
		//
		// Example:
		//  .Where(qb.Gt(map[string]any{"id": 1})) => "id > 1"
		Gt(v map[string]any) db.Query

		// GtOrEq is syntactic sugar for use with Where/Having/Set methods
		//
		// Example:
		//  .Where(qb.GtOrEq(map[string]any{"id": 1})) => "id >= 1"
		GtOrEq(v map[string]any) db.Query

		// And conjunction Query
		And(args ...db.Query) db.Query

		// Or conjunction Query
		Or(args ...db.Query) db.Query
	}

	// SelectQuery defines the interface for building SELECT SQL queries.
	SelectQuery interface {
		db.Query
		SetError[SelectQuery]

		// Prefix adds an expression to the beginning of the query
		Prefix(sql string, args ...any) SelectQuery

		// PrefixQuery adds an expression to the very beginning of the query
		PrefixQuery(query db.Query) SelectQuery

		// Distinct adds a DISTINCT clause to the query
		Distinct() SelectQuery

		// Columns adds result columns to the query
		Columns(columns ...string) SelectQuery

		// Column adds a result column to the query.
		// Unlike Columns, Column accepts args which will be bound to placeholders in
		// the columns string.
		//
		// Example:
		//  qb.Column("IF(col IN ("+qb.Placeholders(3)+"), 1, 0) as col", 1, 2, 3)
		Column(column any, args ...any) SelectQuery

		// RemoveColumns remove all columns from query.
		// Must add a new column with Column or Columns methods, otherwise
		// return a error
		RemoveColumns() SelectQuery

		// From sets the FROM clause of the query
		From(table string, alias ...string) SelectQuery

		// FromSelect sets a subquery into the FROM clause of the query
		FromSelect(from SelectQuery, alias string) SelectQuery

		// Join adds a JOIN clause to the query
		Join(join string, rest ...any) SelectQuery

		// LeftJoin adds a LEFT JOIN clause to the query
		LeftJoin(join string, rest ...any) SelectQuery

		// RightJoin adds a RIGHT JOIN clause to the query
		RightJoin(join string, rest ...any) SelectQuery

		// InnerJoin adds a INNER JOIN clause to the query
		InnerJoin(join string, rest ...any) SelectQuery

		// CrossJoin adds a CROSS JOIN clause to the query
		CrossJoin(join string, rest ...any) SelectQuery

		// AndWhere adds an expression to the WHERE clause of the query,
		// combining it with existing expressions using AND in the generated SQL
		AndWhere(query db.Query) SelectQuery

		// OrWhere adds an expression to the WHERE clause of the query,
		// combining it with existing expressions using OR in the generated SQL
		OrWhere(query db.Query) SelectQuery

		// AndHaving adds an expression to the HAVING clause of the query,
		// combining it with existing expressions using AND in the generated SQL
		AndHaving(query db.Query) SelectQuery

		// OrHaving adds an expression to the HAVING clause of the query,
		// combining it with existing expressions using OR in the generated SQL
		OrHaving(query db.Query) SelectQuery

		// GroupBy adds GROUP BY expressions to the query
		GroupBy(groupBys ...string) SelectQuery

		// OrderByBefore adds ORDER BY expressions to the beginning of the ORDER BY clause of the query
		OrderByBefore(orderBys ...string) SelectQuery

		// OrderByAfter adds ORDER BY expressions to the end of the ORDER BY clause of the query
		OrderByAfter(orderBys ...string) SelectQuery

		// Limit sets a LIMIT clause on the query
		Limit(limit uint64) SelectQuery

		// Offset sets a OFFSET clause on the query
		Offset(offset uint64) SelectQuery

		// Suffix adds an expression to the end of the query
		Suffix(sql string, args ...any) SelectQuery

		// SuffixQuery adds an expression to the end of the query
		SuffixQuery(query db.Query) SelectQuery
	}

	// InsertQuery defines the interface for building INSERT SQL queries.
	InsertQuery interface {
		db.Query
		SetError[InsertQuery]

		// Prefix adds an expression to the beginning of the query
		Prefix(sql string, args ...any) InsertQuery

		// PrefixQuery adds an expression to the very beginning of the query
		PrefixQuery(query db.Query) InsertQuery

		// Into sets the INTO clause of the query
		Into(table string) InsertQuery

		// Columns adds insert columns to the query
		Columns(columns ...string) InsertQuery

		// Values adds a single row's values to the query
		Values(values ...any) InsertQuery

		// Suffix adds an expression to the end of the query
		Suffix(sql string, args ...any) InsertQuery

		// SuffixQuery adds an expression to the end of the query
		SuffixQuery(query db.Query) InsertQuery

		// SetMap set columns and values for insert builder from a map of column name and value
		// note that it will reset all previous columns and values was set if any
		SetMap(clauses map[string]any) InsertQuery

		// Select set Select clause for insert query
		// If Values and Select are used, then Select has higher priority
		Select(sb SelectQuery) InsertQuery
	}

	// UpdateQuery defines the interface for building UPDATE SQL queries.
	UpdateQuery interface {
		db.Query
		SetError[UpdateQuery]

		// Prefix adds an expression to the beginning of the query
		Prefix(sql string, args ...any) UpdateQuery

		// PrefixQuery adds an expression to the very beginning of the query
		PrefixQuery(query db.Query) UpdateQuery

		// Table sets the table to be updated
		Table(table string) UpdateQuery

		// Set adds SET clauses to the query
		Set(column string, value any) UpdateQuery

		// SetMap is a convenience method which calls .Set for each key/value pair in clauses
		SetMap(clauses map[string]any) UpdateQuery

		// From adds FROM clause to the query
		// FROM is valid construct in postgresql only.
		From(from string) UpdateQuery

		// FromSelect sets a subquery into the FROM clause of the query
		FromSelect(from SelectQuery, alias string) UpdateQuery

		// AndWhere adds an expression to the WHERE clause of the query,
		// combining it with existing expressions using AND in the generated SQL
		AndWhere(query db.Query) UpdateQuery

		// OrWhere adds an expression to the WHERE clause of the query,
		// combining it with existing expressions using OR in the generated SQL
		OrWhere(query db.Query) UpdateQuery

		// OrderByBefore adds ORDER BY expressions to the beginning of the ORDER BY clause of the query
		OrderByBefore(orderBys ...string) UpdateQuery

		// OrderByAfter adds ORDER BY expressions to the end of the ORDER BY clause of the query
		OrderByAfter(orderBys ...string) UpdateQuery

		// Limit sets a LIMIT clause on the query
		Limit(limit uint64) UpdateQuery

		// Offset sets a OFFSET clause on the query
		Offset(offset uint64) UpdateQuery

		// Suffix adds an expression to the end of the query
		Suffix(sql string, args ...any) UpdateQuery

		// SuffixQuery adds an expression to the end of the query
		SuffixQuery(query db.Query) UpdateQuery
	}

	// DeleteQuery defines the interface for building DELETE SQL queries.
	DeleteQuery interface {
		db.Query
		SetError[DeleteQuery]

		// Prefix adds an expression to the beginning of the query
		Prefix(sql string, args ...any) DeleteQuery

		// PrefixQuery adds an expression to the very beginning of the query
		PrefixQuery(query db.Query) DeleteQuery

		// From sets the table to be deleted from
		From(table string) DeleteQuery

		// AndWhere adds an expression to the WHERE clause of the query,
		// combining it with existing expressions using AND in the generated SQL
		AndWhere(query db.Query) DeleteQuery

		// OrWhere adds an expression to the WHERE clause of the query,
		// combining it with existing expressions using OR in the generated SQL
		OrWhere(query db.Query) DeleteQuery

		// OrderByBefore adds ORDER BY expressions to the beginning of the ORDER BY clause of the query
		OrderByBefore(orderBys ...string) DeleteQuery

		// OrderByAfter adds ORDER BY expressions to the end of the ORDER BY clause of the query
		OrderByAfter(orderBys ...string) DeleteQuery

		// Limit sets a LIMIT clause on the query
		Limit(limit uint64) DeleteQuery

		// Offset sets a OFFSET clause on the query
		Offset(offset uint64) DeleteQuery

		// Suffix adds an expression to the end of the query
		Suffix(sql string, args ...any) DeleteQuery

		// SuffixQuery adds an expression to the end of the query
		SuffixQuery(query db.Query) DeleteQuery
	}

	// CaseQuery defines the interface for building SQL CASE expressions.
	CaseQuery interface {
		db.Query
		SetError[CaseQuery]

		// When adds "WHEN ... THEN ..." part to CASE construct
		When(when any, then any) CaseQuery

		// Else what sets optional "ELSE ..." part for CASE construct
		Else(expr any) CaseQuery
	}

	// SetError provides a mechanism to attach an error to a builder.
	// If an error is set, subsequent calls to ToSql will return this error
	// instead of generating SQL. This is useful for aborting query construction
	// due to validation failures or hook logic.
	SetError[T any] interface {
		// SetError marks the builder as failed and attaches the given error
		// Any subsequent call to ToSql will return this error instead of generating SQL
		// Useful for aborting query construction from hook or validation logic
		SetError(err error) T
	}
)
