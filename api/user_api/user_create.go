package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models/model_type"
	"myblog_server/models/response"
	"myblog_server/service"
	"myblog_server/utils/device"
)

type UserCreateRequest struct {
	NickName string          `json:"nick_name" binding:"required" msg:"请输入昵称"`  // 昵称
	UserName string          `json:"user_name" binding:"required" msg:"请输入用户名"` // 用户名
	Password string          `json:"password" binding:"required" msg:"请输入密码"`   // 密码
	Role     model_type.Role `json:"role" binding:"required" msg:"请选择权限"`       // 权限  1 管理员  2 普通用户  3 游客
}

// UserCreateView 管理员创建用户
func (UserApi) UserCreateView(c *gin.Context) {
	serviceApp := service.ServiceApp
	var cr UserCreateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		response.FailWithError(err, &cr, c)
		return
	}

	device := device.GetLoginDevice(c)

	err := serviceApp.UserService.CreateUser(cr.UserName, cr.NickName, cr.Password, cr.Role, "", c.ClientIP(), device)
	if err != nil {
		global.Log.Error(err)
		//response.FailWithMessage(err.Error(), c)
		response.FailWithMessage("该用户名已注册！", c)
		return
	}
	response.OkWithMessage(fmt.Sprintf("用户%s创建成功!", cr.UserName), c)
	return
}
