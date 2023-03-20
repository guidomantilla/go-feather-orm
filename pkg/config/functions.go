package config

import (
	"database/sql"
	"fmt"

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

	zap.L().Info("server starting up - setting up DB connection")

	driverName := environment.GetValue(targetPrefix + DatasourceDriver).AsString()
	driver := feather_sql.UnknownDriverName.ValueOf(driverName)
	if driver == feather_sql.UnknownDriverName {
		zap.L().Fatal("server starting up - error setting up DB connection: invalid driver name")
	}

	username := environment.GetValue(targetPrefix + DatasourceUsername).AsString()
	password := environment.GetValue(targetPrefix + DatasourcePassword).AsString()
	server := environment.GetValue(targetPrefix + DatasourceServer).AsString()
	service := environment.GetValue(targetPrefix + DatasourceService).AsString()
	url := environment.GetValue(targetPrefix + DatasourceUrl).AsString()

	_datasourceContext = datasource.NewDefaultRelationalDatasourceContext(driver, paramHolder, url, username, password, server, service)

	_datasource = datasource.NewDefaultRelationalDatasource(_datasourceContext, sql.Open)
	return _datasource, _datasourceContext
}

func Stop() error {

	var err error
	var database *sql.DB

	zap.L().Info("server shutting down - closing DB")

	if database, err = _datasource.GetDatabase(); err != nil {
		zap.L().Error(fmt.Sprintf("server shutting down - error closing DB: %s", err.Error()))
		return err
	}

	if err = database.Close(); err != nil {
		zap.L().Error(fmt.Sprintf("server shutting down - error closing DB: %s", err.Error()))
		return err
	}

	zap.L().Info("server shutting down - DB closed")
	return nil
}
