package user_api

import (
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/response"
	"myblog_server/plugins/email"
	"myblog_server/service"
	"myblog_server/utils/jwt"
	"myblog_server/utils/pwd"
)

// UpdatePasswordRequest 修改当前登录用户的密码
type UpdatePasswordRequest struct {
	OldPwd string `json:"old_pwd" binding:"required" msg:"请输入旧密码"` // 旧密码
	Pwd    string `json:"pwd" binding:"required" msg:"请输入新密码"`     // 新密码
}

// UserUpdatePassword 修改登录人的密码
func (UserApi) UserUpdatePassword(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.Claims)

	var cr UpdatePasswordRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		response.FailWithError(err, &cr, c)
		return
	}
	var user models.User
	err := global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		response.FailWithMessage("用户不存在", c)
		return
	}

	if len(cr.Pwd) < 4 {
		response.FailWithMessage("密码强度太弱", c)
		return
	}

	// 判断密码是否一致
	if !pwd.CheckPwd(user.Password, cr.OldPwd) {
		response.FailWithMessage("密码错误", c)
		return
	}
	hashPwd := pwd.HashPwd(cr.Pwd)
	err = global.DB.Model(&user).Update("password", hashPwd).Error
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage("密码修改失败", c)
		return
	}
	response.OkWithMessage("密码修改成功", c)

	// ⚠️系统日志记录
	logContent := "密码修改成功"
	global.DB.Create(&models.Log{
		UserName: user.UserName,
		NickName: user.NickName,
		IP:       user.IP,
		Address:  user.Address,
		Device:   user.Device,
		Level:    "info",
		Content:  logContent,
	})

	// 🥤密码更新提醒
	sendApi := email.SendEmailApi{}
	err = sendApi.SendUpdatePwd(user.Email)
	if err != nil {
		global.Log.Error("邮箱发送失败", err)
	}

	// *🥤密码修改成功后，原先的token自动注销，需要重新登录，前台则跳转到登录页
	token := c.Request.Header.Get("token")
	err = service.ServiceApp.UserService.Logout(claims, token)
	if err != nil {
		response.FailWithMessage("注销失败，未知错误！", c)
	}
	return
}
