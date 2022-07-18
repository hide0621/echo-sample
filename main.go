package main

import (
	"net/http"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
)

type Content struct {
	gorm.Model
	Text string `json:"text"`
}

// DBのインスタンスをグローバル変数に格納
var (
	db *gorm.DB
)

func main() {
	// DB接続処理
	var err error
	// 第一引数はドライバー、第二引数は使用するテーブル
	db, err = gorm.Open("sqlite3", "gorm.db")
	if err != nil {
		panic("failed to connect database")
	}
	// サーバーが終了したらDB接続も終了する
	defer db.Close()

	// echoパッケージによりサーバーのインスタンスを作成
	e := echo.New()

	// ルーティングの設定
	e.GET("/contents", getALlContent)
	e.POST("/create", createContent)

	// サーバー起動
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

//echoパッケージ版のハンドラ関数の実装
func createContent(c echo.Context) error {
	// レコード登録
	db.Create(&Content{Text: "newText"})
	return c.String(http.StatusOK, "record created")
}

//echoパッケージ版のハンドラ関数の実装
func getALlContent(c echo.Context) error {
	var content Content
	// contentテーブルのレコードを全件取得
	db.Find(&content)
	// 取得したデータをJSONにして返却
	return c.JSON(http.StatusOK, content)
}
