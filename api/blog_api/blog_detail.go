package blog_api

import (
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/response"
	"myblog_server/service"
	"strings"
)

// BlogDetailView 更新分类（名称和封面）
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

	// 对结果再次封装
	blogResponse, err := blogService.GetBlogDetail(blog)
	if err != nil {
		response.FailWithMessage("获取数据错误", c)
		return
	}

	response.OkWithData(blogResponse, c)
	return
}
