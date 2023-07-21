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
	"time"
)

type LogResponse struct {
	ID        uint      `json:"id,select($any)" structs:"-"`         // 主键ID
	CreatedAt time.Time `json:"created_at,select($any)" structs:"-"` // 创建时间
	UpdatedAt time.Time `json:"-" structs:"-"`                       // 更新时间
	UserName  string    `json:"user_name"`                           // 用户名
	NickName  string    `json:"nick_name"`                           // 昵称
	Email     string    `json:"email"`                               // 邮箱
	IP        string    `json:"ip"`                                  // ip
	Address   string    `json:"address"`                             // 地址
	Device    string    `json:"device"`                              // 登录设备
	Level     string    `json:"level"`                               // 日志水平
	Content   string    `json:"content"`                             // 日志内容
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
		//Debug:    true,
		Likes: []string{"level"}, // 按照音乐名查询
	})
	if err != nil {
		global.Log.Warn("获取数据错误：", err)
	}

	var users []LogResponse
	// 如果不是管理员，则对系统日志的信息进行脱敏处理
	for _, value := range list {
		if model_type.Role(claims.Role) != model_type.PermissionAdmin {
			value.UserName = desensitization.DesensitizationUserName(value.UserName)
			value.Email = desensitization.DesensitizationEmail(value.Email)
			value.IP = "****"
		}
		users = append(users, LogResponse{
			ID:        value.ID,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
			UserName:  value.UserName,
			NickName:  value.NickName,
			Email:     value.Email,
			IP:        value.IP,
			Address:   value.Address,
			Device:    value.Device,
			Level:     value.Level,
			Content:   value.Content,
		})
	}
	response.OkWithList(users, count, c)
}
