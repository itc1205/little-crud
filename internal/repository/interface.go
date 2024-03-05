package repository

import "context"

type GoodsRepository interface {
	Create(ctx context.Context, e *Goods) error
	Delete(ctx context.Context, id int32) error
	Get(ctx context.Context, id int32) (*Goods, error)
	List(ctx context.Context, offset, limit int32) ([]*Goods, error)
	Update(ctx context.Context, e *Goods) error
}
