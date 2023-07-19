package comment_api

import (
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/response"
	"myblog_server/service/comment_service"
)

// CommentListView 管理员查看所有评论
func (CommentApi) CommentListView(c *gin.Context) {
	// 查询所有评论的数量
	var count int64
	global.DB.Model(&models.Comment{}).Count(&count)

	// 查询所有评论
	var comments []models.Comment
	// 设置查询条件
	query := global.DB.Order("created_at DESC")
	// 按照条件将查询的数据存入comments中
	query.Find(&comments)

	var responseComment = comment_service.CommentService{}
	result := responseComment.ResponseCommentService(comments)

	response.OkWithList(result, count, c)
}
