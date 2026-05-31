package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tomoEng11/go-backend-practice/api/protected"
	"github.com/tomoEng11/go-backend-practice/api/public"
	mymw "github.com/tomoEng11/go-backend-practice/internal/middleware"
	"github.com/tomoEng11/go-backend-practice/internal/server"
)

func main() {
	// Echoサーバーを起動する
	e := echo.New()

	// 共通ミドルウェアを登録する
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	// 公開エンドポイント（認証不要）
	publicServer := server.NewPublicServer()
	public.RegisterHandlers(e.Group("/public"), publicServer)

	// 認証必須エンドポイント
	protectedServer := server.NewProtectedServer()
	protected.RegisterHandlers(e.Group("/api", mymw.AuthMiddleware), protectedServer)

	// サーバーを起動する
	e.Logger.Fatal(e.Start(":8080"))
}
