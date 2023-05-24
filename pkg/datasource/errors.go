package datasource

import (
	"errors"
	"go.uber.org/multierr"
)

func ErrDBConnectionFailed(errs ...error) error {
	return errors.New("db connection failed: " + multierr.Combine(errs...).Error())
}
