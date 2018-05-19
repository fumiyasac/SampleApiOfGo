package factories

import (
	"github.com/fumiyasac/SampleApi/entities"
)

// TopMessageFactory ... JSONにマッピングする構造体の宣言
type TopMessageFactory struct {
	Title         string                     `json:"title" example:"タイトル情報が入ります"`
	Version       string                     `json:"version" example:"1.x"`
	Detail        string                     `json:"detail" example:"詳細情報が入ります"`
	Announcements []entities.TopAnnouncement `json:"announcements"`
}
