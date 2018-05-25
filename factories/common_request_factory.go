package factories

// RequestUserJSONFactory ... JSONにマッピングする構造体の宣言
type RequestUserJSONFactory struct {
	Username    string `json:"username" binding:"required" example:"usernameis0123"`
	Password    string `json:"password" binding:"required" example:"passwordis0123"`
	Mailaddress string `json:"mailaddress" binding:"required" example:"0123@example.com"`
}
