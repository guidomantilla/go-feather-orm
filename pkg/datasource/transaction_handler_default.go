package datasource

import (
	"context"
	"fmt"

	feather_commons_log "github.com/guidomantilla/go-feather-commons/pkg/log"
)

type DefaultTransactionHandler struct {
	Datasource Datasource
}

func NewTransactionHandler(datasource Datasource) *DefaultTransactionHandler {

	if datasource == nil {
		feather_commons_log.Fatal("starting up - error setting up transactionHandler: datasource is nil")
		return nil
	}

	return &DefaultTransactionHandler{
		Datasource: datasource,
	}
}

func (handler *DefaultTransactionHandler) HandleTransaction(ctx context.Context, fn TransactionHandlerFunction) error {

	db, err := handler.Datasource.GetDatabase()
	if err != nil {
		feather_commons_log.Error(err.Error())
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		feather_commons_log.Error(err.Error())
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			feather_commons_log.Error(fmt.Sprintf("recovering from panic: %v", p))
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
		feather_commons_log.Error(err.Error())
	}
}
