package middleware

import (
	"github.com/gin-gonic/gin"
	"myblog_server/models/model_type"
	"myblog_server/models/response"
	"myblog_server/service/redis_service"
	"myblog_server/utils/jwt"
)

// JwtAuth 登录之后才有权限访问
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			response.FailWithMessage("未携带token", c)
			c.Abort()
			return
		}
		// 解析jwt判断是否过期等信息
		claims, err := jwt.ParseToken(token)

		if err != nil {
			response.FailWithMessage("token错误", c)
			c.Abort()
			return
		}
		// 判断是否在redis中
		if redis_service.CheckLogout(token) {
			response.FailWithMessage("token已失效", c)
			c.Abort()
			return
		}
		// 登录的用户
		c.Set("claims", claims)
	}
}

// 在本项目中，用户和游客的api权限是 一样的，
// 但是看到的不同，前者是自己的界面，后者是管理员的界面（仅仅是关键数据无法操作和获取）

// JwtAdmin 关键操作需要管理员登录的权限
func JwtAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			response.FailWithMessage("未携带token", c)
			c.Abort()
			return
		}
		claims, err := jwt.ParseToken(token)
		if err != nil {
			response.FailWithMessage("token错误", c)
			c.Abort()
			return
		}
		// 判断是否在redis中
		if redis_service.CheckLogout(token) {
			response.FailWithMessage("token已失效", c)
			c.Abort()
			return
		}
		// 登录的用户
		if claims.Role != int(model_type.PermissionAdmin) {
			response.FailWithMessage("权限错误", c)
			c.Abort()
			return
		}

		c.Set("claims", claims)
	}
}
