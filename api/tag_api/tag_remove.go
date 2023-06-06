package tag_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/response"
)

// TagRemoveView 删除标签（如果下面存在博客则不删除）
func (TagApi) TagRemoveView(c *gin.Context) {
	db := global.DB

	// 绑定请求参数
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}

	// 获取请求--标签ID列表
	var list []models.Tag
	count := db.Find(&list, cr.IDList).RowsAffected
	if count == 0 {
		response.FailWithMessage("标签ID不存在", c)
		return
	}

	// 判断标签下是否存在博客（只要有一个标签存在博客，就返回错误）
	var tags []models.Tag
	err = global.DB.Preload("Blogs").Find(&tags, cr.IDList).Error
	if err != nil {
		// 处理查询错误
		response.FailWithError(err, &cr, c)
		return
	}

	for _, tag := range tags {
		// 如果标签下不存在博客，则继续下一个，直到结束
		if len(tag.Blogs) == 0 {
			continue
		} else {
			response.FailWithMessage("所选标签列表中存在博客，无法删除！", c)
			return
		}
	}

	// 否则删除
	global.DB.Delete(&list)
	response.OkWithMessage(fmt.Sprintf("共删除 %d 个标签", count), c)
}
