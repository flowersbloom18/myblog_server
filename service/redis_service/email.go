package redis_service

import (
	"myblog_server/global"
	"myblog_server/utils"
	"time"
)

const prefixEmail = "authCode_"

// SetAuthCode 设置邮箱验证码（token=email+token）
func SetAuthCode(token string, value string, diff time.Duration) error {
	err := global.Redis.Set(prefixEmail+token, value, diff).Err()
	return err
}

// CheckAuthCode 检查验证码
func CheckAuthCode(token string) (string, error) {
	// 获取所有的key，然后从中筛选满足条件的key
	keys := global.Redis.Keys(prefixEmail + "*").Val()
	if utils.InList(prefixEmail+token, keys) {
		global.Log.Info("prefixEmail+token= ", prefixEmail+token)
		// 获取指定 key 对应的value
		value, err := global.Redis.Get(prefixEmail + token).Result()
		if err != nil {
			return "", err
		}
		return value, nil
	} else {
		// 找不到的话，返回空
		return "", nil
	}
}
