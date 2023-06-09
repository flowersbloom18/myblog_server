package config

type TianApi struct {
	DouYinHot  string `yaml:"dou_yin_hot"` // 抖音热搜
	NetWorkHot string `yaml:"network_hot"` // 全网热搜
	WeiBoHot   string `yaml:"wei_bo_hot"`  // 微博热搜
	BulletIn   string `yaml:"bullet_in"`   // 每日简报
	ZaoAn      string `yaml:"zao_an"`      // 早安
	WanAn      string `yaml:"wan_an"`      // 晚安
	LiShi      string `yaml:"li_shi"`      // 历史的今天
}
