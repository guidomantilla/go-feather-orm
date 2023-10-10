package datasource

import (
	"strings"

	feather_commons_log "github.com/guidomantilla/go-feather-commons/pkg/log"

	feather_sql "github.com/guidomantilla/go-feather-sql/pkg/sql"
)

type DefaultDatasourceContext struct {
	driverName  feather_sql.DriverName
	paramHolder feather_sql.ParamHolder
	url         string
	server      string
	service     string
}

func NewDefaultDatasourceContext(driverName feather_sql.DriverName, paramHolder feather_sql.ParamHolder,
	url string, username string, password string, server string, service string) *DefaultDatasourceContext {

	driverName = feather_sql.UndefinedDriverName.ValueFromCardinal(int(driverName))
	if driverName == feather_sql.UndefinedDriverName {
		feather_commons_log.Fatal("starting up - error setting up datasourceContext: driverName undefined")
	}

	paramHolder = feather_sql.UndefinedParamHolder.ValueFromCardinal(int(paramHolder))
	if paramHolder == feather_sql.UndefinedParamHolder {
		feather_commons_log.Fatal("starting up - error setting up datasourceContext: paramHolder undefined")
	}

	if strings.TrimSpace(url) == "" {
		feather_commons_log.Fatal("starting up - error setting up datasourceContext: url is empty")
	}

	if strings.TrimSpace(username) == "" {
		feather_commons_log.Fatal("starting up - error setting up datasourceContext: username is empty")
	}

	if strings.TrimSpace(password) == "" {
		feather_commons_log.Fatal("starting up - error setting up datasourceContext: password is empty")
	}

	if strings.TrimSpace(server) == "" {
		feather_commons_log.Fatal("starting up - error setting up datasourceContext: server is empty")
	}

	if strings.TrimSpace(service) == "" {
		feather_commons_log.Fatal("starting up - error setting up datasourceContext: service is empty")
	}

	url = strings.Replace(url, ":username", username, 1)
	url = strings.Replace(url, ":password", password, 1)
	url = strings.Replace(url, ":server", server, 1)
	url = strings.Replace(url, ":service", service, 1)

	return &DefaultDatasourceContext{
		driverName:  driverName,
		paramHolder: paramHolder,
		url:         url,
		server:      server,
		service:     service,
	}
}

func (context *DefaultDatasourceContext) GetDriverName() feather_sql.DriverName {
	return context.driverName
}

func (context *DefaultDatasourceContext) GetParamHolder() feather_sql.ParamHolder {
	return context.paramHolder
}

func (context *DefaultDatasourceContext) GetUrl() string {
	return context.url
}

func (context *DefaultDatasourceContext) GetServer() string {
	return context.server
}

func (context *DefaultDatasourceContext) GetService() string {
	return context.service
}
