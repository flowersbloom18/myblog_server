package blog_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models/response"
	"myblog_server/service"
	"myblog_server/utils/jwt"
)

type BlogCreateRequest struct {
	Title      string   ` json:"title" binding:"required" msg:"请输入标题"`         // 标题
	Content    string   `json:"content" binding:"required" msg:"请输入内容"`        // 内容
	Cover      string   `json:"cover" `                                        // 封面
	IsComment  bool     `json:"is_comment" binding:"required" msg:"请确认是否开启评论"` // 是否开启评论
	IsPublish  bool     `json:"is_publish" binding:"required" msg:"请确认是否发布"`   // 是否发布
	IsTop      bool     `json:"is_top" binding:"required" msg:"请确认是否置顶"`       // 是否置顶
	CategoryID uint     `json:"category_id" binding:"required" msg:"请输入分类id"`  // 分类ID
	Tags       []string `json:"tags"`                                          // 标签列表
}

/*
	title ,content,cover string,
	is_comment,is_publish,is_top bool,
	category_id,user_id uint,link string,

*/

func (BlogApi) BlogCreateView(c *gin.Context) {
	_claims, _ := c.Get("claims") // 当前登录用户解析后的信息
	claims := _claims.(*jwt.Claims)

	serviceApp := service.ServiceApp

	var cr BlogCreateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		response.FailWithError(err, &cr, c)
		return
	}

	// 创建的service
	err := serviceApp.BlogService.CreateBlog(
		cr.Title, cr.Content, cr.Cover, cr.IsComment, cr.IsPublish, cr.IsTop,
		cr.CategoryID, claims.UserID, cr.Tags)
	if err != nil {
		global.Log.Error("创建博客失败: ", err.Error())
		response.FailWithMessage(fmt.Sprintf("博客创建失败:%s", err), c)
		return
	}
	// 响应成功
	response.OkWithMessage("博客创建成功!", c)
	return
}
