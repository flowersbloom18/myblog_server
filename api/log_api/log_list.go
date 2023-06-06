package log_api

import (
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/model_type"
	"myblog_server/models/response"
	"myblog_server/service/common"
	"myblog_server/utils/desensitization"
	"myblog_server/utils/jwt"
)

type LogResponse struct {
	Log models.Log
}

// LogView 获取系统日志
func (LoginApi) LogView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.Claims)

	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}

	list, count, err := common.ComList(models.Log{}, common.Option{
		PageInfo: cr,
		Debug:    true,
	})
	if err != nil {
		global.Log.Warn("获取数据错误：", err)
	}

	var users []LogResponse
	// 如果不是管理员，则对系统日志的信息进行脱敏处理
	for _, value := range list {
		if model_type.Role(claims.Role) != model_type.PermissionAdmin {
			value.UserName = desensitization.DesensitizationUserName(value.UserName)
		}
		users = append(users, LogResponse{
			Log: value,
		})
	}
	response.OkWithList(users, count, c)
}
