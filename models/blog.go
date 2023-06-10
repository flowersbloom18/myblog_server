package models

import "time"

// Blog 博客表
type Blog struct {
	MODEL
	Title      string    `gorm:"size:30" json:"title"`                                 // 标题
	Abstract   string    `gorm:"size:200" json:"abstract"`                             // 摘要
	Content    string    `gorm:"type:mediumtext" json:"content"`                       // 内容
	Cover      string    `json:"cover"`                                                // 封面
	ReadNum    int       `gorm:"type:int(6);default:0;" json:"read_num"`               // 阅读数量
	CommentNum int       `gorm:"type:int(6);default:0;" json:"comment_num"`            // 评论数量
	LikeNum    int       `gorm:"type:int(6);default:0;" json:"like_num"`               // 点赞数量
	CollectNum int       `gorm:"type:int(6);default:0;" json:"collect_num"`            // 收藏数量
	IsPublish  bool      `json:"is_publish"`                                           // 是否发布
	IsTop      bool      `json:"is_top"`                                               // 是否置顶
	TopTime    time.Time `json:"top_time"`                                             // 置顶时间
	CategoryID uint      `json:"category_id" gorm:"type:int(6);foreignKey:CategoryID"` // 分类ID

	// 定义与 CategoryM 关联的字段，通过这种关联关系，可以在查询博客时同时加载关联的分类信息。
	//Category Category `gorm:"foreignKey:CategoryID" json:"-"` // 一对多关系，一个博客对应一个分类，获取分类信息//json:"-"
	UserID uint   `json:"user_id"` // 发布人ID
	Link   string `json:"link" `   // 博客链接

	Tags []Tag `gorm:"many2many:blog_tags" json:"tags"` // ⚠️多对多关系，一个博客可以有多个标签???
}

//title,abstract,content,cover,read_num,comment_num,like_num,collect_num,
//is_comment,is_publish,is_top,top_time,category_id,user_id,link,// tags[]
