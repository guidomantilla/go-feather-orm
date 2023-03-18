package transaction

import (
	"context"
	"database/sql"
	"fmt"

	"go.uber.org/zap"

	"github.com/guidomantilla/go-feather-sql/pkg/datasource"
)

var (
	_ RelationalTransactionHandler = (*DefaultDBTransactionHandler)(nil)
)

type RelationalTransactionContext struct{}

type RelationalTransactionHandlerFunction func(ctx context.Context, tx *sql.Tx) error

type RelationalTransactionHandler interface {
	HandleTransaction(ctx context.Context, fn RelationalTransactionHandlerFunction) error
}

func BuildRelationalTransactionHandler(relationalDatasource datasource.RelationalDatasource) RelationalTransactionHandler {
	return NewRelationalTransactionHandler(relationalDatasource)
}

type DefaultDBTransactionHandler struct {
	relationalDatasource datasource.RelationalDatasource
}

func NewRelationalTransactionHandler(relationalDatasource datasource.RelationalDatasource) *DefaultDBTransactionHandler {
	return &DefaultDBTransactionHandler{
		relationalDatasource: relationalDatasource,
	}
}

func (handler *DefaultDBTransactionHandler) HandleTransaction(ctx context.Context, fn RelationalTransactionHandlerFunction) error {

	db, err := handler.relationalDatasource.GetDatabase()
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
			fmt.Println(p)
			handleError(tx.Rollback())
		} else if err != nil {
			handleError(tx.Rollback())
		} else {
			handleError(tx.Commit())
		}
	}()

	txCtx := context.WithValue(ctx, RelationalTransactionContext{}, tx)
	err = fn(txCtx, tx)
	return err
}

//

func handleError(err error) {
	if err != nil {
		zap.L().Error(err.Error())
	}
}
