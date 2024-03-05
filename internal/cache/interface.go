package cache

import (
	"context"

	"github.com/itc1205/little-crud/internal/repository"
)

type Cache interface {
	// Gets cache entry
	Get(ctx context.Context, id int32) (*repository.Goods, error)
	// Creates new cache entry
	Create(ctx context.Context, goods *repository.Goods) error
	// Invalidates whole cache
	Invalidate(ctx context.Context) error
}
