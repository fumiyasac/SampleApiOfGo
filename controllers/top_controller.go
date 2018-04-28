package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//TopController ... TopControllerの構造体宣言
type TopController struct{}

//Index ... Index(トップページ)のテンプレートへ引き渡す値の定義
func (ctrl *TopController) Index(c *gin.Context) {
	content := gin.H{
		"title":      "サイトのタイトルが入ります",
		"sampleText": "こちらはGolangのテストサイトになります。",
	}
	c.HTML(http.StatusOK, "index", content)
}
