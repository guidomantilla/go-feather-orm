package feather_sql

import (
	feather_sql_parsing "github.com/guidomantilla/go-feather-sql/pkg/feather-sql-parsing"
)

const (
	UnknownDriverName DriverName = iota - 1
	OracleDriverName
	MysqlDriverName
	PostgresDriverName
)

type DriverNameContext struct{}

type DriverName int

func (enum DriverName) String() string {

	switch enum {
	case UnknownDriverName:
		return "unknown"
	case OracleDriverName:
		return "oracle"
	case MysqlDriverName:
		return "mysql"
	case PostgresDriverName:
		return "pgx"
	}
	return "unknown"
}

func (enum DriverName) ValueOf(driverName string) DriverName {

	switch driverName {
	case UnknownDriverName.String():
		return UnknownDriverName
	case OracleDriverName.String():
		return OracleDriverName
	case MysqlDriverName.String():
		return MysqlDriverName
	case PostgresDriverName.String():
		return PostgresDriverName
	}
	return UnknownDriverName
}

//

type ParamHolder int

const (
	NamedParamHolder ParamHolder = iota
	NumberedParamHolder
	QuestionedParamHolder
)

func (enum ParamHolder) EvalNameValue() feather_sql_parsing.EvalColumnFunc {
	switch enum {
	case NamedParamHolder:
		return feather_sql_parsing.EvalNameValueNamed
	case NumberedParamHolder:
		return feather_sql_parsing.EvalNameValueNumbered
	case QuestionedParamHolder:
		return feather_sql_parsing.EvalNameValueQuestioned
	}
	return nil
}

func (enum ParamHolder) EvalValueOnly() feather_sql_parsing.EvalColumnFunc {
	switch enum {
	case NamedParamHolder:
		return feather_sql_parsing.EvalValueOnlyNamed
	case NumberedParamHolder:
		return feather_sql_parsing.EvalValueOnlyNumbered
	case QuestionedParamHolder:
		return feather_sql_parsing.EvalValueOnlyQuestioned
	}
	return nil
}
