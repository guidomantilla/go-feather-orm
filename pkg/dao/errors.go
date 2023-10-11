package dao

import (
	"errors"

	feather_commons_errors "github.com/guidomantilla/go-feather-commons/pkg/errors"
)

func ErrWriteContextFailed(errs ...error) error {
	return errors.New("db write context failed: " + feather_commons_errors.ErrJoin(errs...).Error())
}

func ErrReadContextFailed(errs ...error) error {
	return errors.New("db read context failed: " + feather_commons_errors.ErrJoin(errs...).Error())
}

func ErrReadRowContextFailed(errs ...error) error {
	return errors.New("db read row context failed: " + feather_commons_errors.ErrJoin(errs...).Error())
}

func ErrContextFailed(errs ...error) error {
	return errors.New("db context failed: " + feather_commons_errors.ErrJoin(errs...).Error())
}

func ErrMutateFailed(errs ...error) error {
	return errors.New("db mutation failed: " + feather_commons_errors.ErrJoin(errs...).Error())
}

func ErrQueryOneFailed(errs ...error) error {
	return errors.New("db query one failed: " + feather_commons_errors.ErrJoin(errs...).Error())
}

func ErrQueryManyFailed(errs ...error) error {
	return errors.New("db query many failed: " + feather_commons_errors.ErrJoin(errs...).Error())
}
