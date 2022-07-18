package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
)

//以下の書き方は構文のようなもの
func main() {

	// echoパッケージにより簡単にサーバーのインスタンスを作成
	e := echo.New()

	// ルーティングの設定
	e.GET("/", helloworld)

	//サーバー起動
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))

}

//"Hello, World"と吐き出すサーバーの実装
func helloworld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World")
}
