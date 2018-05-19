package factories

// RequestUserJSONFactory ... JSONにマッピングする構造体の宣言
type RequestUserJSONFactory struct {
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Mailaddress string `json:"mailaddress" binding:"required"`
}
