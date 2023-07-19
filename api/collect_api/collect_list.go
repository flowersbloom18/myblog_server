package collect_api

import (
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/response"
	"myblog_server/utils/jwt"
)

func (CollectApi) CollectListView(c *gin.Context) {
	_claims, _ := c.Get("claims") // 断言，然后调用
	claims := _claims.(*jwt.Claims)
	userID := claims.UserID

	var cr models.PageInfo // contentRequest内容请求 -->cr
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}
	var collects []models.Collect
	global.DB.Find(&collects, "user_id=?", userID)

	var blogs []models.Blog
	for _, v := range collects {

		var blog models.Blog
		err := global.DB.Take(&blog, "id=?", v.BlogID).Error
		if err != nil {
			global.Log.Warn("博客不存在")
			continue
		}
		blogs = append(blogs, blog)
	}
	count := int64(len(collects))
	response.OkWithList(blogs, count, c)
}
