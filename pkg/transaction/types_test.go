package transaction

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"

	"github.com/guidomantilla/go-feather-sql/pkg/datasource"
)

func TestDefaultDBTransactionHandler_HandleTransaction(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	errGetDatabasePath := func() *DefaultDBTransactionHandler {
		datasource := datasource.NewMockRelationalDatasource(ctrl)
		datasource.EXPECT().GetDatabase().Return(nil, errors.New("some_error"))
		return NewRelationalTransactionHandler(datasource)
	}

	errBeginPath := func() *DefaultDBTransactionHandler {
		db, sqlMock, _ := sqlmock.New()
		sqlMock.ExpectBegin().WillReturnError(errors.New("some_error"))
		datasource := datasource.NewMockRelationalDatasource(ctrl)
		datasource.EXPECT().GetDatabase().Return(db, nil)
		return NewRelationalTransactionHandler(datasource)
	}

	errDeferPath := func() *DefaultDBTransactionHandler {
		db, sqlMock, _ := sqlmock.New()
		sqlMock.ExpectBegin()
		datasource := datasource.NewMockRelationalDatasource(ctrl)
		datasource.EXPECT().GetDatabase().Return(db, nil)
		return NewRelationalTransactionHandler(datasource)
	}

	panicDeferPath := func() *DefaultDBTransactionHandler {
		db, sqlMock, _ := sqlmock.New()
		sqlMock.ExpectBegin()
		sqlMock.ExpectRollback()
		datasource := datasource.NewMockRelationalDatasource(ctrl)
		datasource.EXPECT().GetDatabase().Return(db, nil)
		return NewRelationalTransactionHandler(datasource)
	}

	happyPath := func() *DefaultDBTransactionHandler {
		db, sqlMock, _ := sqlmock.New()
		sqlMock.ExpectBegin()
		datasource := datasource.NewMockRelationalDatasource(ctrl)
		datasource.EXPECT().GetDatabase().Return(db, nil)
		return NewRelationalTransactionHandler(datasource)
	}

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
