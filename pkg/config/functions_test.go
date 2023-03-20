package config

import (
	"reflect"
	"testing"

	"github.com/guidomantilla/go-feather-commons/pkg/environment"

	"github.com/guidomantilla/go-feather-sql/pkg/datasource"
	feather_sql "github.com/guidomantilla/go-feather-sql/pkg/sql"
)

func TestInit(t *testing.T) {
	type args struct {
		targetPrefix string
		environment  environment.Environment
		paramHolder  feather_sql.ParamHolder
	}
	tests := []struct {
		name  string
		args  args
		want  datasource.RelationalDatasource
		want1 datasource.RelationalDatasourceContext
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Init(tt.args.targetPrefix, tt.args.environment, tt.args.paramHolder)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Init() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Init() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestStop(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Stop(); (err != nil) != tt.wantErr {
				t.Errorf("Stop() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
