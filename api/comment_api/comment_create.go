package comment_api

import (
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/model_type"
	"myblog_server/models/response"
	"myblog_server/utils/jwt"
)

type CommentRequest struct {
	Content  string              `json:"content" binding:"required" msg:"请输入内容"`     // 评论内容
	PageType model_type.PageType `json:"page_type" binding:"required" msg:"请输入页面类型"` // 评论页面的类型
	Page     string              `json:"page" binding:"required" msg:"请输入页面路径"`      // 评论页面
	FatherID uint                `json:"father_id"`                                  // 父级ID
	PanelID  uint                `json:"panel_id"`                                   // 面板ID
}

// CreateCommentView 用户发表评论，但是如果全局禁止评论，则无法发表。
func (CommentApi) CreateCommentView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.Claims)

	userID := claims.UserID // 用户id
	isAdmin := false        // 是否为管理员
	if claims.Role == 1 {
		isAdmin = true
	}

	var commentOpen models.CommentOpen
	err := global.DB.Take(&commentOpen).Error
	if err != nil {
		response.FailWithMessage("评论开启出现未知错误", c)
		return
	}

	if !commentOpen.IsOpen {
		response.FailWithMessage("当前时间禁止评论,请稍后再试。", c)
		return
	}

	var cr CommentRequest
	err = c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithError(err, &cr, c)
		return
	}

	err = global.DB.Create(&models.Comment{
		Content:  cr.Content,  // 内容
		UserID:   userID,      // 用户id
		PageType: cr.PageType, // 评论页面类型（博客、友链、关于）
		Page:     cr.Page,     // 评论页面路径
		IsAdmin:  isAdmin,     // 是否为管理员
		FatherID: cr.FatherID, // 父级ID
		PanelID:  cr.PanelID,  // 面板ID
	}).Error
	if err != nil {
		response.FailWithMessage("评论数据创建出错", c)
		return
	}

	response.OkWithMessage("评论成功", c)
}
