package user_api

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/response"
	"myblog_server/utils/jwt"
	"strings"
)

type UserUpdateNicknameRequest struct {
	NickName string `json:"nick_name" structs:"nick_name"`
	Avatar   string `json:"avatar" structs:"avatar"`
}

// UserUpdateNickName 修改当前登录人的昵称，头像
func (UserApi) UserUpdateNickName(c *gin.Context) {
	var cr UserUpdateNicknameRequest

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.Claims)

	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}

	var newMaps = map[string]interface{}{}
	// 结构体转map，方便对数据去除空字符串
	maps := structs.Map(cr)
	for key, v := range maps {
		if val, ok := v.(string); ok && strings.TrimSpace(val) != "" {
			newMaps[key] = val
		}
	}

	var userModel models.User
	err = global.DB.Debug().Take(&userModel, claims.UserID).Error
	if err != nil {
		response.FailWithMessage("用户不存在", c)
		return
	}
	// 更新信息用map最好
	err = global.DB.Model(&userModel).Updates(newMaps).Error

	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage("修改用户信息失败", c)
		return
	}
	response.OkWithMessage("修改个人信息成功", c)

}
