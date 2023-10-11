package dao

import (
	"context"
	"fmt"
	"strings"

	feather_commons_log "github.com/guidomantilla/go-feather-commons/pkg/log"

	feather_sql_datasource "github.com/guidomantilla/go-feather-sql/pkg/datasource"
	feather_sql "github.com/guidomantilla/go-feather-sql/pkg/sql"
)

type DefaultCrudDao struct {
	//datasourceContext datasource.DatasourceContext
	driverName        feather_sql.DriverName
	paramHolder       feather_sql.ParamHolder
	table             string
	statementCreate   string
	statementUpdate   string
	statementDelete   string
	statementFindById string
	statementFindAll  string
}

func NewDefaultCrudDao(datasourceContext feather_sql_datasource.DatasourceContext, table string, model any) *DefaultCrudDao {

	if datasourceContext == nil {
		feather_commons_log.Fatal(fmt.Sprintf("starting up - error setting up %s dao: datasourceContext is nil", table))
	}

	if strings.TrimSpace(table) == "" {
		feather_commons_log.Fatal(fmt.Sprintf("starting up - error setting up %s dao: table is empty", table))
	}

	if model == nil {
		feather_commons_log.Fatal(fmt.Sprintf("starting up - error setting up %s dao: model is nil", table))
	}

	driverName := datasourceContext.GetDriverName()
	paramHolder := datasourceContext.GetParamHolder()

	statementCreate, err := feather_sql.CreateInsertSQL(table, model, driverName, paramHolder)
	if err != nil {
		feather_commons_log.Fatal(fmt.Sprintf("starting up - error setting up %s dao: %s", table, err.Error()))
	}

	statementUpdate, err := feather_sql.CreateUpdateSQL(table, model, driverName, paramHolder, feather_sql.PkColumnFilter)
	if err != nil {
		feather_commons_log.Fatal(fmt.Sprintf("starting up - error setting up %s dao: %s", table, err.Error()))
	}

	statementDelete, err := feather_sql.CreateDeleteSQL(table, model, driverName, paramHolder, feather_sql.PkColumnFilter)
	if err != nil {
		feather_commons_log.Fatal(fmt.Sprintf("starting up - error setting up %s dao: %s", table, err.Error()))
	}

	statementFindById, err := feather_sql.CreateSelectSQL(table, model, driverName, paramHolder, feather_sql.PkColumnFilter)
	if err != nil {
		feather_commons_log.Fatal(fmt.Sprintf("starting up - error setting up %s dao: %s", table, err.Error()))
	}

	statementFindAll, err := feather_sql.CreateSelectSQL(table, model, driverName, paramHolder, nil)
	if err != nil {
		feather_commons_log.Fatal(fmt.Sprintf("starting up - error setting up %s dao: %s", table, err.Error()))
	}

	return &DefaultCrudDao{
		table:             table,
		driverName:        driverName,
		paramHolder:       paramHolder,
		statementCreate:   statementCreate,
		statementUpdate:   statementUpdate,
		statementDelete:   statementDelete,
		statementFindById: statementFindById,
		statementFindAll:  statementFindAll,
	}
}

func (dao *DefaultCrudDao) Save(ctx context.Context, args ...any) (*int64, error) {
	return MutateOne(ctx, dao.statementCreate, args...)
}

func (dao *DefaultCrudDao) Update(ctx context.Context, args ...any) error {
	_, err := MutateOne(ctx, dao.statementUpdate, args...)
	return err
}

func (dao *DefaultCrudDao) Delete(ctx context.Context, id any) error {
	_, err := MutateOne(ctx, dao.statementDelete, id)
	return err
}

func (dao *DefaultCrudDao) FindById(ctx context.Context, id any, dest ...any) error {
	return QueryOne(ctx, dao.statementFindById, id, dest...)
}

func (dao *DefaultCrudDao) ExistsById(ctx context.Context, id any, dest ...any) bool {
	return Exists(ctx, dao.statementFindById, id, dest...)
}

func (dao *DefaultCrudDao) FindAll(ctx context.Context, fn ReadFunction) error {
	return QueryMany(ctx, dao.statementFindAll, fn)
}
