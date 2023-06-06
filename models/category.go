package models

// Category 分类表
type Category struct {
	MODEL
	Name  string `gorm:"size:18" json:"name"` // 分类名称
	Cover string `json:"cover"`               // 封面

	// 方便查询，通过preload的方式
	Blogs []Blog `gorm:"foreignKey:CategoryID" json:"-"` // 一对多关系，一个分类可以对应多个博客,不进行json转换
}
