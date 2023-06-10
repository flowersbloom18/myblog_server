package routers

import (
	"myblog_server/api"
	"myblog_server/middleware"
)

func (router RouterGroup) User() {
	app := api.ApiGroupApp.UserApi
	// 用户自己注册
	router.POST("register", app.UserRegisterView)
	// 管理员创建 任意权限的用户⚠️需要添加中间件（管理员登录才行）
	router.POST("create_user", middleware.JwtAdmin(), app.UserCreateView)

	// 用户登录(通过邮箱或用户名）
	router.POST("user_login", app.UserLoginView)
	// 用户退出登录
	router.POST("logout", middleware.JwtAuth(), app.LogoutView)
	// 获取批量用户
	router.GET("users", middleware.JwtAuth(), app.UserListView)
	// 删除批量用户
	router.DELETE("users", middleware.JwtAdmin(), app.UserRemoveView)

	// 修改当前登录用户的密码
	router.PUT("update_password", middleware.JwtAuth(), app.UserUpdatePassword)
	// 修改当前登录用户的昵称和头像
	router.PUT("update_nick_name", middleware.JwtAuth(), app.UserUpdateNickName)
	// 管理员修改指定用户的昵称和权限
	router.PUT("update_role", middleware.JwtAdmin(), app.UserUpdateRoleView)

	// 当前登录的用户绑定邮箱
	router.PUT("user_bind_email", middleware.JwtAuth(), app.UserBindEmailView)
	// 用户忘记密码⚠️

	// 当前登录用户的个人信息
	router.GET("user_info", middleware.JwtAuth(), app.UserInfoView)
	// 用户忘记密码，通过邮箱+验证码即可
	router.PUT("user_forget_password", app.UserForgetPasswordView)

}
