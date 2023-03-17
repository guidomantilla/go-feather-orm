package feather_relational_datasource

import (
	"database/sql"

	"go.uber.org/zap"
)

type DefaultRelationalDatasource struct {
	driver   string
	url      string
	database *sql.DB
	openFunc OpenDatasourceFunc
}

func NewDefaultRelationalDatasource(datasourceContext RelationalDatasourceContext, openFunc OpenDatasourceFunc) *DefaultRelationalDatasource {

	return &DefaultRelationalDatasource{
		driver:   datasourceContext.GetDriverName().String(),
		url:      datasourceContext.GetUrl(),
		database: nil,
		openFunc: openFunc,
	}
}

func (datasource *DefaultRelationalDatasource) GetDatabase() (*sql.DB, error) {

	var err error

	if datasource.database == nil {
		if datasource.database, err = datasource.openFunc(datasource.driver, datasource.url); err != nil {
			zap.L().Error(err.Error())
			return nil, err
		}
	}

	if err = datasource.database.Ping(); err != nil {
		if datasource.database, err = datasource.openFunc(datasource.driver, datasource.url); err != nil {
			zap.L().Error(err.Error())
			return nil, err
		}
	}

	return datasource.database, nil
}
