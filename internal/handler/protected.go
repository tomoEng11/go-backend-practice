package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tomoEng11/go-backend-practice/gen/protected"
)

// ProtectedHandler implements protected.ServerInterface
type ProtectedHandler struct{}

func NewProtectedHandler() *ProtectedHandler {
	return &ProtectedHandler{}
}

// GetChannels implements protected.ServerInterface
func (h *ProtectedHandler) GetChannels(ctx echo.Context) error {
	channels := []protected.Channel{
		{Id: 1, Name: "Go Backend Channel"},
		{Id: 2, Name: "Echo Practice Channel"},
	}
	return ctx.JSON(http.StatusOK, channels)
}

// GetUsers implements protected.ServerInterface
func (h *ProtectedHandler) GetUsers(ctx echo.Context) error {
	users := []protected.User{
		{Id: 1, Name: "Alice"},
		{Id: 2, Name: "Bob"},
	}
	return ctx.JSON(http.StatusOK, users)
}
