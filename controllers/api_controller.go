package controllers

import (
	"github.com/gin-gonic/gin"
)

//APIController ... APIControllerの構造体宣言
type APIController struct{}

//Top ... Top(トップ部分)のエンドポイント定義
func (ctrl *APIController) Top(c *gin.Context) {
	content := gin.H{
		"message": "Hello, サンプルアプリトップページへ！",
	}
	c.JSON(200, content)
}
