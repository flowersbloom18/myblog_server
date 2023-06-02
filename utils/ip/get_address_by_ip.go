package ip

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"myblog_server/global"
	"net/http"
)

type IPRequest struct {
	Resultcode string `json:"resultcode"`
	Reason     string `json:"reason"`
	Result     struct {
		Country  string `json:"Country"`  //国家
		Province string `json:"Province"` //省份
		City     string `json:"City"`     //城市
		District string `json:"District"`
		Isp      string `json:"Isp"`
	} `json:"result"`
	ErrorCode int `json:"error_code"`
}

// GetAddressByIp 将ip转为对应的地址
func GetAddressByIp(ip string) (address string) {
	key := global.Config.Juhe.Key
	url := fmt.Sprintf("http://apis.juhe.cn/ip/ipNewV3?ip=" + ip + "&key=" + key)
	// 发送GET请求
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("发送请求失败:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应数据
	body, err := ioutil.ReadAll(resp.Body)

	// 将json数据反序列化为结构体数据
	var result IPRequest
	err = json.Unmarshal(body, &result)
	if err != nil {
		//解析失败也是未知，解析失败有可能是接口出错了。
		address = "未知地址"
		global.Log.Error("数据解析失败:", err)
	}
	//判断是否响应成功，若响应失败，🥤有可能是次数用完了。就给一个未知地址。🥤也可能是接口出现错误
	if result.Resultcode == "200" && result.Reason == "查询成功" {
		// 判断是否为内网IP，
		if result.Result.Isp == "保留IP" {
			address = "内网IP"
		} else {
			// 判断省份是否为空，若空则只获取国家
			if result.Result.Province == "" {
				address = result.Result.Country
			} else {
				// 判断城市是否为空，若空则只获取国家和省份
				if result.Result.City == "" {
					address = result.Result.Country + "-" + result.Result.Province
				} else if result.Result.City == result.Result.Province {
					// 判断省份是否跟城市相同，若相同，则只保留省份，否则全部保留。
					address = result.Result.Country + "-" + result.Result.Province
				} else {
					address = result.Result.Country + "-" + result.Result.Province + "-" + result.Result.City
				}
			}
		}

	} else {
		address = "未知地址"
		global.Log.Warn("查询失败")
	}

	return address
}
