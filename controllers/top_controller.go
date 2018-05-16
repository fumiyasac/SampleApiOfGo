package controllers

import (
	"net/http"

	"github.com/fumiyasac/SampleApi/factories"
	"github.com/fumiyasac/SampleApi/repositories"
	"github.com/gin-gonic/gin"
)

// TopController ... 構造体宣言
type TopController struct{}

// TopMessageRepository ... interfaceの宣言
type TopMessageRepository interface {
	GetTopMessage() factories.TopMessageFactory
}

// GetMessage godoc
// @Summary このAPIに関するメッセージを表示する
// @Description TOP用の挨拶メッセージとこのAPIに関する概要や現行バージョン・更新履歴を表示する
// @Accept  json
// @Produce  json
// @Success 200 {object} factories.TopMessageFactory
// @Router /top [get]
func (ctrl TopController) GetMessage(c *gin.Context) {

	repository := repositories.NewTopMessageRepository()
	topMessage := repository.GetTopMessage()

	c.JSON(http.StatusOK, gin.H{
		"top_message": topMessage,
	})
}
