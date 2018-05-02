package controllers

//パッケージの宣言
import (
	"github.com/fumiyasac/SampleApi/models"
	"github.com/gin-gonic/gin"
)

//APIController ... APIControllerの構造体宣言
type APIController struct{}

//Top ... Top(トップ部分)のエンドポイント定義
func (ctrl *APIController) Top(c *gin.Context) {

	//TODO: できればmodelへ移動したい
	TopMessage := models.TopMessage{
		Title:       "Go言語で作るサンプルAPIテンプレート",
		Version:     "1.0",
		Description: "このAPIは自分でつくるサンプルアプリに合わせてカスタマイズする用のサンプルになります。適宜調整をして使ってみてください。",
		Announcements: []models.Announcement{
			models.Announcement{
				ID:         2,
				Title:      "APIバージョン1.1を現在開発中です。",
				DateString: "2018-05-02",
				Category:   "おしらせ",
			},
			models.Announcement{
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
