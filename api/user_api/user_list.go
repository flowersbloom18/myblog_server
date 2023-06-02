package user_api

import (
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/response"
	"myblog_server/service/common"
)

func (UserApi) UserListView(c *gin.Context) {
	var cr models.PageInfo // contentRequest内容请求 -->cr
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}

	list, count, err := common.ComList(models.UserModel{}, common.Option{
		PageInfo: cr,
		Debug:    true,
	})
	if err != nil {
		global.Log.Warn("获取数据错误：", err)
	}

	response.OkWithList(list, count, c)
}
