package routers

import (
	"myblog_server/api"
	"myblog_server/middleware"
)

func (router RouterGroup) Comment() {
	api := api.ApiGroupApp.CommentAPI

	// 创建评论
	router.POST("comment", middleware.JwtAuth(), api.CreateCommentView)

	// 查看所有评论（管理员）
	router.GET("comments", api.CommentListView)
	// 查看博客下的评论
	router.GET("comments_blog", api.CommentBlogListView)
	// 查看友链下的评论
	router.GET("comments_friendlink", api.CommentFriendLinkListView)
	// 查看关于下的评论
	router.GET("comments_about", api.CommentAboutListView)
	// 查看用户的评论
	router.GET("comments_user", middleware.JwtAuth(), api.CommentUserListView)

	// 一键开启/关闭评论【首先获取开关的状态，然后修改】
	router.GET("comment_open", api.CommentStatusView)
	router.PUT("comment_open", api.CommentOpenView)

	// 删除评论
	router.DELETE("comments", middleware.JwtAuth(), api.CommentRemoveView)
}
