package router

import (
	"bubble/pkg/router/routes"
	"github.com/gin-gonic/gin"
)

func New(app *gin.Engine) {
	// 创建多版本 API 路由组
	api1 := app.Group("/api/v1")
	//api2 := app.Group("/api/v2")

	// 注册账户相关的路由
	routes.RegisterAccountRoutes(api1)
	// 注册待办事项相关的路由
	routes.RegisterTodoRoutes(api1)
}
