package datasource

import (
	"database/sql"

	"go.uber.org/zap"
)

type DefaultDatasource struct {
	driver   string
	url      string
	database *sql.DB
	openFunc OpenDatasourceFunc
}

func NewDefaultDatasource(datasourceContext DatasourceContext, openFunc OpenDatasourceFunc) *DefaultDatasource {

	if datasourceContext == nil {
		zap.L().Fatal("starting up - error setting up datasource: datasourceContext is nil")
	}

	if openFunc == nil {
		zap.L().Fatal("starting up - error setting up datasource: openFunc is nil")
	}

	return &DefaultDatasource{
		driver:   datasourceContext.GetDriverName().String(),
		url:      datasourceContext.GetUrl(),
		database: nil,
		openFunc: openFunc,
	}
}

func (datasource *DefaultDatasource) GetDatabase() (*sql.DB, error) {

	var err error

	if datasource.database == nil {
		if datasource.database, err = datasource.openFunc(datasource.driver, datasource.url); err != nil {
			zap.L().Error(err.Error())
			return nil, ErrDBConnectionFailed(err)
		}
	}

	if err = datasource.database.Ping(); err != nil {
		if datasource.database, err = datasource.openFunc(datasource.driver, datasource.url); err != nil {
			zap.L().Error(err.Error())
			return nil, ErrDBConnectionFailed(err)
		}
	}

	return datasource.database, nil
}
