package main

//パッケージの宣言
import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

//メイン処理部分
func main() {
	//GInの読み込み
	r := gin.Default()

	//接続テスト
	getDB()

	//APIのエンドポイントの設定
	v1 := r.Group("api/v1")
	{
		//TOP画面用の文言を出す
		v1.GET("/top", getHello)
	}

	//実行
	r.Run(":8080")
}

//ハローメッセージの組み立て
func getHello(c *gin.Context) {
	content := gin.H{
		"message": "Hello, サンプルアプリトップページへ！",
	}
	c.JSON(200, content)
}

//Databaseとの接続を行う
func getDB() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:root/reizoukoseiri")
	if err != nil {
		log.Fatal("Open Database Error: %v", err)
	} else {
		log.Println("Success!")
	}
	return db
}
