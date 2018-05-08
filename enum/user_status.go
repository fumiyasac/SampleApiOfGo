package enum

// UserStatusEnum ... ユーザーの状態に関する表示のEnum定義
type UserStatusEnum int

// Enumで定義される
const (
	Subscribe UserStatusEnum = iota
	Unsubscribe
)

// GetUserStatusName ... ステータスコードからステータス名を取得する
func (u UserStatusEnum) GetUserStatusName() string {
	switch u {
	case Subscribe:
		return "有効"
	case Unsubscribe:
		return "無効"
	default:
		return "不明"
	}
}
