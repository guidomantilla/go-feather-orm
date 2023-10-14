package dao

import (
	"context"
	"database/sql"
	"errors"

	feather_commons_log "github.com/guidomantilla/go-feather-commons/pkg/log"
	"github.com/jmoiron/sqlx"

	feather_sql_datasource "github.com/guidomantilla/go-feather-sql/pkg/datasource"
)

const (
	ErrorClosingStatement = "Error closing the statement"
	ErrorClosingResultSet = "Error closing the result set"
)

type QueryFunction func(statement *sqlx.NamedStmt) error

type MutateFunction func(result sql.Result) error

//

func CloseStatement(statement *sqlx.NamedStmt) {
	if err := statement.Close(); err != nil {
		feather_commons_log.Error(ErrorClosingStatement)
	}
}

func CloseResultSet(rows *sqlx.Rows) {
	if err := rows.Close(); err != nil {
		feather_commons_log.Error(ErrorClosingResultSet)
	}
}

func QueryContext(ctx context.Context, sqlStatement string, fn QueryFunction) error {

	tx, ok := ctx.Value(feather_sql_datasource.TransactionCtxKey{}).(*sqlx.Tx)
	if !ok {
		return ErrContextFailed(errors.New(sqlStatement), errors.New("transaction not found in context"))
	}

	statement, err := tx.PrepareNamed(sqlStatement)
	if err != nil {
		return ErrContextFailed(errors.New(sqlStatement), err)
	}
	defer CloseStatement(statement)

	if err = fn(statement); err != nil {
		return ErrContextFailed(errors.New(sqlStatement), err)
	}
	return nil
}

func MutateContext[T any](ctx context.Context, sqlStatement string, target *T, fn MutateFunction) error {

	tx, ok := ctx.Value(feather_sql_datasource.TransactionCtxKey{}).(*sqlx.Tx)
	if !ok {
		return ErrContextFailed(errors.New(sqlStatement), errors.New("transaction not found in context"))
	}

	result, err := tx.NamedExecContext(ctx, sqlStatement, target)
	if err != nil {
		return ErrContextFailed(errors.New(sqlStatement), err)
	}

	if err = fn(result); err != nil {
		return ErrContextFailed(errors.New(sqlStatement), err)
	}
	return nil
}

//

type Entity interface {
	any
}

func MutateOne[T Entity](ctx context.Context, sqlStatement string, target *T) error {
	return MutateContext[T](ctx, sqlStatement, target, func(result sql.Result) error {
		count, err := result.RowsAffected()
		if err != nil {
			return err
		}
		if count > int64(1) {
			return errors.New("more than one affected")
		}
		return nil
	})
}

func QueryOne[T Entity](ctx context.Context, sqlStatement string, target *T) error {
	return QueryContext(ctx, sqlStatement, func(statement *sqlx.NamedStmt) error {
		return statement.GetContext(ctx, target, target)
	})
}

func QueryMany[T Entity](ctx context.Context, sqlStatement string, args *T) ([]T, error) {

	var dest []T
	err := QueryContext(ctx, sqlStatement, func(statement *sqlx.NamedStmt) error {
		return statement.SelectContext(ctx, &dest, args)
	})

	if err != nil {
		return nil, err
	}

	return dest, nil
}

func Exists[T Entity](ctx context.Context, sqlStatement string, target *T) bool {
	if err := QueryOne[T](ctx, sqlStatement, target); err != nil {
		return false
	}
	return true
}
