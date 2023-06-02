package log_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/response"
)

func (LoginApi) LogRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}

	var list []models.LogModel
	count := global.DB.Find(&list, cr.IDList).RowsAffected
	if count == 0 {
		response.FailWithMessage("登录日志不存在", c)
		return
	}
	global.DB.Delete(&list)
	response.OkWithMessage(fmt.Sprintf("共删除 %d 个日志", count), c)
}
