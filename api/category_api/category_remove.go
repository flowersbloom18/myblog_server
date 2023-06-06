package category_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/response"
)

// CategoryRemoveView 更新分类（名称和封面）
func (CategoryApi) CategoryRemoveView(c *gin.Context) {
	db := global.DB

	// 绑定请求参数
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}

	// 获取请求--分类ID列表
	var list []models.Category
	count := db.Find(&list, cr.IDList).RowsAffected
	if count == 0 {
		response.FailWithMessage("分类ID不存在", c)
		return
	}

	// 检查分类下是否存在文章,如果存在，则统一回复该分类下存在博客无法删除
	// 🥤方案一：通过where
	//for _, value := range cr.IDList {
	//	// 只要找到一个，直接退出
	//	var blogsCount int64
	//	err = db.Model(&models.BlogModel{}).Where("category_id = ?", value).Count(&blogsCount).Error
	//	if err != nil {
	//		response.FailWithError(err, &cr, c)
	//		return
	//	}
	//	// 如果博客数量大于0则存在
	//	if blogsCount > 0 {
	//		response.FailWithMessage("所选分类列表中存在博客，无法删除！", c)
	//		return
	//	} else {
	//		continue
	//	}
	//}

	// 🥤方案二：通过preload
	var categories []models.Category
	err = global.DB.Preload("Blogs").Find(&categories, cr.IDList).Error
	if err != nil {
		// 处理查询错误
		response.FailWithError(err, &cr, c)
		return
	}
	for _, category := range categories {
		blogCount := global.DB.Model(&category).Association("Blogs").Count()
		if blogCount > 0 {
			response.FailWithMessage("所选分类列表中存在博客，无法删除！", c)
			return
		} else {
			continue
		}
	}

	// 否则删除
	global.DB.Delete(&list)
	response.OkWithMessage(fmt.Sprintf("共删除 %d 个分类", count), c)
}
