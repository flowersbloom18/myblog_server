package blog_api

import (
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/response"
	"sync"
)

// BlogLikeView 博客点赞
func (BlogApi) BlogLikeView(c *gin.Context) {
	var mutex sync.Mutex
	db := global.DB
	blogID := c.Param("id")

	// 使用互斥锁保护点赞操作，同一时间只有一个请求等待，别的都不行
	mutex.Lock()
	defer mutex.Unlock()

	// 更新MySQL数据库中的点赞数量
	var blog models.Blog
	if err := db.First(&blog, blogID).Error; err != nil {
		response.FailWithMessage("博客不存在", c)
		return
	}

	blog.LikeNum += 1

	// 保存
	if err := db.Save(&blog).Error; err != nil {
		response.FailWithMessage("更新点赞失败", c)
		return
	}

	response.OkWithMessage("点赞成功", c)
}
