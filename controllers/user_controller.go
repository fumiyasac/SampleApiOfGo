package controllers

//パッケージの宣言
import (
	"strconv"

	"github.com/fumiyasac/SampleApi/models"
	"github.com/gin-gonic/gin"
)

//UserController ... UserControllerの構造体宣言
type UserController struct{}

//GetUser ...
func (ctrl UserController) GetUser(c *gin.Context) {

	var id int
	id, _ = strconv.Atoi(c.Param("id"))
	repository := models.NewUserRepository()
	user := repository.GetByID(id)

	JSONContent := gin.H{
		"user": user,
	}
	c.JSON(200, JSONContent)
}
