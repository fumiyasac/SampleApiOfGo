package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin/binding"

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
	GetLists(id int) (factories.SingleUserFactory, bool)
	GetByID(id int) (factories.SingleUserFactory, bool)
	Create(username string, password string, mailaddress string) bool
	UpdateByID(id int, username string, password string, mailaddress string) bool
	DeleteByID(id int) bool
}

// UserValidator ... interfaceの宣言
type UserValidator interface {
	IsValidForUserCreate(username string, password string, mailaddress string) (bool, string)
}

// GetUsers godoc
// @Summary ユーザー一覧を表示する
// @Description userテーブルに登録されているデータを全件取得する
// @Accept json
// @Produce json
// @Success 200 {object} factories.UserItemsFactory
// @Failure 400 {object} factories.APIErrorMessageFactory
// @Router /users [get]
func (ctrl UserController) GetUsers(c *gin.Context) {

	userRepository := repositories.NewUserRepository()
	userData, fetchResult := userRepository.GetLists()

	if fetchResult != true {
		errorMessage := factories.APIErrorMessageFactory{
			Message: constants.ErrorMessageOfUserNotFound,
		}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	c.JSON(http.StatusOK, userData)
}

// GetUser godoc
// @Summary ユーザーデータを1件表示する
// @Description userテーブルに登録されているデータのうちパラメーター(id)に該当するものを1件取得する
// @Accept json
// @Produce json
// @Param id path int true "ユーザーID"
// @Success 200 {object} factories.SingleUserFactory
// @Failure 400 {object} factories.APIErrorMessageFactory
// @Router /users/{id} [get]
func (ctrl UserController) GetUser(c *gin.Context) {

	var id int

	id, paramError := strconv.Atoi(c.Param("id"))

	if paramError != nil {
		errorMessage := factories.APIErrorMessageFactory{
			Message: constants.ErrorMessageOfInvalidID,
		}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	userRepository := repositories.NewUserRepository()
	userData, fetchResult := userRepository.GetByID(id)

	if fetchResult != true {
		errorMessage := factories.APIErrorMessageFactory{
			Message: constants.ErrorMessageOfUserNotFound,
		}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	c.JSON(http.StatusOK, userData)
}

// CreateUser godoc
// @Summary ユーザーデータを1件新規登録する
// @Description userテーブルへ新規データを1件登録する
// @Accept json
// @Produce json
// @Param request body factories.RequestUserJSONFactory true "username: ユーザー名 ・ password: パスワード ・ mailaddress: メールアドレス"
// @Success 201 {object} factories.APISuccessMessageFactory
// @Failure 400 {object} factories.APIErrorMessageFactory
// @Router /users [post]
func (ctrl UserController) CreateUser(c *gin.Context) {

	var requestUserJSONFactory factories.RequestUserJSONFactory
	var username string
	var password string
	var mailaddress string

	JSONError := c.ShouldBindWith(&requestUserJSONFactory, binding.JSON)
	if JSONError != nil {
		errorMessage := factories.APIErrorMessageFactory{
			Message: constants.ErrorMessageOfInvalidRequest,
		}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	username = requestUserJSONFactory.Username
	password = requestUserJSONFactory.Password
	mailaddress = requestUserJSONFactory.Mailaddress

	userRequestValidator := validators.NewUserRequestValidator()
	validatorResult, validatorErrorMessage := userRequestValidator.IsValidForUserCreate(username, password, mailaddress)

	if validatorResult != true {
		errorMessage := factories.APIErrorMessageFactory{
			Message: validatorErrorMessage,
		}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	userRepository := repositories.NewUserRepository()
	createResult := userRepository.Create(username, password, mailaddress)

	if createResult != true {
		errorMessage := factories.APIErrorMessageFactory{
			Message: constants.ErrorMessageOfFailureForUserCreate,
		}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	successMessage := factories.APISuccessMessageFactory{
		Message: constants.SuccessForUserCreate,
	}
	c.JSON(http.StatusCreated, successMessage)
}

// UpdateUser godoc
// @Summary ユーザーデータを1件更新する
// @Description userテーブルに登録されているデータのうちパラメーター(id)に該当するものを1件更新する
// @Accept json
// @Produce json
// @Param id path int true "ユーザーID"
// @Param request body factories.RequestUserJSONFactory true "username: ユーザー名 ・ password: パスワード ・ mailaddress: メールアドレス"
// @Success 200 {object} factories.APISuccessMessageFactory
// @Failure 400 {object} factories.APIErrorMessageFactory
// @Router /users/{id} [put]
func (ctrl UserController) UpdateUser(c *gin.Context) {

	var id int
	var requestUserJSONFactory factories.RequestUserJSONFactory
	var username string
	var password string
	var mailaddress string

	id, paramError := strconv.Atoi(c.Param("id"))
	if paramError != nil {
		errorMessage := factories.APIErrorMessageFactory{
			Message: constants.ErrorMessageOfInvalidID,
		}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	JSONError := c.ShouldBindWith(&requestUserJSONFactory, binding.JSON)
	if JSONError != nil {
		errorMessage := factories.APIErrorMessageFactory{
			Message: constants.ErrorMessageOfInvalidRequest,
		}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	username = requestUserJSONFactory.Username
	password = requestUserJSONFactory.Password
	mailaddress = requestUserJSONFactory.Mailaddress

	userRequestValidator := validators.NewUserRequestValidator()
	validatorResult, validatorErrorMessage := userRequestValidator.IsValidForUserCreate(username, password, mailaddress)

	if validatorResult != true {
		errorMessage := factories.APIErrorMessageFactory{
			Message: validatorErrorMessage,
		}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	userRepository := repositories.NewUserRepository()
	updateResult := userRepository.UpdateByID(id, username, password, mailaddress)

	if updateResult != true {
		errorMessage := factories.APIErrorMessageFactory{
			Message: constants.ErrorMessageOfFailureForUserUpdate,
		}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	successMessage := factories.APISuccessMessageFactory{
		Message: constants.SuccessForUserUpdate,
	}
	c.JSON(http.StatusOK, successMessage)
}

// DeleteUser godoc
// @Summary ユーザーデータを1件削除する
// @Description userテーブルに登録されているデータのうちパラメーター(id)に該当するものを1件削除する
// @Accept json
// @Produce json
// @Param id path int true "ユーザーID"
// @Success 200 {object} factories.APISuccessMessageFactory
// @Failure 400 {object} factories.APIErrorMessageFactory
// @Router /users/{id} [delete]
func (ctrl UserController) DeleteUser(c *gin.Context) {

	var id int

	id, paramError := strconv.Atoi(c.Param("id"))
	if paramError != nil {
		errorMessage := factories.APIErrorMessageFactory{
			Message: constants.ErrorMessageOfInvalidID,
		}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	userRepository := repositories.NewUserRepository()
	deleteResult := userRepository.DeleteByID(id)

	if deleteResult != true {
		errorMessage := factories.APIErrorMessageFactory{
			Message: constants.ErrorMessageOfFailureForUserDelete,
		}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	successMessage := factories.APISuccessMessageFactory{
		Message: constants.SuccessForUserDelete,
	}
	c.JSON(http.StatusOK, successMessage)
}
