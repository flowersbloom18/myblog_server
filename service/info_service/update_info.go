package info_service

import (
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/utils/info"
)

// UpdateInfoService 更新所有api信息（先更新一个）
func (InfoService) UpdateInfoService() {

	tianApi := global.Config.TianApi
	// 日志标注
	success := true
	content := ""
	// 获取最新的Content，紧接着写入数据库，如果存在则更新，否则创建
	for id := 1; id <= 7; id++ {
		if id == 1 { // 抖音热搜
			url := tianApi.DouYinHot
			if !updateInfo(url, id) {
				success = false
				content += "抖音热搜"
			}
		} else if id == 2 { // 全网热搜
			url := tianApi.NetWorkHot
			if !updateInfo(url, id) {
				success = false
				content += "+全网热搜"
			}
		} else if id == 3 { // 微博热搜
			url := tianApi.WeiBoHot
			if !updateInfo(url, id) {
				success = false
				content += "+微博热搜"
			}
		} else if id == 4 { // 每日简报
			url := tianApi.BulletIn
			if !updateInfo(url, id) {
				success = false
				content += "+每日简报"
			}
		} else if id == 5 { // 早安
			url := tianApi.ZaoAn
			if !updateInfo(url, id) {
				success = false
				content += "+早安"
			}
		} else if id == 6 { // 晚安
			url := tianApi.WanAn
			if !updateInfo(url, id) {
				success = false
				content += "+晚安"
			}
		} else if id == 7 { // 历史的今天
			today := info.GetMonthDay()
			url := tianApi.LiShi + today
			if !updateInfo(url, id) {
				success = false
				content += "抖音热搜"
			}
		} else {
			global.Log.Error("未知错误")
		}
	}

	// 统一记录日志
	if success {
		global.DB.Create(&models.Log{
			UserName: "系统日志",
			NickName: "系统日志",
			Level:    "info",
			Content:  "信息更新成功",
		})
		global.Log.Info("信息更新成功")
	} else {
		global.DB.Create(&models.Log{
			UserName: "系统日志",
			NickName: "系统日志",
			Level:    "info",
			Content:  "信息更新失败，失败部分为：" + content,
		})
		global.Log.Warn("信息更新失败，失败部分为：" + content)
	}
}

func updateInfo(url string, typeID int) bool {
	db := global.DB

	// 获取响应结果
	responseData, err := info.GetHttpResponse(url)
	// 查询数据，如果有，则原基础更新。否则或许数据创建后，在更新。
	var info models.Info
	count := db.Take(&info, "type_id=?", typeID).RowsAffected

	if count != 0 {
		err = db.Model(&info).Update("content", responseData).Error
		if err != nil {
			global.Log.Error("Info更新出错，err=", err)
			return false
		}
		return true
	} else if count == 0 {
		// 创建数据
		err = db.Create(&models.Info{
			Content: responseData,
			TypeId:  typeID,
		}).Error
		if err != nil {
			global.Log.Error("Info创建出错，err=", err)
			return false
		}
	}
	return true
}
