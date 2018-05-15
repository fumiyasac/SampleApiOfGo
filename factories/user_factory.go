package factories

import (
	"time"
)

// SingleUserFactory ... JSONにマッピングする構造体の宣言
type SingleUserFactory struct {
	ID          int       `json:"id"`
	Username    string    `json:"username"`
	MailAddress string    `json:"mailaddress"`
	Password    string    `json:"password"`
	UserStatus  string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// UserItemsFactory ... JSONにマッピングする構造体の宣言
type UserItemsFactory struct {
	Items []SingleUserFactory `json:"items"`
}
