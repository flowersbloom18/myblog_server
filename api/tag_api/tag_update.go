package tag_api

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/response"
)

type TagUpdateRequest struct {
	Name  string `json:"name"`
	Cover string `json:"cover"`
}

// TagUpdateView   更新标签（名称和封面）
func (TagApi) TagUpdateView(c *gin.Context) {
	db := global.DB
	id := c.Param("id")

	var cr TagUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithError(err, &cr, c)
		return
	}

	var tag models.Tag
	// 获取对应id的数据
	err = db.Take(&tag, "id = ?", id).Error
	if err != nil {
		global.Log.Warn("标签不存在")
		response.FailWithMessage("标签不存在", c)
		return
	}

	// 检查新创建的标签名称是否存在，不存在则更新
	var tag1 models.Tag
	// 标签名称不同则需要检测
	if tag.Name != cr.Name {
		err = db.First(&tag1, "name = ?", cr.Name).Error
		// 查询到数据则err为空，说明标签已经存在。反之不存在，即可更新数据。
		if err == nil {
			global.Log.Warn("标签已存在", err)
			response.FailWithMessage("标签已存在", c)
			return
		}
	}

	// 更新
	maps := structs.Map(&cr) // 结构体转map
	err = global.DB.Model(&tag).Updates(maps).Error

	//err = global.DB.Model(&tag).Updates(map[string]interface{}{
	//	"name":  cr.Name,
	//	"cover": cr.Cover,
	//}).Error

	if err != nil {
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
	return
}
