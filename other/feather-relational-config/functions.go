package feather_relational_config

import (
	"database/sql"
	"fmt"

	feather_relational_datasource "github.com/guidomantilla/go-feather-sql/pkg/feather-relational-datasource"
	feather_sql "github.com/guidomantilla/go-feather-sql/pkg/feather-sql"
	"go.uber.org/zap"

	"oracle-api-orchestrator/pkg/environment"
)

var _singletonDatasource feather_relational_datasource.RelationalDatasource

var _singletonDatasourceContext feather_relational_datasource.RelationalDatasourceContext

func InitDB(environment environment.Environment) (feather_relational_datasource.RelationalDatasource, feather_relational_datasource.RelationalDatasourceContext) {

	zap.L().Info("server starting up - setting up DB connection")

	driver := environment.GetValue(DATASOURCE_DRIVER).AsString()
	if driver != feather_sql.OracleDriverName.String() {
		zap.L().Fatal("server starting up - error setting up DB connection: invalid driver name")
	}

	username := environment.GetValue(DATASOURCE_USERNAME).AsString()
	password := environment.GetValue(DATASOURCE_PASSWORD).AsString()
	server := environment.GetValue(DATASOURCE_SERVER).AsString()
	service := environment.GetValue(DATASOURCE_SERVICE).AsString()
	url := environment.GetValue(DATASOURCE_URL).AsString()

	_singletonDatasourceContext = feather_relational_datasource.BuildRelationalDatasourceContext(feather_sql.OracleDriverName, feather_sql.NumberedParamHolder,
		url, username, password, server, service)

	_singletonDatasource = feather_relational_datasource.BuildRelationalDatasource(_singletonDatasourceContext, sql.Open)
	return _singletonDatasource, _singletonDatasourceContext
}

func StopDB() error {

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
