package main

import (
	"fmt"
	"myblog_server/core"
	"myblog_server/global"
	"myblog_server/utils/jwt"
)

func main() {
	core.InitConf()
	global.Log = core.InitLogger()

	token, err := jwt.GenToken(jwt.PayLoad{
		UserID:   1,
		Role:     1,
		NickName: "张三",
	})
	fmt.Println(token, err)

	claims, err := jwt.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuaWNrX25hbWUiOiJ4eHgiLCJyb2xlIjoxLCJ1c2VyX2lkIjoxLCJhdmF0YXIiOiIiLCJleHAiOjE2ODU1MDYxMTUuMzU3NzYxLCJpc3MiOiJmbG93ZXJzYmxvb20ifQ.Wmo1dDO5_ph9Li5O9dEX2h1aYdS7V0zYvz0Yfccqef0")
	fmt.Println(claims, err)

}
