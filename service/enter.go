package service

import (
	"myblog_server/service/blog_service"
	"myblog_server/service/category_service"
	"myblog_server/service/comment_service"
	"myblog_server/service/friendlink_service"
	"myblog_server/service/info_service"
	"myblog_server/service/music_service"
	"myblog_server/service/tag_service"
	"myblog_server/service/user_service"
)

type ServiceGroup struct {
	UserService       user_service.UserService
	CategoryService   category_service.CategoryService
	TagService        tag_service.TagService
	BlogService       blog_service.BlogService
	InfoService       info_service.InfoService
	FriendLinkService friendlink_service.FriendLinkService
	MusicService      music_service.MusicService
	CommentService    comment_service.CommentService
}

var ServiceApp = ServiceGroup{}
