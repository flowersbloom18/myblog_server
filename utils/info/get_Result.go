package info

import (
	"encoding/json"
	"fmt"
	"myblog_server/global"
)

// GetInfoResult 将结果转为前端可读的json
func GetInfoResult(content string, id int) (any, error) {

	// 1、将json数据转为结构体数据
	// 2、返回数据列表

	if id == 1 { // 1-抖音热搜
		var response DouYinHotResponse
		err := json.Unmarshal([]byte(content), &response)
		if err != nil {
			// 处理解析错误
			global.Log.Error("JSON解析错误:", err)
			return "", err
		}
		// key为空判断
		if response.Code == 230 {
			return "", fmt.Errorf("key错误或为空")
		}
		lists := response.Result.List
		return lists, nil
	} else if id == 2 { // 2-全网热搜
		var response NetWorkHotResponse
		err := json.Unmarshal([]byte(content), &response)
		if err != nil {
			// 处理解析错误
			global.Log.Error("JSON解析错误:", err)
			return "", err
		}
		// key为空判断
		if response.Code == 230 {
			return "", fmt.Errorf("key错误或为空")
		}
		lists := response.Result.List
		return lists, nil
	} else if id == 3 { // 3-微博热搜
		var response WeiBoHotResponse
		err := json.Unmarshal([]byte(content), &response)
		if err != nil {
			// 处理解析错误
			global.Log.Error("JSON解析错误:", err)
			return "", err
		}
		// key为空判断
		if response.Code == 230 {
			return "", fmt.Errorf("key错误或为空")
		}
		lists := response.Result.List
		return lists, nil
	} else if id == 4 { // 4-每日简报
		var response BulletInResponse
		err := json.Unmarshal([]byte(content), &response)
		if err != nil {
			// 处理解析错误
			global.Log.Error("JSON解析错误:", err)
			return "", err
		}
		// key为空判断
		if response.Code == 230 {
			return "", fmt.Errorf("key错误或为空")
		}
		lists := response.Result.List
		return lists, nil
	} else if id == 5 { // 5-早安
		var response ZaoAnResponse
		err := json.Unmarshal([]byte(content), &response)
		if err != nil {
			// 处理解析错误
			global.Log.Error("JSON解析错误:", err)
			return "", err
		}
		// key为空判断
		if response.Code == 230 {
			return "", fmt.Errorf("key错误或为空")
		}
		lists := response.Result.Content
		return lists, nil
	} else if id == 6 { // 6-晚安
		var response WanAnResponse
		err := json.Unmarshal([]byte(content), &response)
		if err != nil {
			// 处理解析错误
			global.Log.Error("JSON解析错误:", err)
			return "", err
		}
		// key为空判断
		if response.Code == 230 {
			return "", fmt.Errorf("key错误或为空")
		}
		lists := response.Result.Content
		return lists, nil
	} else if id == 7 { // 7-历史的今天
		var response LiShiResponse
		err := json.Unmarshal([]byte(content), &response)
		if err != nil {
			// 处理解析错误
			global.Log.Error("JSON解析错误:", err)
			return "", err
		}
		// key为空判断
		if response.Code == 230 {
			return "", fmt.Errorf("key错误或为空")
		}
		lists := response.Result.List
		return lists, nil
	} else {
		global.Log.Error("参数错误")
		return "", fmt.Errorf("参数错误")
	}
}
