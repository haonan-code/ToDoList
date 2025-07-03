package cmd

import (
	"bubble/configs"
	"bubble/internal/db"
	"bubble/pkg/router"
	"github.com/gin-gonic/gin"
)

func Start() {
	// 读取配置信息
	config.InitConfig()

	// todo 初始化 Logger

	// todo 初始化 gin 实例
	app := gin.Default()

	// todo 初始化中间件

	// 初始化数据库连接并自动迁移模型
	db.New()

	// todo 初始化 Redis 连接

	// 注册路由
	router.New(app)

	// 启动服务
	app.Run(":9090")
}
