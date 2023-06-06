package email

// SendEmailApi 发送邮箱的API
type SendEmailApi struct{}

// SendBindEmailContent 绑定邮箱
func (SendEmailApi) SendBindEmailContent(receiveEmail, nickName, authCode string) error {
	return SendEmail(receiveEmail, nickName, authCode, "电子邮件验证码："+authCode, BindEmailContent())
}

// SendForgetPwd 忘记密码，找回密码
func (SendEmailApi) SendForgetPwd(receiveEmail string, authCode string) error {
	return SendEmail(receiveEmail, "", authCode, "电子邮件验证码："+authCode, BindForgetPwd())
}

// SendUpdatePwd 密码更新提醒
func (SendEmailApi) SendUpdatePwd(receiveEmail string) error {
	return SendEmail(receiveEmail, "", "", "密码更新提醒", BindUpdatePwd())
}

func BindEmailContent() string {
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
				FlowersBloom 收到了将 <a style="font-weight: bold;" target="_blank" rel="noopener">{{.email}}</a> 绑定到用户昵称为 <a style="font-weight: bold;" target="_blank" rel="noopener">{{.name}}</a> 的请求。<br><br>请使用此验证码完成绑定邮箱的设置：<br>
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
	return htmlContent
}
func BindForgetPwd() string {
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
				<div style="font-size: 24px;text-align: center">请验证您的绑定邮箱 </div>
			</div>
			<div style="font-family: Roboto-Regular,Helvetica,Arial,sans-serif; font-size: 14px; color: rgba(0,0,0,0.87); line-height: 20px;padding-top: 20px; text-align: left;">
				FlowersBloom 收到了 <a style="font-weight: bold;" target="_blank" rel="noopener">{{.email}}</a> 重置密码的请求。<br><br>请使用此验证码完成重置设置：<br>
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
	return htmlContent
}
func BindUpdatePwd() string {
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
				<div style="font-size: 24px;text-align: center">密码更新提醒</div>
			</div>
			<div style="font-family: Roboto-Regular,Helvetica,Arial,sans-serif; font-size: 14px; color: rgba(0,0,0,0.87); line-height: 20px;padding-top: 20px; text-align: left;">
				FlowersBloom 邮箱 <a style="font-weight: bold;text-decoration: underline;" target="_blank" rel="noopener">{{.email}}</a> 的用户您好，您账号的密码已经修改。
				<br><br>如果您不认识
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
	return htmlContent
}
