package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tomoEng11/go-backend-practice/api"
)

// Server implements api.ServerInterface
type Server struct {
	// 将来的にDBやサービスの依存を追加できる
}

// NewServer creates a new Server instance
func NewServer() *Server {
	return &Server{}
}

// GetHealth implements api.ServerInterface
func (s *Server) GetHealth(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, api.HealthResponse{
		Status: "OK",
	})
}

// GetChannels implements api.ServerInterface
func (s *Server) GetChannels(ctx echo.Context) error {
	channels := []api.Channel{
		{Id: 1, Name: "Go Backend Channel"},
		{Id: 2, Name: "Echo Practice Channel"},
	}
	return ctx.JSON(http.StatusOK, channels)
}
