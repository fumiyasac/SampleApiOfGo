package validators

import (
	"regexp"
)

// ユーザー文字列の正規表現パターン
var inputStringPattern = regexp.MustCompile(`^[a-zA-Z0-9]*$`)
var inputEmailPattern = regexp.MustCompile(`^(?i:[^ @"<>]+|".*")@(?i:[a-z1-9.])+.(?i:[a-z])+$`)

// UserRequestValidator ... 構造体宣言
type UserRequestValidator struct {
}

// NewUserRequestValidator ... 構造体の新規作成メソッド
func NewUserRequestValidator() UserRequestValidator {
	return UserRequestValidator{}
}

// IsValidForUserCreate ... 新規登録時のバリデーションを行う
func (v UserRequestValidator) IsValidForUserCreate(username string, password string, mailaddress string) (bool, string) {

	if !(len(username) >= 6 && 100 >= len(username)) {
		return false, "ユーザー名の文字数が不正です。"
	}

	if !(inputStringPattern.Match([]byte(username))) {
		return false, "ユーザー名の形式が不正です。"
	}

	if !(len(password) >= 6 && 100 >= len(password)) {
		return false, "パスワードの文字数が不正です。"
	}

	if !(inputStringPattern.Match([]byte(password))) {
		return false, "パスワードの形式が不正です。"
	}

	if !(inputEmailPattern.Match([]byte(mailaddress))) {
		return false, "メールアドレスの形式が不正です。"
	}

	return true, ""
}
