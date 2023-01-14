package repository

import (
	"context"

	"github.com/RIckyBan/webapp-from-scratch/backend/app/entity"
	"github.com/oklog/ulid"
)

type CartRepository interface {
	GetItems(ctx context.Context, userID ulid.ULID) ([]*entity.Item, error)
	AddItem(ctx context.Context, item *entity.Item, quantity int64) error
}
