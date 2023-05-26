package datasource

import (
	"database/sql"
	"errors"
	"reflect"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"

	feather_sql "github.com/guidomantilla/go-feather-sql/pkg/sql"
)

func TestNewDefaultDatasource(t *testing.T) {
	openFunc := OpenDatasourceFunc(func(driverName, dataSourceUrl string) (*sql.DB, error) {
		return nil, nil
	})
	datasourceCtx := &DefaultDatasourceContext{
		driverName:  feather_sql.UndefinedDriverName,
		paramHolder: feather_sql.NamedParamHolder,
		url:         "some_usersome_passsome_serversome_service",
	}
	datasource := &DefaultDatasource{
		driver:   datasourceCtx.driverName.String(),
		url:      datasourceCtx.url,
		database: nil,
		openFunc: openFunc,
	}

	type args struct {
		datasourceContext DatasourceContext
		openFunc          OpenDatasourceFunc
	}
	tests := []struct {
		name string
		args args
		want *DefaultDatasource
	}{
		{
			name: "Happy Path",
			args: args{
				datasourceContext: datasourceCtx,
				openFunc:          openFunc,
			},
			want: datasource,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDefaultDatasource(tt.args.datasourceContext, tt.args.openFunc)
			if !reflect.DeepEqual(reflect.ValueOf(got.openFunc), reflect.ValueOf(tt.want.openFunc)) {
				t.Errorf("NewDefaultDatasource() = %v, want %v", got.openFunc, tt.want.openFunc)
			}
			if !reflect.DeepEqual(got.url, tt.want.url) {
				t.Errorf("NewDefaultDatasource() = %v, want %v", got.url, tt.want.url)
			}
			if !reflect.DeepEqual(got.driver, tt.want.driver) {
				t.Errorf("NewDefaultDatasource() = %v, want %v", got.driver, tt.want.driver)
			}
			if !reflect.DeepEqual(got.database, tt.want.database) {
				t.Errorf("NewDefaultDatasource() = %v, want %v", got.database, tt.want.database)
			}
		})
	}
}

func TestDefaultDatasource_GetDatabase(t *testing.T) {

	datasourceCtx := &DefaultDatasourceContext{
		driverName:  feather_sql.UndefinedDriverName,
		paramHolder: feather_sql.NamedParamHolder,
		url:         "some_usersome_passsome_serversome_service",
	}

	errOpenFuncPath := func() *DefaultDatasource {
		return &DefaultDatasource{
			driver:   datasourceCtx.driverName.String(),
			url:      datasourceCtx.url,
			database: nil,
			openFunc: func(driverName, dataSourceUrl string) (*sql.DB, error) {
				return nil, errors.New("some_error")
			},
		}
	}

	errOpenFuncPath2 := func() *DefaultDatasource {

		var err error
		var db *sql.DB
		var mock sqlmock.Sqlmock
		if db, mock, err = sqlmock.New(sqlmock.MonitorPingsOption(true)); err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		mock.ExpectPing().WillReturnError(errors.New("some error"))
		return &DefaultDatasource{
			driver:   datasourceCtx.driverName.String(),
			url:      datasourceCtx.url,
			database: db,
			openFunc: func(driverName, dataSourceUrl string) (*sql.DB, error) {
				return nil, errors.New("some_error")
			},
		}
	}

	var err error
	var db *sql.DB
	var mock sqlmock.Sqlmock
	if db, mock, err = sqlmock.New(); err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	mock.ExpectPing()
	happyPath := func() *DefaultDatasource {
		return &DefaultDatasource{
			driver: datasourceCtx.driverName.String(),
			url:    datasourceCtx.url,
			openFunc: func(driverName, dataSourceUrl string) (*sql.DB, error) {
				return db, nil
			},
		}
	}

	tests := []struct {
		name       string
		datasource *DefaultDatasource
		want       *sql.DB
		wantErr    bool
	}{
		{
			name:       "Err openFunc Path",
			datasource: errOpenFuncPath(),
			want:       nil,
			wantErr:    true,
		},
		{
			name:       "Err openFunc Path1",
			datasource: errOpenFuncPath2(),
			want:       nil,
			wantErr:    true,
		},
		{
			name:       "HappyPath",
			datasource: happyPath(),
			want:       db,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.datasource.GetDatabase()
			if (err != nil) != tt.wantErr {
				t.Errorf("DefaultDatasource.GetDatabase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultDatasource.GetDatabase() = %v, want %v", got, tt.want)
			}
		})
	}
}
