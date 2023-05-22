package dao

import (
	"context"
	"reflect"
	"testing"

	feather_sql_datasource "github.com/guidomantilla/go-feather-sql/pkg/datasource"
)

func TestNewDefaultCrudDao(t *testing.T) {
	type args struct {
		datasourceContext feather_sql_datasource.RelationalDatasourceContext
		table             string
		model             any
	}
	tests := []struct {
		name string
		args args
		want *DefaultCrudDao
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDefaultCrudDao(tt.args.datasourceContext, tt.args.table, tt.args.model); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDefaultCrudDao() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultCrudDao_Save(t *testing.T) {
	type args struct {
		ctx  context.Context
		args []any
	}
	tests := []struct {
		name    string
		dao     *DefaultCrudDao
		args    args
		want    *int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.dao.Save(tt.args.ctx, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("DefaultCrudDao.Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultCrudDao.Save() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultCrudDao_Update(t *testing.T) {
	type args struct {
		ctx  context.Context
		args []any
	}
	tests := []struct {
		name    string
		dao     *DefaultCrudDao
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.dao.Update(tt.args.ctx, tt.args.args...); (err != nil) != tt.wantErr {
				t.Errorf("DefaultCrudDao.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDefaultCrudDao_Delete(t *testing.T) {
	type args struct {
		ctx context.Context
		id  any
	}
	tests := []struct {
		name    string
		dao     *DefaultCrudDao
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.dao.Delete(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DefaultCrudDao.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDefaultCrudDao_FindById(t *testing.T) {
	type args struct {
		ctx  context.Context
		id   any
		args []any
	}
	tests := []struct {
		name    string
		dao     *DefaultCrudDao
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.dao.FindById(tt.args.ctx, tt.args.id, tt.args.args...); (err != nil) != tt.wantErr {
				t.Errorf("DefaultCrudDao.FindById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDefaultCrudDao_ExistsById(t *testing.T) {
	type args struct {
		ctx  context.Context
		id   any
		args []any
	}
	tests := []struct {
		name string
		dao  *DefaultCrudDao
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dao.ExistsById(tt.args.ctx, tt.args.id, tt.args.args...); got != tt.want {
				t.Errorf("DefaultCrudDao.ExistsById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultCrudDao_FindAll(t *testing.T) {
	type args struct {
		ctx context.Context
		fn  ReadFunction
	}
	tests := []struct {
		name    string
		dao     *DefaultCrudDao
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.dao.FindAll(tt.args.ctx, tt.args.fn); (err != nil) != tt.wantErr {
				t.Errorf("DefaultCrudDao.FindAll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
