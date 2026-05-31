package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tomoEng11/go-backend-practice/api/protected"
)

// ProtectedServer implements protected.ServerInterface
type ProtectedServer struct{}

func NewProtectedServer() *ProtectedServer {
	return &ProtectedServer{}
}

// GetChannels implements protected.ServerInterface
func (s *ProtectedServer) GetChannels(ctx echo.Context) error {
	channels := []protected.Channel{
		{Id: 1, Name: "Go Backend Channel"},
		{Id: 2, Name: "Echo Practice Channel"},
	}
	return ctx.JSON(http.StatusOK, channels)
}

// GetUsers implements protected.ServerInterface
func (s *ProtectedServer) GetUsers(ctx echo.Context) error {
	users := []protected.User{
		{Id: 1, Name: "Alice"},
		{Id: 2, Name: "Bob"},
	}
	return ctx.JSON(http.StatusOK, users)
}
