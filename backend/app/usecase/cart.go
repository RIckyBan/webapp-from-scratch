package usecase

import (
	"context"

	"github.com/RIckyBan/webapp-from-scratch/backend/app/entity"
	"github.com/RIckyBan/webapp-from-scratch/backend/app/repository"
	"github.com/oklog/ulid"
)

type CartService struct {
	cartRepo repository.CartRepository
}

func NewCartService() *CartService {
	return &CartService{}
}

func (s *CartService) GetAllItemsInCart(ctx context.Context, userID ulid.ULID) ([]*entity.Item, error) {
	items, err := s.cartRepo.GetAllItems(ctx, userID)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (s *CartService) AddItemToCart(ctx context.Context, itemID ulid.ULID, quantity int64) error {
	err := s.cartRepo.AddItem(ctx, itemID, quantity)
	if err != nil {
		return err
	}

	return nil
}
