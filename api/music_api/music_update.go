package music_api

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/response"
)

type MusicUpdateRequest struct {
	Name   string `json:"name"`
	Author string `json:"author" `
	Url    string `json:"url"`
	Cover  string `json:"cover"`
	Status bool   `json:"status"`
	Sort   int    `json:"sort"`
}

func (MusicApi) MusicUpdateView(c *gin.Context) {
	// 获取需要更新的音乐ID
	id := c.Param("id")

	db := global.DB

	var cr MusicUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithError(err, &cr, c)
		return
	}

	var Music models.Music
	// 获取对应id的数据
	err = db.Take(&Music, "id = ?", id).Error
	if err != nil {
		global.Log.Warn("音乐不存在")
		response.FailWithMessage("音乐不存在", c)
		return
	}

	// 检查新创建的音乐名称是否存在，不存在则更新【音乐重复值判断】
	var Music1 models.Music
	// ⚠️如果说新增音乐跟当前数据库【不同】则进行重复值判断，否则不用。
	if Music.Name != cr.Name {
		err = db.First(&Music1, "name = ?", cr.Name).Error
		// 查询到数据则err为空，说明音乐已经存在。反之不存在，即可更新数据。
		if err == nil {
			global.Log.Warn("音乐已存在", err)
			response.FailWithMessage("音乐已存在", c)
			return
		}
	}

	// 更新
	maps := structs.Map(&cr)
	err = global.DB.Model(&Music).Updates(maps).Error
	if err != nil {
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
	return
}
