package datasource

import (
	"context"

	"github.com/jmoiron/sqlx"

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

type OpenDatasourceFunc func(driverName, datasourceUrl string) (*sqlx.DB, error)

type Datasource interface {
	GetDatabase() (*sqlx.DB, error)
}

//

type TransactionCtxKey struct{}

type TransactionHandlerFunction func(ctx context.Context, tx *sqlx.Tx) error

type TransactionHandler interface {
	HandleTransaction(ctx context.Context, fn TransactionHandlerFunction) error
}
