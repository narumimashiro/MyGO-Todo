package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"go_modules"
)

func main() {
	// echoインスタンス生成
	e := echo.New()

	// テンプレートの設定
	e.Renderer = go_modules.NewTemplate()

	// ルーティング
	e.GET("/", handleHome)

	// サーバーの起動
	e.Start(":3000")

	// add task endpoint
	http.HandleFunc("/add", HandleAdd)
	http.ListenAndServe(":3000", nil)
}

func handleHome(c echo.Context) error {
	data := map[string]string{
		"PageTitle": "MyGo Todo",
	}
	return c.Render(http.StatusOK, "index.html", data)
}