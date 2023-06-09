package info_api

import (
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/response"
	"myblog_server/service"
	"myblog_server/utils/info"
	"strconv"
)

type CategoryListResponse struct {
	Category models.Category `json:"category"`
	BlogNum  uint            `json:"blog_num"` // 响应数据时候，转为json数据的key名称
}

func (InfoApi) InfoListView(c *gin.Context) {
	infoService := service.ServiceApp.InfoService
	tianApi := global.Config.TianApi
	url := ""

	// 获取请求的数据参数,并转为整数类型
	_id := c.Param("id")
	id, err := strconv.Atoi(_id)
	if err != nil {
		global.Log.Error("error=", err)
		response.FailWithMessage("请求参数有误。", c)
		return
	}

	if id == 1 {
		url = tianApi.DouYinHot // 抖音热搜
	} else if id == 2 {
		url = tianApi.NetWorkHot // 全网热搜
	} else if id == 3 {
		url = tianApi.WeiBoHot // 微博热搜
	} else if id == 4 {
		url = tianApi.BulletIn // 每日简报
	} else if id == 5 {
		url = tianApi.ZaoAn // 早安
	} else if id == 6 {
		url = tianApi.WanAn // 晚安
	} else if id == 7 {
		today := info.GetMonthDay()
		url = tianApi.LiShi + today // 历史的今天
	} else {
		global.Log.Error("请求数据不存在")
		response.FailWithMessage("请求数据不存在。", c)
		return
	}

	// 查找数据
	content, err := infoService.GetInfoService(url, id)
	if err != nil {
		global.Log.Error("error=", err)
		return
	}

	// 解析
	result, err := info.GetInfoResult(content, id)
	if err != nil {
		global.Log.Error("error=", err)
		return
	}

	// 定时更新数据

	response.OkWithData(result, c)
}
