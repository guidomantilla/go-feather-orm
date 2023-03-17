package feather_sql_builder

import (
	feather_sql "github.com/guidomantilla/go-feather-sql/pkg/feather-sql"
	feather_sql_parsing "github.com/guidomantilla/go-feather-sql/pkg/feather-sql-parsing"
	feather_sql_reflection "github.com/guidomantilla/go-feather-sql/pkg/feather-sql-reflection"
)

func CreateSelectSQL(_ feather_sql.DriverName, table string, value any, fn01 feather_sql_reflection.ColumnFilterFunc, fn02 feather_sql_parsing.EvalColumnFunc) (string, error) {

	empty := ""
	separator := ", "

	var err error
	var sequence string
	if sequence, _, err = feather_sql_parsing.ParseColumnAsNameValueSequence(value, empty, empty, separator, 0, feather_sql_reflection.ColumnFilter, feather_sql_parsing.EvalNameOnly); err != nil {
		return "", err
	}

	if fn01 == nil {
		return "SELECT " + sequence + " FROM " + table, nil
	}

	separator = " AND "
	var whereSequence string
	if whereSequence, _, err = feather_sql_parsing.ParseColumnAsNameValueSequence(value, empty, empty, separator, 0, fn01, fn02); err != nil {
		return "", err
	}

	return "SELECT " + sequence + " FROM " + table + " WHERE " + whereSequence, nil
}

func CreateInsertSQL(driverName feather_sql.DriverName, table string, value any, fn02 feather_sql_parsing.EvalColumnFunc) (string, error) {

	initChar := "("
	endChar := ")"
	separator := ", "
	empty := ""

	var err error
	var nameSequence string
	if nameSequence, _, err = feather_sql_parsing.ParseColumnAsNameValueSequence(value, initChar, endChar, separator, 0, feather_sql_reflection.NonePkGeneratedColumnFilter, feather_sql_parsing.EvalNameOnly); err != nil {
		return "", err
	}

	var cont int
	var valueSequence string
	if valueSequence, cont, err = feather_sql_parsing.ParseColumnAsNameValueSequence(value, initChar, endChar, separator, 0, feather_sql_reflection.NonePkGeneratedColumnFilter, fn02); err != nil {
		return "", err
	}

	var returning string
	if driverName == feather_sql.OracleDriverName {
		var pkNameSequence string
		if pkNameSequence, _, err = feather_sql_parsing.ParseColumnAsNameValueSequence(value, empty, empty, separator, cont, feather_sql_reflection.PkGeneratedColumnFilter, feather_sql_parsing.EvalNameOnly); err != nil {
			return "", err
		}

		var pkValueSequence string
		if pkValueSequence, _, err = feather_sql_parsing.ParseColumnAsNameValueSequence(value, empty, empty, separator, cont, feather_sql_reflection.PkGeneratedColumnFilter, fn02); err != nil {
			return "", err
		}
		returning = " RETURNING " + pkNameSequence + " INTO " + pkValueSequence
	}

	return "INSERT INTO " + table + " " + nameSequence + " VALUES " + valueSequence + returning, nil
}

func CreateUpdateSQL(_ feather_sql.DriverName, table string, value any, fn01 feather_sql_reflection.ColumnFilterFunc, fn02 feather_sql_parsing.EvalColumnFunc) (string, error) {

	separator := ", "
	empty := ""

	var err error
	var cont int
	var nameSequence string
	if nameSequence, cont, err = feather_sql_parsing.ParseColumnAsNameValueSequence(value, empty, empty, separator, 0, feather_sql_reflection.ColumnFilter, fn02); err != nil {
		return "", err
	}

	var valueSequence string
	if valueSequence, _, err = feather_sql_parsing.ParseColumnAsNameValueSequence(value, empty, empty, separator, cont, fn01, fn02); err != nil {
		return "", err
	}

	return "UPDATE " + table + " SET " + nameSequence + " WHERE " + valueSequence, nil
}

func CreateDeleteSQL(_ feather_sql.DriverName, table string, value any, fn01 feather_sql_reflection.ColumnFilterFunc, fn02 feather_sql_parsing.EvalColumnFunc) (string, error) {

	separator := ", "
	empty := ""

	var err error
	var valueSequence string
	if valueSequence, _, err = feather_sql_parsing.ParseColumnAsNameValueSequence(value, empty, empty, separator, 0, fn01, fn02); err != nil {
		return "", err
	}

	return "DELETE FROM " + table + " WHERE " + valueSequence, nil
}
