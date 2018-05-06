package entities

// TopAnnouncement ... TOPお知らせ用の構造体
type TopAnnouncement struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	DateString string `json:"dateString"`
	Category   string `json:"category"`
}
