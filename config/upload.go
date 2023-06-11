package config

// Upload 本地配置
type Upload struct {
	Size int    `yaml:"size" json:"size"` // 附件上传的大小
	Path string `yaml:"path" json:"path"` // 附件上传的目录
}
