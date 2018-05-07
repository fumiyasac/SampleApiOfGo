package entities

// User ... 登録されているユーザー情報用の構造体
type User struct {
	ID       int    `xorm:"'id'"`
	Username string `xorm:"'username'"`
	Password string `xorm:"'password'"`
}
