package config

type SiteInfo struct {
	CreateAt          string `yaml:"create_at" json:"create_at"`                   // 建站时间
	BeiAn             string `yaml:"bei_an" json:"bei_an"`                         // 备案信息
	Title             string `yaml:"title" json:"title"`                           // 站点标题
	LogoLight         string `yaml:"logo_light" json:"logo_light"`                 // 站点logo，默认，白天模式
	LogoDark          string `yaml:"logo_dark" json:"logo_dark"`                   // 站点logo，夜间模式
	Favicon           string `yaml:"favicon" json:"favicon"`                       // 网站图标
	CopyrightProtocol string `yaml:"copyright_protocol" json:"copyright_protocol"` // 许可协议
	CopyRightInfo     string `yaml:"copyright_info" json:"copyright_info"`         // 版权信息
	ServerName        string `yaml:"server_name" json:"server_name"`               // 服务商名称
	QQImage           string `yaml:"qq_image" json:"qq_image"`                     // qq照片
	GiteeUrl          string `yaml:"gitee_url" json:"gitee_url"`                   // gitee地址
	GithubUrl         string `yaml:"github_url" json:"github_url"`                 // github地址
}
