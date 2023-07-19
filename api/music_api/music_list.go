package music_api

import (
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/response"
	"myblog_server/service/common"
)

func (MusicApi) MusicListView(c *gin.Context) {
	var cr models.PageInfo // contentRequest内容请求 -->cr
	// 默认按照sort,升序排序
	cr.Sort = "sort asc"
	err := c.ShouldBindQuery(&cr)

	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}

	list, count, err := common.ComList(models.Music{}, common.Option{
		PageInfo: cr,
		//Debug:    true,
		Likes: []string{"name"}, // 按照音乐名查询
	})
	if err != nil {
		global.Log.Warn("获取数据错误：", err)
	}

	response.OkWithList(list, count, c)
}
