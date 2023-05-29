package core

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/fs"
	"log"
	"myblog_server/config"
	"myblog_server/global"
	"os"
)

const ConfigFile = "settings.yaml"

// InitConf 读取yaml文件的配置
func InitConf() {
	c := &config.Config{}

	yamlConf, err := os.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf error: %s", err).(any))
	}

	err = yaml.Unmarshal(yamlConf, c) //⚠️将一个 YAML 格式的配置文件解析成一个结构体对象;反序列化
	if err != nil {
		log.Fatalf("config Init Unmarshal: %v", err) //输出日志并终止程序运行
	}
	log.Println("😁yaml文件初始化成功😁")

	//⚠️全局配置
	global.Config = c
}

// SetYaml 写入yaml文件
func SetYaml() error {
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		return err
	}
	err = os.WriteFile(ConfigFile, byteData, fs.ModePerm)
	if err != nil {
		return err
	}
	global.Log.Info("配置文件修改成功")
	return nil
}
