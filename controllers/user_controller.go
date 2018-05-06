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
	id, err := strconv.Atoi(c.Param("id"))

	var JSONContent gin.H
	if err != nil {
		JSONContent = gin.H{
			"message": "Parameter is invalid.",
		}
		c.JSON(http.StatusBadRequest, JSONContent)
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
			"message": "User not found.",
		}
	}
	c.JSON(http.StatusOK, JSONContent)
}
