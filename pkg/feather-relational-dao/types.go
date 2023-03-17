package feather_relational_dao

import (
	"context"
)

type CrudDao interface {
	Save(ctx context.Context, args ...any) (*int64, error)
	Update(ctx context.Context, args ...any) error
	Delete(ctx context.Context, id any) error
	FindById(ctx context.Context, id any, args ...any) error
	FindAll(ctx context.Context, fn ReadFunction) error
	ExistsById(ctx context.Context, id any, args ...any) bool
}
