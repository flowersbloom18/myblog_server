package comment_api

import (
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/response"
)

// CommentBlogListView 博客下评论展示【那友链、关于呢？一个一个来】
func (CommentApi) CommentBlogListView(c *gin.Context) {
	var cr models.PageInfo // contentRequest内容请求 -->cr
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}

	if cr.Key == "" {
		response.FailWithMessage("请求参数的页面为空", c)
		return
	}
	// 根据查询页面查询博客
	var comments []models.Comment
	query := global.DB.Debug().
		Where("page_type = 1").    // 博客类型
		Where("page = ?", cr.Key). // 页面
		Order("created_at DESC")

	query.Find(&comments)

	count := query.RowsAffected

	response.OkWithList(comments, count, c)
}
