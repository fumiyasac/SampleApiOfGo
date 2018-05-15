package repositories

import (
	"github.com/fumiyasac/SampleApi/constants"
	"github.com/fumiyasac/SampleApi/database"
	"github.com/fumiyasac/SampleApi/entities"
	"github.com/fumiyasac/SampleApi/factories"
	"github.com/go-xorm/xorm"
	"golang.org/x/crypto/bcrypt"
)

// UserRepository ... 構造体宣言
type UserRepository struct {
}

var engine *xorm.Engine

func init() {

	// xormを利用するための宣言
	engine = database.UseEngine()
}

// NewUserRepository ... 構造体の新規作成メソッド
func NewUserRepository() UserRepository {
	return UserRepository{}
}

// GetLists ... ユーザー一覧情報を取得する
// @GET("api/v1/users")
func (repo UserRepository) GetLists() (factories.UserItemsFactory, bool) {

	var users []entities.User
	var userFactories []factories.SingleUserFactory
	var userItemsFactory factories.UserItemsFactory

	err := engine.Find(&users)
	if err != nil {
		return userItemsFactory, false
	}

	for _, user := range users {
		userFactories = append(userFactories, factories.SingleUserFactory{
			ID:          user.ID,
			Username:    user.Username,
			MailAddress: user.MailAddress,
			Password:    user.Password,
			UserStatus:  constants.GetUserStatusNameFromStatusCode(user.StatusCode),
			CreatedAt:   user.CreatedAt,
			UpdatedAt:   user.UpdatedAt,
		})
	}

	userItemsFactory = factories.UserItemsFactory{
		Items: userFactories,
	}
	return userItemsFactory, true
}

// GetByID ... 引数のidに該当するユーザー情報を取得する
// @GET("api/v1/users/:id")
func (repo UserRepository) GetByID(id int) (factories.SingleUserFactory, bool) {

	var user = entities.User{}
	var userFactory factories.SingleUserFactory

	result, _ := engine.Where("id = ?", id).Get(&user)
	if result {
		userFactory = factories.SingleUserFactory{
			ID:          user.ID,
			Username:    user.Username,
			MailAddress: user.MailAddress,
			Password:    user.Password,
			UserStatus:  constants.GetUserStatusNameFromStatusCode(user.StatusCode),
			CreatedAt:   user.CreatedAt,
			UpdatedAt:   user.UpdatedAt,
		}
	}
	return userFactory, result
}

// Create ... 受け取った値をユーザー情報として新規登録する
// @POST("api/v1/users")
func (repo UserRepository) Create(username string, password string, mailaddress string) bool {

	var user = entities.User{}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return false
	}

	user = entities.User{
		Username:    username,
		MailAddress: mailaddress,
		Password:    string(hash),
		StatusCode:  constants.UserSubscribed.GetRawValue(),
	}

	affected, _ := engine.Insert(&user)
	return (affected > 0)
}

// UpdateByID ... 引数のidに該当するユーザー情報を更新する
// @PUT("api/v1/users/:id")
func (repo UserRepository) UpdateByID(id int, username string, password string, mailaddress string) bool {

	var user = entities.User{}

	result, _ := engine.Where("id = ?", id).Get(&user)
	if result != true {
		return false
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return false
	}

	user = entities.User{
		Username:    username,
		MailAddress: mailaddress,
		Password:    string(hash),
	}

	affected, _ := engine.Where("id = ?", id).Update(&user)
	return (affected > 0)
}

// DeleteByID ... 引数のidに該当するユーザー情報を物理削除する
// @DELETE("api/v1/users/:id")
func (repo UserRepository) DeleteByID(id int) bool {

	var user = entities.User{}

	result, _ := engine.Where("id = ?", id).Get(&user)
	if result != true {
		return false
	}

	affected, _ := engine.Where("id = ?", id).Delete(&user)
	return (affected > 0)
}
