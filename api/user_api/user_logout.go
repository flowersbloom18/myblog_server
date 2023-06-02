package user_api

import (
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models/response"
	"myblog_server/service"
	"myblog_server/utils/jwt"
)

// LogoutView 用户退出登录
func (UserApi) LogoutView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.Claims) // 断言

	token := c.Request.Header.Get("token")

	// 手动注销token
	err := service.ServiceApp.UserService.Logout(claims, token)

	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage("退出失败", c)
		return
	}

	response.OkWithMessage("退出成功", c)

}
