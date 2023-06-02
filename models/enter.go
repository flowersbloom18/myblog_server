package models

import "time"

type MODEL struct {
	ID        uint      `gorm:"primarykey" json:"id,select($any)" structs:"-"` // 主键ID
	CreatedAt time.Time `json:"created_at,select($any)" structs:"-"`           // 创建时间
	UpdatedAt time.Time `json:"-" structs:"-"`                                 // 更新时间
}

// PageInfo 分页查询
type PageInfo struct {
	Page  int    `form:"page"`  // 第几页🥤
	Key   string `form:"key"`   //
	Limit int    `form:"limit"` // 一页限制几条🥤
	Sort  string `form:"sort"`  // 排序方式
}

// RemoveRequest 单个删除/部分删除
type RemoveRequest struct {
	IDList []uint `json:"id_list"`
}
