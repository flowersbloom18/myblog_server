package tag_api

import (
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/response"
	"myblog_server/service"
)

// GetBlogsByTag 找到该分类下的所有博客
func (TagApi) GetBlogsByTag(c *gin.Context) {
	db := global.DB
	name := c.Param("name")

	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		//response.FailWithMessage("请求参数有错误",c)
		response.FailWithError(err, &cr, c)
		return
	}

	// 限制每页查询10条
	pageSize := 10

	var tag models.Tag
	//	⚠️gorm分页查询标签下的博客的时候获得博客下的标签		Preload("Blogs.Tags")
	err = db.Where("name=?", name).Preload("Blogs.Tags").Offset((cr.Page - 1) * pageSize).Limit(pageSize).First(&tag).Error
	if err != nil {
		global.Log.Warn("err=", err)
	}

	count := int64(len(tag.Blogs))

	//对响应的结果中的tags进行优化
	serviceApp := service.ServiceApp.BlogService
	results, err := serviceApp.GetBlogList(tag.Blogs)
	if err != nil {
		global.Log.Warn("err=", err)
		response.FailWithMessage("查询出错-1", c)
		return
	}
	response.OkWithList(results, count, c)
	return
}
