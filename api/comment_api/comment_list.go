package comment_api

import (
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/response"
)

// CommentListView 管理员查看所有评论
func (CommentApi) CommentListView(c *gin.Context) {

	var cr models.PageInfo // contentRequest内容请求 -->cr
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}

	// 查询所有评论，分页查询，根据评论类型查询
	var comments []models.Comment
	query := global.DB.Debug().
		Order("created_at DESC").
		Offset((cr.Page - 1) * cr.Limit).
		Limit(cr.Limit)

	if cr.Key != "" {
		query = query.Where("page_type = ?", cr.Key)
	}

	// 🥤后期可优化，根据博客名称进行模糊查询
	//if cr.Title != "" {
	//	query = query.Where("title LIKE ?", "%"+cr.Title+"%")
	//}

	query.Find(&comments)

	count := query.RowsAffected

	response.OkWithList(comments, count, c)
}
