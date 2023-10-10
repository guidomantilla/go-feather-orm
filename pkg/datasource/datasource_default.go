package datasource

import (
	"database/sql"
	"fmt"

	feather_commons_log "github.com/guidomantilla/go-feather-commons/pkg/log"
)

type DefaultDatasource struct {
	driver   string
	url      string
	server   string
	service  string
	database *sql.DB
	openFunc OpenDatasourceFunc
}

func NewDefaultDatasource(datasourceContext DatasourceContext, openFunc OpenDatasourceFunc) *DefaultDatasource {

	if datasourceContext == nil {
		feather_commons_log.Fatal("starting up - error setting up datasource: datasourceContext is nil")
	}

	if openFunc == nil {
		feather_commons_log.Fatal("starting up - error setting up datasource: openFunc is nil")
	}

	return &DefaultDatasource{
		driver:   datasourceContext.GetDriverName().String(),
		url:      datasourceContext.GetUrl(),
		server:   datasourceContext.GetServer(),
		service:  datasourceContext.GetService(),
		database: nil,
		openFunc: openFunc,
	}
}

func (datasource *DefaultDatasource) GetDatabase() (*sql.DB, error) {

	var err error

	if datasource.database == nil {
		if datasource.database, err = datasource.openFunc(datasource.driver, datasource.url); err != nil {
			feather_commons_log.Error(err.Error())
			return nil, ErrDBConnectionFailed(err)
		}
		feather_commons_log.Info(fmt.Sprintf("connection - connected to %s@%s/%s", datasource.driver, datasource.server, datasource.service))
	}

	if err = datasource.database.Ping(); err != nil {
		if datasource.database, err = datasource.openFunc(datasource.driver, datasource.url); err != nil {
			feather_commons_log.Error(err.Error())
			return nil, ErrDBConnectionFailed(err)
		}
		feather_commons_log.Info(fmt.Sprintf("connection - reconnected to %s@%s/%s", datasource.driver, datasource.server, datasource.service))
	}

	return datasource.database, nil
}
