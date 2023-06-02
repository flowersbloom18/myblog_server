package user_service

import (
	"myblog_server/global"
	"myblog_server/service/redis_service"
	"myblog_server/utils/jwt"
	"time"
)

func (UserService) Logout(claims *jwt.Claims, token string) error {
	exp := claims.ExpiresAt
	now := time.Now()
	diff := exp.Time.Sub(now)
	global.Log.Info("过期时间为：", diff)
	return redis_service.Logout(token, diff)
}
