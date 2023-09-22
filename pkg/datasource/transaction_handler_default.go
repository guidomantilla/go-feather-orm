package datasource

import (
	"context"
	"fmt"
	"log/slog"
)

type DefaultTransactionHandler struct {
	Datasource Datasource
}

func NewTransactionHandler(datasource Datasource) *DefaultTransactionHandler {

	if datasource == nil {
		slog.Error("starting up - error setting up transactionHandler: datasource is nil")
		return nil
	}

	return &DefaultTransactionHandler{
		Datasource: datasource,
	}
}

func (handler *DefaultTransactionHandler) HandleTransaction(ctx context.Context, fn TransactionHandlerFunction) error {

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
