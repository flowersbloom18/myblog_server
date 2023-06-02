package email

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"myblog_server/global"
	"strings"
)

// SendEmail 发送邮箱（超时处理）
func SendEmail(receiveEmail, nickName, authCode, sendTitle, htmlContent string) error {
	// 数据初始化
	host := global.Config.Email.Host           // 服务器地址
	port := global.Config.Email.Port           // 服务器端口
	sendEmail := global.Config.Email.SendEmail // 发送人邮箱
	pwd := global.Config.Email.Password        // 发送人密钥
	logoEmail := global.Config.Email.LogoEmail // 邮箱Logo

	sendName := global.Config.Email.SendName // 发送人昵称

	// 创建邮件模板
	message := gomail.NewMessage()

	message.SetHeader("From", message.FormatAddress(sendEmail, sendName)) // 发送人邮箱和昵称
	message.SetHeader("To", receiveEmail)                                 // 收件人邮箱
	message.SetHeader("Subject", sendTitle)                               // 发送的标题

	// 定义要替换的变量
	variables := map[string]interface{}{
		"logo-email": logoEmail,    // logo-email
		"email":      receiveEmail, // 接收人邮箱
		"name":       nickName,     // 用户昵称
		"auth_code":  authCode,     // 验证码
	}

	// 使用模板引擎替换变量
	htmlContent = ReplaceVariables(htmlContent, variables)
	// 设置邮件内容
	message.SetBody("text/html", htmlContent)

	// 配置SMTP服务器信息
	dialer := gomail.NewDialer(host, port, sendEmail, pwd)

	// 发送邮件
	err := dialer.DialAndSend(message)
	if err != nil {
		return err // 错误
	} else {
		return nil //成功
	}
}

// ReplaceVariables 替换HTML模板中的变量
func ReplaceVariables(html string, variables map[string]interface{}) string {
	for key, value := range variables {
		placeholder := "{{." + key + "}}"
		html = strings.ReplaceAll(html, placeholder, fmt.Sprintf("%v", value))
	}
	return html
}
