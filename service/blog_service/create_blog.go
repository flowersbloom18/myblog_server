package blog_service

import (
	"fmt"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/utils/link"
	"time"
)

const Cover = "/uploads/loading.gif"

func (BlogService) CreateBlog(title, content, cover string, isPublish, isTop bool, categoryId, userId uint, tagString []string) error {
	db := global.DB

	// title,abstract,content,cover,
	// read_num,comment_num,like_num,collect_num,
	// is_comment,is_publish,is_top,top_time,
	// category_id,user_id,link,
	// tag[]

	// 重复值判断
	var blogExist []models.Blog
	count := global.DB.Find(&blogExist, "title = ?", title).RowsAffected
	if count > 0 {
		return fmt.Errorf("重复的博客")
	}

	// 1、摘要获取的字数
	const MaxSummaryLength = 150

	var abstract string
	if len(content) <= MaxSummaryLength {
		abstract = content
	} else {
		abstract = content[:MaxSummaryLength]
	}

	// 2、封面如果为空则给个
	if cover == "" {
		cover = Cover
	}

	// 3、生成链接
	link := link.GetLink(title)

	blog := models.Blog{
		Title:      title,      // 标题
		Abstract:   abstract,   // 摘要
		Content:    content,    // 内容
		Cover:      cover,      // 封面
		ReadNum:    0,          // 阅读数
		CommentNum: 0,          // 评论数
		LikeNum:    0,          // 点赞数
		CollectNum: 0,          // 收藏数
		IsPublish:  isPublish,  // 是否发布
		IsTop:      isTop,      // 是否置顶
		TopTime:    time.Now(), // 置顶时间
		CategoryID: categoryId, // 分类ID
		UserID:     userId,     // 博主ID
		Link:       link,       // 博客链接®®®®®®
		//Tags:       []models.Tag{}, //标签？®®®®®®
	}

	// []string -->  []models.Tag
	var tags []models.Tag
	for _, tag := range tagString {
		// 将标签添加到切片中
		tags = append(tags, models.Tag{
			Name: tag,
		})
	}

	// 新增的标签，如果数据库不存在，则创建，如果创建则仅仅建立多对多关系即可。
	// 遍历标签列表
	for _, tag := range tags {
		// 检查标签是否已存在于数据库中
		var existingTag models.Tag
		err := db.First(&existingTag, "name = ?", tag.Name).Error
		if err != nil {
			// 标签不存在，创建新标签并关联到博客
			err = db.Create(&tag).Error
			if err != nil {
				// 处理错误
				return fmt.Errorf("%s", err)
			} else {
				// 建立多对多关联关系
				blog.Tags = append(blog.Tags, tag)
			}
		} else {
			// 标签已存在，直接关联到博客
			blog.Tags = append(blog.Tags, existingTag)
		}
	}

	err := db.Create(&blog).Error
	if err != nil {
		global.Log.Error("创建博客失败: ", err.Error())
		return fmt.Errorf("创建分类失败: %s", err.Error())
	}
	return nil
}
