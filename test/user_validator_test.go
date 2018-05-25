package test

import (
	"testing"

	"github.com/fumiyasac/SampleApi/validators"
)

// TestCaseIsValidForUserCreate ... ユーザー文字列の正規表現パターン
func TestCaseIsValidForUserCreate(t *testing.T) {
	userRequestValidator := validators.NewUserRequestValidator()

	// Case1 ユーザー名の文字数が不足しているケース
	actualOfCase1, messageOfCase1 := userRequestValidator.IsValidForUserCreate("test", "password1234", "test@example.com")
	expectedOfCase1 := false

	// Case2 ユーザー名の文字が半角英数字以外であるケース
	actualOfCase2, messageOfCase2 := userRequestValidator.IsValidForUserCreate("テストユーザー", "password1234", "test@example.com")
	expectedOfCase2 := false

	// Case3 パスワードの文字数が不足しているケース
	actualOfCase3, messageOfCase3 := userRequestValidator.IsValidForUserCreate("testuser1234", "test", "test@example.com")
	expectedOfCase3 := false

	// Case4 パスワードの文字が半角英数字以外であるケース
	actualOfCase4, messageOfCase4 := userRequestValidator.IsValidForUserCreate("testuser1234", "テストパスワード", "test@example.com")
	expectedOfCase4 := false

	// Case5 メールアドレスの形式がおかしなケース
	actualOfCase5, messageOfCase5 := userRequestValidator.IsValidForUserCreate("testuser1234", "password1234", "testexample.com")
	expectedOfCase5 := false

	if actualOfCase1 != expectedOfCase1 {
		t.Errorf("Case1 期待する値:(%v), 実際の値:(%v), メッセージ:(%v)", expectedOfCase1, actualOfCase1, messageOfCase1)
	}

	if actualOfCase2 != expectedOfCase2 {
		t.Errorf("Case2 期待する値:(%v), 実際の値:(%v), メッセージ:(%v)", expectedOfCase2, actualOfCase2, messageOfCase2)
	}

	if actualOfCase3 != expectedOfCase3 {
		t.Errorf("Case3 期待する値:(%v), 実際の値:(%v), メッセージ:(%v)", expectedOfCase3, actualOfCase3, messageOfCase3)
	}

	if actualOfCase4 != expectedOfCase4 {
		t.Errorf("Case4 期待する値:(%v), 実際の値:(%v), メッセージ:(%v)", expectedOfCase4, actualOfCase4, messageOfCase4)
	}

	if actualOfCase5 != expectedOfCase5 {
		t.Errorf("Case5 期待する値:(%v), 実際の値:(%v), メッセージ:(%v)", expectedOfCase5, actualOfCase5, messageOfCase5)
	}
}
