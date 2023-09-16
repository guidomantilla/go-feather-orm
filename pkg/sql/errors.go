package sql

import (
	"errors"
)

var (
	ErrAnyIsNil              = errors.New("value (any - interface{}) is nil")
	ErrAnyNotPointerOrStruct = errors.New("value (any - interface{}) not a pointer, not a struct")
	ErrAnyNotTagged          = errors.New("value (any - interface{}) not tagged")
	ErrColumnFilterFuncIsNil = errors.New("columnFilterFunc is nil")
	ErrEvalColumnFuncIsNil   = errors.New("evalColumnFunc is nil")
)

func ErrSQLGenerationFailed(errs ...error) error {
	return errors.New("sql generation failed: " + errors.Join(errs...).Error())
}
