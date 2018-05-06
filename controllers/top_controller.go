package controllers

//パッケージの宣言
import (
	"github.com/fumiyasac/SampleApi/models"
	"github.com/gin-gonic/gin"
)

//TopController ... GuideControllerの構造体宣言
type TopController struct{}

//GetMessage ...
func (ctrl *TopController) GetMessage(c *gin.Context) {

	TopMessage := models.TopMessage{
		Title:   "Go言語で作るサンプルAPIテンプレート",
		Version: "1.0",
		Detail:  "このAPIは自分でつくるサンプルアプリに合わせてカスタマイズする用のサンプルになります。適宜調整をして使ってみてください。",
		Announcements: []models.TopAnnouncement{
			models.TopAnnouncement{
				ID:         2,
				Title:      "APIバージョン1.1を現在開発中です。",
				DateString: "2018-05-02",
				Category:   "おしらせ",
			},
			models.TopAnnouncement{
				ID:         1,
				Title:      "APIバージョン1.0を公開しました。",
				DateString: "2018-05-02",
				Category:   "おしらせ",
			},
		},
	}

	JSONContent := gin.H{
		"top_message": TopMessage,
	}
	c.JSON(200, JSONContent)
}
