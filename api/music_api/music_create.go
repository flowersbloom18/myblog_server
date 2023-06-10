package music_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models/response"
	"myblog_server/service"
)

// name，author，url，cover，status，sort
type MusicCreateRequest struct {
	Name   string `json:"name" binding:"required" msg:"请输入音乐名称"`
	Author string `json:"author" `
	Url    string `json:"url" binding:"required" msg:"请输入音乐的url"`
	Cover  string `json:"cover" binding:"required" msg:"请输入音乐的封面"`
	Status bool   `json:"status"`
	Sort   int    `json:"sort"`
}

func (MusicApi) MusicCreateView(c *gin.Context) {
	musicService := service.ServiceApp.MusicService

	var cr MusicCreateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		response.FailWithError(err, &cr, c)
		return
	}

	// 创建的service
	err := musicService.CreateMusic(cr.Name, cr.Author, cr.Url, cr.Cover, cr.Status, cr.Sort)
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 响应成功
	response.OkWithMessage(fmt.Sprintf("音乐'%s'创建成功!", cr.Name), c)
	return
}
