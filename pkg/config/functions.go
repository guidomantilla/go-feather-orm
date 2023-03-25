package config

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/guidomantilla/go-feather-commons/pkg/environment"
	"go.uber.org/zap"

	"github.com/guidomantilla/go-feather-sql/pkg/datasource"
	feather_sql "github.com/guidomantilla/go-feather-sql/pkg/sql"
)

const (
	DatasourceDriver   = "DATASOURCE_DRIVER"
	DatasourceUsername = "DATASOURCE_USERNAME"
	DatasourcePassword = "DATASOURCE_PASSWORD"
	DatasourceServer   = "DATASOURCE_SERVER"
	DatasourceService  = "DATASOURCE_SERVICE"
	DatasourceUrl      = "DATASOURCE_URL"
)

var (
	_datasource        datasource.RelationalDatasource
	_datasourceContext datasource.RelationalDatasourceContext
)

func Init(targetPrefix string, environment environment.Environment, paramHolder feather_sql.ParamHolder) (datasource.RelationalDatasource, datasource.RelationalDatasourceContext) {

	zap.L().Info("starting up - setting up DB connection")

	if strings.TrimSpace(targetPrefix) == "" {
		zap.L().Fatal("starting up - error setting up DB config: targetPrefix is empty")
	}

	if environment == nil {
		zap.L().Fatal("starting up - error setting up DB config: environment is nil")
	}

	paramHolder = feather_sql.UndefinedParamHolder.ValueFromCardinal(int(paramHolder))
	if paramHolder == feather_sql.UndefinedParamHolder {
		zap.L().Fatal("starting up - error setting up DB config: invalid param holder")
	}

	driverName := environment.GetValue(targetPrefix + DatasourceDriver).AsString()
	driver := feather_sql.UndefinedDriverName.ValueFromName(driverName)
	if driver == feather_sql.UndefinedDriverName {
		zap.L().Fatal("starting up - error setting up DB config: invalid driver name")
	}

	url := environment.GetValue(targetPrefix + DatasourceUrl).AsString()
	if strings.TrimSpace(url) == "" {
		zap.L().Fatal("starting up - error setting up DB config: url is empty")
	}

	username := environment.GetValue(targetPrefix + DatasourceUsername).AsString()
	if strings.TrimSpace(username) == "" {
		zap.L().Fatal("starting up - error setting up DB config: username is empty")
	}

	password := environment.GetValue(targetPrefix + DatasourcePassword).AsString()
	if strings.TrimSpace(password) == "" {
		zap.L().Fatal("starting up - error setting up DB config: password is empty")
	}

	server := environment.GetValue(targetPrefix + DatasourceServer).AsString()
	if strings.TrimSpace(server) == "" {
		zap.L().Fatal("starting up - error setting up DB config: server is empty")
	}

	service := environment.GetValue(targetPrefix + DatasourceService).AsString()
	if strings.TrimSpace(service) == "" {
		zap.L().Fatal("starting up - error setting up DB config: service is empty")
	}

	_datasourceContext = datasource.NewDefaultRelationalDatasourceContext(driver, paramHolder, url, username, password, server, service)

	_datasource = datasource.NewDefaultRelationalDatasource(_datasourceContext, sql.Open)
	return _datasource, _datasourceContext
}

func Stop() error {

	var err error
	var database *sql.DB

	zap.L().Info("shutting down - closing DB")

	if database, err = _datasource.GetDatabase(); err != nil {
		zap.L().Error(fmt.Sprintf("shutting down - error closing DB: %s", err.Error()))
		return err
	}

	if err = database.Close(); err != nil {
		zap.L().Error(fmt.Sprintf("shutting down - error closing DB: %s", err.Error()))
		return err
	}

	zap.L().Info("shutting down - DB closed")
	return nil
}
