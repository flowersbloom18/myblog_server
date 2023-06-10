package comment_api

import (
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/response"
)

// CommentFriendLinkListView 友链下评论展示
func (CommentApi) CommentFriendLinkListView(c *gin.Context) {

	// 查询友链
	var comments []models.Comment
	query := global.DB.Debug().
		Where("page_type = 2"). // 友链类型
		Order("created_at DESC")

	query.Find(&comments)

	count := query.RowsAffected

	response.OkWithList(comments, count, c)
}
