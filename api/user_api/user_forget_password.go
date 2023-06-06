package user_api

import (
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/response"
	"myblog_server/plugins/email"
	"myblog_server/service/redis_service"
	"myblog_server/utils/pwd"
	"myblog_server/utils/random"
	"time"
)

type BindForgetPwdRequest struct {
	Email    string  `json:"email" binding:"required,email" msg:"邮箱非法"`
	Code     *string `json:"code"`
	Password string  `json:"password"`
}

// UserForgetPasswordView 用户忘记密码，通过邮箱重置🥤
func (UserApi) UserForgetPasswordView(c *gin.Context) {

	// 用户忘记密码， 第一次输入是 邮箱
	var cr BindForgetPwdRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithError(err, &cr, c)
		return
	}

	if cr.Code == nil {
		// ⚠️优化了邮箱不存在、两次邮箱输入不一致的情况，以及恢复密码后，删除该key

		// 判断邮箱是否存在，不存在正是需要的
		var userModel models.User
		err = global.DB.Take(&userModel, "email = ?", cr.Email).Error
		// 如果err!=nil，表示系统找不到邮箱
		if err != nil {
			global.Log.Warn("邮箱不存在")
			response.FailWithMessage("邮箱不存在，请重新输入", c)
			return
		}

		// 第一次，后台发验证码
		// token是唯一的，把值改为邮箱_验证码，判断邮箱中数据是否一致即可
		// 生成4位验证码
		code := random.Code(6)

		// 写入redis(5分钟内有效）
		fiveMinutes := 5 * time.Minute
		err = redis_service.SetAuthCode(cr.Email, code, fiveMinutes)
		if err != nil {
			global.Log.Error(err)
			response.FailWithMessage("redis写入出错", c)
			return
		}

		// 发送验证码
		sendApi := email.SendEmailApi{}
		err = sendApi.SendForgetPwd(cr.Email, code)

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
	var userModel models.User
	err = global.DB.Take(&userModel, "email = ?", cr.Email).Error
	// 如果err!=nil，表示系统找不到邮箱
	if err != nil {
		global.Log.Warn("邮箱不存在")
		response.FailWithMessage("邮箱不存在，请重新输入", c)
		return
	}

	global.Log.Info("cr.Email= ", cr.Email)

	// 获取验证码(这是邮箱_code）
	code, err := redis_service.CheckAuthCode(cr.Email)

	// ⚠️内部错误，或者找不到数据都是,两次输入邮箱不一致
	if err != nil {
		global.Log.Error("两次邮箱不一致", err)
		response.FailWithMessage("两次邮箱不一致，请重新输入", c)
		return
	}

	// 判断数据是否为空
	if code == "" {
		response.FailWithMessage("验证码已过期，请重新获取", c)
		return
	}

	// 校验验证码
	if code != *cr.Code {
		response.FailWithMessage("验证码错误，请重新输入", c)
		return
	}

	// 修改用户的邮箱
	var user models.User

	// 查询对应邮箱的用户并将信息存储到user中
	err = global.DB.Take(&user, "email = ?", cr.Email).Error
	if err != nil {
		response.FailWithMessage("用户不存在", c)
		return
	}

	if len(cr.Password) < 4 {
		response.FailWithMessage("密码强度太低", c)
		return
	}
	hashPwd := pwd.HashPwd(cr.Password)
	//第一次的邮箱，和第二次的邮箱也要做一致性校验
	err = global.DB.Model(&user).Updates(map[string]any{
		"password": hashPwd,
	}).Error
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage("更新密码失败", c)
		return
	}
	// 完成绑定
	response.OkWithMessage("更新密码成功", c)

	// 系统日志记录
	logContent := "密码修改成功"
	global.DB.Create(&models.Log{
		UserName: user.UserName,
		NickName: user.NickName,
		IP:       user.IP,
		Address:  user.Address,
		Device:   user.Device,
		Level:    "info",
		Content:  logContent,
	})

	// 🥤密码更新提醒
	sendApi := email.SendEmailApi{}
	err = sendApi.SendUpdatePwd(user.Email)
	if err != nil {
		global.Log.Error("邮箱发送失败", err)
	}

	// 删除 key-value
	err = global.Redis.Del("authCode_" + cr.Email).Err()
	if err != nil {
		//response.FailWithMessage("验证码删除错误", c)
		global.Log.Error("验证码删除错误", err)
	}
	return
}
