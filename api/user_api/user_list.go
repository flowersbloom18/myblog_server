package user_api

import (
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/model_type"
	"myblog_server/models/response"
	"myblog_server/service/common"
	"myblog_server/utils/desensitization"
	"myblog_server/utils/jwt"
)

type UserResponse struct {
	User   models.User `json:"user"`
	RoleID int         `json:"role_id"`
}

type UserListRequest struct {
	models.PageInfo
	Role int `json:"role" form:"role"`
}

func (UserApi) UserListView(c *gin.Context) {
	//var cr models.PageInfo // contentRequest内容请求 -->cr
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.Claims)

	var cr UserListRequest
	err := c.ShouldBindQuery(&cr) // 将请求参数绑定到结构体中
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}

	// 查询对应权限的用户
	list, count, err := common.ComList(models.User{Role: model_type.Role(cr.Role)}, common.Option{
		PageInfo: cr.PageInfo,           // 第几页、一页多少条、哪个用户（key）
		Likes:    []string{"nick_name"}, // 模糊查询关键字,根据昵称查询
		//Debug:    true,
	})
	if err != nil {
		global.Log.Warn("获取数据错误：", err)
	}

	var users []UserResponse
	// 如果不是管理员，则对信息进行脱敏处理
	for _, value := range list {
		if model_type.Role(claims.Role) != model_type.PermissionAdmin {
			value.UserName = desensitization.DesensitizationUserName(value.UserName)
			value.Email = desensitization.DesensitizationEmail(value.Email)
			value.IP = "****"
		}
		users = append(users, UserResponse{
			User:   value,
			RoleID: int(value.Role),
		})
	}
	response.OkWithList(users, count, c)
}
