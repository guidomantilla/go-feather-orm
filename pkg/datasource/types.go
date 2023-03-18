package datasource

import (
	"database/sql"

	feather_sql "github.com/guidomantilla/go-feather-sql/pkg/sql"
)

type RelationalDatasourceContext interface {
	GetDriverName() feather_sql.DriverName
	GetParamHolder() feather_sql.ParamHolder
	GetUrl() string
}

func BuildRelationalDatasourceContext(driverName feather_sql.DriverName, paramHolder feather_sql.ParamHolder,
	url string, username string, password string, server string, service string) RelationalDatasourceContext {

	return NewRelationalDatasourceContext(driverName, paramHolder, url, username, password, server, service)
}

//

type OpenDatasourceFunc func(driverName, datasourceUrl string) (*sql.DB, error)

type RelationalDatasource interface {
	GetDatabase() (*sql.DB, error)
}

func BuildRelationalDatasource(datasourceContext RelationalDatasourceContext, openFunc OpenDatasourceFunc) RelationalDatasource {
	return NewDefaultRelationalDatasource(datasourceContext, openFunc)
}
