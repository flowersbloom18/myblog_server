package attachment_api

import (
	"github.com/gin-gonic/gin"
	"myblog_server/models"
	"myblog_server/models/response"
	"myblog_server/service/common"
)

// AttachmentListView 附件列表
func (AttachmentApi) AttachmentListView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}
	list, count, err := common.ComList(models.Attachment{}, common.Option{
		PageInfo: cr,
		Debug:    false,
		Likes:    []string{"name"}, // 根据名字查找
	})
	// 注意，需要给返回的数据修复一下，目的是只显示文件名称而不带扩展名
	response.OkWithList(list, count, c)
}
