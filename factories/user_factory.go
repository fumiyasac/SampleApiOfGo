package factories

// SingleUserFactory ... JSONにマッピングする構造体の宣言
type SingleUserFactory struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
