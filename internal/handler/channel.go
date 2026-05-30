package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Channel struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetChannels(c echo.Context) error {
	channels := []Channel{
		{ID: 1, Name: "Go Backend Channel"},
		{ID: 2, Name: "Echo Practice Channel"},
	}
	return c.JSON(http.StatusOK, channels)
}

func GetChannel(c echo.Context) error {

	channel := Channel{ID: 10, Name: "Hence"}
	return c.JSON(http.StatusOK, channel)
}
