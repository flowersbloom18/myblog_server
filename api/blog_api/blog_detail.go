package blog_api

import (
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/response"
	"myblog_server/service"
	"strings"
)

// BlogDetailView 博客详情
func (BlogApi) BlogDetailView(c *gin.Context) {
	blogService := service.ServiceApp.BlogService
	db := global.DB

	// 获取需要更新的博客链接
	link := c.Param("link")
	// 去除前缀斜杠
	link = strings.TrimPrefix(link, "/")

	// 链接是否存在！
	var blog models.Blog
	// 预加载博客的标签
	err := db.Preload("Tags").Take(&blog, "link=?", link).Error
	if err != nil {
		response.FailWithMessage("博客不存在！", c)
		return
	}

	// 收藏数量统计
	var collect []models.Collect
	_count := db.Where("blog_id=?", blog.ID).Find(&collect).RowsAffected
	count := int(_count)
	// 如果博客中的收藏数量不等于真实的收藏数量，则更新数据。
	if count != blog.CollectNum {
		blog.CollectNum = count
		db.Save(&blog)
	}

	// 评论数量统计
	var comments []models.Comment
	result := global.DB.Debug().
		Where("page_type = 1").       // 博客类型
		Where("page = ?", blog.Link). // 页面
		Find(&comments)
	commentNum := int(result.RowsAffected)
	// 如果博客中的评论数量不等于真实的评论数量，则更新数据。
	if commentNum != blog.CommentNum {
		blog.CommentNum = commentNum
		db.Save(&blog)
	}

	// 博客阅读数量+1
	blog.ReadNum += 1
	if err := db.Save(&blog).Error; err != nil {
		global.Log.Error("博客保存失败。")
		//return
	}

	// 对结果再次封装
	blogResponse, err := blogService.GetBlogDetail(blog)
	if err != nil {
		response.FailWithMessage("获取数据错误", c)
		return
	}

	response.OkWithData(blogResponse, c)
	return
}
