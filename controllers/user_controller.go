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
	Create(username string, password string) bool
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

// CreateUser ... 新規ユーザーを登録する
func (ctrl UserController) CreateUser(c *gin.Context) {

	var username string
	var password string
	var JSONContent gin.H

	username = c.PostForm("username")
	password = c.PostForm("password")
	if len(username) == 0 || len(password) == 0 {
		JSONContent = gin.H{
			"message": "Unprocessable Entity.",
		}
		c.JSON(http.StatusUnprocessableEntity, JSONContent)
		return
	}

	repository := repositories.NewUserRepository()
	result := repository.Create(username, password)
	if result {
		JSONContent = gin.H{
			"success": "New User Cleated.",
		}
		c.JSON(http.StatusCreated, JSONContent)
	} else {
		JSONContent = gin.H{
			"error": "Internal Sever Error.",
		}
		c.JSON(http.StatusInternalServerError, JSONContent)
	}
}
