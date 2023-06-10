package comment_api

import (
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/response"
)

// CommentAboutListView 关于下评论展示
func (CommentApi) CommentAboutListView(c *gin.Context) {

	// 查询关于
	var comments []models.Comment
	query := global.DB.Debug().
		Where("page_type = 3"). // 关于类型
		Order("created_at DESC")

	query.Find(&comments)

	count := query.RowsAffected

	response.OkWithList(comments, count, c)
}
