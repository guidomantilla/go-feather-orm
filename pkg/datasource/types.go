package datasource

import (
	"context"
	"database/sql"

	feather_sql "github.com/guidomantilla/go-feather-sql/pkg/sql"
)

var (
	_ DatasourceContext  = (*DefaultDatasourceContext)(nil)
	_ Datasource         = (*DefaultDatasource)(nil)
	_ TransactionHandler = (*DefaultTransactionHandler)(nil)
)

type DatasourceContext interface {
	GetDriverName() feather_sql.DriverName
	GetParamHolder() feather_sql.ParamHolder
	GetUrl() string
	GetServer() string
	GetService() string
}

//

type OpenDatasourceFunc func(driverName, datasourceUrl string) (*sql.DB, error)

type Datasource interface {
	GetDatabase() (*sql.DB, error)
}

//

type TransactionCtxKey struct{}

type TransactionHandlerFunction func(ctx context.Context, tx *sql.Tx) error

type TransactionHandler interface {
	HandleTransaction(ctx context.Context, fn TransactionHandlerFunction) error
}
