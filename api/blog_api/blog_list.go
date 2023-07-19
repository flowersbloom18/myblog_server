package blog_api

import (
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/response"
	"myblog_server/service"
	"myblog_server/service/common"
)

func (BlogApi) BlogListView(c *gin.Context) {
	blogService := service.ServiceApp.BlogService

	var cr models.PageInfo // contentRequest内容请求 -->cr
	err := c.ShouldBindQuery(&cr)

	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}

	list, count, err := common.ComList(models.Blog{}, common.Option{
		PageInfo: cr,
		//Debug:    true,
		Likes:   []string{"title"}, // 按照博客名查询,如何添加按照其它方式查询。
		Preload: []string{"Tags"},  // 预加载，体现博客跟标签的多对多关系// 大写⚠️
	})
	if err != nil {
		global.Log.Warn("获取数据错误：", err)
		return
	}

	//response.OkWithList(list, count, c)

	// 对结果再次封装
	blogResponse, err := blogService.GetBlogList(list)
	if err != nil {
		response.FailWithMessage("获取数据错误", c)
		return
	}

	response.OkWithList(blogResponse, count, c)
	return
}
