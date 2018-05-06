package controllers

import (
	"net/http"
	"strconv"

	"github.com/fumiyasac/SampleApi/repositories"
	"github.com/gin-gonic/gin"
)

// UserController ... 構造体宣言
type UserController struct{}

// GetUser ... idに該当するユーザーを表示する
func (ctrl UserController) GetUser(c *gin.Context) {

	var id int
	var JSONContent gin.H

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		JSONContent = gin.H{
			"message": "Unprocessable Entity.",
		}
		c.JSON(http.StatusUnprocessableEntity, JSONContent)
		return
	}

	repository := repositories.NewUserRepository()
	user := repository.GetByID(id)
	if user != nil {
		JSONContent = gin.H{
			"user": user,
		}
	} else {
		JSONContent = gin.H{
			"message": "User Not Found.",
		}
	}
	c.JSON(http.StatusOK, JSONContent)
}
