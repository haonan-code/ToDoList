package main

import (
	"bubble/config"
	"bubble/db"
	_ "bubble/docs"
	"bubble/routers"
)

// @title			待办事项 API 文档
// @version		1.0
// @description	这是详细介绍待办事项的 API 文档
//
// @contact.name	huang
// @contact.email	nanguatou10@gmail.com
func main() {
	// 读取配置信息
	config.InitConfig()
	// 初始化数据库
	db.InitDatabase()
	// 注册路由
	r := routers.SetupRouter()
	r.Run(":9090")
}
