package repositories

import (
	"github.com/fumiyasac/SampleApi/entities"
)

// TopMessageRepository ... 構造体宣言
type TopMessageRepository struct {
}

// NewTopMessageRepository ... 構造体の新規作成メソッド
func NewTopMessageRepository() TopMessageRepository {
	return TopMessageRepository{}
}

// GetTopMessage ... TOPメッセージを取得する
// @Get("api/v1/top")
func (m TopMessageRepository) GetTopMessage() entities.TopMessage {

	// TOPお知らせの定義
	var TopAnnouncements = []entities.TopAnnouncement{
		entities.TopAnnouncement{
			ID:         2,
			Title:      "APIバージョン1.1を現在開発中です。",
			DateString: "2018-05-02",
			Category:   "おしらせ",
		},
		entities.TopAnnouncement{
			ID:         1,
			Title:      "APIバージョン1.0を公開しました。",
			DateString: "2018-05-02",
			Category:   "おしらせ",
		},
	}

	// TOPメッセージの定義
	var TopMessage = entities.TopMessage{
		Title:         "Go言語で作るサンプルAPIテンプレート",
		Version:       "1.0",
		Detail:        "このAPIは自分でつくるサンプルアプリに合わせてカスタマイズする用のサンプルになります。適宜調整をして使ってみてください。",
		Announcements: TopAnnouncements,
	}
	return TopMessage
}
