package blog_api

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/response"
	"strings"
)

// BlogUpdateRequest 【暂且忽略，记录数量的情况】
type BlogUpdateRequest struct {
	Title      string   `json:"title" binding:"required" msg:"请输入标题"`         // 标题
	Content    string   `json:"content" binding:"required" msg:"请输入内容"`       // 内容
	Cover      string   `json:"cover" `                                       // 封面
	IsComment  bool     `json:"is_comment"`                                   // 是否开启评论
	IsPublish  bool     `json:"is_publish"`                                   // 是否发布
	IsTop      bool     `json:"is_top"`                                       // 是否置顶
	CategoryID uint     `json:"category_id" binding:"required" msg:"请输入分类id"` // 分类ID
	Tags       []string `json:"tags"`                                         // 标签列表
}

//{
//	"Title     "  :"",
//	"Content   "  :"",
//	"Cover     "  :"",
//	"IsComment "  :"",
//	"IsPublish "  :"",
//	"IsTop     "  :"",
//	"CategoryID"  :"",
//	"Tags      "  :"",
//}

// BlogUpdateView 更新分类（名称和封面）
func (BlogApi) BlogUpdateView(c *gin.Context) {
	db := global.DB

	var cr BlogUpdateRequest
	err := c.ShouldBindJSON(&cr) // 请求的json数据绑定到结构体中
	if err != nil {
		response.FailWithError(err, &cr, c)
		return
	}

	// 获取需要更新的博客链接
	link := c.Param("link")
	// 去除前缀斜杠
	link = strings.TrimPrefix(link, "/")

	// 链接是否存在！
	var blog models.Blog
	err = db.Take(&blog, "link=?", link).Error
	if err != nil {
		response.FailWithMessage("博客不存在！", c)
		return
	}

	// 判断更新博客的标题是否重复
	// 如果标题变化，则需要判断，反之不需要
	if cr.Title != blog.Title {
		err = db.Take(models.Blog{}, "title=?", cr.Title).Error
		if err == nil {
			response.FailWithMessage("博客标题重复！", c)
			return
		}
	}
	// 字符串切片，转为实体切片
	var tags []models.Tag
	for _, tag := range cr.Tags {
		// 将标签添加到切片中
		tags = append(tags, models.Tag{
			Name: tag,
		})
	}

	// 【判断是存在的还是新增】新增的标签，如果数据库不存在，则创建，如果创建则仅仅建立多对多关系即可。
	for _, tag := range tags {
		// 检查标签是否已存在于数据库中
		var existingTag models.Tag
		err := db.First(&existingTag, "name = ?", tag.Name).Error
		if err != nil {
			// 标签不存在，创建新标签并关联到博客
			err = db.Create(&tag).Error
			if err != nil {
				// 处理错误
				response.FailWithMessage(fmt.Sprintf("%s", err), c)
				//return
			} else {
				// 建立多对多关联关系
				blog.Tags = append(blog.Tags, tag)
			}
		} else {
			// 标签已存在，直接关联到博客
			blog.Tags = append(blog.Tags, existingTag)
		}
	}

	// 更新数据需要同步更新摘要
	if cr.Content != blog.Content {
		blog.Abstract = "？？摘要被更新"
	}

	// 更新数据
	maps := structs.Map(&cr)
	err = global.DB.Model(&blog).Updates(maps).Error
	if err != nil {
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
	return

	//
	//
	////var cr
	//err := c.ShouldBindJSON(&cr)
	//if err != nil {
	//	response.FailWithError(err, &cr, c)
	//	return
	//}
	//
	//var category models.Category
	//// 获取对应id的数据
	//err = db.Take(&category, "id = ?", id).Error
	//if err != nil {
	//	global.Log.Warn("分类不存在")
	//	response.FailWithMessage("分类不存在", c)
	//	return
	//}
	//
	//// 检查新创建的分类名称是否存在，不存在则更新【分类重复值判断】
	//var category1 models.Category
	//// ⚠️如果说新增分类跟当前数据库【不同】则进行重复值判断，否则不用。
	//if category.Name != cr.Name {
	//	err = db.First(&category1, "name = ?", cr.Name).Error
	//	// 查询到数据则err为空，说明分类已经存在。反之不存在，即可更新数据。
	//	if err == nil {
	//		global.Log.Warn("分类已存在", err)
	//		response.FailWithMessage("分类已存在", c)
	//		return
	//	}
	//}
	//
	//// 更新

}
