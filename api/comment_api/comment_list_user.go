package comment_api

import (
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/response"
	"myblog_server/service/comment_service"
	"myblog_server/utils/jwt"
)

// CommentUserListView 用户的评论列表【我的评论】
func (CommentApi) CommentUserListView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.Claims)

	userID := claims.UserID // 用户id

	// 分页查询，根据当前登录用户的ID，查询类型
	var comments []models.Comment
	query := global.DB.Debug().Where("user_id = ?", userID).
		Order("created_at DESC")

	query.Find(&comments)

	count := query.RowsAffected

	var responseComment = comment_service.CommentService{}
	result := responseComment.ResponseCommentService(comments)
	response.OkWithList(result, count, c)
}
