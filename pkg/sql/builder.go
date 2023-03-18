package sql

func CreateSelectSQL(_ DriverName, table string, value any, fn01 ColumnFilterFunc, fn02 EvalColumnFunc) (string, error) {

	empty := ""
	separator := ", "

	var err error
	var sequence string
	if sequence, _, err = ParseColumnAsNameValueSequence(value, empty, empty, separator, 0, ColumnFilter, EvalNameOnly); err != nil {
		return "", err
	}

	if fn01 == nil {
		return "SELECT " + sequence + " FROM " + table, nil
	}

	separator = " AND "
	var whereSequence string
	if whereSequence, _, err = ParseColumnAsNameValueSequence(value, empty, empty, separator, 0, fn01, fn02); err != nil {
		return "", err
	}

	return "SELECT " + sequence + " FROM " + table + " WHERE " + whereSequence, nil
}

func CreateInsertSQL(driverName DriverName, table string, value any, fn02 EvalColumnFunc) (string, error) {

	initChar := "("
	endChar := ")"
	separator := ", "
	empty := ""

	var err error
	var nameSequence string
	if nameSequence, _, err = ParseColumnAsNameValueSequence(value, initChar, endChar, separator, 0, NonePkGeneratedColumnFilter, EvalNameOnly); err != nil {
		return "", err
	}

	var cont int
	var valueSequence string
	if valueSequence, cont, err = ParseColumnAsNameValueSequence(value, initChar, endChar, separator, 0, NonePkGeneratedColumnFilter, fn02); err != nil {
		return "", err
	}

	var returning string
	if driverName == OracleDriverName {
		var pkNameSequence string
		if pkNameSequence, _, err = ParseColumnAsNameValueSequence(value, empty, empty, separator, cont, PkGeneratedColumnFilter, EvalNameOnly); err != nil {
			return "", err
		}

		var pkValueSequence string
		if pkValueSequence, _, err = ParseColumnAsNameValueSequence(value, empty, empty, separator, cont, PkGeneratedColumnFilter, fn02); err != nil {
			return "", err
		}
		returning = " RETURNING " + pkNameSequence + " INTO " + pkValueSequence
	}

	return "INSERT INTO " + table + " " + nameSequence + " VALUES " + valueSequence + returning, nil
}

func CreateUpdateSQL(_ DriverName, table string, value any, fn01 ColumnFilterFunc, fn02 EvalColumnFunc) (string, error) {

	separator := ", "
	empty := ""

	var err error
	var cont int
	var nameSequence string
	if nameSequence, cont, err = ParseColumnAsNameValueSequence(value, empty, empty, separator, 0, ColumnFilter, fn02); err != nil {
		return "", err
	}

	var valueSequence string
	if valueSequence, _, err = ParseColumnAsNameValueSequence(value, empty, empty, separator, cont, fn01, fn02); err != nil {
		return "", err
	}

	return "UPDATE " + table + " SET " + nameSequence + " WHERE " + valueSequence, nil
}

func CreateDeleteSQL(_ DriverName, table string, value any, fn01 ColumnFilterFunc, fn02 EvalColumnFunc) (string, error) {

	separator := ", "
	empty := ""

	var err error
	var valueSequence string
	if valueSequence, _, err = ParseColumnAsNameValueSequence(value, empty, empty, separator, 0, fn01, fn02); err != nil {
		return "", err
	}

	return "DELETE FROM " + table + " WHERE " + valueSequence, nil
}
