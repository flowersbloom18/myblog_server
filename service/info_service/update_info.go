package info_service

import (
	"encoding/json"
	"fmt"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/utils/info"
	"strings"
)

// UpdateInfoService 更新所有api信息（先更新一个）
func (InfoService) UpdateInfoService() {

	tianApi := global.Config.TianApi
	key := tianApi.Key
	url := ""
	// 日志标注
	success := true
	content := ""
	// 获取最新的Content，紧接着写入数据库，如果存在则更新，否则创建
	for id := 1; id <= 7; id++ {
		if id == 1 { // 抖音热搜
			url = strings.Replace(tianApi.DouYinHot, "你的APIKEY", key, -1)
			ok, result := updateInfo(url, id)
			if !ok {
				success = false
				content += "抖音热搜" + "-->" + result
			}
		} else if id == 2 { // 全网热搜
			url = strings.Replace(tianApi.NetWorkHot, "你的APIKEY", key, -1)
			ok, result := updateInfo(url, id)
			if !ok {
				success = false
				content += "+全网热搜" + "-->" + result
			}
		} else if id == 3 { // 微博热搜
			url = strings.Replace(tianApi.WeiBoHot, "你的APIKEY", key, -1)
			ok, result := updateInfo(url, id)
			if !ok {
				success = false
				content += "+微博热搜" + "-->" + result
			}
		} else if id == 4 { // 每日简报
			url = strings.Replace(tianApi.BulletIn, "你的APIKEY", key, -1)
			ok, result := updateInfo(url, id)
			if !ok {
				success = false
				content += "+每日简报" + "-->" + result
			}
		} else if id == 5 { // 早安
			url = strings.Replace(tianApi.ZaoAn, "你的APIKEY", key, -1)
			ok, result := updateInfo(url, id)
			if !ok {
				success = false
				content += "+早安" + "-->" + result
			}
		} else if id == 6 { // 晚安
			url = strings.Replace(tianApi.WanAn, "你的APIKEY", key, -1)
			ok, result := updateInfo(url, id)
			if !ok {
				success = false
				content += "+晚安" + "-->" + result
			}
		} else if id == 7 { // 历史的今天
			today := info.GetMonthDay()
			//?key=key&date=today
			url = strings.Replace(tianApi.LiShi, "你的APIKEY", key, -1)
			url = strings.Replace(url, "0101", today, -1)
			ok, result := updateInfo(url, id)
			if !ok {
				success = false
				content += "历史的今天" + "-->" + result
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

func updateInfo(url string, typeID int) (ok bool, result string) {
	db := global.DB

	// 获取响应结果
	responseData, err := info.GetHttpResponse(url)

	// 分析响应结果是否为 {"code":200,"msg":"success"}
	var data map[string]interface{}
	err = json.Unmarshal([]byte(responseData), &data)
	if err != nil {
		global.Log.Warn("解析 JSON 出错:", err)
		return false, fmt.Sprintf("解析 JSON 出错:%s", err)
	}

	// 需要类型断言
	code, ok2 := data["code"].(float64)
	if !ok2 {
		fmt.Println("code 字段类型错误或不存在")
		return
	}
	msg, ok2 := data["msg"].(string)
	if !ok2 {
		fmt.Println("msg 字段类型错误或不存在")
		return
	}

	if code != 200 { //
		return false, fmt.Sprintf("{响应状态码:%.f,响应信息:%s}", code, msg)
	}

	// 查询数据，如果有，则原基础更新。否则或许数据创建后，在更新。
	var info models.Info
	count := db.Take(&info, "type_id=?", typeID).RowsAffected

	if count != 0 {
		// 更新数据
		err = db.Model(&info).Update("content", responseData).Error
		if err != nil {
			global.Log.Error("Info更新出错，err=", err)
			return false, fmt.Sprintf("Info更新出错，err=%s", err)
		}
		return true, fmt.Sprintf("{响应状态码:%.f,响应信息:%s}", code, msg)
	} else if count == 0 {
		// 创建数据
		err = db.Create(&models.Info{
			Content: responseData,
			TypeId:  typeID,
		}).Error
		if err != nil {
			global.Log.Error("Info创建出错，err=", err)
			return false, fmt.Sprintf("Info创建出错，err=%s", err)
		}
	}
	return true, fmt.Sprintf("{响应状态码:%.f,响应信息:%s}", code, msg)
}
