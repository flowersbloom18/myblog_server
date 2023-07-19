package collect_api

import (
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/response"
	"myblog_server/utils/jwt"
	"strconv"
)

func (CollectApi) CollectCreateView(c *gin.Context) {
	_blogID := c.Param("blogID")
	blogID, err := strconv.Atoi(_blogID)

	_claims, _ := c.Get("claims") // 断言，然后调用
	claims := _claims.(*jwt.Claims)
	userID := claims.UserID

	// 1、要收藏的博客必须存在
	// 2、一个用户只能收藏一次任意一篇博客。
	if err != nil {
		global.Log.Error("请求数据出错，err:", err)
		response.FailWithMessage("请求数据出错", c)
		return
	}
	// 查询博客是否存在
	err = global.DB.Take(&models.Blog{}, "id=?", blogID).Error
	if err != nil {
		global.Log.Error("该博客不存在，err:", err)
		response.FailWithMessage("该博客不存在", c)
		return
	}

	var collect models.Collect
	// 查询博客是否被当前用户收藏过
	count := global.DB.Where("blog_id = ?", blogID).Where("user_id = ?", userID).Take(&collect).RowsAffected

	// 收藏过的话，博客收藏列表+1，反之-1【】
	var blog models.Blog
	err = global.DB.Where("id = ?", blogID).Take(&blog).Error
	if err != nil {
		response.FailWithMessage("数据查询失败", c)
		return
	}

	// 收藏过，则删除，否则继续收藏
	if count != 0 {
		global.DB.Delete(&collect)

		global.DB.Model(blog).Update("CollectNum", blog.CollectNum-1)

		response.OkWithMessage("取消收藏成功", c)
		return
	} else {
		global.DB.Create(&models.Collect{ // 创建
			BlogID: blogID,
			UserID: int(userID),
		})
		global.DB.Model(blog).Update("CollectNum", blog.CollectNum+1)

		response.OkWithMessage("收藏博客成功", c)
	}

}
