package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tomoEng11/go-backend-practice/gen/public"
)

// PublicHandler implements public.ServerInterface
type PublicHandler struct{}

func NewPublicHandler() *PublicHandler {
	return &PublicHandler{}
}

// GetHealth implements public.ServerInterface
func (h *PublicHandler) GetHealth(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, public.HealthResponse{
		Status: "ok",
	})
}
