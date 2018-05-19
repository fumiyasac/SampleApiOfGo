package factories

// TopMessageFactory ... JSONにマッピングする構造体の宣言
type TopMessageFactory struct {
	Title         string                   `json:"title" example:"タイトル情報が入ります"`
	Version       string                   `json:"version" example:"1.x"`
	Detail        string                   `json:"detail" example:"詳細情報が入ります"`
	Announcements []TopAnnouncementFactory `json:"announcements"`
}

// TopAnnouncementFactory ... TOPお知らせ用の構造体
type TopAnnouncementFactory struct {
	ID         int    `json:"id" example:"1"`
	Title      string `json:"title" example:"お知らせのタイトル情報が入ります"`
	DateString string `json:"dateString" example:"2018-05-02"`
	Category   string `json:"category" example:"お知らせのカテゴリ情報が入ります"`
}
