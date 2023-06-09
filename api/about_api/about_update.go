package about_api

import (
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/response"
)

type AboutRequest struct {
	Content string `json:"content"`
}

func (AboutApi) UpdateAboutView(c *gin.Context) {
	var cr AboutRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithError(err, &cr, c) // 错误，参数，上下文
		return
	}

	var about models.About
	err = global.DB.First(&about).Error // 获取第一条数据
	if err != nil {
		global.Log.Info("更新失败")
		response.FailWithMessage("更新失败", c)
		return
	}

	err = global.DB.Model(about).Update("content", cr.Content).Error
	if err != nil {
		response.FailWithMessage("更新失败", c)
		return
	}

	response.OkWithMessage("更新成功", c)
	return
}
