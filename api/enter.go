package api

import (
	"myblog_server/api/about_api"
	"myblog_server/api/blog_api"
	"myblog_server/api/category_api"
	"myblog_server/api/collect_api"
	"myblog_server/api/comment_api"
	"myblog_server/api/friendlink_api"
	"myblog_server/api/info_api"
	"myblog_server/api/log_api"
	"myblog_server/api/music_api"
	"myblog_server/api/tag_api"
	"myblog_server/api/user_api"
)

type ApiGroup struct {
	UserApi       user_api.UserApi
	LoginApi      log_api.LoginApi
	CategoryApi   category_api.CategoryApi
	TagApi        tag_api.TagApi
	BlogApi       blog_api.BlogApi
	InfoApi       info_api.InfoApi
	AboutAPI      about_api.AboutApi
	FriendLinkAPI friendlink_api.FriendLinkApi
	MusicAPI      music_api.MusicApi
	CollectApi    collect_api.CollectApi
	CommentAPI    comment_api.CommentApi
}

var ApiGroupApp = &ApiGroup{}
