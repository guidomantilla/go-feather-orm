package dao

import (
	"fmt"

	feather_commons_errors "github.com/guidomantilla/go-feather-commons/pkg/errors"
)

func ErrContextFailed(errs ...error) error {
	return fmt.Errorf("db context failed: %s", feather_commons_errors.ErrJoin(errs...).Error())
}
