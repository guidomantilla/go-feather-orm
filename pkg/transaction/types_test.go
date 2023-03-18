package transaction

import (
	"context"
	"reflect"
	"testing"

	"github.com/guidomantilla/go-feather-sql/pkg/datasource"
)

func TestBuildRelationalTransactionHandler(t *testing.T) {
	type args struct {
		relationalDatasource datasource.RelationalDatasource
	}
	tests := []struct {
		name string
		args args
		want RelationalTransactionHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildRelationalTransactionHandler(tt.args.relationalDatasource); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildRelationalTransactionHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRelationalTransactionHandler(t *testing.T) {
	type args struct {
		relationalDatasource datasource.RelationalDatasource
	}
	tests := []struct {
		name string
		args args
		want *DefaultDBTransactionHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRelationalTransactionHandler(tt.args.relationalDatasource); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRelationalTransactionHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultDBTransactionHandler_HandleTransaction(t *testing.T) {
	type args struct {
		ctx context.Context
		fn  RelationalTransactionHandlerFunction
	}
	tests := []struct {
		name    string
		handler *DefaultDBTransactionHandler
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.handler.HandleTransaction(tt.args.ctx, tt.args.fn); (err != nil) != tt.wantErr {
				t.Errorf("DefaultDBTransactionHandler.HandleTransaction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_handleError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handleError(tt.args.err)
		})
	}
}
