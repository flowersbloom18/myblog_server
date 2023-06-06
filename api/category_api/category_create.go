package category_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models/response"
	"myblog_server/service"
)

type CategoryCreateRequest struct {
	Name  string `json:"name" binding:"required" msg:"请输入分类名称"`
	Cover string `json:"cover"`
}

func (CategoryApi) CategoryCreateView(c *gin.Context) {
	serviceApp := service.ServiceApp

	var cr CategoryCreateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		response.FailWithError(err, &cr, c)
		return
	}

	// 创建的service
	err := serviceApp.CategoryService.CreateCategory(cr.Name, cr.Cover)
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 响应成功
	response.OkWithMessage(fmt.Sprintf("分类'%s'创建成功!", cr.Name), c)
	return
}
