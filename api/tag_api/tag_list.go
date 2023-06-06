package tag_api

import (
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/response"
	"myblog_server/service/common"
)

type TagListResponse struct {
	Tag     models.Tag `json:"tag"`
	BlogNum uint       `json:"blog_num"` // 响应数据时候，转为json数据的key名称
}

func (TagApi) TagListView(c *gin.Context) {
	var cr models.PageInfo // contentRequest内容请求 -->cr
	err := c.ShouldBindQuery(&cr)

	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}

	list, count, err := common.ComList(models.Tag{}, common.Option{
		PageInfo: cr,
		Debug:    true,
		Likes:    []string{"name"},  // 按照标签名查询
		Preload:  []string{"Blogs"}, // 搜索标签的时候预加载博客
	})
	if err != nil {
		global.Log.Warn("获取数据错误：", err)
	}

	//response.OkWithList(list, count, c)

	// 对响应数据进一步封装，获取标签下博客的个数。
	var result []TagListResponse
	for _, value := range list {
		//if len(value.Blogs) == 0 {
		//	blog.Tags = []Tag{}
		//}
		result = append(result, TagListResponse{
			Tag:     value,
			BlogNum: uint(len(value.Blogs)),
		})
	}
	response.OkWithList(result, count, c)
}
