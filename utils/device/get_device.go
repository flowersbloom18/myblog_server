package device

import (
	"github.com/gin-gonic/gin"
	"strings"
)

// GetLoginDevice 获取登录设备
func GetLoginDevice(c *gin.Context) string {
	userAgent := c.Request.UserAgent()
	if strings.Contains(userAgent, "Android") {
		return "Android"
	} else if strings.Contains(userAgent, "iPhone") {
		return "iPhone"
	} else if strings.Contains(userAgent, "iPad") {
		return "iPad"
	} else if strings.Contains(userAgent, "Macintosh") {
		return "Mac"
	} else if strings.Contains(userAgent, "Windows") {
		return "Windows"
	} else if strings.Contains(userAgent, "Linux") {
		return "Linux"
	} else {
		return "未知设备"
	}
}
