package feather_relational_datasource

import (
	"strings"

	feather_sql "github.com/guidomantilla/go-feather-sql/pkg/feather-sql"
)

type DefaultRelationalDatasourceContext struct {
	driverName  feather_sql.DriverName
	paramHolder feather_sql.ParamHolder
	url         string
}

func NewRelationalDatasourceContext(driverName feather_sql.DriverName, paramHolder feather_sql.ParamHolder,
	url string, username string, password string, server string, service string) *DefaultRelationalDatasourceContext {

	url = strings.Replace(url, ":username", username, 1)
	url = strings.Replace(url, ":password", password, 1)
	url = strings.Replace(url, ":server", server, 1)
	url = strings.Replace(url, ":service", service, 1)

	return &DefaultRelationalDatasourceContext{
		driverName:  driverName,
		paramHolder: paramHolder,
		url:         url,
	}
}

func (context *DefaultRelationalDatasourceContext) GetDriverName() feather_sql.DriverName {
	return context.driverName
}

func (context *DefaultRelationalDatasourceContext) GetParamHolder() feather_sql.ParamHolder {
	return context.paramHolder
}

func (context *DefaultRelationalDatasourceContext) GetUrl() string {
	return context.url
}
