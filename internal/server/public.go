package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tomoEng11/go-backend-practice/api/public"
)

// PublicServer implements public.ServerInterface
type PublicServer struct{}

func NewPublicServer() *PublicServer {
	return &PublicServer{}
}

// GetHealth implements public.ServerInterface
func (s *PublicServer) GetHealth(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, public.HealthResponse{
		Status: "ok",
	})
}
