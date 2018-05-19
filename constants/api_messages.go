package constants

// ユーザー登録に関する成功メッセージ
const (
	SuccessForUserCreate = "新規ユーザーの登録に成功しました。"
	SuccessForUserUpdate = "既存ユーザーの更新に成功しました。"
	SuccessForUserDelete = "既存ユーザーの削除に成功しました。"
)

// ユーザー登録に関するエラーメッセージ
const (
	ErrorMessageOfInvalidID            = "ユーザーIDが不正です。"
	ErrorMessageOfInvalidRequest       = "ユーザーのリクエストが不正です。"
	ErrorMessageOfUserNotFound         = "該当のユーザーが見つかりませんでした。"
	ErrorMessageOfFailureForUserCreate = "新規ユーザーの登録に失敗しました。"
	ErrorMessageOfFailureForUserUpdate = "新規ユーザーの更新に失敗しました。"
	ErrorMessageOfFailureForUserDelete = "新規ユーザーの削除に失敗しました。"
)
