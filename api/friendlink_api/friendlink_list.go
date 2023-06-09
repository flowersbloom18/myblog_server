package friendlink_api

import (
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/response"
	"myblog_server/service/common"
	"time"
)

type FriendLinkListResponse struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Logo        string    `json:"logo"`
	Url         string    `json:"url"`
	IsTop       bool      `json:"is_top" `
	TopTime     time.Time `json:"top_time"`
}

func (FriendLinkApi) FriendLinkListView(c *gin.Context) {
	var cr models.PageInfo // contentRequest内容请求 -->cr
	err := c.ShouldBindQuery(&cr)

	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}

	list, count, err := common.ComList(models.FriendLink{}, common.Option{
		PageInfo: cr,
		Debug:    true,
		Likes:    []string{"name"}, // 按照友链名查询
	})
	if err != nil {
		global.Log.Warn("获取数据错误：", err)
	}

	response.OkWithList(list, count, c)
}
