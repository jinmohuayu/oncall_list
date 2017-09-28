package config

import (
	"os"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"log"
)

// config 配置
type Config struct {
	Port  int    // 监听端口
	Mysql string `yaml:"mysql"` // mysql 连接字符串
}

// Get 获取配置
func (c *Config) Get() error {

	// 打开文件
	file, err := os.OpenFile("config.yaml", os.O_RDONLY, 0666)
	if err != nil {
		return nil
	}
	defer file.Close()

	// 读取文件
	buffer, err := ioutil.ReadAll(file)
	if err != nil {
		return nil
	}

	//将yaml形式的字符串解析成struct类型
	err = yaml.Unmarshal(buffer, c)
	if err != nil {
		log.Fatal("error: %v", err)
	}

	return nil
}
