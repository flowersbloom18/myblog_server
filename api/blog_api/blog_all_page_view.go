package blog_api

import (
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/response"
)

// BlogAllPageView 博客总浏览量查询
func (BlogApi) BlogAllPageView(c *gin.Context) {

	var blogs []models.Blog
	err := global.DB.Find(&blogs).Error
	if err != nil {
		global.Log.Warn("查询出错，err=", err)
		response.FailWithMessage("博客总浏览量查询出错了", c)
		return
	}

	// 博客总浏览量
	var readNums = 0
	for _, v := range blogs {
		readNums += v.ReadNum
	}
	//fmt.Println("博客总浏览量为", readNums)

	response.OkWithData(readNums, c)
	return
}
