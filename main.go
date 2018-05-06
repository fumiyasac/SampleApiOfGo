package main

//パッケージの宣言
import (
	"github.com/fumiyasac/SampleApi/controllers"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

//メイン処理部分
func main() {

	//Ginの読み込み
	Router := gin.Default()

	//静的htmlファイル表示の設定
	Router.Static("guide", "./static")

	//APIのエンドポイント設定
	APIRouter := Router.Group("api")
	{
		//バージョン情報
		v1 := APIRouter.Group("v1")
		{
			//API用のコントローラー呼び出し
			TopController := new(controllers.TopController)
			UserController := new(controllers.UserController)

			//エンドポイントへのリクエスト
			v1.GET("/top", TopController.GetMessage)

			v1.GET("/user/:id", UserController.GetUser)
		}
	}

	//実行
	Router.Run(":8080")
}
