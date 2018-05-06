package controllers

//パッケージの宣言
import (
	"github.com/fumiyasac/SampleApi/models"
	"github.com/gin-gonic/gin"
)

//TopController ... TopControllerの構造体宣言
type TopController struct{}

//GetMessage ...
func (ctrl TopController) GetMessage(c *gin.Context) {

	repository := models.NewTopMessageRepository()
	topMessage := repository.GetTopMessage()

	JSONContent := gin.H{
		"top_message": topMessage,
	}
	c.JSON(200, JSONContent)
}
