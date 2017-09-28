package main

import (
	"log"
	"os"
	"runtime/debug"

	"git.elenet.me/DA/oncall/config"
	"git.elenet.me/DA/oncall_list/data/mysql"
	"git.elenet.me/DA/oncall_list/web"
)

func main() {

	logger := log.New(os.Stdout, "", log.LstdFlags)
	log.Print("oncall_list start")

	// 当程序意外退出时记录日志
	defer func(_logger *log.Logger) {
		logWhenPanic(_logger)
	}(logger)

	// 读取配置
	conf := config.Config{}
	err := conf.Get()
	if err != nil {
		log.Fatal(err)
	}

	// 连接mysql
	db, err := mysql.GetDB(conf.Mysql)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 启动web 服务器
	server := web.NewServer(&web.ServerConfig{DB: db, Port: conf.Port}, logger)
	server.StartAndWait()
}

// logWhenPanic 当程序意外退出时记录日志
func logWhenPanic(logger *log.Logger) {

	// 本地日志也记录以防万一
	logger.Print("程序发生了意外退出")

	// 捕获panic异常
	if err := recover(); err != nil {
		logger.Print("致命错误:", err)
	}

	// 打印栈
	stack := string(debug.Stack())
	logger.Print(stack)
}
