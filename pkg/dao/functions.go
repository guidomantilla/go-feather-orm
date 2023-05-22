package dao

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"go.uber.org/zap"

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

	driverName := ctx.Value(feather_sql.DriverNameContext{}).(feather_sql.DriverName)

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
		return nil, err
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
		defer closeResultSet(rows)

		if err = fn(rows); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func ReadRowContext(ctx context.Context, sqlStatement string, key any, dest ...any) error {

	var err error
	err = Context(ctx, sqlStatement, func(statement *sql.Stmt) error {

		row := statement.QueryRow(key)
		if err = row.Scan(dest...); err != nil {
			if err.Error() == "sql: no rows in result set" {
				return fmt.Errorf("row with key %v not found", key)
			}
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

//

func Context(ctx context.Context, sqlStatement string, fn Function) error {

	var err error
	var statement *sql.Stmt
	var tx = ctx.Value(feather_sql_transaction.RelationalTransactionCtxKey{}).(*sql.Tx)
	if statement, err = tx.Prepare(sqlStatement); err != nil {
		return err
	}
	defer closeStatement(statement)

	if err = fn(statement); err != nil {
		return err
	}
	return nil
}

func closeStatement(statement *sql.Stmt) {
	if err := statement.Close(); err != nil {
		zap.L().Error(ErrorClosingStatement)
	}
}

func closeResultSet(rows *sql.Rows) {
	if err := rows.Close(); err != nil {
		zap.L().Error(ErrorClosingResultSet)
	}
}
