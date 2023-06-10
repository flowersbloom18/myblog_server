package comment_api

import (
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/response"
	"myblog_server/service"
	"myblog_server/utils/jwt"
)

// CommentRemoveView 用户删除评论。如果删除博客，则删除博客下的所有评论
func (CommentApi) CommentRemoveView(c *gin.Context) {
	// 用户可以删除自己发表过的评论。管理员可以删除所有评论。

	// 如果面板id为0，且需要删除的是此面板数据，则会删除此面板下的所有数据。

	// 如果删除，父级id，则它下列的数据不删，然后如果获取评论找不到父级id，则返回（已删除评论）
	commentService := service.ServiceApp.CommentService

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.Claims)
	role := claims.Role
	userID := claims.UserID

	db := global.DB
	// 绑定请求参数
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}

	// 获取请求--评论ID列表
	var list []models.Comment
	count := db.Find(&list, cr.IDList).RowsAffected
	if count == 0 {
		response.FailWithMessage("评论ID不存在", c)
		return
	}

	err = commentService.RemoveCommentService(role, userID, list)
	if err != nil {
		response.FailWithError(err, "", c)
		return
	}
	response.OkWithMessage("评论删除成功", c)
}
