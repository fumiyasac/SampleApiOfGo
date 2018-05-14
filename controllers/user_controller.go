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
	UpdateByID(id int, username string, password string, mailaddress string) bool
	DeleteByID(id int) bool
}

// UserValidator ... interfaceの宣言
type UserValidator interface {
	IsValidForUserCreate(username string, password string, mailaddress string) (bool, string)
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
	validatorResult, validatorErrorMessage := userRequestValidator.IsValidForUserCreate(username, password, mailaddress)

	if validatorResult != true {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": validatorErrorMessage,
		})
		return
	}

	userRepository := repositories.NewUserRepository()
	createResult := userRepository.Create(username, password, mailaddress)

	if createResult != true {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": constants.ErrorMessageOfFailureForUserCreate,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": constants.SuccessForUserCreate,
	})
}

// UpdateUser ... idに該当するユーザーを更新する
func (ctrl UserController) UpdateUser(c *gin.Context) {

	var id int
	var username string
	var password string
	var mailaddress string

	id, paramError := strconv.Atoi(c.Param("id"))
	if paramError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": constants.ErrorMessageOfInvalidID,
		})
		return
	}

	username = c.PostForm("username")
	password = c.PostForm("password")
	mailaddress = c.PostForm("mailaddress")

	userRequestValidator := validators.NewUserRequestValidator()
	validatorResult, validatorErrorMessage := userRequestValidator.IsValidForUserCreate(username, password, mailaddress)

	if validatorResult != true {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": validatorErrorMessage,
		})
		return
	}

	userRepository := repositories.NewUserRepository()
	updateResult := userRepository.UpdateByID(id, username, password, mailaddress)

	if updateResult != true {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": constants.ErrorMessageOfFailureForUserUpdate,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": constants.SuccessForUserUpdate,
	})
}

// DeleteUser ... idに該当するユーザーを削除する
func (ctrl UserController) DeleteUser(c *gin.Context) {

	var id int

	id, paramError := strconv.Atoi(c.Param("id"))
	if paramError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": constants.ErrorMessageOfInvalidID,
		})
		return
	}

	userRepository := repositories.NewUserRepository()
	deleteResult := userRepository.DeleteByID(id)

	if deleteResult != true {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": constants.ErrorMessageOfFailureForUserDelete,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": constants.SuccessForUserDelete,
	})
}
