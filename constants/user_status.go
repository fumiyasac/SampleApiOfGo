package constants

// UserStatus ... ユーザーの状態に関する表示の定義
type UserStatus int

// UserStatusの定義される
const (
	UserSubscribed UserStatus = iota
	UserUnsubscribed
)

// GetString ... UserStatusからステータス名を取得する
func (u UserStatus) GetString() string {
	switch u {
	case UserSubscribed:
		return "有効"
	case UserUnsubscribed:
		return "無効"
	default:
		return "不明"
	}
}

// GetUserStatusNameFromStatusCode ... 引数のステータスコードからステータス名を取得する
func GetUserStatusNameFromStatusCode(statusCode int) string {
	switch UserStatus(statusCode) {
	case UserSubscribed:
		return "有効"
	case UserUnsubscribed:
		return "無効"
	default:
		return "不明"
	}
}
