package service

import (
	"myblog_server/service/user_service"
)

type ServiceGroup struct {
	UserService user_service.UserService
}

var ServiceApp = ServiceGroup{}
