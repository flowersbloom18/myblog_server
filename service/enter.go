package service

import (
	"myblog_server/service/blog_service"
	"myblog_server/service/category_service"
	"myblog_server/service/tag_service"
	"myblog_server/service/user_service"
)

type ServiceGroup struct {
	UserService     user_service.UserService
	CategoryService category_service.CategoryService
	TagService      tag_service.TagService
	BlogService     blog_service.BlogService
}

var ServiceApp = ServiceGroup{}
