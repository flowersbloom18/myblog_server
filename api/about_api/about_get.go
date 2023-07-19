package about_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/response"
)

func (AboutApi) GetAboutView(c *gin.Context) {
	db := global.DB
	var about models.About

	// 查询关于表的第一行数据
	result := db.First(&about)

	if result.Error != nil {
		// 处理查询错误
		global.Log.Info(fmt.Sprintf("查询关于信息失败:%s", result.Error))
	}

	// 判断查询结果是否为空
	if result.RowsAffected == 0 {
		// 创建关于表数据
		result := db.Create(&about)
		if result.Error != nil {
			// 处理创建错误
			fmt.Println("创建关于信息失败:", result.Error)
			return
		}
		about.Content = "欢迎大家来常来我的网站看呀，嘻嘻。\n系统自动生成的默认内容"
	}
	response.OkWithData(about, c)
}
