package dao

import (
	"errors"

	"go.uber.org/multierr"
)

func ErrWriteContextFailed(errs ...error) error {
	return errors.New("db write context failed: " + multierr.Combine(errs...).Error())
}

func ErrReadContextFailed(errs ...error) error {
	return errors.New("db read context failed: " + multierr.Combine(errs...).Error())
}

func ErrReadRowContextFailed(errs ...error) error {
	return errors.New("db read row context failed: " + multierr.Combine(errs...).Error())
}

func ErrContextFailed(errs ...error) error {
	return errors.New("db context failed: " + multierr.Combine(errs...).Error())
}

func ErrSaveFailed(errs ...error) error {
	return errors.New("db save failed: " + multierr.Combine(errs...).Error())
}

func ErrUpdateFailed(errs ...error) error {
	return errors.New("db update failed: " + multierr.Combine(errs...).Error())
}

func ErrDeleteFailed(errs ...error) error {
	return errors.New("db delete failed: " + multierr.Combine(errs...).Error())
}

func ErrFindByIdFailed(errs ...error) error {
	return errors.New("db find by id failed: " + multierr.Combine(errs...).Error())
}

func ErrFindAllFailed(errs ...error) error {
	return errors.New("db find all failed: " + multierr.Combine(errs...).Error())
}
