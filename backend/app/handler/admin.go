package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo"

	"github.com/RIckyBan/webapp-from-scratch/backend/app/usecase"
)

type adminHandler struct {
	usecase *usecase.AdminService
}

func NewAdminHandler(usecase *usecase.AdminService) *adminHandler {
	return &adminHandler{
		usecase: usecase,
	}
}

func (h *adminHandler) HandleGetAllUsers(c echo.Context) error {
	users, err := h.usecase.GetAllUsers(context.Background())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &Response{
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, &Response{
		Message: "success",
		Data:    users,
	})
}
