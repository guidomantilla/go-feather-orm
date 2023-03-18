package feather_relational_datasource

import (
	"reflect"
	"testing"

	feather_sql "github.com/guidomantilla/go-feather-sql/pkg/feather-sql"
)

func TestNewRelationalDatasourceContext(t *testing.T) {
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
		want *DefaultRelationalDatasourceContext
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRelationalDatasourceContext(tt.args.driverName, tt.args.paramHolder, tt.args.url, tt.args.username, tt.args.password, tt.args.server, tt.args.service); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRelationalDatasourceContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultRelationalDatasourceContext_GetDriverName(t *testing.T) {
	tests := []struct {
		name    string
		context *DefaultRelationalDatasourceContext
		want    feather_sql.DriverName
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.context.GetDriverName(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultRelationalDatasourceContext.GetDriverName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultRelationalDatasourceContext_GetParamHolder(t *testing.T) {
	tests := []struct {
		name    string
		context *DefaultRelationalDatasourceContext
		want    feather_sql.ParamHolder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.context.GetParamHolder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultRelationalDatasourceContext.GetParamHolder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultRelationalDatasourceContext_GetUrl(t *testing.T) {
	tests := []struct {
		name    string
		context *DefaultRelationalDatasourceContext
		want    string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.context.GetUrl(); got != tt.want {
				t.Errorf("DefaultRelationalDatasourceContext.GetUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
