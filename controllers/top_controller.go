package controllers

import (
	"net/http"

	"github.com/fumiyasac/SampleApi/repositories"
	"github.com/gin-gonic/gin"
)

// TopController ... 構造体宣言
type TopController struct{}

// GetMessage ... TOPメッセージを表示する
func (ctrl TopController) GetMessage(c *gin.Context) {

	var JSONContent gin.H

	repository := repositories.NewTopMessageRepository()
	topMessage := repository.GetTopMessage()

	JSONContent = gin.H{
		"top_message": topMessage,
	}
	c.JSON(http.StatusOK, JSONContent)
}
