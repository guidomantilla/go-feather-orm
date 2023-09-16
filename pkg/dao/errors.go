package dao

import (
	"errors"
)

func ErrWriteContextFailed(errs ...error) error {
	return errors.New("db write context failed: " + errors.Join(errs...).Error())
}

func ErrReadContextFailed(errs ...error) error {
	return errors.New("db read context failed: " + errors.Join(errs...).Error())
}

func ErrReadRowContextFailed(errs ...error) error {
	return errors.New("db read row context failed: " + errors.Join(errs...).Error())
}

func ErrContextFailed(errs ...error) error {
	return errors.New("db context failed: " + errors.Join(errs...).Error())
}

func ErrSaveFailed(errs ...error) error {
	return errors.New("db save failed: " + errors.Join(errs...).Error())
}

func ErrUpdateFailed(errs ...error) error {
	return errors.New("db update failed: " + errors.Join(errs...).Error())
}

func ErrDeleteFailed(errs ...error) error {
	return errors.New("db delete failed: " + errors.Join(errs...).Error())
}

func ErrFindByIdFailed(errs ...error) error {
	return errors.New("db find by id failed: " + errors.Join(errs...).Error())
}

func ErrFindAllFailed(errs ...error) error {
	return errors.New("db find all failed: " + errors.Join(errs...).Error())
}
