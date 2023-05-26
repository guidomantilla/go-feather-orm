package datasource

import (
	"database/sql"

	feather_sql "github.com/guidomantilla/go-feather-sql/pkg/sql"
)

var (
	_ DatasourceContext = (*DefaultDatasourceContext)(nil)
	_ Datasource        = (*DefaultDatasource)(nil)
)

type DatasourceContext interface {
	GetDriverName() feather_sql.DriverName
	GetParamHolder() feather_sql.ParamHolder
	GetUrl() string
}

//

type OpenDatasourceFunc func(driverName, datasourceUrl string) (*sql.DB, error)

type Datasource interface {
	GetDatabase() (*sql.DB, error)
}
