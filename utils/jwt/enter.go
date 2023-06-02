package jwt

import (
	"github.com/dgrijalva/jwt-go/v4"
)

// PayLoad jwt中payload数据
type PayLoad struct {
	NickName string `json:"nick_name"` // 昵称
	Role     int    `json:"role"`      // 权限  1 管理员  2 普通用户  3 游客
	UserID   uint   `json:"user_id"`   // 用户id
	Avatar   string `json:"avatar"`
}

var MySecret []byte

type Claims struct {
	PayLoad
	jwt.StandardClaims
}

/*
JWT包含三部分
	1. Header（头部）：包含了 JWT 的类型和使用的算法，通常是固定的，使用 Base64 编码。

	2. Payload（负载）：包含了 JWT 的具体内容，比如用户的 ID、角色、权限等信息，也可以自定义其他信息，使用 Base64 编码。

	3. Signature（签名）：由 Header 和 Payload 通过指定的算法生成的签名，用于验证 JWT 的合法性。
*/
