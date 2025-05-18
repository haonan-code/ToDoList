package main

import (
	"bubble/dao"
	_ "bubble/docs"
	"bubble/routers"
)

//	@title			待办事项 API 文档
//	@version		1.0
//	@description	这是详细介绍待办事项的 API 文档
//
//	@contact.name	huang
//	@contact.email	nanguatou10@gmail
func main() {
	// 创建数据库
	// sql: CREATE DATABASE bubble;
	// 连接数据库 & 模型绑定
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.Close() // 程序退出关闭数据库连接
	// 注册路由
	r := routers.SetupRouter()
	r.Run(":9090")
}
