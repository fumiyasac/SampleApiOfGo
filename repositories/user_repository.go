package repositories

import (
	"github.com/fumiyasac/SampleApi/constants"
	"github.com/fumiyasac/SampleApi/database"
	"github.com/fumiyasac/SampleApi/entities"
	"github.com/fumiyasac/SampleApi/factories"
	"github.com/go-xorm/xorm"
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
// @Get("api/v1/user/:id")
func (repo UserRepository) GetByID(id int) (factories.SingleUserFactory, bool) {
	var user = entities.User{}
	var userFactory factories.SingleUserFactory
	result, _ := engine.Where("id = ?", id).Get(&user)
	if result {
		userFactory = factories.SingleUserFactory{
			ID:         user.ID,
			Username:   user.Username,
			Password:   user.Password,
			UserStatus: constants.GetUserStatusNameFromStatusCode(user.StatusCode),
			CreatedAt:  user.CreatedAt,
			UpdatedAt:  user.UpdatedAt,
		}
	}
	return userFactory, result
}
