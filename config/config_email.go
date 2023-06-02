package config

type Email struct {
	Host      string `json:"host" yaml:"host"`             // 地址
	Port      int    `json:"port" yaml:"port"`             // 端口号
	SendEmail string `json:"send_email" yaml:"send_email"` // 发件人邮箱
	Password  string `json:"password" yaml:"password"`     // 密钥
	SendName  string `json:"send_name" yaml:"send_name"`   // 发件人名字
	LogoEmail string `json:"logo_email" yaml:"logo_email"` // 邮箱Logo
}
