package transaction

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"

	"jukebox-app/pkg/datasource"
)

func Test_NewRelationalTransactionHandler(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	relationalDataSource := datasource.NewMockRelationalDataSource(ctrl)
	relationalDataSource.EXPECT().GetDriverName().Return("some_database")

	handler := NewRelationalTransactionHandler(relationalDataSource)

	assert.NotNil(t, handler)
	assert.Equal(t, "some_database", handler.GetDriverName())
}

func Test_handleError(t *testing.T) {

	zapCore, logs := observer.New(zap.InfoLevel)
	logger := zap.New(zapCore)
	zap.ReplaceGlobals(logger)

	err := errors.New("some_err")
	handleError(err)

	filteredLogs := logs.Filter(func(entry observer.LoggedEntry) bool {
		return entry.Message == err.Error()
	})

	assert.Len(t, filteredLogs.All(), 1)
	assert.Equal(t, err.Error(), filteredLogs.All()[0].Message)
}

func Test_HandleTransaction_Ok(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	db, sqlMock, _ := sqlmock.New()
	sqlMock.ExpectBegin()

	relationalDataSource := datasource.NewMockRelationalDataSource(ctrl)
	relationalDataSource.EXPECT().GetDatabase().Return(db, nil)

	handler := NewRelationalTransactionHandler(relationalDataSource)

	err := handler.HandleTransaction(func(tx *sql.Tx) error {
		return nil
	})

	assert.Nil(t, err)
}

func Test_HandleTransaction_Begin_Err(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	db, sqlMock, _ := sqlmock.New()
	sqlMock.ExpectBegin().WillReturnError(errors.New("some_error"))

	relationalDataSource := datasource.NewMockRelationalDataSource(ctrl)
	relationalDataSource.EXPECT().GetDatabase().Return(db, nil)

	handler := NewRelationalTransactionHandler(relationalDataSource)

	err := handler.HandleTransaction(func(tx *sql.Tx) error {
		return nil
	})

	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

func Test_HandleTransaction_Defer_Err(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	db, sqlMock, _ := sqlmock.New()
	sqlMock.ExpectBegin()

	relationalDataSource := datasource.NewMockRelationalDataSource(ctrl)
	relationalDataSource.EXPECT().GetDatabase().Return(db, nil)

	handler := NewRelationalTransactionHandler(relationalDataSource)

	err := handler.HandleTransaction(func(tx *sql.Tx) error {
		return errors.New("some_error")
	})

	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}

func Test_HandleTransaction_Defer_Panic(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	db, sqlMock, _ := sqlmock.New()
	sqlMock.ExpectBegin()
	sqlMock.ExpectRollback()

	relationalDataSource := datasource.NewMockRelationalDataSource(ctrl)
	relationalDataSource.EXPECT().GetDatabase().Return(db, nil)

	handler := NewRelationalTransactionHandler(relationalDataSource)

	err := handler.HandleTransaction(func(tx *sql.Tx) error {
		panic("some_panic")
	})

	assert.Nil(t, err)
}

func Test_HandleTransaction_GetDatabase_Err(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	relationalDataSource := datasource.NewMockRelationalDataSource(ctrl)
	relationalDataSource.EXPECT().GetDatabase().Return(nil, errors.New("some_error"))

	handler := NewRelationalTransactionHandler(relationalDataSource)

	err := handler.HandleTransaction(func(tx *sql.Tx) error {
		return nil
	})

	assert.Error(t, err)
	assert.Equal(t, "some_error", err.Error())
}
