package music_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/response"
)

func (MusicApi) MusicRemoveView(c *gin.Context) {
	db := global.DB

	// 绑定请求参数
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}

	// 获取请求--音乐ID列表
	var list []models.Music
	count := db.Find(&list, cr.IDList).RowsAffected
	if count == 0 {
		response.FailWithMessage("音乐ID不存在", c)
		return
	}

	// 删除
	global.DB.Delete(&list)
	response.OkWithMessage(fmt.Sprintf("共删除 %d 个音乐", count), c)
}
