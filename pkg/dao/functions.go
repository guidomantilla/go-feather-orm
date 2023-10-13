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

type QueryFunction func(statement *sqlx.Stmt) error

type MutateFunction func(result sql.Result) error

//

func CloseStatement(statement *sqlx.Stmt) {
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

	statement, err := tx.Preparex(sqlStatement)
	if err != nil {
		return ErrContextFailed(errors.New(sqlStatement), err)
	}
	defer CloseStatement(statement)

	if err = fn(statement); err != nil {
		return ErrContextFailed(errors.New(sqlStatement), err)
	}
	return nil
}

func MutateContext(ctx context.Context, sqlStatement string, target any, fn MutateFunction) error {

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

func MutateOne(ctx context.Context, sqlStatement string, target any) error {
	return MutateContext(ctx, sqlStatement, target, func(result sql.Result) error {
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

func QueryOne(ctx context.Context, sqlStatement string, dest any, args ...any) error {
	return QueryContext(ctx, sqlStatement, func(statement *sqlx.Stmt) error {
		return statement.GetContext(ctx, dest, args...)
	})
}

func QueryMany(ctx context.Context, sqlStatement string, dest any, args ...any) error {
	return QueryContext(ctx, sqlStatement, func(statement *sqlx.Stmt) error {
		return statement.SelectContext(ctx, dest, args...)
	})
}

func Exists(ctx context.Context, sqlStatement string, dest any, args ...any) bool {
	if err := QueryOne(ctx, sqlStatement, dest, args); err != nil {
		return false
	}
	return true
}
