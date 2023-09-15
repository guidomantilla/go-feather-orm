package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"strings"

	feather_sql "github.com/guidomantilla/go-feather-sql/pkg/sql"
	feather_sql_transaction "github.com/guidomantilla/go-feather-sql/pkg/transaction"
)

const (
	ErrorClosingStatement = "Error closing the statement"
	ErrorClosingResultSet = "Error closing the result set"
)

type Function func(statement *sql.Stmt) error

type ReadFunction func(rows *sql.Rows) error

func WriteContext(ctx context.Context, sqlStatement string, args ...any) (*int64, error) {

	var ok bool
	var driverName feather_sql.DriverName
	if driverName, ok = ctx.Value(feather_sql.DriverNameContext{}).(feather_sql.DriverName); !ok {
		return nil, ErrWriteContextFailed(errors.New(sqlStatement), errors.New("driver name not found in context"))
	}

	var err error
	var serial int64
	err = Context(ctx, sqlStatement, func(statement *sql.Stmt) error {

		var result sql.Result
		if result, err = statement.Exec(args...); err != nil {
			return err
		}

		if strings.Index(strings.ToLower(sqlStatement), "insert") == 0 && driverName != feather_sql.OracleDriverName {
			if serial, err = result.LastInsertId(); err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return nil, ErrWriteContextFailed(errors.New(sqlStatement), err)
	}

	return &serial, nil
}

func ReadContext(ctx context.Context, sqlStatement string, fn ReadFunction) error {

	var err error
	err = Context(ctx, sqlStatement, func(statement *sql.Stmt) error {

		var rows *sql.Rows
		if rows, err = statement.Query(); err != nil {
			return err
		}
		defer CloseResultSet(rows)

		if err = fn(rows); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return ErrReadContextFailed(errors.New(sqlStatement), err)
	}

	return nil
}

func ReadRowContext(ctx context.Context, sqlStatement string, key any, dest ...any) error {

	var err error
	err = Context(ctx, sqlStatement, func(statement *sql.Stmt) error {

		row := statement.QueryRow(key)
		if err = row.Scan(dest...); err != nil {
			if err.Error() == "db_column: no rows in result set" {
				return fmt.Errorf("row with key %v not found", key)
			}
			return err
		}
		return nil
	})
	if err != nil {
		return ErrReadRowContextFailed(errors.New(sqlStatement), err)
	}

	return nil
}

//

func Context(ctx context.Context, sqlStatement string, fn Function) error {

	var ok bool
	var tx *sql.Tx
	if tx, ok = ctx.Value(feather_sql_transaction.TransactionCtxKey{}).(*sql.Tx); !ok {
		return ErrContextFailed(errors.New(sqlStatement), errors.New("transaction not found in context"))
	}

	var err error
	var statement *sql.Stmt
	if statement, err = tx.Prepare(sqlStatement); err != nil {
		return ErrContextFailed(errors.New(sqlStatement), err)
	}
	defer CloseStatement(statement)

	if err = fn(statement); err != nil {
		return ErrContextFailed(errors.New(sqlStatement), err)
	}
	return nil
}

func CloseStatement(statement *sql.Stmt) {
	if err := statement.Close(); err != nil {
		slog.Error(ErrorClosingStatement)
	}
}

func CloseResultSet(rows *sql.Rows) {
	if err := rows.Close(); err != nil {
		slog.Error(ErrorClosingResultSet)
	}
}
