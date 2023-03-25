package datasource

import (
	"strings"

	"go.uber.org/zap"

	feather_sql "github.com/guidomantilla/go-feather-sql/pkg/sql"
)

type DefaultRelationalDatasourceContext struct {
	driverName  feather_sql.DriverName
	paramHolder feather_sql.ParamHolder
	url         string
}

func NewDefaultRelationalDatasourceContext(driverName feather_sql.DriverName, paramHolder feather_sql.ParamHolder,
	url string, username string, password string, server string, service string) *DefaultRelationalDatasourceContext {

	driverName = feather_sql.UndefinedDriverName.ValueFromCardinal(int(driverName))
	if driverName == feather_sql.UndefinedDriverName {
		zap.L().Fatal("starting up - error setting up datasourceContext: driverName undefined")
	}

	paramHolder = feather_sql.UndefinedParamHolder.ValueFromCardinal(int(paramHolder))
	if paramHolder == feather_sql.UndefinedParamHolder {
		zap.L().Fatal("starting up - error setting up datasourceContext: paramHolder undefined")
	}

	if strings.TrimSpace(url) == "" {
		zap.L().Fatal("starting up - error setting up datasourceContext: url is empty")
	}

	if strings.TrimSpace(username) == "" {
		zap.L().Fatal("starting up - error setting up datasourceContext: username is empty")
	}

	if strings.TrimSpace(password) == "" {
		zap.L().Fatal("starting up - error setting up datasourceContext: password is empty")
	}

	if strings.TrimSpace(server) == "" {
		zap.L().Fatal("starting up - error setting up datasourceContext: server is empty")
	}

	if strings.TrimSpace(service) == "" {
		zap.L().Fatal("starting up - error setting up datasourceContext: service is empty")
	}

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
