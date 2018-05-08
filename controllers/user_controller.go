package controllers

import (
	"net/http"
	"strconv"

	"github.com/fumiyasac/SampleApi/factories"
	"github.com/fumiyasac/SampleApi/repositories"
	"github.com/gin-gonic/gin"
)

// UserController ... 構造体宣言
type UserController struct{}

// UserRepository ... interfaceの宣言
type UserRepository interface {
	GetByID(id int) (factories.SingleUserFactory, bool)
}

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
	user, result := repository.GetByID(id)
	if result {
		JSONContent = gin.H{
			"user": user,
		}
		c.JSON(http.StatusOK, JSONContent)
	} else {
		JSONContent = gin.H{
			"error": "User Not Found.",
		}
		c.JSON(http.StatusNotFound, JSONContent)
	}
}
