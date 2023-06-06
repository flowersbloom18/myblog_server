package category_api

import (
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/response"
	"myblog_server/service/common"
)

type CategoryListResponse struct {
	Category models.Category `json:"category"`
	BlogNum  uint            `json:"blog_num"` // 响应数据时候，转为json数据的key名称
}

func (CategoryApi) CategoryListView(c *gin.Context) {
	var cr models.PageInfo // contentRequest内容请求 -->cr
	err := c.ShouldBindQuery(&cr)

	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}

	list, count, err := common.ComList(models.Category{}, common.Option{
		PageInfo: cr,
		Debug:    true,
		Likes:    []string{"name"}, // 按照分类名查询
		Preload:  []string{"Blogs"},
	})
	if err != nil {
		global.Log.Warn("获取数据错误：", err)
	}

	// 对响应数据进一步封装，获取分类下博客的个数。
	var result []CategoryListResponse
	for _, value := range list {
		result = append(result, CategoryListResponse{
			Category: value,
			BlogNum:  uint(len(value.Blogs)),
		})
	}

	response.OkWithList(result, count, c)
}
