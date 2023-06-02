package api

import (
	"myblog_server/api/log_api"
	"myblog_server/api/user_api"
)

type ApiGroup struct {
	UserApi  user_api.UserApi
	LoginApi log_api.LoginApi
}

var ApiGroupApp = &ApiGroup{}
