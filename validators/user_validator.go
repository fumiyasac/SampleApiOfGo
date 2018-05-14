package validators

// UserRequestValidator ... 構造体宣言
type UserRequestValidator struct {
}

// NewUserRequestValidator ... 構造体の新規作成メソッド
func NewUserRequestValidator() UserRequestValidator {
	return UserRequestValidator{}
}

// IsValidForUserCreate ... 新規登録時のバリデーションを行う
func (v UserRequestValidator) IsValidForUserCreate(username string, password string) bool {

	

	return true
}
