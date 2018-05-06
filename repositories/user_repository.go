package repositories

import (
	"github.com/fumiyasac/SampleApi/database"
	"github.com/fumiyasac/SampleApi/entities"
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
func (m UserRepository) GetByID(id int) *entities.User {
	var user = entities.User{ID: id}
	result, _ := engine.Get(&user)
	if result {
		return &user
	}
	return nil
}
