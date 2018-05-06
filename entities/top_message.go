package entities

// TopMessage ... TOPメッセージ用の構造体
type TopMessage struct {
	Title         string            `json:"title"`
	Version       string            `json:"version"`
	Detail        string            `json:"detail"`
	Announcements []TopAnnouncement `json:"announcements"`
}
