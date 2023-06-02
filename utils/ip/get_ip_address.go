package ip

import "github.com/gin-gonic/gin"

// GetAddrByGin 获取ip和由ip转换后的地址
func GetAddrByGin(c *gin.Context) (ip, addr string) {
	ip = c.ClientIP()
	addr = GetAddressByIp(ip)
	return ip, addr
}
