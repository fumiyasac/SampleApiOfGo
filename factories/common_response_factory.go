package factories

// APIErrorMessageFactory ... JSONにマッピングする構造体の宣言
type APIErrorMessageFactory struct {
	Message string `json:"error" example:"APIの処理に失敗しました。"`
}

// APISuccessMessageFactory ... JSONにマッピングする構造体の宣言
type APISuccessMessageFactory struct {
	Message string `json:"success" example:"APIの処理に成功しました。"`
}
