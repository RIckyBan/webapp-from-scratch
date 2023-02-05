package handler

import (
	"strconv"

	"github.com/RIckyBan/webapp-from-scratch/backend/app/common"
	"github.com/RIckyBan/webapp-from-scratch/backend/app/usecase"
	"github.com/labstack/echo/v4"
)

type cartHandler struct {
	usecase *usecase.CartService
}

func NewCartHandler(usecase *usecase.CartService) *cartHandler {
	return &cartHandler{usecase: usecase}
}

func (h *cartHandler) HandleAddItemsToCart(c echo.Context) error {
	itemID := common.ParseULID(c.Param("itemID"))
	quantity, err := strconv.ParseInt(c.Param("quantity"), 10, 64)
	if err != nil {
		return err
	}

	err = h.usecase.AddItemToCart(c.Request().Context(), itemID, quantity)
	if err != nil {
		return err
	}
	return nil
}
