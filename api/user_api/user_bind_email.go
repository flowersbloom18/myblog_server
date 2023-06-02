package user_api

import (
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/response"
	"myblog_server/plugins/email"
	"myblog_server/service/redis_service"
	"myblog_server/utils/jwt"
	"myblog_server/utils/random"
	"strings"
	"time"
)

type BindEmailRequest struct {
	Email string  `json:"email" binding:"required,email" msg:"邮箱非法"` //绑定的邮箱
	Code  *string `json:"code"`                                      // 邮箱验证码
}

// UserBindEmailView 用户绑定邮箱
// 1、绑定邮箱可以使用邮箱登录（⚠️一个邮箱只能被注册一次！）
// 2、如果忘记密码，是否可以通过邮箱重置密码呢？⚠️🥤
func (UserApi) UserBindEmailView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.Claims)

	// 用户绑定邮箱， 第一次输入是 邮箱
	// 后台会给这个邮箱发验证码
	var cr BindEmailRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithError(err, &cr, c)
		return
	}

	// 获取用户的token，确定唯一性
	token := c.Request.Header.Get("token")

	if cr.Code == nil {
		// ⚠️优化了邮箱重复值、两次邮箱输入不一致的情况，以及邮箱绑定成功后，删除该key

		// 1、发送验证码之前必须要判断该邮箱是否被其它用户绑定，若绑定，则无法使用 ☑️
		// 2、不重复后，输入验证码和密码之后，两次邮箱输入不一致会怎样？所以要在验证码输入前后都进行邮箱重复判断 ☑️
		// 3、验证码在5分钟后会失效，且只能被使用一次 ，两次输入的邮箱必须相同，否则就会在验证码发送后使用为别人的邮箱☑️
		// 解决方案，第一次发送邮箱后，设置email信息，后续判断email是否相同即可。
		// 当使用code后，设置code为空，如果检测到为空，则表示code被使用，需要重新获取

		// 判断邮箱是否存在，不存在正是需要的
		var userModel models.UserModel
		err = global.DB.Take(&userModel, "email = ?", cr.Email).Error
		// 如果err==nil，表示系统存在改邮箱
		if err == nil {
			global.Log.Warn("邮箱已存在")
			response.FailWithMessage("邮箱已存在，请重新输入", c)
			return
		}

		// 第一次，后台发验证码
		// token是唯一的，把值改为邮箱_验证码，判断邮箱中数据是否一致即可
		// 生成4位验证码
		code := random.Code(4)
		codeEmail := cr.Email + "_" + code

		// 写入redis(5分钟内有效）
		fiveMinutes := 5 * time.Minute
		err = redis_service.SetAuthCode(token, codeEmail, fiveMinutes)
		if err != nil {
			global.Log.Error(err)
			response.FailWithMessage("redis写入出错", c)
			return
		}

		// 发送验证码

		//err = email.SendEmail(cr.Email, claims.NickName, code)
		sendApi := email.SendEmailApi{}
		err = sendApi.SendBindEmailContent(cr.Email, claims.NickName, code)

		if err != nil {
			global.Log.Error("邮箱不存在", err)
			response.OkWithMessage("邮箱不存在", c)
			return
		}
		response.OkWithMessage("验证码已发送，请查收", c)
		return
	}
	// 第二次，用户输入邮箱，验证码，密码

	// ⚠️判断邮箱是否存在
	var userModel models.UserModel
	err = global.DB.Take(&userModel, "email = ?", cr.Email).Error
	// 如果err==nil，表示系统存在改邮箱
	if err == nil {
		global.Log.Warn("邮箱已存在")
		response.FailWithMessage("邮箱已存在", c)
		return
	}

	// 获取验证码(这是邮箱_code）
	newCode, err := redis_service.CheckAuthCode(token)
	// 内部错误，或者找不到数据都是验证码过期了。
	if err != nil {
		global.Log.Error("验证码已过期，请重新获取", err)
		response.FailWithMessage("验证码已过期，请重新获取", c)
		return
	}

	// 判断数据是否为空
	if newCode == "" {
		response.FailWithMessage("验证码已过期，请重新获取", c)
		return
	}

	// 分别获取value中的email和code
	email := strings.Split(newCode, "_")[0]

	// 判断两次邮箱是否一致
	if email != cr.Email {
		response.FailWithMessage("两次邮箱不一致！", c)
		return
	}
	code := strings.Split(newCode, "_")[1]

	// 校验验证码
	if code != *cr.Code {
		response.FailWithMessage("验证码错误，请重新输入", c)
		return
	}

	// 修改用户的邮箱
	var user models.UserModel

	// 查询对应ID的用户并将信息存储到user中
	err = global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		response.FailWithMessage("用户不存在", c)
		return
	}

	err = global.DB.Model(&user).Updates(map[string]any{
		"email": cr.Email,
	}).Error
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage("绑定邮箱失败", c)
		return
	}
	// 完成绑定
	response.OkWithMessage("邮箱绑定成功", c)

	// 完成绑定则value设置为空

	// 删除 key-value
	err = global.Redis.Del("authCode_" + token).Err()
	if err != nil {
		//response.FailWithMessage("验证码删除错误", c)
		global.Log.Error("验证码删除错误", err)
	}
	return
}
