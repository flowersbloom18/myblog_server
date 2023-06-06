package blog_service

import (
	"myblog_server/global"
	"myblog_server/models"
)

// GetBlogDetail 将数据库的数据进行再次封装，
func (BlogService) GetBlogDetail(blog models.Blog) ([]BlogResponse, error) {

	var blogResponse []BlogResponse
	// 对结果进行处理。
	// 第一层循环获得标签列表，第二层循环获取单个标签

	// 将标签结构体转为字符切片
	var tags []string
	for _, v2 := range blog.Tags {
		tags = append(tags, v2.Name)
	}

	// 根据分类id和用户id来获取昵称
	var category models.Category
	var user models.User

	err := global.DB.Take(&category, "id=?", blog.CategoryID).Error
	if err != nil {
		global.Log.Warn("找不到该分类")
	}
	err = global.DB.Take(&user, "id=?", blog.UserID).Error
	if err != nil {
		global.Log.Warn("找不到该用户")
	}

	blogResponse = append(blogResponse, BlogResponse{
		ID:       blog.ID,
		Title:    blog.Title,
		Abstract: blog.Abstract,
		Cover:    blog.Cover,
		Content:  blog.Content,

		Author:   user.UserName,
		Category: category.Name,
		Tags:     tags,

		ReadNum:    blog.ReadNum,
		CommentNum: blog.CommentNum,
		LikeNum:    blog.LikeNum,
		CollectNum: blog.CollectNum,

		IsComment: blog.IsComment,
		IsPublish: blog.IsPublish,
		IsTop:     blog.IsTop,

		CreatedAt: blog.CreatedAt,
		Link:      blog.Link,
	})
	return blogResponse, nil
}

// 删除！
