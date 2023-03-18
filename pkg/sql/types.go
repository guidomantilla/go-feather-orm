package sql

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
