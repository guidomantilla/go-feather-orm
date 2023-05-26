package dao

import (
	"context"
	"database/sql"
	"errors"
	"reflect"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"

	feather_sql_transaction "github.com/guidomantilla/go-feather-sql/pkg/transaction"
)

func TestWriteContext(t *testing.T) {
	type args struct {
		ctx          context.Context
		sqlStatement string
		args         []any
	}
	tests := []struct {
		name    string
		args    args
		want    *int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := WriteContext(tt.args.ctx, tt.args.sqlStatement, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteContext() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WriteContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadContext(t *testing.T) {
	sqlStatement := "some_sql_statement"
	errContextPath := func() context.Context {
		var err error
		var db *sql.DB
		var mock sqlmock.Sqlmock
		if db, mock, err = sqlmock.New(); err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		mock.ExpectBegin()
		mock.ExpectPrepare(sqlStatement).WillReturnError(errors.New("some_error"))
		tx, _ := db.Begin()

		txCtx := context.WithValue(context.TODO(), feather_sql_transaction.TransactionCtxKey{}, tx)

		return txCtx
	}
	errQueryPath := func() context.Context {
		var err error
		var db *sql.DB
		var mock sqlmock.Sqlmock
		if db, mock, err = sqlmock.New(); err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		mock.ExpectBegin()
		mock.ExpectPrepare(sqlStatement)
		mock.ExpectQuery(sqlStatement).WillReturnError(errors.New("some_error"))
		tx, _ := db.Begin()

		txCtx := context.WithValue(context.TODO(), feather_sql_transaction.TransactionCtxKey{}, tx)

		return txCtx
	}
	happyPath := func() context.Context {
		var err error
		var db *sql.DB
		var mock sqlmock.Sqlmock
		if db, mock, err = sqlmock.New(); err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		mock.ExpectBegin()
		mock.ExpectPrepare(sqlStatement)
		mock.ExpectQuery(sqlStatement).WillReturnRows(
			sqlmock.NewRows([]string{"id", "uuid", "title", "content"}).
				AddRow("1", "bea1b24d-0627-4ea0-aa2b-8af4c6c2a41c", "test", "test").
				AddRow("1", "bea1b24d-0627-4ea0-aa2b-8af4c6c2a41c", "test", "test"),
		)
		tx, _ := db.Begin()

		txCtx := context.WithValue(context.TODO(), feather_sql_transaction.TransactionCtxKey{}, tx)

		return txCtx
	}

	type args struct {
		ctx          context.Context
		sqlStatement string
		fn           ReadFunction
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Err Context Path",
			args: args{
				ctx:          errContextPath(),
				sqlStatement: sqlStatement,
				fn:           nil,
			},
			wantErr: true,
		},
		{
			name: "Err Query Path",
			args: args{
				ctx:          errQueryPath(),
				sqlStatement: sqlStatement,
				fn:           nil,
			},
			wantErr: true,
		},
		{
			name: "Err Fn Path",
			args: args{
				ctx:          happyPath(),
				sqlStatement: sqlStatement,
				fn: func(rows *sql.Rows) error {
					return errors.New("some_error")
				},
			},
			wantErr: true,
		},
		{
			name: "Happy Path",
			args: args{
				ctx:          happyPath(),
				sqlStatement: sqlStatement,
				fn: func(rows *sql.Rows) error {
					return nil
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ReadContext(tt.args.ctx, tt.args.sqlStatement, tt.args.fn); (err != nil) != tt.wantErr {
				t.Errorf("ReadContext() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestReadRowContext(t *testing.T) {

	sqlStatement := "some_sql_statement"
	errContextPath := func() context.Context {
		var err error
		var db *sql.DB
		var mock sqlmock.Sqlmock
		if db, mock, err = sqlmock.New(); err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		mock.ExpectBegin()
		mock.ExpectPrepare(sqlStatement).WillReturnError(errors.New("some_error"))
		tx, _ := db.Begin()

		txCtx := context.WithValue(context.TODO(), feather_sql_transaction.TransactionCtxKey{}, tx)

		return txCtx
	}
	errRowScanPath := func() context.Context {
		var err error
		var db *sql.DB
		var mock sqlmock.Sqlmock
		if db, mock, err = sqlmock.New(); err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		mock.ExpectBegin()
		mock.ExpectPrepare(sqlStatement)
		mock.ExpectQuery(sqlStatement).WillReturnError(errors.New("some_error"))

		tx, _ := db.Begin()

		txCtx := context.WithValue(context.TODO(), feather_sql_transaction.TransactionCtxKey{}, tx)

		return txCtx
	}
	errRowScanNoRowsPath := func() context.Context {
		var err error
		var db *sql.DB
		var mock sqlmock.Sqlmock
		if db, mock, err = sqlmock.New(); err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		mock.ExpectBegin()
		mock.ExpectPrepare(sqlStatement)
		mock.ExpectQuery(sqlStatement).WillReturnError(errors.New("db_column: no rows in result set"))

		tx, _ := db.Begin()

		txCtx := context.WithValue(context.TODO(), feather_sql_transaction.TransactionCtxKey{}, tx)

		return txCtx
	}
	happyPath := func() context.Context {
		var err error
		var db *sql.DB
		var mock sqlmock.Sqlmock
		if db, mock, err = sqlmock.New(); err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		mock.ExpectBegin()
		mock.ExpectPrepare(sqlStatement)
		mock.ExpectQuery(sqlStatement).WillReturnRows(
			sqlmock.NewRows([]string{"id", "uuid", "title", "content"}).
				AddRow("1", "bea1b24d-0627-4ea0-aa2b-8af4c6c2a41c", "test", "test").
				AddRow("1", "bea1b24d-0627-4ea0-aa2b-8af4c6c2a41c", "test", "test"),
		)
		tx, _ := db.Begin()

		txCtx := context.WithValue(context.TODO(), feather_sql_transaction.TransactionCtxKey{}, tx)

		return txCtx
	}

	var id, uuid, title, content string

	type args struct {
		ctx          context.Context
		sqlStatement string
		key          any
		dest         []any
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Err Context Path",
			args: args{
				ctx:          errContextPath(),
				sqlStatement: sqlStatement,
				key:          nil,
				dest:         nil,
			},
			wantErr: true,
		},
		{
			name: "Err RowScan Path",
			args: args{
				ctx:          errRowScanPath(),
				sqlStatement: sqlStatement,
				key:          nil,
				dest:         nil,
			},
			wantErr: true,
		},
		{
			name: "Err RowScan No Rows Path",
			args: args{
				ctx:          errRowScanNoRowsPath(),
				sqlStatement: sqlStatement,
				key:          nil,
				dest:         nil,
			},
			wantErr: true,
		},
		{
			name: "Happy Path",
			args: args{
				ctx:          happyPath(),
				sqlStatement: sqlStatement,
				key:          nil,
				dest:         []any{&id, &uuid, &title, &content},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ReadRowContext(tt.args.ctx, tt.args.sqlStatement, tt.args.key, tt.args.dest...); (err != nil) != tt.wantErr {
				t.Errorf("ReadRowContext() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestContext(t *testing.T) {

	sqlStatement := "some_sql_statement"

	errPreparePath := func() context.Context {
		var err error
		var db *sql.DB
		var mock sqlmock.Sqlmock
		if db, mock, err = sqlmock.New(); err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		mock.ExpectBegin()
		mock.ExpectPrepare(sqlStatement).WillReturnError(errors.New("some_error"))
		tx, _ := db.Begin()

		txCtx := context.WithValue(context.TODO(), feather_sql_transaction.TransactionCtxKey{}, tx)

		return txCtx
	}

	happyPath := func() context.Context {
		var err error
		var db *sql.DB
		var mock sqlmock.Sqlmock
		if db, mock, err = sqlmock.New(); err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		mock.ExpectBegin()
		mock.ExpectPrepare(sqlStatement)
		mock.ExpectExec(sqlStatement).WillReturnResult(sqlmock.NewResult(1, 1))

		tx, _ := db.Begin()

		txCtx := context.WithValue(context.TODO(), feather_sql_transaction.TransactionCtxKey{}, tx)

		return txCtx
	}

	type args struct {
		ctx          context.Context
		sqlStatement string
		fn           Function
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Err Prepare Path",
			args: args{
				ctx:          errPreparePath(),
				sqlStatement: sqlStatement,
				fn:           nil,
			},
			wantErr: true,
		},
		{
			name: "Err Fn Path",
			args: args{
				ctx:          happyPath(),
				sqlStatement: sqlStatement,
				fn: func(statement *sql.Stmt) error {
					return errors.New("some_error")
				},
			},
			wantErr: true,
		},
		{
			name: "Happy Path",
			args: args{
				ctx:          happyPath(),
				sqlStatement: sqlStatement,
				fn: func(statement *sql.Stmt) error {
					return nil
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Context(tt.args.ctx, tt.args.sqlStatement, tt.args.fn); (err != nil) != tt.wantErr {
				t.Errorf("Context() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_closeStatement(t *testing.T) {
	var err error
	var db *sql.DB
	var mock sqlmock.Sqlmock
	if db, mock, err = sqlmock.New(); err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	sqlStatement := "some_sql_statement"

	mock.ExpectBegin()
	mock.ExpectPrepare(sqlStatement).WillReturnCloseError(errors.New("some_error"))

	tx, _ := db.Begin()
	statement, _ := tx.Prepare(sqlStatement)

	type args struct {
		statement *sql.Stmt
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Happy Path",
			args: args{statement: statement},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			closeStatement(tt.args.statement)
		})
	}
}

func Test_closeResultSet(t *testing.T) {

	var err error
	var db *sql.DB
	var mock sqlmock.Sqlmock
	if db, mock, err = sqlmock.New(); err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	value := "something"

	mock.ExpectBegin()
	mock.ExpectPrepare(value)

	mock.ExpectQuery(value).WillReturnRows(sqlmock.NewRows([]string{"id", "uuid", "title", "content"}).
		AddRow("1", "bea1b24d-0627-4ea0-aa2b-8af4c6c2a41c", "test", "test").
		AddRow("1", "bea1b24d-0627-4ea0-aa2b-8af4c6c2a41c", "test", "test").
		CloseError(errors.New("some_error")))

	tx, _ := db.Begin()
	statement, _ := tx.Prepare(value)
	rows, _ := statement.Query()

	type args struct {
		rows *sql.Rows
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Happy Path",
			args: args{rows: rows},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			closeResultSet(tt.args.rows)
		})
	}
}
