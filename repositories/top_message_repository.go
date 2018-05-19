package repositories

import (
	"github.com/fumiyasac/SampleApi/factories"
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
func (repo TopMessageRepository) GetTopMessage() factories.TopMessageFactory {

	// TopAnnouncementFactoryの配列データ作成
	var topAnnouncementFactories = []factories.TopAnnouncementFactory{
		factories.TopAnnouncementFactory{
			ID:         2,
			Title:      "APIバージョン1.1を現在開発中です。",
			DateString: "2018-05-02",
			Category:   "おしらせ",
		},
		factories.TopAnnouncementFactory{
			ID:         1,
			Title:      "APIバージョン1.0を公開しました。",
			DateString: "2018-05-02",
			Category:   "おしらせ",
		},
	}

	// TopMessageFactoryモデルの作成
	var topMessageFactory = factories.TopMessageFactory{
		Title:         "Go言語で作るサンプルAPIテンプレート",
		Version:       "1.0",
		Detail:        "このAPIは自分でつくるサンプルアプリに合わせてカスタマイズする用のサンプルになります。適宜調整をして使ってみてください。",
		Announcements: topAnnouncementFactories,
	}
	return topMessageFactory
}
