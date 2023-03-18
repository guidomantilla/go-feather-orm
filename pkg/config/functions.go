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
	_singletonDatasource        datasource.RelationalDatasource
	_singletonDatasourceContext datasource.RelationalDatasourceContext
)

func Init(targetPrefix string, environment environment.Environment) (datasource.RelationalDatasource, datasource.RelationalDatasourceContext) {

	zap.L().Info("server starting up - setting up DB connection")

	driver := environment.GetValue(targetPrefix + DatasourceDriver).AsString()
	if driver != feather_sql.OracleDriverName.String() {
		zap.L().Fatal("server starting up - error setting up DB connection: invalid driver name")
	}

	username := environment.GetValue(targetPrefix + DatasourceUsername).AsString()
	password := environment.GetValue(targetPrefix + DatasourcePassword).AsString()
	server := environment.GetValue(targetPrefix + DatasourceServer).AsString()
	service := environment.GetValue(targetPrefix + DatasourceService).AsString()
	url := environment.GetValue(targetPrefix + DatasourceUrl).AsString()

	_singletonDatasourceContext = datasource.BuildRelationalDatasourceContext(feather_sql.OracleDriverName, feather_sql.NumberedParamHolder,
		url, username, password, server, service)

	_singletonDatasource = datasource.BuildRelationalDatasource(_singletonDatasourceContext, sql.Open)
	return _singletonDatasource, _singletonDatasourceContext
}

func Stop() error {

	var err error
	var database *sql.DB

	zap.L().Info("server shutting down - closing DB")

	if database, err = _singletonDatasource.GetDatabase(); err != nil {
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
