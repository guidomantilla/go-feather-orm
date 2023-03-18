package feather_sql_reflection

import "errors"

var (
	ErrAnyIsNil              = errors.New("value (any - interface{}) is nil")
	ErrAnyNotPointerOrStruct = errors.New("value (any - interface{}) not a pointer, not a struct")
)
