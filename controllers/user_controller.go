package controllers

import (
	"net/http"
	"strconv"

	"github.com/fumiyasac/SampleApi/constants"
	"github.com/fumiyasac/SampleApi/factories"
	"github.com/fumiyasac/SampleApi/repositories"
	"github.com/fumiyasac/SampleApi/validators"
	"github.com/gin-gonic/gin"
)

// UserController ... 構造体宣言
type UserController struct{}

// UserRepository ... interfaceの宣言
type UserRepository interface {
	GetByID(id int) (factories.SingleUserFactory, bool)
	Create(username string, password string, mailaddress string) bool
}

// UserValidator ... interfaceの宣言
type UserValidator interface {
	IsValidForUserCreate(username string, password string, mailaddress string) bool
}

// GetUser ... idに該当するユーザーを表示する
func (ctrl UserController) GetUser(c *gin.Context) {

	var id int

	id, paramError := strconv.Atoi(c.Param("id"))

	if paramError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": constants.ErrorMessageOfInvalidID,
		})
		return
	}

	userRepository := repositories.NewUserRepository()
	userData, fetchResult := userRepository.GetByID(id)

	if fetchResult != true {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": constants.ErrorMessageOfUserNotFound,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": userData,
	})
}

// CreateUser ... 新規ユーザーを登録する
func (ctrl UserController) CreateUser(c *gin.Context) {

	var username string
	var password string
	var mailaddress string

	username = c.PostForm("username")
	password = c.PostForm("password")
	mailaddress = c.PostForm("mailaddress")

	userRequestValidator := validators.NewUserRequestValidator()
	validatorResult := userRequestValidator.IsValidForUserCreate(username, password, mailaddress)

	if validatorResult != true {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": constants.ErrorMessageOfInvalidValueForUserCreate,
		})
		return
	}

	userRepository := repositories.NewUserRepository()
	fetchResult := userRepository.Create(username, password, mailaddress)

	if fetchResult != true {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": constants.ErrorMessageOfInvalidValueForUserCreate,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": constants.SuccessForUserCreate,
	})
}
