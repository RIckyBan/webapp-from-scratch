package handler

import (
	"errors"
	"net/http"
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

func (h *cartHandler) HandleAddItemsToCart(c echo.Context) error {
	userID, err := common.ParseULID(c.QueryParam("userID"))
	if err != nil {
		return errors.New("invalid user id")
	}

	itemID, err := common.ParseULID(c.QueryParam("itemID"))
	if err != nil {
		return errors.New("invalid item id")
	}

	quantity, err := strconv.Atoi(c.QueryParam("quantity"))
	if err != nil {
		return err
	}

	err = h.usecase.AddItemToCart(c.Request().Context(), userID, itemID, quantity)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &Response{
		Message: "success",
		Data:    nil,
	})
}
