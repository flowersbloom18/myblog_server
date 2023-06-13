package main

import (
	"encoding/json"
	"fmt"
	"myblog_server/global"
	"myblog_server/utils/info"
)

func main() {
	url := "https://apis.tianapi.com/douyinhot/index?key=你的APIKEY"
	// 获取响应结果
	responseData, err := info.GetHttpResponse(url)
	if err != nil {
		global.Log.Warn("err=", err)
		return
	}
	fmt.Println("响应结果为：", responseData)
	// {"code":230,"msg":"key错误或为空"}

	// 转为map集合，方便获取数据
	var data map[string]interface{}
	err = json.Unmarshal([]byte(responseData), &data)
	if err != nil {
		fmt.Println("解析 JSON 出错:", err)
		return
	}

	// 需要类型断言
	code := data["code"].(int)
	msg := data["msg"].(string)

	fmt.Println("Code:", code)
	fmt.Println("Msg:", msg)
}
