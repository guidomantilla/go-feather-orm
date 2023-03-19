package sql

import (
	"errors"
	"strconv"
	"strings"
)

func ParseColumnAsNameValueSequence(value any, initChar string, endChar string, separator string, cont int, fn01 ColumnFilterFunc, fn02 EvalColumnFunc) (string, int, error) {

	var err error
	var columnNames []string
	if columnNames, err = RetrieveColumnNames(value, fn01); err != nil {
		return "", 0, err
	}

	if len(columnNames) == 0 {
		return "", 0, errors.New("value (any - interface{}) not tagged")
	}

	if fn02 == nil {
		return "", 0, ErrEvalColumnFuncIsNil
	}

	sequence := initChar
	for i, columnName := range columnNames {
		sequence += fn02(columnName, cont+i+1, separator)
	}

	last := strings.LastIndex(sequence, separator)
	sequence = sequence[0:last] + endChar

	return sequence, len(columnNames), nil
}

//

type EvalColumnFunc func(name string, cont int, separator string) string

func EvalNameOnly(name string, _ int, separator string) string {
	return name + separator
}

// Named

func EvalNameValueNamed(name string, _ int, separator string) string {
	return name + " = :" + name + separator
}

func EvalValueOnlyNamed(name string, _ int, separator string) string {
	return ":" + name + separator
}

// Numbered

func EvalNameValueNumbered(name string, cont int, separator string) string {
	return name + " = :" + strconv.Itoa(cont) + separator
}

func EvalValueOnlyNumbered(_ string, cont int, separator string) string {
	return ":" + strconv.Itoa(cont) + separator
}

// Questioned

func EvalNameValueQuestioned(name string, _ int, separator string) string {
	return name + " = :?" + separator
}

func EvalValueOnlyQuestioned(_ string, _ int, separator string) string {
	return ":?" + separator
}
