package models

//TopMessage ... メインメッセージ用の構造体宣言
type TopMessage struct {
	Title         string         `json:"title"`
	Version       string         `json:"version"`
	Description   string         `json:"description"`
	Announcements []Announcement `json:"announcements"`
}

//Announcement ... お知らせ用の構造体宣言
type Announcement struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	DateString string `json:"dateString"`
	Category   string `json:"category"`
}

//Get ... メインメッセージ + n個のお知らせデータを取得

/* ↓↓↓ 悩ましい: これだとエラーになる(Controllerのところをまとめただけ) ↓↓↓
--------
func (model *TopMessage) Get() (topMessage TopMessage) {
	topMessage = TopMessage{
		Title:       "Go言語で作るサンプルAPIテンプレート",
		Version:     "1.0",
		Description: "このAPIは自分でつくるサンプルアプリに合わせてカスタマイズする用のサンプルになります。適宜調整をして使ってみてください。",
		Announcements: []Announcement{
			Announcement{
				ID:         2,
				Title:      "APIバージョン1.1を現在開発中です。",
				DateString: "2018-05-02",
				Category:   "おしらせ",
			},
			Announcement{
				ID:         1,
				Title:      "APIバージョン1.0を公開しました。",
				DateString: "2018-05-02",
				Category:   "おしらせ",
			},
		},
	}
	return
}
--------
*/
