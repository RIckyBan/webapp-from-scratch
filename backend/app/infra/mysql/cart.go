package mysql

import (
	"context"
	"database/sql"

	"github.com/RIckyBan/webapp-from-scratch/backend/app/common"
	"github.com/RIckyBan/webapp-from-scratch/backend/app/entity"
	"github.com/RIckyBan/webapp-from-scratch/backend/app/repository"
	"github.com/RIckyBan/webapp-from-scratch/backend/db/models"
	"github.com/oklog/ulid"
	"github.com/volatiletech/sqlboiler/boil"
)

type cartRepository struct {
	db *sql.DB
}

func NewCartRepository(db *sql.DB) repository.CartRepository {
	return &cartRepository{db: db}
}

func (r *cartRepository) GetItems(ctx context.Context, userID ulid.ULID) ([]*entity.Item, error) {
	var itemEntities []*entity.Item

	items, err := models.Phones().All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		screenSize, ok := item.ScreenSizeInch.Float64()
		if !ok {
			return nil, err
		}
		weightG, ok := item.WeightG.Float64()
		if !ok {
			return nil, err
		}

		itemEntities = append(itemEntities, &entity.Item{

			ID:              common.ParseULID(item.ID),
			Brand:           item.Brand,
			Model:           item.Model,
			OperatingSystem: item.OperatingSystem,
			StorageGB:       int64(item.StorageGB),
			RAMGB:           int64(item.RAMGB),
			Color:           item.Color,
			ScreenSizeInch:  screenSize,
			WeightG:         weightG,
			Price:           int64(item.Price),
			ReleaseDate:     item.ReleaseDate,
			Description:     item.Description,
		})
	}

	return itemEntities, nil
}

func (r *cartRepository) AddItem(ctx context.Context, item *entity.Item, quantity int64) error {
	cartModel := models.Cart{
		ID:       item.ID.String(),
		PhoneID:  item.ID.String(),
		Quantity: int(quantity),
	}

	err := cartModel.Upsert(ctx, r.db, boil.Infer())
	if err != nil {
		return err
	}

	return nil
}
