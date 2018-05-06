package models

//TopMessage ... TOPメッセージ用の構造体宣言
type TopMessage struct {
	Title         string            `json:"title"`
	Version       string            `json:"version"`
	Detail        string            `json:"detail"`
	Announcements []TopAnnouncement `json:"announcements"`
}

// TopMessageRepository ...
type TopMessageRepository struct {
}

// NewTopMessageRepository  ...
func NewTopMessageRepository() TopMessageRepository {
	return TopMessageRepository{}
}

//
func (m TopMessageRepository) GetTopMessage() TopMessage {
	var TopMessage = TopMessage{
		Title:   "Go言語で作るサンプルAPIテンプレート",
		Version: "1.0",
		Detail:  "このAPIは自分でつくるサンプルアプリに合わせてカスタマイズする用のサンプルになります。適宜調整をして使ってみてください。",
		Announcements: []TopAnnouncement{
			TopAnnouncement{
				ID:         2,
				Title:      "APIバージョン1.1を現在開発中です。",
				DateString: "2018-05-02",
				Category:   "おしらせ",
			},
			TopAnnouncement{
				ID:         1,
				Title:      "APIバージョン1.0を公開しました。",
				DateString: "2018-05-02",
				Category:   "おしらせ",
			},
		},
	}
	return TopMessage
}
