package user_api

import (
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/response"
	"myblog_server/utils/jwt"
)

func (UserApi) UserInfoView(c *gin.Context) {

	_claims, _ := c.Get("claims") // 断言，然后调用
	claims := _claims.(*jwt.Claims)

	var userInfo models.User
	err := global.DB.Take(&userInfo, claims.UserID).Error
	if err != nil {
		response.FailWithMessage("用户不存在", c)
		return
	}
	//filter.Select("info", userInfo)的意义是
	//使用json-filter库中的Select方法，从userInfo这个结构体中选择出名为"info"的字段，并返回该字段的值。
	response.OkWithData(filter.Select("info", userInfo), c)
}
