package sql

func CreateSelectSQL(table string, value any, _ DriverName, paramHolder ParamHolder, fn01 ColumnFilterFunc) (string, error) {

	empty := ""
	separator := ", "

	var err error
	var sequence string
	if sequence, _, err = ParseColumnAsNameValueSequence(value, empty, empty, separator, 0, ColumnFilter, EvalNameOnly); err != nil {
		return "", ErrSQLGenerationFailed(err)
	}

	if fn01 == nil {
		return "SELECT " + sequence + " FROM " + table, nil
	}

	separator = " AND "
	var whereSequence string
	if whereSequence, _, err = ParseColumnAsNameValueSequence(value, empty, empty, separator, 0, fn01, paramHolder.EvalNameValue()); err != nil {
		return "", ErrSQLGenerationFailed(err)
	}

	return "SELECT " + sequence + " FROM " + table + " WHERE " + whereSequence, nil
}

func CreateInsertSQL(table string, value any, driverName DriverName, paramHolder ParamHolder) (string, error) {

	initChar := "("
	endChar := ")"
	separator := ", "
	empty := ""

	var err error
	var nameSequence string
	if nameSequence, _, err = ParseColumnAsNameValueSequence(value, initChar, endChar, separator, 0, NonePkGeneratedColumnFilter, EvalNameOnly); err != nil {
		return "", ErrSQLGenerationFailed(err)
	}

	var cont int
	var valueSequence string
	if valueSequence, cont, err = ParseColumnAsNameValueSequence(value, initChar, endChar, separator, 0, NonePkGeneratedColumnFilter, paramHolder.EvalValueOnly()); err != nil {
		return "", ErrSQLGenerationFailed(err)
	}

	var returning string
	if driverName == OracleDriverName {
		var pkNameSequence string
		if pkNameSequence, _, err = ParseColumnAsNameValueSequence(value, empty, empty, separator, cont, PkGeneratedColumnFilter, EvalNameOnly); err != nil {
			return "", ErrSQLGenerationFailed(err)
		}

		var pkValueSequence string
		if pkValueSequence, _, err = ParseColumnAsNameValueSequence(value, empty, empty, separator, cont, PkGeneratedColumnFilter, paramHolder.EvalValueOnly()); err != nil {
			return "", ErrSQLGenerationFailed(err)
		}
		returning = " RETURNING " + pkNameSequence + " INTO " + pkValueSequence
	}

	return "INSERT INTO " + table + " " + nameSequence + " VALUES " + valueSequence + returning, nil
}

func CreateUpsertSQL(table string, value any, driverName DriverName, paramHolder ParamHolder) (string, error) {
	return "TO IMPLEMENT", nil
}

func CreateUpdateSQL(table string, value any, _ DriverName, paramHolder ParamHolder, fn01 ColumnFilterFunc) (string, error) {

	separator := ", "
	empty := ""

	var err error
	var cont int
	var nameSequence string
	if nameSequence, cont, err = ParseColumnAsNameValueSequence(value, empty, empty, separator, 0, ColumnFilter, paramHolder.EvalNameValue()); err != nil {
		return "", ErrSQLGenerationFailed(err)
	}

	var whereSequence string
	if whereSequence, _, err = ParseColumnAsNameValueSequence(value, empty, empty, separator, cont, fn01, paramHolder.EvalNameValue()); err != nil {
		return "", ErrSQLGenerationFailed(err)
	}

	return "UPDATE " + table + " SET " + nameSequence + " WHERE " + whereSequence, nil
}

func CreateDeleteSQL(table string, value any, _ DriverName, paramHolder ParamHolder, fn01 ColumnFilterFunc) (string, error) {

	separator := ", "
	empty := ""

	var err error
	var valueSequence string
	if valueSequence, _, err = ParseColumnAsNameValueSequence(value, empty, empty, separator, 0, fn01, paramHolder.EvalNameValue()); err != nil {
		return "", ErrSQLGenerationFailed(err)
	}

	return "DELETE FROM " + table + " WHERE " + valueSequence, nil
}
