package transaction

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"

	feather_sql_datasource "github.com/guidomantilla/go-feather-sql/pkg/datasource"
)

var (
	_ TransactionHandler = (*DefaultDBTransactionHandler)(nil)
)

type TransactionCtxKey struct{}

type TransactionHandlerFunction func(ctx context.Context, tx *sql.Tx) error

type TransactionHandler interface {
	HandleTransaction(ctx context.Context, fn TransactionHandlerFunction) error
}

type DefaultDBTransactionHandler struct {
	Datasource feather_sql_datasource.Datasource
}

func NewTransactionHandler(datasource feather_sql_datasource.Datasource) *DefaultDBTransactionHandler {

	if datasource == nil {
		slog.Error("starting up - error setting up transactionHandler: datasource is nil")
		return nil
	}

	return &DefaultDBTransactionHandler{
		Datasource: datasource,
	}
}

func (handler *DefaultDBTransactionHandler) HandleTransaction(ctx context.Context, fn TransactionHandlerFunction) error {

	db, err := handler.Datasource.GetDatabase()
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			slog.Error(fmt.Sprintf("recovering from panic: %v", p))
			handleError(tx.Rollback())
		} else if err != nil {
			handleError(tx.Rollback())
		} else {
			handleError(tx.Commit())
		}
	}()

	txCtx := context.WithValue(ctx, TransactionCtxKey{}, tx)
	err = fn(txCtx, tx)
	return err
}

//

func handleError(err error) {
	if err != nil {
		slog.Error(err.Error())
	}
}
