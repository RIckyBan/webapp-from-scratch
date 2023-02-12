package handler

import (
	"errors"
	"net/http"

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

func (h *cartHandler) HandleGetAllItemsInCart(c echo.Context) error {
	userID, err := common.ParseULID(c.QueryParam("userID"))
	if err != nil {
		return errors.New("invalid user id")
	}

	items, err := h.usecase.GetAllItemsInCart(c.Request().Context(), userID)
	if err != nil {
		return c.JSON(http.StatusOK, &Response{
			Message: "error",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, &Response{
		Message: "success",
		Data:    items,
	})
}

func (h *cartHandler) HandleAddItemToCart(c echo.Context) error {
	req := new(AddItemToCartRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, &Response{
			Message: err.Error(),
			Data:    nil,
		})
	}

	userID, err := common.ParseULID(req.UserID)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest, &Response{
				Message: "invalid user id",
				Data:    nil,
			},
		)
	}

	itemID, err := common.ParseULID(req.ItemID)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest, &Response{
				Message: "invalid item id",
				Data:    nil,
			},
		)
	}

	err = h.usecase.AddItemToCart(c.Request().Context(), userID, itemID, req.Quantity)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &Response{
		Message: "success",
		Data:    nil,
	})
}
