package blog_api

import (
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/response"
	"myblog_server/service"
)

// BlogContentView 博客内容用来针对博主后台编辑使用的
func (BlogApi) BlogContentView(c *gin.Context) {
	blogService := service.ServiceApp.BlogService

	blogID := c.Param("id")
	var blog models.Blog
	if err := global.DB.First(&blog, blogID).Error; err != nil {
		response.FailWithMessage("博客不存在", c)
		return
	}

	// 预加载博客的标签
	err := global.DB.Preload("Tags").Take(&blog, "id=?", blogID).Error
	if err != nil {
		response.FailWithMessage("博客不存在！", c)
		return
	}

	// 对结果再次封装
	blogResponse, err := blogService.GetBlogList([]models.Blog{blog})
	if err != nil {
		response.FailWithMessage("获取数据错误", c)
		return
	}

	response.OkWithData(blogResponse, c)
	return
}
