package sql

import "strings"

const (
	UndefinedDriverName DriverName = iota
	OracleDriverName
	MysqlDriverName
	PostgresDriverName
)

type DriverNameContext struct{}

type DriverName int

func (enum DriverName) String() string {

	switch enum {
	case UndefinedDriverName:
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

func (enum DriverName) ValueFromName(driverName string) DriverName {

	switch strings.ToLower(driverName) {
	case "oracle":
		return OracleDriverName
	case "mysql":
		return MysqlDriverName
	case "pgx":
		return PostgresDriverName
	}
	return UndefinedDriverName
}

func (enum DriverName) ValueFromCardinal(value int) DriverName {

	switch value {
	case int(OracleDriverName):
		return OracleDriverName
	case int(MysqlDriverName):
		return MysqlDriverName
	case int(PostgresDriverName):
		return PostgresDriverName
	}
	return UndefinedDriverName
}

//

type ParamHolder int

const (
	UndefinedParamHolder ParamHolder = iota
	NamedParamHolder
	NumberedParamHolder
	QuestionedParamHolder
)

func (enum ParamHolder) EvalNameValue() EvalColumnFunc {
	switch enum {
	case NamedParamHolder:
		return EvalNameValueNamed
	case NumberedParamHolder:
		return EvalNameValueNumbered
	case QuestionedParamHolder:
		return EvalNameValueQuestioned
	}
	return nil
}

func (enum ParamHolder) EvalValueOnly() EvalColumnFunc {
	switch enum {
	case NamedParamHolder:
		return EvalValueOnlyNamed
	case NumberedParamHolder:
		return EvalValueOnlyNumbered
	case QuestionedParamHolder:
		return EvalValueOnlyQuestioned
	}
	return nil
}

func (enum ParamHolder) ValueFromCardinal(value int) ParamHolder {

	switch value {
	case int(NamedParamHolder):
		return NamedParamHolder
	case int(NumberedParamHolder):
		return NumberedParamHolder
	case int(QuestionedParamHolder):
		return QuestionedParamHolder
	}
	return UndefinedParamHolder
}
