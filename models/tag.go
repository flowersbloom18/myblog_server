package models

// Tag 标签表
type Tag struct {
	MODEL
	Name  string `gorm:"size:18" json:"name"` // 标签名称
	Cover string `json:"cover"`               // 封面

	// json:"-"在返回json数据时不转换
	Blogs []Blog `gorm:"many2many:blog_tags" json:"-"` // 多对多关系，一个标签可以对应多个博客，方便preload预加载
}
