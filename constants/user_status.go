package constants

// UserStatus ... ユーザーの状態に関する表示の定義
type UserStatus int

// UserStatusの定義される
const (
	UserSubscribed UserStatus = iota
	UserUnsubscribed
	UserUnknown
)

// GetRawValue ... UserStatusから値を取得する
func (u UserStatus) GetRawValue() int {
	switch u {
	case UserSubscribed:
		return 0
	case UserUnsubscribed:
		return 1
	default:
		return 2
	}
}

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
		return UserSubscribed.GetString()
	case UserUnsubscribed:
		return UserUnsubscribed.GetString()
	default:
		return UserUnknown.GetString()
	}
}
