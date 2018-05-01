package controllers

import (
	"github.com/gin-gonic/gin"
)

//TopController ... TopControllerの構造体宣言
type TopController struct{}

//Index ... Index(トップページ)のテンプレートへ引き渡す値の定義
func (ctrl *TopController) Index(c *gin.Context) {
	content := gin.H{
		"sampleText": "こちらはGolangのテストサイトになります。",
	}
	c.HTML(200, "index.html", content)
}
