package attachment_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/response"
)

// AttachmentRemoveView  默认传入上下文，
func (AttachmentApi) AttachmentRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr) //将请求的id列表（json数据）绑定到结构体
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}

	var lists []models.Attachment

	//  `Find` 方法，该方法的第一个参数是一个指向 `lists` 的指针，第二个参数是 `cr.IDList`，表示要查找的记录的 ID 列表。
	// `Find` 方法会在数据库中查找符合条件的记录，并将结果存储在 `lists` 中。
	count := global.DB.Find(&lists, cr.IDList).RowsAffected
	if count == 0 {
		response.FailWithMessage("附件不存在", c)
		return
	}
	global.DB.Delete(&lists)
	response.OkWithMessage(fmt.Sprintf("共删除 %d 个附件", count), c)

}
