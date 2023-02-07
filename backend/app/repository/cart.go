package repository

import (
	"context"

	"github.com/RIckyBan/webapp-from-scratch/backend/app/entity"
	"github.com/oklog/ulid"
)

type CartRepository interface {
	GetAllItems(ctx context.Context, userID ulid.ULID) ([]*entity.Item, error)
	AddItem(ctx context.Context, userID ulid.ULID, itemID ulid.ULID, quantity int) error
}
