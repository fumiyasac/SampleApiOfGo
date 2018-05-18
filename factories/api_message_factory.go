package factories

// APIErrorMessageFactory ... JSONにマッピングする構造体の宣言
type APIErrorMessageFactory struct {
	Message string `json:"error"`
}

// APISuccessMessageFactory ... JSONにマッピングする構造体の宣言
type APISuccessMessageFactory struct {
	Message string `json:"success"`
}
