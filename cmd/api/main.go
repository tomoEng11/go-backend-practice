package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tomoEng11/go-backend-practice/api"
	"github.com/tomoEng11/go-backend-practice/internal/server"
)

func main() {
	// Echoサーバーを起動する
	e := echo.New()

	// Middlewareを登録する
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	// OpenAPI ServerInterfaceを実装したサーバーを作成
	s := server.NewServer()

	// OpenAPIで定義したルートを自動登録
	api.RegisterHandlers(e, s)

	// サーバーを起動する
	e.Logger.Fatal(e.Start(":8080"))

}
