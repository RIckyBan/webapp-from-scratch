package mysql

import (
	"context"
	"database/sql"

	"github.com/RIckyBan/webapp-from-scratch/backend/app/common"
	"github.com/RIckyBan/webapp-from-scratch/backend/app/entity"
	"github.com/RIckyBan/webapp-from-scratch/backend/app/repository"
	"github.com/RIckyBan/webapp-from-scratch/backend/db/models"
	"github.com/oklog/ulid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type cartRepository struct {
	db *sql.DB
}

func NewCartRepository(db *sql.DB) repository.CartRepository {
	return &cartRepository{db: db}
}

func (r *cartRepository) GetAllItems(ctx context.Context, userID ulid.ULID) ([]*entity.Item, error) {
	var itemEntities []*entity.Item

	items, err := models.Phones(
		qm.InnerJoin("carts on phones.id = carts.phone_id"),
		qm.Where("carts.user_id = ?", userID.String()),
	).All(ctx, r.db)
	if err != sql.ErrNoRows && err != nil {
		return nil, err
	}

	for _, item := range items {
		itemID, err := common.ParseULID(item.ID)
		if err != nil {
			return nil, err
		}

		screenSize, ok := item.ScreenSizeInch.Float64()
		if !ok {
			return nil, err
		}
		weightG, ok := item.WeightG.Float64()
		if !ok {
			return nil, err
		}

		itemEntities = append(itemEntities, &entity.Item{
			ID:              itemID,
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

func (r *cartRepository) AddItem(ctx context.Context, userID ulid.ULID, itemID ulid.ULID, quantity int) error {
	cartModel := models.Cart{
		ID:       common.GenerateULID().String(),
		UserID:   userID.String(),
		PhoneID:  itemID.String(),
		Quantity: quantity,
	}

	err := cartModel.Upsert(ctx, r.db, boil.Infer(), boil.Infer())
	if err != nil {
		return err
	}

	return nil
}
