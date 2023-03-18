package transaction

import (
	"database/sql"

	"go.uber.org/zap"

	"jukebox-app/pkg/datasource"
)

var (
	_ RelationalTransactionHandler    = (*DefaultDBTransactionHandler)(nil)
	_ datasource.RelationalDataSource = (*DefaultDBTransactionHandler)(nil)
)

type RelationalTransactionContext struct{}

type RelationalTransactionHandlerFunction func(tx *sql.Tx) error

type RelationalTransactionHandler interface {
	HandleTransaction(fn RelationalTransactionHandlerFunction) error
}

type DefaultDBTransactionHandler struct {
	datasource.RelationalDataSource
}

func (handler *DefaultDBTransactionHandler) HandleTransaction(fn RelationalTransactionHandlerFunction) error {

	db, err := handler.GetDatabase()
	if err != nil {
		zap.L().Error(err.Error())
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		zap.L().Error(err.Error())
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			handleError(tx.Rollback())
		} else if err != nil {
			handleError(tx.Rollback())
		} else {
			handleError(tx.Commit())
		}
	}()

	err = fn(tx)
	return err
}

//

func NewRelationalTransactionHandler(relationalDatasource datasource.RelationalDataSource) *DefaultDBTransactionHandler {
	return &DefaultDBTransactionHandler{
		RelationalDataSource: relationalDatasource,
	}
}

func handleError(err error) {
	if err != nil {
		zap.L().Error(err.Error())
	}
}
