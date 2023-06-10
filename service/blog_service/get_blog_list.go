package blog_service

import (
	"myblog_server/global"
	"myblog_server/models"
	"time"
)

type BlogResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`      // 标题
	Abstract  string    `json:"abstract"`   // 摘要
	Cover     string    `json:"cover"`      // 封面
	Content   string    `json:content`      // 内容
	CreatedAt time.Time `json:"created_at"` // 创建时间

	ReadNum    int `json:"read_num"`    // 阅读数
	CommentNum int `json:"comment_num"` // 评论数
	LikeNum    int `json:"like_num"`    // 点赞数
	CollectNum int `json:"collect_num"` // 收藏数

	IsComment bool `json:"is_comment"` // 是否评论
	IsPublish bool `json:"is_publish"` // 是否发布
	IsTop     bool `json:"is_top"`     // 是否置顶
	//TopTime   time.Time `json:"top_time"`   // 置顶时间

	Category string   `json:"category"` // 分类名称
	Author   string   `json:"author"`   // 管理员昵称
	Link     string   `json:"link"`     // 博客链接
	Tags     []string `json:"tags"`
}

// GetBlogList  将数据库的数据进行再次封装，
func (BlogService) GetBlogList(list []models.Blog) ([]BlogResponse, error) {

	var blogResponse []BlogResponse
	// 对结果进行处理。
	// 第一层循环获得标签列表，第二层循环获取单个标签
	for _, v1 := range list {

		// 将标签结构体转为字符切片
		var tags []string
		// 需要判断博客的标签是否为空，若为空，则直接返回空字符数组，否则将数据存进tags
		if len(v1.Tags) == 0 {
			//global.Log.Info("v1.Tags=", v1.Tags)
			tags = []string{} // 空
		} else {
			for _, v2 := range v1.Tags {
				tags = append(tags, v2.Name)
			}
		}

		// 根据分类id和用户id来获取昵称
		var category models.Category
		var user models.User

		err := global.DB.Take(&category, "id=?", v1.CategoryID).Error
		if err != nil {
			global.Log.Warn("找不到该分类")
		}
		err = global.DB.Take(&user, "id=?", v1.UserID).Error
		if err != nil {
			global.Log.Warn("找不到该用户")
		}

		blogResponse = append(blogResponse, BlogResponse{
			ID:       v1.ID,
			Title:    v1.Title,
			Abstract: v1.Abstract,
			Cover:    v1.Cover,
			Content:  v1.Content,

			Author:   user.UserName,
			Category: category.Name,
			Tags:     tags,

			ReadNum:    v1.ReadNum,
			CommentNum: v1.CommentNum,
			LikeNum:    v1.LikeNum,
			CollectNum: v1.CollectNum,

			IsPublish: v1.IsPublish,
			IsTop:     v1.IsTop,

			CreatedAt: v1.CreatedAt,
			Link:      v1.Link,
		})
	}
	return blogResponse, nil
}
