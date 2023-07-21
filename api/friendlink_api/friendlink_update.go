package friendlink_api

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/response"
	"time"
)

type FriendLinkUpdateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Logo        string `json:"logo"`
	Url         string `json:"url"`
	IsTop       bool   `json:"is_top" `

	// 不对参数作用
	TopTime time.Time `json:"top_time"` // 这里只是作为一个属性方便使用
}

func (FriendLinkApi) FriendLinkUpdateView(c *gin.Context) {
	// 获取需要更新的友链ID
	id := c.Param("id")

	db := global.DB

	var cr FriendLinkUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithError(err, &cr, c)
		return
	}

	var friendlink models.FriendLink
	// 获取对应id的数据
	err = db.Take(&friendlink, "id = ?", id).Error
	if err != nil {
		global.Log.Warn("友链不存在")
		response.FailWithMessage("友链不存在", c)
		return
	}

	// 检查新创建的友链名称是否存在，不存在则更新【友链重复值判断】
	var friendlink1 models.FriendLink
	// ⚠️如果说新增友链跟当前数据库【不同】则进行重复值判断，否则不用。
	if friendlink.Name != cr.Name {
		err = db.First(&friendlink1, "name = ?", cr.Name).Error
		// 查询到数据则err为空，说明友链已经存在。反之不存在，即可更新数据。
		if err == nil {
			global.Log.Warn("友链已存在", err)
			response.FailWithMessage("友链已存在", c)
			return
		}
	}

	// 判断置顶状态是否改变，如果改变，且为true则修改。
	// 置顶排序的问题
	if cr.IsTop != friendlink1.IsTop {
		if cr.IsTop == true {
			global.Log.Warn("置顶啦！")
			cr.TopTime = time.Now()
		} else {
			// 取消置顶，时间归为0值
			cr.TopTime = time.Time{}
		}
	}

	// 更新
	maps := structs.Map(&cr)
	// 如果开启置顶
	if cr.IsTop {
		maps["top_time"] = time.Now()
	}
	err = global.DB.Model(&friendlink).Updates(maps).Error
	if err != nil {
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
	return
}
