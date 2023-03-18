package repository

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"

	"jukebox-app/pkg/datasource"
	"jukebox-app/pkg/transaction"
)

func Test_RelationalContext_Ok(t *testing.T) {

	var err error
	var db *sql.DB
	var mock sqlmock.Sqlmock
	if db, mock, err = sqlmock.New(); err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	sqlStatement := "some_sql_statement"

	mock.ExpectBegin()
	mock.ExpectPrepare(sqlStatement)
	mock.ExpectExec(sqlStatement).WillReturnResult(sqlmock.NewResult(1, 1))

	openFunc := datasource.OpenDataSourceFunc(func(driverName, dataSourceUrl string) (*sql.DB, error) {
		return db, nil
	})
	dataSource := datasource.NewRelationalDataSource("some_driver_name", "some_username", "some_password", ":username_:password", openFunc)

	database, _ := dataSource.GetDatabase()
	tx, _ := database.Begin()

	txCtx := context.WithValue(context.Background(), transaction.RelationalTransactionContext{}, tx)

	err = RelationalContext(txCtx, sqlStatement, func(statement *sql.Stmt) error {

		return nil
	})

	assert.Nil(t, err)
}

func Test_RelationalContext_Prepare_Err(t *testing.T) {

	var err error
	var db *sql.DB
	var mock sqlmock.Sqlmock
	if db, mock, err = sqlmock.New(); err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	sqlStatement := "some_sql_statement"

	mock.ExpectBegin()
	mock.ExpectPrepare(sqlStatement).WillReturnError(errors.New("some_error"))

	openFunc := datasource.OpenDataSourceFunc(func(driverName, dataSourceUrl string) (*sql.DB, error) {
		return db, nil
	})
	dataSource := datasource.NewRelationalDataSource("some_driver_name", "some_username", "some_password", ":username_:password", openFunc)

	database, _ := dataSource.GetDatabase()
	tx, _ := database.Begin()

	txCtx := context.WithValue(context.Background(), transaction.RelationalTransactionContext{}, tx)

	err = RelationalContext(txCtx, sqlStatement, func(statement *sql.Stmt) error {

		return nil
	})

	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

func Test_RelationalContext_Err(t *testing.T) {

	var err error
	var db *sql.DB
	var mock sqlmock.Sqlmock
	if db, mock, err = sqlmock.New(); err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	sqlStatement := "some_sql_statement"

	mock.ExpectBegin()
	mock.ExpectPrepare(sqlStatement)
	mock.ExpectExec(sqlStatement).WillReturnResult(sqlmock.NewResult(1, 1))

	openFunc := datasource.OpenDataSourceFunc(func(driverName, dataSourceUrl string) (*sql.DB, error) {
		return db, nil
	})
	dataSource := datasource.NewRelationalDataSource("some_driver_name", "some_username", "some_password", ":username_:password", openFunc)

	database, _ := dataSource.GetDatabase()
	tx, _ := database.Begin()

	txCtx := context.WithValue(context.Background(), transaction.RelationalTransactionContext{}, tx)

	err = RelationalContext(txCtx, sqlStatement, func(statement *sql.Stmt) error {

		return errors.New("some_error")
	})

	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

//

func Test_RelationalQueryContext_Ok(t *testing.T) {

	var err error
	var db *sql.DB
	var mock sqlmock.Sqlmock
	if db, mock, err = sqlmock.New(); err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	sqlStatement := "some_sql_statement"

	mock.ExpectBegin()
	mock.ExpectPrepare(sqlStatement)
	mock.ExpectQuery(sqlStatement).WillReturnRows(
		sqlmock.NewRows([]string{"id", "uuid", "title", "content"}).
			AddRow("1", "bea1b24d-0627-4ea0-aa2b-8af4c6c2a41c", "test", "test").
			AddRow("1", "bea1b24d-0627-4ea0-aa2b-8af4c6c2a41c", "test", "test"),
	)

	openFunc := datasource.OpenDataSourceFunc(func(driverName, dataSourceUrl string) (*sql.DB, error) {
		return db, nil
	})
	dataSource := datasource.NewRelationalDataSource("some_driver_name", "some_username", "some_password", ":username_:password", openFunc)

	database, _ := dataSource.GetDatabase()
	tx, _ := database.Begin()

	txCtx := context.WithValue(context.Background(), transaction.RelationalTransactionContext{}, tx)

	err = RelationalQueryContext(txCtx, sqlStatement, func(rows *sql.Rows) error {
		return nil
	})

	assert.Nil(t, err)
}

func Test_RelationalQueryContext_Prepare_Err(t *testing.T) {

	var err error
	var db *sql.DB
	var mock sqlmock.Sqlmock
	if db, mock, err = sqlmock.New(); err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	sqlStatement := "some_sql_statement"

	mock.ExpectBegin()
	mock.ExpectPrepare(sqlStatement).WillReturnError(errors.New("some_error"))

	openFunc := datasource.OpenDataSourceFunc(func(driverName, dataSourceUrl string) (*sql.DB, error) {
		return db, nil
	})
	dataSource := datasource.NewRelationalDataSource("some_driver_name", "some_username", "some_password", ":username_:password", openFunc)

	database, _ := dataSource.GetDatabase()
	tx, _ := database.Begin()

	txCtx := context.WithValue(context.Background(), transaction.RelationalTransactionContext{}, tx)

	err = RelationalQueryContext(txCtx, sqlStatement, func(rows *sql.Rows) error {
		return nil
	})

	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

func Test_RelationalQueryContext_Query_Err(t *testing.T) {

	var err error
	var db *sql.DB
	var mock sqlmock.Sqlmock
	if db, mock, err = sqlmock.New(); err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	sqlStatement := "some_sql_statement"

	mock.ExpectBegin()
	mock.ExpectPrepare(sqlStatement)
	mock.ExpectQuery(sqlStatement).WillReturnError(errors.New("some_error"))

	openFunc := datasource.OpenDataSourceFunc(func(driverName, dataSourceUrl string) (*sql.DB, error) {
		return db, nil
	})
	dataSource := datasource.NewRelationalDataSource("some_driver_name", "some_username", "some_password", ":username_:password", openFunc)

	database, _ := dataSource.GetDatabase()
	tx, _ := database.Begin()

	txCtx := context.WithValue(context.Background(), transaction.RelationalTransactionContext{}, tx)

	err = RelationalQueryContext(txCtx, sqlStatement, func(rows *sql.Rows) error {
		return nil
	})

	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

func Test_RelationalQueryContext_Err(t *testing.T) {

	var err error
	var db *sql.DB
	var mock sqlmock.Sqlmock
	if db, mock, err = sqlmock.New(); err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	sqlStatement := "some_sql_statement"

	mock.ExpectBegin()
	mock.ExpectPrepare(sqlStatement)
	mock.ExpectQuery(sqlStatement).WillReturnRows(
		sqlmock.NewRows([]string{"id", "uuid", "title", "content"}).
			AddRow("1", "bea1b24d-0627-4ea0-aa2b-8af4c6c2a41c", "test", "test").
			AddRow("1", "bea1b24d-0627-4ea0-aa2b-8af4c6c2a41c", "test", "test"),
	)

	openFunc := datasource.OpenDataSourceFunc(func(driverName, dataSourceUrl string) (*sql.DB, error) {
		return db, nil
	})
	dataSource := datasource.NewRelationalDataSource("some_driver_name", "some_username", "some_password", ":username_:password", openFunc)

	database, _ := dataSource.GetDatabase()
	tx, _ := database.Begin()

	txCtx := context.WithValue(context.Background(), transaction.RelationalTransactionContext{}, tx)

	err = RelationalQueryContext(txCtx, sqlStatement, func(rows *sql.Rows) error {
		return errors.New("some_error")
	})

	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

//

func Test_RelationalWriteContext_Ok(t *testing.T) {

	var err error
	var db *sql.DB
	var mock sqlmock.Sqlmock
	if db, mock, err = sqlmock.New(); err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	sqlStatement := "insert_sql_statement"

	mock.ExpectBegin()
	mock.ExpectPrepare(sqlStatement)
	mock.ExpectExec(sqlStatement).WillReturnResult(sqlmock.NewResult(1, 1))

	openFunc := datasource.OpenDataSourceFunc(func(driverName, dataSourceUrl string) (*sql.DB, error) {
		return db, nil
	})
	dataSource := datasource.NewRelationalDataSource("some_driver_name", "some_username", "some_password", ":username_:password", openFunc)

	database, _ := dataSource.GetDatabase()
	tx, _ := database.Begin()

	txCtx := context.WithValue(context.Background(), transaction.RelationalTransactionContext{}, tx)

	var id *int64
	id, err = RelationalWriteContext(txCtx, sqlStatement, "", "")

	assert.Nil(t, err)
	assert.NotNil(t, id)
	assert.Equal(t, int64(1), *id)
}

func Test_RelationalWriteContext_Exec_Err(t *testing.T) {

	var err error
	var db *sql.DB
	var mock sqlmock.Sqlmock
	if db, mock, err = sqlmock.New(); err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	sqlStatement := "insert_sql_statement"

	mock.ExpectBegin()
	mock.ExpectPrepare(sqlStatement)
	mock.ExpectExec(sqlStatement).WillReturnError(errors.New("some_error"))

	openFunc := datasource.OpenDataSourceFunc(func(driverName, dataSourceUrl string) (*sql.DB, error) {
		return db, nil
	})
	dataSource := datasource.NewRelationalDataSource("some_driver_name", "some_username", "some_password", ":username_:password", openFunc)

	database, _ := dataSource.GetDatabase()
	tx, _ := database.Begin()

	txCtx := context.WithValue(context.Background(), transaction.RelationalTransactionContext{}, tx)

	var id *int64
	id, err = RelationalWriteContext(txCtx, sqlStatement, "", "")

	assert.NotNil(t, err)
	assert.Nil(t, id)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

func Test_RelationalWriteContext_LastInsertId_Err(t *testing.T) {

	var err error
	var db *sql.DB
	var mock sqlmock.Sqlmock
	if db, mock, err = sqlmock.New(); err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	sqlStatement := "insert_sql_statement"

	mock.ExpectBegin()
	mock.ExpectPrepare(sqlStatement)
	mock.ExpectExec(sqlStatement).WillReturnResult(sqlmock.NewErrorResult(errors.New("some_error")))

	openFunc := datasource.OpenDataSourceFunc(func(driverName, dataSourceUrl string) (*sql.DB, error) {
		return db, nil
	})
	dataSource := datasource.NewRelationalDataSource("some_driver_name", "some_username", "some_password", ":username_:password", openFunc)

	database, _ := dataSource.GetDatabase()
	tx, _ := database.Begin()

	txCtx := context.WithValue(context.Background(), transaction.RelationalTransactionContext{}, tx)

	var id *int64
	id, err = RelationalWriteContext(txCtx, sqlStatement, "", "")

	assert.NotNil(t, err)
	assert.Nil(t, id)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

//

func Test_RelationalQueryRowContext_Ok(t *testing.T) {

	var err error
	var db *sql.DB
	var mock sqlmock.Sqlmock
	if db, mock, err = sqlmock.New(); err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	sqlStatement := "some_sql_statement"

	mock.ExpectBegin()
	mock.ExpectPrepare(sqlStatement)
	mock.ExpectQuery(sqlStatement).WillReturnRows(
		sqlmock.NewRows([]string{"id", "uuid", "title", "content"}).
			AddRow("1", "bea1b24d-0627-4ea0-aa2b-8af4c6c2a41c", "test", "test").
			AddRow("1", "bea1b24d-0627-4ea0-aa2b-8af4c6c2a41c", "test", "test"),
	)

	openFunc := datasource.OpenDataSourceFunc(func(driverName, dataSourceUrl string) (*sql.DB, error) {
		return db, nil
	})
	dataSource := datasource.NewRelationalDataSource("some_driver_name", "some_username", "some_password", ":username_:password", openFunc)

	database, _ := dataSource.GetDatabase()
	tx, _ := database.Begin()

	txCtx := context.WithValue(context.Background(), transaction.RelationalTransactionContext{}, tx)

	var id, uuid, title, content string
	err = RelationalQueryRowContext(txCtx, sqlStatement, "", &id, &uuid, &title, &content)

	assert.Nil(t, err)
	assert.NotEmpty(t, id)
	assert.NotEmpty(t, uuid)
	assert.NotEmpty(t, title)
	assert.NotEmpty(t, content)
}

func Test_RelationalQueryRowContext_Prepare_Err(t *testing.T) {

	var err error
	var db *sql.DB
	var mock sqlmock.Sqlmock
	if db, mock, err = sqlmock.New(); err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	sqlStatement := "some_sql_statement"

	mock.ExpectBegin()
	mock.ExpectPrepare(sqlStatement).WillReturnError(errors.New("some_error"))

	openFunc := datasource.OpenDataSourceFunc(func(driverName, dataSourceUrl string) (*sql.DB, error) {
		return db, nil
	})
	dataSource := datasource.NewRelationalDataSource("some_driver_name", "some_username", "some_password", ":username_:password", openFunc)

	database, _ := dataSource.GetDatabase()
	tx, _ := database.Begin()

	txCtx := context.WithValue(context.Background(), transaction.RelationalTransactionContext{}, tx)

	var id, uuid, title, content string
	err = RelationalQueryRowContext(txCtx, sqlStatement, "", &id, &uuid, &title, &content)

	assert.NotNil(t, err)
	assert.Empty(t, id)
	assert.Empty(t, uuid)
	assert.Empty(t, title)
	assert.Empty(t, content)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

func Test_RelationalQueryRowContext_Scan_Err(t *testing.T) {

	var err error
	var db *sql.DB
	var mock sqlmock.Sqlmock
	if db, mock, err = sqlmock.New(); err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	sqlStatement := "some_sql_statement"

	mock.ExpectBegin()
	mock.ExpectPrepare(sqlStatement)
	mock.ExpectQuery(sqlStatement).WillReturnError(errors.New("some_error"))

	openFunc := datasource.OpenDataSourceFunc(func(driverName, dataSourceUrl string) (*sql.DB, error) {
		return db, nil
	})
	dataSource := datasource.NewRelationalDataSource("some_driver_name", "some_username", "some_password", ":username_:password", openFunc)

	database, _ := dataSource.GetDatabase()
	tx, _ := database.Begin()

	txCtx := context.WithValue(context.Background(), transaction.RelationalTransactionContext{}, tx)

	var id, uuid, title, content string
	err = RelationalQueryRowContext(txCtx, sqlStatement, "", &id, &uuid, &title, &content)

	assert.NotNil(t, err)
	assert.Empty(t, id)
	assert.Empty(t, uuid)
	assert.Empty(t, title)
	assert.Empty(t, content)
	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

func Test_RelationalQueryRowContext_Scan_No_Rows_Err(t *testing.T) {

	var err error
	var db *sql.DB
	var mock sqlmock.Sqlmock
	if db, mock, err = sqlmock.New(); err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	sqlStatement := "some_sql_statement"

	mock.ExpectBegin()
	mock.ExpectPrepare(sqlStatement)
	mock.ExpectQuery(sqlStatement).WillReturnError(errors.New("sql: no rows in result set"))

	openFunc := datasource.OpenDataSourceFunc(func(driverName, dataSourceUrl string) (*sql.DB, error) {
		return db, nil
	})
	dataSource := datasource.NewRelationalDataSource("some_driver_name", "some_username", "some_password", ":username_:password", openFunc)

	database, _ := dataSource.GetDatabase()
	tx, _ := database.Begin()

	txCtx := context.WithValue(context.Background(), transaction.RelationalTransactionContext{}, tx)

	var id, uuid, title, content string
	err = RelationalQueryRowContext(txCtx, sqlStatement, "1", &id, &uuid, &title, &content)

	assert.NotNil(t, err)
	assert.Empty(t, id)
	assert.Empty(t, uuid)
	assert.Empty(t, title)
	assert.Empty(t, content)
	assert.Error(t, err)
	assert.Equal(t, "row with key 1 not found", err.Error())
}

//

func Test_closeStatement_Err(t *testing.T) {

	zapCore, logs := observer.New(zap.InfoLevel)
	logger := zap.New(zapCore)
	zap.ReplaceGlobals(logger)

	var err error
	var db *sql.DB
	var mock sqlmock.Sqlmock
	if db, mock, err = sqlmock.New(); err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	sqlStatement := "some_sql_statement"

	mock.ExpectBegin()
	mock.ExpectPrepare(sqlStatement).WillReturnCloseError(errors.New("some_error"))

	openFunc := datasource.OpenDataSourceFunc(func(driverName, dataSourceUrl string) (*sql.DB, error) {
		return db, nil
	})
	dataSource := datasource.NewRelationalDataSource("some_driver_name", "some_username", "some_password", ":username_:password", openFunc)

	database, _ := dataSource.GetDatabase()
	tx, _ := database.Begin()
	statement, _ := tx.Prepare(sqlStatement)

	closeStatement(statement)

	assert.Len(t, logs.All(), 1)
	assert.Equal(t, Error_Closing_Statement, logs.All()[0].Message)

	assert.Nil(t, err)
}

func Test_closeResultSet_Err(t *testing.T) {

	zapCore, logs := observer.New(zap.InfoLevel)
	logger := zap.New(zapCore)
	zap.ReplaceGlobals(logger)

	var err error
	var db *sql.DB
	var mock sqlmock.Sqlmock
	if db, mock, err = sqlmock.New(); err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	sqlStatement := "some_sql_statement"

	mock.ExpectBegin()
	mock.ExpectPrepare(sqlStatement)
	mock.ExpectQuery(sqlStatement).WillReturnRows(
		sqlmock.NewRows([]string{"id", "uuid", "title", "content"}).
			AddRow("1", "bea1b24d-0627-4ea0-aa2b-8af4c6c2a41c", "test", "test").
			AddRow("1", "bea1b24d-0627-4ea0-aa2b-8af4c6c2a41c", "test", "test").
			CloseError(errors.New("some_error")),
	)

	openFunc := datasource.OpenDataSourceFunc(func(driverName, dataSourceUrl string) (*sql.DB, error) {
		return db, nil
	})
	dataSource := datasource.NewRelationalDataSource("some_driver_name", "some_username", "some_password", ":username_:password", openFunc)

	database, _ := dataSource.GetDatabase()
	tx, _ := database.Begin()
	statement, _ := tx.Prepare(sqlStatement)
	rows, _ := statement.Query()

	closeResultSet(rows)

	assert.Len(t, logs.All(), 1)
	assert.Equal(t, Error_Closing_ResultSet, logs.All()[0].Message)

	assert.Nil(t, err)
}

//
