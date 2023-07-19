package collect_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/response"
	"myblog_server/utils/jwt"
)

func (CollectApi) CollectRemoveView(c *gin.Context) {
	// 1、批量取消收藏，暂时不用该功能。
	db := global.DB
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.Claims)
	userID := claims.UserID

	// 绑定请求参数
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}
	global.Log.Info("id_list=", cr.IDList)

	// 获取请求--收藏ID列表
	var list []models.Collect

	// 当前用户，且是这篇博客
	for _, v := range cr.IDList {
		var collect models.Collect
		count := db.Where("blog_id=?", v).Where("user_id=?", userID).Take(&collect).RowsAffected
		if count == 0 { // 如果数据为0，则查不到，跳到下一条数据
			continue
		}
		list = append(list, collect)
	}

	count := len(list)
	if count == 0 { // 数据不存在
		response.FailWithMessage("数据不存在", c)
		return
	}

	// 删除
	global.DB.Delete(&list)
	response.OkWithMessage(fmt.Sprintf("共删除 %d 条数据", count), c)
}

// 单机取消，批量？？？？？
