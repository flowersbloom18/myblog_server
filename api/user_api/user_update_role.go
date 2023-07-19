package user_api

import (
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/model_type"
	"myblog_server/models/response"
	"myblog_server/utils/jwt"
)

type UserRole struct {
	Role     model_type.Role `json:"role" binding:"required,oneof=1 2 3" msg:"权限参数错误"`
	NickName string          `json:"nick_name"` // 防止用户昵称非法，管理员有能力修改
	UserID   uint            `json:"user_id" binding:"required" msg:"用户id错误"`
}

// UserUpdateRoleView 用户权限变更(权限和昵称）
func (UserApi) UserUpdateRoleView(c *gin.Context) {
	// 获取当前登录的用户的id，如果需要修改的id中包含自己，则禁止修改。
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.Claims) // 断言

	var cr UserRole
	if err := c.ShouldBindJSON(&cr); err != nil {
		global.Log.Error("参数绑定错误：", err)
		response.FailWithError(err, &cr, c)
		return
	}
	var user models.User
	err := global.DB.Take(&user, cr.UserID).Error
	if err != nil {
		response.FailWithMessage("用户id错误，用户不存在", c)
		return
	}

	// *🥤系统禁止修改当前登录用户自身的状态！
	if claims.UserID == cr.UserID {
		response.FailWithMessage("系统禁止修改当前登录用户自身的状态！", c)
		return
	} else {
		err = global.DB.Model(&user).Updates(map[string]any{
			"role":      cr.Role,
			"nick_name": cr.NickName,
		}).Error
		if err != nil {
			global.Log.Error(err)
			response.FailWithMessage("修改权限失败", c)
			return
		}

		//⚠️系统日志记录
		logContent := "修改权限成功"
		global.DB.Create(&models.Log{
			UserName: user.UserName,
			NickName: user.NickName,
			Email:    user.Email,
			IP:       user.IP,
			Address:  user.Address,
			Device:   user.Device,
			Level:    "Info",
			Content:  logContent,
		})
		response.OkWithMessage("修改权限成功", c)
	}
}
