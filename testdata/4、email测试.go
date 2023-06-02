package main

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"myblog_server/core"
	"myblog_server/global"
	"strings"
)

// 邮箱发送测试
func main() {
	sendEmail()
}

func sendEmail() error {
	// 数据初始化
	core.InitConf()
	host := global.Config.Email.Host           // 服务器地址
	port := global.Config.Email.Port           // 服务器端口
	sendEmail := global.Config.Email.SendEmail // 发送人邮箱
	pwd := global.Config.Email.Password        // 发送人密钥
	logoEmail := global.Config.Email.LogoEmail // 邮箱Logo
	sendName := global.Config.Email.SendName   // 发送人昵称

	authCode := "8848"                 // 验证码
	receiveEmail := "xxxx@qq.com"      // ⚠️接收人邮箱
	sendTitle := "电子邮件验证码：" + authCode // 发送的标题
	receiveName := "张三"                // 接收者昵称

	// 创建邮件模板
	message := gomail.NewMessage()

	message.SetHeader("From", message.FormatAddress(sendEmail, sendName)) // 发送人邮箱和昵称
	message.SetHeader("To", receiveEmail)                                 // 收件人邮箱
	message.SetHeader("Subject", sendTitle)                               // 发送的标题

	// 创建HTML内容
	htmlContent := `
		<!DOCTYPE html>
		<html>
		<head>
			<meta charset="UTF-8">
			<title>平台验证码</title>
			<style>
				body {
					font-family: Arial, sans-serif;
					background-color: #f2f2f2;
				}
		
				.container {
					max-width: 600px;
					margin: 0 auto;
					padding: 20px;
					background-color: #ffffff;
					border: 1px solid #e6e6e6;
				}
		
			</style>
		</head>
		<body>
		<div class="container">
			<img src="{{.logo-email}}" alt="" width="100%">
			<div style="">
				<div style="font-size: 24px;text-align: center">请验证您的绑定邮箱</div>
			</div>
			<div style="font-family: Roboto-Regular,Helvetica,Arial,sans-serif; font-size: 14px; color: rgba(0,0,0,0.87); line-height: 20px;padding-top: 20px; text-align: left;">
				FlowersBloom 收到了将 <a style="font-weight: bold;" target="_blank" rel="noopener">{{.email}}</a> 绑定到用户名为 <a style="font-weight: bold;" target="_blank" rel="noopener">{{.name}}</a> 的请求。<br><br>请使用此验证码完成绑定邮箱的设置：<br>
				<div style="text-align: center; font-size: 36px; margin-top: 20px; line-height: 44px;">
					<span style="border-bottom:1px dashed #ccc;z-index:1" t="7" onclick="return false;">{{.auth_code}}</span></div>
				<br>此验证码将在 5 分钟后失效。<br><br>如果您不认识
				<a style="font-weight: bold;" target="_blank" rel="noopener">{{.email}}</a>，可以放心地忽略这封电子邮件。
			</div>
		
			<div style="text-align: left;">
				<div style="font-family: Roboto-Regular,Helvetica,Arial,sans-serif;color: rgba(0,0,0,0.54); font-size: 11px; line-height: 18px; padding-top: 12px; text-align: center;">
					<div style="direction: ltr;">
						© 2023 FlowersBloom
					</div>
				</div>
			</div>
		</div>
		</body>
		</html>
	`

	// 定义要替换的变量
	variables := map[string]interface{}{
		"logo-email": logoEmail,    // logo-email
		"email":      receiveEmail, // 接收人邮箱
		"name":       receiveName,  // 接收人用户名
		"auth_code":  authCode,     // 验证码
	}

	// 使用模板引擎替换变量
	htmlContent = replaceVariables(htmlContent, variables)

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

// 替换HTML模板中的变量
func replaceVariables(html string, variables map[string]interface{}) string {
	for key, value := range variables {
		placeholder := "{{." + key + "}}"
		html = strings.ReplaceAll(html, placeholder, fmt.Sprintf("%v", value))
	}
	return html
}
