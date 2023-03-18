package feather_relational_datasource

import (
	"database/sql"
	"reflect"
	"testing"
)

func TestNewDefaultRelationalDatasource(t *testing.T) {
	type args struct {
		datasourceContext RelationalDatasourceContext
		openFunc          OpenDatasourceFunc
	}
	tests := []struct {
		name string
		args args
		want *DefaultRelationalDatasource
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDefaultRelationalDatasource(tt.args.datasourceContext, tt.args.openFunc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDefaultRelationalDatasource() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultRelationalDatasource_GetDatabase(t *testing.T) {
	tests := []struct {
		name       string
		datasource *DefaultRelationalDatasource
		want       *sql.DB
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.datasource.GetDatabase()
			if (err != nil) != tt.wantErr {
				t.Errorf("DefaultRelationalDatasource.GetDatabase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultRelationalDatasource.GetDatabase() = %v, want %v", got, tt.want)
			}
		})
	}
}
