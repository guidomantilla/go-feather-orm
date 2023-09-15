package datasource

import (
	"log/slog"
	"os"
	"strings"

	feather_sql "github.com/guidomantilla/go-feather-sql/pkg/sql"
)

type DefaultDatasourceContext struct {
	driverName  feather_sql.DriverName
	paramHolder feather_sql.ParamHolder
	url         string
}

func NewDefaultDatasourceContext(driverName feather_sql.DriverName, paramHolder feather_sql.ParamHolder,
	url string, username string, password string, server string, service string) *DefaultDatasourceContext {

	driverName = feather_sql.UndefinedDriverName.ValueFromCardinal(int(driverName))
	if driverName == feather_sql.UndefinedDriverName {
		slog.Error("starting up - error setting up datasourceContext: driverName undefined")
		os.Exit(1)
	}

	paramHolder = feather_sql.UndefinedParamHolder.ValueFromCardinal(int(paramHolder))
	if paramHolder == feather_sql.UndefinedParamHolder {
		slog.Error("starting up - error setting up datasourceContext: paramHolder undefined")
		os.Exit(1)
	}

	if strings.TrimSpace(url) == "" {
		slog.Error("starting up - error setting up datasourceContext: url is empty")
		os.Exit(1)
	}

	if strings.TrimSpace(username) == "" {
		slog.Error("starting up - error setting up datasourceContext: username is empty")
		os.Exit(1)
	}

	if strings.TrimSpace(password) == "" {
		slog.Error("starting up - error setting up datasourceContext: password is empty")
		os.Exit(1)
	}

	if strings.TrimSpace(server) == "" {
		slog.Error("starting up - error setting up datasourceContext: server is empty")
		os.Exit(1)
	}

	if strings.TrimSpace(service) == "" {
		slog.Error("starting up - error setting up datasourceContext: service is empty")
		os.Exit(1)
	}

	url = strings.Replace(url, ":username", username, 1)
	url = strings.Replace(url, ":password", password, 1)
	url = strings.Replace(url, ":server", server, 1)
	url = strings.Replace(url, ":service", service, 1)

	return &DefaultDatasourceContext{
		driverName:  driverName,
		paramHolder: paramHolder,
		url:         url,
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
