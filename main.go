package main

//パッケージの宣言
import (
	"github.com/fumiyasac/SampleApi/controllers"
	"github.com/gin-gonic/gin"
)

//メイン処理部分
func main() {

	//Ginの読み込み
	router := gin.Default()

	//APIのエンドポイント設定
	webapi := router.Group("api")
	{
		//API用のコントローラー呼び出し
		apiController := new(controllers.APIController)

		//バージョン情報
		v1 := webapi.Group("v1")
		{
			//エンドポイントへのリクエスト
			v1.GET("/top", apiController.Top)
		}
	}

	//実行
	router.Run(":8080")
}
