package comment_service

import (
	"fmt"
	"myblog_server/global"
	"myblog_server/models"
)

func (CommentService) RemoveCommentService(role int, userID uint, list []models.Comment) error {
	// 如果为管理员，则拥有全部删除的权限
	if role == 1 {
		for _, v := range list {
			// 执行删除的逻辑
			err := removeComment(v)
			if err != nil {
				return err
			}
		}
	} else { // 否则，只能删除当前登录用户自己的评论
		for _, v := range list {
			if userID == v.UserID {
				// 执行删除的逻辑
				err := removeComment(v)
				if err != nil {
					return err
				}
			} else {
				return fmt.Errorf("您无法删除他人评论")
			}
		}
	}
	return nil
}
func removeComment(v models.Comment) error {
	if v.PanelID == 0 { // 最上层面板（顶级评论）。需要删除面板下所有评论
		err := global.DB.Where("panel_id=?", v.ID).Delete(&models.Comment{}).Error
		if err != nil {
			global.Log.Warn("评论删除失败,err:", err)
			return fmt.Errorf("评论删除失败")
		}
		// 最后删除该面板
		err = global.DB.Delete(v).Error // 仅仅删除这一条
		if err != nil {
			global.Log.Warn("评论删除失败,err:", err)
			return fmt.Errorf("评论删除失败")
		}
	} else {
		err := global.DB.Delete(v).Error // 仅仅删除这一条
		if err != nil {
			global.Log.Warn("评论删除失败,err:", err)
			return fmt.Errorf("评论删除失败")
		}
	}
	return nil // success
}
