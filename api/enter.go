package api

import (
	"myblog_server/api/blog_api"
	"myblog_server/api/category_api"
	"myblog_server/api/log_api"
	"myblog_server/api/tag_api"
	"myblog_server/api/user_api"
)

type ApiGroup struct {
	UserApi     user_api.UserApi
	LoginApi    log_api.LoginApi
	CategoryApi category_api.CategoryApi
	TagApi      tag_api.TagApi
	BlogApi     blog_api.BlogApi
}

var ApiGroupApp = &ApiGroup{}
