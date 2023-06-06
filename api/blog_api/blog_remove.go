package blog_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/response"
)

// BlogRemoveView 删除博客（并删除标签的关联关系）
func (BlogApi) BlogRemoveView(c *gin.Context) {
	db := global.DB

	// 绑定请求参数
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}

	// 获取请求--博客ID列表
	var blogs []models.Blog
	count := db.Find(&blogs, cr.IDList).RowsAffected
	if count == 0 {
		response.FailWithMessage("博客ID不存在", c)
		return
	}

	// ⚠️开启一个事务，以确保删除操作的原子性
	// 使用事务可以确保多个数据库操作要么全部成功，要么全部失败，保持数据库的一致性
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			global.Log.Warn("err=", err)
			response.FailWithMessage("删除失败-0", c)
			return
		}
	}()

	// 遍历博客列表，依次删除每个博客的关联标签。
	for _, blog := range blogs {
		if err := tx.Model(&blog).Association("Tags").Clear(); err != nil {
			tx.Rollback()
			global.Log.Warn("err=", err)
			response.FailWithMessage("删除博客标签关联关系出错-1", c)
			return
		}
	}

	// 删除博客记录
	if err := tx.Delete(&blogs).Error; err != nil {
		tx.Rollback()
		response.FailWithMessage("删除博客出错-2", c)
		return
	}

	// 如果没有发生错误，提交事务。
	if err := tx.Commit().Error; err != nil {
		response.FailWithMessage("提交事务出错-3", c)
		return
	}

	response.OkWithMessage(fmt.Sprintf("共删除 %d 个博客", count), c)
}
