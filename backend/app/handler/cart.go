package handler

import (
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
	return nil
}
