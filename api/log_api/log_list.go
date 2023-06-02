package log_api

import (
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/response"
	"myblog_server/service/common"
)

// LogView 获取系统日志
func (LoginApi) LogView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}

	list, count, err := common.ComList(models.LogModel{}, common.Option{
		PageInfo: cr,
		Debug:    true,
	})
	if err != nil {
		global.Log.Warn("获取数据错误：", err)
	}

	response.OkWithList(list, count, c)
}
