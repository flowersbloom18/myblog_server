package attachment_api

import (
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/response"
)

type AttachmentUpdateRequest struct {
	ID   uint   `json:"id" binding:"required" msg:"请选择附件id"`
	Name string `json:"name" binding:"required" msg:"请输入附件名称"`
}

func (AttachmentApi) AttachmentUpdateView(c *gin.Context) {
	var cr AttachmentUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithError(err, &cr, c)
		return
	}

	var attachment models.Attachment
	err = global.DB.Take(&attachment, cr.ID).Error
	if err != nil {
		response.FailWithMessage("附件不存在", c)
		return
	}
	err = global.DB.Model(&attachment).Update("name", cr.Name).Error
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("附件名称修改成功", c)
	return
}
