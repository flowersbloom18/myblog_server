package config

type Config struct {
	Mysql    Mysql    `yaml:"mysql"`
	Logger   Logger   `yaml:"logger"`
	System   System   `yaml:"system"`
	Juhe     Juhe     `yaml:"juhe"`
	Jwt      Jwt      `yaml:"jwt"`
	Redis    Redis    `yaml:"redis"`
	Email    Email    `yaml:"email"`
	TianApi  TianApi  `yaml:"tianapi"`
	QiNiu    QiNiu    `yaml:"qi_niu"`
	Upload   Upload   `yaml:"upload"`
	SiteInfo SiteInfo `yaml:"site_info"`
}
