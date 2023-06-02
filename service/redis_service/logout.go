package redis_service

import (
	"myblog_server/global"
	"myblog_server/utils"
	"time"
)

const prefixLogout = "logout_"

// Logout 针对退出登录的操作
func Logout(token string, diff time.Duration) error {
	err := global.Redis.Set(prefixLogout+token, "", diff).Err()
	return err
}

func CheckLogout(token string) bool {
	keys := global.Redis.Keys(prefixLogout + "*").Val()
	if utils.InList(prefixLogout+token, keys) {
		return true
	}
	return false
}
