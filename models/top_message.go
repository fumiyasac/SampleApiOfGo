package models

//TopMessage ... TOPメッセージ用の構造体宣言
type TopMessage struct {
	Title         string            `json:"title"`
	Version       string            `json:"version"`
	Detail        string            `json:"detail"`
	Announcements []TopAnnouncement `json:"announcements"`
}
