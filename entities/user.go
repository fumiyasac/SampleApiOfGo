package entities

// User ... 登録されているユーザー情報用の構造体
type User struct {
	ID       int    `json:"id" xorm:"'id'"`
	Username string `json:"username" xorm:"'username'"`
	Password string `json:"password" xorm:"'password'"`
}
