package main

import (
	"bubble/db"
	_ "bubble/docs"
	"bubble/routers"
)

// @title			待办事项 API 文档
// @version		1.0
// @description	这是详细介绍待办事项的 API 文档
//
// @contact.name	huang
// @contact.email	nanguatou10@gmail
func main() {
	// 初始化数据库
	err := db.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// 注册路由
	r := routers.SetupRouter()
	r.Run(":9090")
}
