package datasource

import (
	"database/sql"
	"log/slog"
	"os"
)

type DefaultDatasource struct {
	driver   string
	url      string
	database *sql.DB
	openFunc OpenDatasourceFunc
}

func NewDefaultDatasource(datasourceContext DatasourceContext, openFunc OpenDatasourceFunc) *DefaultDatasource {

	if datasourceContext == nil {
		slog.Error("starting up - error setting up datasource: datasourceContext is nil")
		os.Exit(1)
	}

	if openFunc == nil {
		slog.Error("starting up - error setting up datasource: openFunc is nil")
		os.Exit(1)
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
			slog.Error(err.Error())
			return nil, ErrDBConnectionFailed(err)
		}
	}

	if err = datasource.database.Ping(); err != nil {
		if datasource.database, err = datasource.openFunc(datasource.driver, datasource.url); err != nil {
			slog.Error(err.Error())
			return nil, ErrDBConnectionFailed(err)
		}
	}

	return datasource.database, nil
}
