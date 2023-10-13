package sql

import (
	"reflect"
	"strings"
)

const (
	TagColumnKey      = "db"
	TagPkValue        = "pk"
	TagUqValue        = "uq"
	TagGeneratedValue = "generated"
)

func RetrieveColumnNames(value any, columnFilterFunc ColumnFilterFunc) ([]string, error) {

	var err error
	var reflectedValue *reflect.Value
	if reflectedValue, err = RetrieveReflectedStruct(value); err != nil {
		return nil, err
	}

	if columnFilterFunc == nil {
		return nil, ErrColumnFilterFuncIsNil
	}

	fields := RetrieveFields(*reflectedValue)

	columnNames := make([]string, 0)
	for _, field := range fields {
		reflectedTagKey := field.Tag.Get(TagColumnKey)
		if reflectedTagKey == "-" || reflectedTagKey == "" {
			continue
		}

		values := strings.Split(reflectedTagKey, ",")
		/*
			if values == nil || len(values) == 0 || (len(values) == 1 && values[0] == "") {
				continue
			}
		*/
		if columnFilterFunc(values) {
			columnNames = append(columnNames, values[0])
		}
	}

	return columnNames, nil
}

type ColumnFilterFunc func(values []string) bool

func KeyColumnFilter(values []string) bool {
	return PkColumnFilter(values) || UqColumnFilter(values)
}

func NonePkGeneratedColumnFilter(values []string) bool {
	return !PkGeneratedColumnFilter(values)
}

func NoneUqGeneratedColumnFilter(values []string) bool {
	return !UqGeneratedColumnFilter(values)
}

func PkGeneratedColumnFilter(values []string) bool {
	return PkColumnFilter(values) && len(values) >= 3 && values[2] == TagGeneratedValue
}

func UqGeneratedColumnFilter(values []string) bool {
	return UqColumnFilter(values) && len(values) >= 3 && values[2] == TagGeneratedValue
}

func GeneratedColumnFilter(values []string) bool {
	return ColumnFilter(values) && (len(values) >= 2 && values[1] == TagGeneratedValue || len(values) >= 3 && values[2] == TagGeneratedValue)
}

//

func PkColumnFilter(values []string) bool {
	return len(values) >= 2 && values[0] != TagPkValue && values[0] != TagUqValue && values[1] == TagPkValue
}

func UqColumnFilter(values []string) bool {
	return len(values) >= 2 && values[0] != TagPkValue && values[0] != TagUqValue && values[1] == TagUqValue
}

func ColumnFilter(values []string) bool {
	return len(values) >= 1 && values[0] != TagPkValue && values[0] != TagUqValue
}

func NoneColumnFilter(_ []string) bool {
	return false
}

func RetrieveFields(reflectedValue reflect.Value) []*reflect.StructField {

	fields := make([]*reflect.StructField, 0)
	reflectedType := reflectedValue.Type()
	numFields := reflectedType.NumField()

	for i := 0; i < numFields; i++ {
		reflectedField := reflectedType.Field(i)
		if !reflectedField.Anonymous && len(reflectedField.PkgPath) > 0 {
			continue
		}

		fields = append(fields, &reflectedField)
	}

	return fields
}

func RetrieveReflectedStruct(value any) (*reflect.Value, error) {

	if value == nil {
		return nil, ErrAnyIsNil
	}

	reflectedValue := reflect.ValueOf(value)
	if reflectedValue.Kind() == reflect.Ptr {
		reflectedValue = reflectedValue.Elem()
		return &reflectedValue, nil
	}

	if reflectedValue.Kind() == reflect.Struct {
		return &reflectedValue, nil
	}

	return nil, ErrAnyNotPointerOrStruct
}
