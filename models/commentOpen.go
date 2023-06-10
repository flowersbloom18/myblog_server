package models

// CommentOpen 全局评论一键开关
type CommentOpen struct {
	MODEL
	IsOpen bool `json:"is_open"` // 是否开启
}
