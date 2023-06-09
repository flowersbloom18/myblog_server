package info

// 封装响应结果

// DouYinHotResponse 【1、抖音热搜】
type DouYinHotResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	//"hotindex": 11556814,
	//"label": 5,
	//"word": "EDG：猎鲨开始了"
	Result struct {
		List []struct {
			Word     string `json:"word"`
			HotIndex int    `json:"hotindex"`
		} `json:"list"`
	} `json:"result"`
}

// NetWorkHotResponse 【2、全网热搜】
type NetWorkHotResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`

	Result struct {
		List []struct {
			Title  string `json:"title"`
			HotNum int    `json:"hotnum"`
			Digest string `json:"digest"`
		} `json:"list"`
	} `json:"result"`
}

// WeiBoHotResponse 【3、微博热搜】
type WeiBoHotResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	//"hotword": "大叔被疑偷拍自证清白后仍遭女子曝光",
	//"hotwordnum": " 2323958",
	//"hottag": "新"
	Result struct {
		List []struct {
			HotWord    string `json:"hotword"`
			HotWordNum string `json:"hotwordnum"`
			HotTag     string `json:"hottag"`
		} `json:"list"`
	} `json:"result"`
}

// BulletInResponse 【4、每日简报】
type BulletInResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	//{
	//	"mtime": "2023-06-09",
	//	"title": "考生喊话取消调休",
	//	"digest": "6月8日，四川峨嵋，高考考生替网友喊话：“取消调休，取消调休”。高考互联网嘴替他来了！网友：让他上清华大学，我没开玩笑。"
	//},
	Result struct {
		List []struct {
			Title  string `json:"title"`
			Digest string `json:"digest"`
			MTime  string `json:"mtime"`
		} `json:"list"`
	} `json:"result"`
}

// ZaoAnResponse 【5、早安】
type ZaoAnResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	//	"result": {
	//	"content": "每天醒来记得鼓励自己。没有奇迹，只有自己的努力；没有运气，只有自己的坚持，每一份坚持都是成功的累积，相信自己的能力生命中就会有更多的惊喜！早安"
	//}
	Result struct {
		Content string `json:"content"`
	} `json:"result"`
}

// WanAnResponse 【6、晚安】
type WanAnResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`

	Result struct {
		Content string `json:"content"`
	} `json:"result"`
}

// LiShiResponse 【7、历史的今天】
type LiShiResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	//{
	//	"title": "俄国彼得大帝诞辰",
	//	"lsdate": "1672-06-09"
	//},
	Result struct {
		List []struct {
			Title  string `json:"title"`
			LsDate string `json:"lsdate"`
		} `json:"list"`
	} `json:"result"`
}
