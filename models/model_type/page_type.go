package model_type

import "encoding/json"

type PageType int

// 页面类型
const (
	PageBlog       PageType = 1 // 博客
	PageFriendLink PageType = 2 // 友链
	PageAbout      PageType = 3 // 关于
)

func (s PageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s PageType) String() string {
	var str string
	switch s {
	case PageBlog:
		str = "博客"
	case PageFriendLink:
		str = "友链"
	case PageAbout:
		str = "关于"
	default:
		str = "其他"
	}
	return str
}
