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

// GetByID ... 引数のidに該当するユーザー情報を取得する
// @Get("api/v1/users/:id")
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
// @Post("api/v1/users")
func (repo UserRepository) Create(username string, password string, mailaddress string) bool {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return false
	}
	var user = entities.User{
		Username:    username,
		MailAddress: mailaddress,
		Password:    string(hash),
		StatusCode:  constants.UserSubscribed.GetRawValue(),
	}
	affected, _ := engine.Insert(&user)
	return (affected > 0)
}
