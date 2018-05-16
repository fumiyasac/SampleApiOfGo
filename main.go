package main

//パッケージの宣言
import (
	"github.com/fumiyasac/SampleApi/controllers"
	_ "github.com/fumiyasac/SampleApi/docs"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Sample API
// @version 1.0
// @description This is a sample API server. (This is a practice of fumiyasac.)

// @contact.name fumiyasac
// @contact.url https://qiita.com/fumiyasac@github
// @contact.email just1factory@gmail.com

// @host localhost:8080
// @BasePath /api/v1

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

			//TOP用のメッセージ表示用のエンドポイント
			v1.GET("/top", TopController.GetMessage)

			//ユーザー情報のに関する処理のエンドポイント
			v1.GET("/users/:id", UserController.GetUser)
			v1.GET("/users", UserController.GetUsers)
			v1.POST("/users", UserController.CreateUser)
			v1.PUT("/users/:id", UserController.UpdateUser)
			v1.DELETE("/users/:id", UserController.DeleteUser)
		}
	}

	//Swagger用のエンドポイント
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//実行
	Router.Run(":8080")
}
