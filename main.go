package main

import (
	"bubble/dao"
	_ "bubble/docs"
	"bubble/routers"
)

//	@title			示例 API 文档
//	@version		1.0
//	@description	一个简单的 Swagger 示例
//	@host			localhost:8080
//	@BasePath		/api/v1

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
