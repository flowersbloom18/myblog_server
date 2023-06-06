package category_api

import (
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/response"
	"myblog_server/service"
)

// GetBlogsByCategory 找到该分类下的所有博客
func (CategoryApi) GetBlogsByCategory(c *gin.Context) {
	db := global.DB
	name := c.Param("name")

	var cr models.PageInfo // contentRequest内容请求 -->cr
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}

	// 根据分类名称查找id
	var category models.Category
	if err := db.Where("name=?", name).First(&category).Error; err != nil {
		response.FailWithMessage("该分类不存在！", c)
		return
	}
	global.Log.Warn("category.ID=", category.ID)

	// 查询总记录数
	var total int64
	if err := db.Model(&models.Blog{}).Where("category_id = ?", category.ID).Count(&total).Error; err != nil {
		global.Log.Warn("err=", err)
		return
	}

	// 如果为0则，表示查不到数据
	if total == 0 {
		response.OkWithMessage("该分类下不存在数据", c)
		return
	}

	// 第几页
	// 每一页多少条数据；固定10条
	pageSize := 10
	// 分页查询博客
	var blogs []models.Blog
	if err := db.Where("category_id = ?", category.ID).
		Offset((cr.Page - 1) * pageSize). //偏移量，相对于第一页便宜多少个数据
		Limit(pageSize).
		Preload("Tags"). // 预先加载Tags
		Find(&blogs).Error; err != nil {
		global.Log.Warn("err=", err)
		response.FailWithMessage("查询出错-0", c)
		return
	}

	//response.OkWithList(blogs, total, c)

	//对响应的结果中的tags进行优化

	serviceApp := service.ServiceApp.BlogService
	results, err := serviceApp.GetBlogList(blogs)
	if err != nil {
		global.Log.Warn("err=", err)
		response.FailWithMessage("查询出错-1", c)
		return
	}

	response.OkWithList(results, total, c)
	return
}
