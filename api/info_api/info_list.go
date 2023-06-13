package info_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models/response"
	"myblog_server/service"
	"myblog_server/utils/info"
	"strconv"
)

func (InfoApi) InfoListView(c *gin.Context) {
	infoService := service.ServiceApp.InfoService

	// 获取请求的数据参数,并转为整数类型
	_id := c.Param("id")
	id, err := strconv.Atoi(_id)
	if err != nil {
		global.Log.Error("error=", err)
		response.FailWithMessage("请求参数有误。", c)
		return
	}

	// 查找数据
	content, err := infoService.GetInfoService(id)
	if content == "" {
		global.Log.Error("error=", err)
		response.FailWithMessage(fmt.Sprintf("error=%s", err), c)
		return
	}

	// 解析
	result, err := info.GetInfoResult(content, id)
	if err != nil {
		global.Log.Error("error=", err)
		response.FailWithMessage(fmt.Sprintf("error=%s", err), c)
		return
	}
	response.OkWithData(result, c)
}
