package attachment_api

import (
	"github.com/gin-gonic/gin"
	"myblog_server/models"
	"myblog_server/models/model_type"
	"myblog_server/models/response"
	"myblog_server/service/common"
)

// AttachmentListRequest 根据上传位置查询
type AttachmentListRequest struct {
	models.PageInfo
	Location int `json:"image_type" form:"image_type"`
}

// AttachmentListView 附件列表
func (AttachmentApi) AttachmentListView(c *gin.Context) {
	var cr AttachmentListRequest
	err := c.ShouldBindQuery(&cr)

	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}
	list, count, err := common.ComList(models.Attachment{Location: model_type.LocationType(cr.Location)}, common.Option{
		PageInfo: cr.PageInfo,
		//Debug:    false,
		Likes: []string{"name"}, // 根据名字查找
	})
	// 注意，需要给返回的数据修复一下，目的是只显示文件名称而不带扩展名
	response.OkWithList(list, count, c)
}
