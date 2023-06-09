package friendlink_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models/response"
	"myblog_server/service"
)

type FriendLinkCreateRequest struct {
	Name        string `json:"name" binding:"required" msg:"请输入友链名称"`
	Description string `json:"description"`
	Logo        string `json:"logo"`
	Url         string `json:"url" binding:"required" msg:"请输入友链的url"`
	IsTop       bool   `json:"is_top" `
}

func (FriendLinkApi) FriendLinkCreateView(c *gin.Context) {
	friendlinkService := service.ServiceApp.FriendLinkService

	var cr FriendLinkCreateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		response.FailWithError(err, &cr, c)
		return
	}

	// 创建的service
	err := friendlinkService.CreateFriendLink(cr.Name, cr.Description, cr.Logo, cr.Url, cr.IsTop)
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 响应成功
	response.OkWithMessage(fmt.Sprintf("友链'%s'创建成功!", cr.Name), c)
	return
}
