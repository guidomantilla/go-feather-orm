package datasource

import (
	"context"
	"fmt"

	feather_commons_log "github.com/guidomantilla/go-feather-commons/pkg/log"

	feather_sql "github.com/guidomantilla/go-feather-sql/pkg/sql"
)

type DefaultTransactionHandler struct {
	datasource        Datasource
	datasourceContext DatasourceContext
}

func NewTransactionHandler(datasourceContext DatasourceContext, datasource Datasource) *DefaultTransactionHandler {

	if datasourceContext == nil {
		feather_commons_log.Fatal("starting up - error setting up transactionHandler: datasourceContext is nil")
	}

	if datasource == nil {
		feather_commons_log.Fatal("starting up - error setting up transactionHandler: datasource is nil")
		return nil
	}

	return &DefaultTransactionHandler{
		datasource:        datasource,
		datasourceContext: datasourceContext,
	}
}

func (handler *DefaultTransactionHandler) HandleTransaction(ctx context.Context, fn TransactionHandlerFunction) error {

	dbx, err := handler.datasource.GetDatabase()
	if err != nil {
		feather_commons_log.Error(err.Error())
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			feather_commons_log.Error(fmt.Sprintf("recovering from panic: %v", p))
		}
	}()

	tx := dbx.MustBegin()
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

	driverNameCtx := context.WithValue(ctx, feather_sql.DriverNameCtxKey{}, handler.datasourceContext.GetDriverName())
	txCtx := context.WithValue(driverNameCtx, TransactionCtxKey{}, tx)
	err = fn(txCtx, tx)
	return err
}

//

func handleError(err error) {
	if err != nil {
		feather_commons_log.Error(err.Error())
	}
}
