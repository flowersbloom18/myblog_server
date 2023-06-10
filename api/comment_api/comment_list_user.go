package comment_api

import (
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/response"
	"myblog_server/utils/jwt"
)

// CommentUserListView 用户的评论列表【我的评论】
func (CommentApi) CommentUserListView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.Claims)

	userID := claims.UserID // 用户id
	//isAdmin := false        // 是否为管理员
	//if claims.Role == 1 {
	//	isAdmin = true
	//}

	var cr models.PageInfo // contentRequest内容请求 -->cr
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}

	// 分页查询，根据当前登录用户的ID，查询类型
	var comments []models.Comment
	query := global.DB.Debug().Where("user_id = ?", userID).
		Order("created_at DESC").
		Offset((cr.Page - 1) * cr.Limit).
		Limit(cr.Limit)

	if cr.Key != "" {
		query = query.Where("page_type = ?", cr.Key)
	}

	query.Find(&comments)

	count := query.RowsAffected

	response.OkWithList(comments, count, c)
}
