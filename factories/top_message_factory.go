package factories

import (
	"github.com/fumiyasac/SampleApi/entities"
)

// TopMessageFactory ... JSONにマッピングする構造体の宣言
type TopMessageFactory struct {
	Title         string                     `json:"title"`
	Version       string                     `json:"version"`
	Detail        string                     `json:"detail"`
	Announcements []entities.TopAnnouncement `json:"announcements"`
}
