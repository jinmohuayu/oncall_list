package main

// config 配置
type config struct {
	Port  int    // 监听端口
	Mysql string `yaml:"mysql"` // mysql 连接字符串
}

// Get 获取配置
func (c *config) Get() error {

	// TODO: 改成从外部读取
	c.Port = 9981
	c.Mysql = "user:password@/dbname"

	return nil
}
