package factories

import (
	"time"
)

// SingleUserFactory ... JSONにマッピングする構造体の宣言
type SingleUserFactory struct {
	ID          int       `json:"id" example:"1"`
	Username    string    `json:"username" example:"samplename"`
	MailAddress string    `json:"mailaddress" example:"sample@example.com"`
	Password    string    `json:"password" example:"(crypted strings...)"`
	UserStatus  string    `json:"status" example:"有効"`
	CreatedAt   time.Time `json:"created_at" example:"2018-05-15T12:59:09+09:00"`
	UpdatedAt   time.Time `json:"updated_at" example:"2018-05-15T12:59:09+09:00"`
}

// UserItemsFactory ... JSONにマッピングする構造体の宣言
type UserItemsFactory struct {
	Users []SingleUserFactory `json:"users"`
}
