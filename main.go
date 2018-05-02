package main

//パッケージの宣言
import (
	"github.com/fumiyasac/SampleApi/controllers"
	"github.com/gin-gonic/gin"
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
		//API用のコントローラー呼び出し
		API := new(controllers.APIController)

		//バージョン情報
		V1 := APIRouter.Group("v1")
		{
			//エンドポイントへのリクエスト
			V1.GET("/top", API.Top)
		}
	}

	//実行
	Router.Run(":8080")
}
