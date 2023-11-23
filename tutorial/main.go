package main

import (
	"html/template"
	"io"
	"net/http"
	"github.com/labstack/echo/v4"
)

func main() {
	// インスタンス生成
	e := echo.New()

	// テンプレートエンジンの設定
	// RendererプロパティにHTMLテンプレートを設定している
  e.Renderer = &Template{
		// template.Must はエラーを処理し、エラーがあればプログラムを停止させます
		// template.ParseGlob("templates/*.html") は、指定されたパターンに一致するファイルを読み込んでテンプレートをパースします
    templates: template.Must(template.ParseGlob("templates/*.html")),
  }

	// eインスタンスに対して、GETメソッドでの"/"パスへのハンドラを登録
	e.GET("/", func(c echo.Context) error {
				// mapを定義
				// interfaceはちょっと違うけど、TsのAny型みたいな感じ
        // data := map[string]interface{}{
				data :=map[string]string{
          "PageTitle": "Golang WebApp",
          "SubTitle": "Why Golang",
          "Context": "TextTextTextTextTextTextTextText",
        }
        return c.Render(http.StatusOK, "index.html", data)
    })

    e.Start(":3000")
}

// Template 構造体の定義
type Template struct {
    templates *template.Template
}

// Render 関数の実装
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}