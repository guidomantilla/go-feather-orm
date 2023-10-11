package datasource

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"

	feather_sql "github.com/guidomantilla/go-feather-sql/pkg/sql"
)

func TestDefaultTransactionHandler_HandleTransaction(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	errGetDatabasePath := func() *DefaultTransactionHandler {
		datasourceCtx := NewMockDatasourceContext(ctrl)
		datasource := NewMockDatasource(ctrl)
		datasource.EXPECT().GetDatabase().Return(nil, errors.New("some_error"))
		return NewTransactionHandler(datasourceCtx, datasource)
	}

	errBeginPath := func() *DefaultTransactionHandler {
		db, sqlMock, _ := sqlmock.New()
		sqlMock.ExpectBegin().WillReturnError(errors.New("some_error"))
		datasourceCtx := NewMockDatasourceContext(ctrl)
		datasource := NewMockDatasource(ctrl)
		datasource.EXPECT().GetDatabase().Return(db, nil)
		return NewTransactionHandler(datasourceCtx, datasource)
	}

	errDeferPath := func() *DefaultTransactionHandler {
		db, sqlMock, _ := sqlmock.New()
		sqlMock.ExpectBegin()
		datasourceCtx := NewMockDatasourceContext(ctrl)
		datasourceCtx.EXPECT().GetDriverName().Return(feather_sql.MysqlDriverName)
		datasource := NewMockDatasource(ctrl)
		datasource.EXPECT().GetDatabase().Return(db, nil)
		return NewTransactionHandler(datasourceCtx, datasource)
	}

	panicDeferPath := func() *DefaultTransactionHandler {
		db, sqlMock, _ := sqlmock.New()
		sqlMock.ExpectBegin()
		datasourceCtx := NewMockDatasourceContext(ctrl)
		datasourceCtx.EXPECT().GetDriverName().Return(feather_sql.MysqlDriverName)
		sqlMock.ExpectRollback()
		datasource := NewMockDatasource(ctrl)
		datasource.EXPECT().GetDatabase().Return(db, nil)
		return NewTransactionHandler(datasourceCtx, datasource)
	}

	happyPath := func() *DefaultTransactionHandler {
		db, sqlMock, _ := sqlmock.New()
		sqlMock.ExpectBegin()
		datasourceCtx := NewMockDatasourceContext(ctrl)
		datasourceCtx.EXPECT().GetDriverName().Return(feather_sql.MysqlDriverName)
		datasource := NewMockDatasource(ctrl)
		datasource.EXPECT().GetDatabase().Return(db, nil)
		return NewTransactionHandler(datasourceCtx, datasource)
	}

	type args struct {
		ctx context.Context
		fn  TransactionHandlerFunction
	}
	tests := []struct {
		name    string
		handler *DefaultTransactionHandler
		args    args
		wantErr bool
	}{
		{
			name:    "Err GetDatabase() Path",
			handler: errGetDatabasePath(),
			args: args{
				ctx: context.TODO(),
				fn: func(ctx context.Context, tx *sql.Tx) error {
					return errors.New("some_error")
				},
			},
			wantErr: true,
		},
		{
			name:    "Err Begin() Path",
			handler: errBeginPath(),
			args: args{
				ctx: context.TODO(),
				fn: func(ctx context.Context, tx *sql.Tx) error {
					return nil
				},
			},
			wantErr: true,
		},
		{
			name:    "Err Defer Path",
			handler: errDeferPath(),
			args: args{
				ctx: context.TODO(),
				fn: func(ctx context.Context, tx *sql.Tx) error {
					return errors.New("some_error")
				},
			},
			wantErr: true,
		},
		{
			name:    "Panic Defer Path",
			handler: panicDeferPath(),
			args: args{
				ctx: context.TODO(),
				fn: func(ctx context.Context, tx *sql.Tx) error {
					panic("some_panic")
				},
			},
			wantErr: false,
		},

		{
			name:    "Happy ath",
			handler: happyPath(),
			args: args{
				ctx: context.TODO(),
				fn: func(ctx context.Context, tx *sql.Tx) error {
					return nil
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.handler.HandleTransaction(tt.args.ctx, tt.args.fn); (err != nil) != tt.wantErr {
				t.Errorf("DefaultTransactionHandler.HandleTransaction() error = %v, wantErr %v", err, tt.wantErr)
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
		{
			name: "HappyPath",
			args: args{err: nil},
		},
		{
			name: "UnhappyPath",
			args: args{err: errors.New("some error")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handleError(tt.args.err)
		})
	}
}
