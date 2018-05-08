package entities

import (
	"time"
)

// User ... 登録されているユーザー情報用の構造体
type User struct {
	ID         int       `xorm:"'id'"`
	Username   string    `xorm:"'username'"`
	Password   string    `xorm:"'password'"`
	StatusCode int       `xorm:"'status_code'"`
	CreatedAt  time.Time `xorm:"'created_at'"`
	UpdatedAt  time.Time `xorm:"'updated_at'"`
}
