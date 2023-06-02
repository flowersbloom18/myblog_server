package jwt

import (
	"github.com/dgrijalva/jwt-go/v4"
	"myblog_server/global"
	"time"
)

// GenToken 创建 Token
func GenToken(user PayLoad) (string, error) {
	MySecret = []byte(global.Config.Jwt.Secret)

	// Payload（负载）！声明：包含用户信息、计算后的过期时间、颁发人
	claim := Claims{
		user,
		jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Hour * time.Duration(global.Config.Jwt.Expires))), // 默认2小时过期
			Issuer:    global.Config.Jwt.Issuer,                                                     // 签发人
		},
	}

	// Header（头部）
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	// Signature（签名）使用密钥对JWT进行签名
	return token.SignedString(MySecret)
}
