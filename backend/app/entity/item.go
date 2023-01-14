package entity

import (
	"time"

	"github.com/oklog/ulid"
	"github.com/volatiletech/null/v8"
)

type Item struct {
	ID              ulid.ULID
	Brand           string
	Model           string
	OperatingSystem string
	StorageGB       int64
	RAMGB           int64
	Color           string
	ScreenSizeInch  float64
	WeightG         float64
	Price           int64
	ReleaseDate     time.Time
	Description     null.String
}

func NewItem(
	id ulid.ULID,
	brand string,
	model string,
	operatingSystem string,
	storageGB int64,
	ramgb int64,
	color string,
	screenSizeInch float64,
	weightG float64,
	price int64,
	releaseDate time.Time,
	description null.String,
) *Item {
	return &Item{
		ID:              id,
		Brand:           brand,
		Model:           model,
		OperatingSystem: operatingSystem,
		StorageGB:       storageGB,
		RAMGB:           ramgb,
		Color:           color,
		ScreenSizeInch:  screenSizeInch,
		WeightG:         weightG,
		Price:           price,
		ReleaseDate:     releaseDate,
		Description:     description,
	}
}
