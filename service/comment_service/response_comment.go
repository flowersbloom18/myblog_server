package comment_service

import (
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/model_type"
	"time"
)

// CommentListResponse 评论列表响应
type CommentListResponse struct {
	ID        uint                `json:"id"`         // 评论ID
	CreatedAt time.Time           `json:"created_at"` // 评论时间
	Content   string              `json:"content"`    // 评论内容
	PageType  model_type.PageType `json:"page_type"`  // 评论页面的类型
	Page      string              `json:"page"`       // 评论页面
	IsAdmin   bool                `json:"is_admin"`   // 是否为管理员
	FatherID  uint                `json:"father_id"`  // 父级ID
	PanelID   uint                `json:"panel_id"`   // 面板ID

	UserID    uint   `json:"user_id"`    // 评论用户ID
	IPAddress string `json:"ip_address"` // IP属地【评论创建时候产生的，不会发生改变】
	NickName  string `json:"nick_name"`  // 评论昵称
	Avatar    string `json:"avatar"`     // 用户头像

	FatherNickName string `json:"father_nick_name"` // 被回复者的昵称
}

// ResponseCommentService 根据不同类型查询评论
func (CommentService) ResponseCommentService(list []models.Comment) []CommentListResponse {
	// 对评论列表进一步封装
	var result []CommentListResponse
	for _, v := range list {
		// 查询该行评论的用户头像和昵称,如果查询不到，销户处理
		var user models.User

		var fatherNickName string
		// 如果父级id=0则不用查询父级昵称。
		if v.FatherID == 0 {
			fatherNickName = ""
		} else {
			// 查询父级id，如果不存在，则父级昵称为空，否则继续查询用户id，
			var comment1 models.Comment
			// 查询父级的昵称,如果查询不到，销户处理
			var user2 models.User

			err0 := global.DB.Take(&comment1, "id=?", v.FatherID).Error

			if err0 != nil {
				// 父级评论被删除
				fatherNickName = "评论被删除"
			} else {
				err2 := global.DB.Take(&user2, "id=?", comment1.UserID).Error
				if err2 != nil {
					fatherNickName = ""
				} else {
					fatherNickName = user2.NickName
				}
			}
		}

		err := global.DB.Take(&user, "id=?", v.UserID).Error
		if err != nil {
			global.Log.Warn("该用户不存在")
			// 如果父级不存在，则销户处理，否则存在处理。
			// 用户不存在则用户已注销
			result = append(result, CommentListResponse{
				ID:             v.ID,
				CreatedAt:      v.CreatedAt,
				Content:        v.Content,
				PageType:       v.PageType,
				Page:           v.Page,
				IsAdmin:        v.IsAdmin,
				FatherID:       v.FatherID,
				PanelID:        v.PanelID,
				UserID:         v.UserID,
				IPAddress:      v.IPAddress,
				NickName:       "用户已注销",             // 用户昵称
				Avatar:         "/uploads/已注销.jpeg", // 用户头像
				FatherNickName: fatherNickName,      // 父级昵称
			})

			// 跳到下一个开头
			continue
		}

		result = append(result, CommentListResponse{
			ID:             v.ID,
			CreatedAt:      v.CreatedAt,
			Content:        v.Content,
			PageType:       v.PageType,
			Page:           v.Page,
			IsAdmin:        v.IsAdmin,
			FatherID:       v.FatherID,
			PanelID:        v.PanelID,
			UserID:         v.UserID,
			IPAddress:      v.IPAddress,
			NickName:       user.NickName,  // 用户昵称
			Avatar:         user.Avatar,    // 用户头像
			FatherNickName: fatherNickName, // 父级昵称
		})
	}

	return result
}
