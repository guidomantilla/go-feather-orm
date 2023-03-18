package datasource

import (
	"reflect"
	"testing"

	feather_sql "github.com/guidomantilla/go-feather-sql/pkg/sql"
)

func TestBuildRelationalDatasourceContext(t *testing.T) {
	type args struct {
		driverName  feather_sql.DriverName
		paramHolder feather_sql.ParamHolder
		url         string
		username    string
		password    string
		server      string
		service     string
	}
	tests := []struct {
		name string
		args args
		want RelationalDatasourceContext
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildRelationalDatasourceContext(tt.args.driverName, tt.args.paramHolder, tt.args.url, tt.args.username, tt.args.password, tt.args.server, tt.args.service); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildRelationalDatasourceContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildRelationalDatasource(t *testing.T) {
	type args struct {
		datasourceContext RelationalDatasourceContext
		openFunc          OpenDatasourceFunc
	}
	tests := []struct {
		name string
		args args
		want RelationalDatasource
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildRelationalDatasource(tt.args.datasourceContext, tt.args.openFunc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildRelationalDatasource() = %v, want %v", got, tt.want)
			}
		})
	}
}
