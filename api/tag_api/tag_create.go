package tag_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models/response"
	"myblog_server/service"
)

type TagCreateRequest struct {
	Name  string `json:"name" binding:"required" msg:"请输入标签名称"`
	Cover string `json:"cover"`
}

func (TagApi) TagCreateView(c *gin.Context) {
	serviceApp := service.ServiceApp
	var cr TagCreateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		response.FailWithError(err, &cr, c)
		return
	}

	// 创建的service
	err := serviceApp.TagService.CreateTag(cr.Name, cr.Cover)
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 响应成功
	response.OkWithMessage(fmt.Sprintf("标签'%s'创建成功!", cr.Name), c)
	return
}
