package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/response"
	"myblog_server/utils"
	"myblog_server/utils/jwt"
)

func (UserApi) UserRemoveView(c *gin.Context) {
	// 获取当前登录的用户的id，如果需要删除的id中包含自己，则禁止删除。
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.Claims) // 断言

	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}

	// 当前用户无法删除自己！
	// 判断删除列表中是否存在当前登录用户，如果存在则禁止删除
	if utils.IsExist(claims.UserID, cr.IDList) {
		response.FailWithMessage("当前登录用户无法被删除！", c)
		return
	} else { // 否则可以删除
		var list []models.User

		count := global.DB.Find(&list, cr.IDList).RowsAffected
		if count == 0 {
			response.FailWithMessage("用户信息不存在", c)
			return
		}
		global.DB.Delete(&list)
		response.OkWithMessage(fmt.Sprintf("共删除 %d 个用户", count), c)

		// ⚠️系统日志记录
		var user models.User
		err = global.DB.Take(&user, claims.UserID).Error
		if err != nil {
			global.Log.Warn("用户不存在", err)
		}
		logContent := fmt.Sprintf("用户删除，删除ID列表：%v", cr.IDList)
		global.DB.Create(&models.Log{
			UserName: user.UserName,
			NickName: user.NickName,
			IP:       user.IP,
			Address:  user.Address,
			Device:   user.Device,
			Level:    "info",
			Content:  logContent,
		})
	}
}
