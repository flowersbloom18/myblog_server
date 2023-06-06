package category_api

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/response"
)

type CategoryUpdateRequest struct {
	Name  string `json:"name"`
	Cover string `json:"cover"`
}

// CategoryUpdateView  更新分类（名称和封面）
func (CategoryApi) CategoryUpdateView(c *gin.Context) {
	// 获取需要更新的分类ID
	id := c.Param("id")

	db := global.DB

	var cr CategoryUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithError(err, &cr, c)
		return
	}

	var category models.Category
	// 获取对应id的数据
	err = db.Take(&category, "id = ?", id).Error
	if err != nil {
		global.Log.Warn("分类不存在")
		response.FailWithMessage("分类不存在", c)
		return
	}

	// 检查新创建的分类名称是否存在，不存在则更新【分类重复值判断】
	var category1 models.Category
	// ⚠️如果说新增分类跟当前数据库【不同】则进行重复值判断，否则不用。
	if category.Name != cr.Name {
		err = db.First(&category1, "name = ?", cr.Name).Error
		// 查询到数据则err为空，说明分类已经存在。反之不存在，即可更新数据。
		if err == nil {
			global.Log.Warn("分类已存在", err)
			response.FailWithMessage("分类已存在", c)
			return
		}
	}

	// 更新
	maps := structs.Map(&cr)
	err = global.DB.Model(&category).Updates(maps).Error
	if err != nil {
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
	return
}
