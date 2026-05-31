package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tomoEng11/go-backend-practice/gen/protected"
	"github.com/tomoEng11/go-backend-practice/gen/public"
	"github.com/tomoEng11/go-backend-practice/internal/handler"
	mymw "github.com/tomoEng11/go-backend-practice/internal/middleware"
)

func main() {
	// Echoサーバーを起動する
	e := echo.New()

	// 共通ミドルウェアを登録する
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	// 公開エンドポイント（認証不要）
	publicHandler := handler.NewPublicHandler()
	public.RegisterHandlers(e.Group("/public"), publicHandler)

	// 認証必須エンドポイント
	protectedHandler := handler.NewProtectedHandler()
	protected.RegisterHandlers(e.Group("/api", mymw.AuthMiddleware), protectedHandler)

	// サーバーを起動する
	e.Logger.Fatal(e.Start(":8080"))
}
