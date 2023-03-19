package sql

import "errors"

var (
	ErrAnyIsNil              = errors.New("value (any - interface{}) is nil")
	ErrAnyNotPointerOrStruct = errors.New("value (any - interface{}) not a pointer, not a struct")
	ErrColumnFilterFuncIsNil = errors.New("columnFilterFunc is nil")
	ErrEvalColumnFuncIsNil   = errors.New("evalColumnFunc is nil")
)
