package routers

import (
	"bubble/controller"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

/*
路由层：在路由层中，我们定义HTTP请求的URL路径和HTTP方法，并将其与处理该请求的控制器函数关联起来。
*/

func SetupRouter() *gin.Engine {
	// 创建一个带有默认中间件的新的gin示例，包括Logger中间件和Recovery中间件
	r := gin.Default()
	// 告诉gin框架模板文件引用的静态文件去哪里找
	r.Static("/static", "static")
	// 告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")
	// 为路由绑定处理函数
	r.GET("/", controller.IndexHandler)
	r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 创建路由分组v1
	v1Group := r.Group("v1")
	{
		v1Group.GET("/ping", controller.Ping)
		// 待办事项
		// 添加
		v1Group.POST("/todo", controller.CreateTodo)
		// 查看所有的待办事项
		v1Group.GET("/todo", controller.GetTodoList)
		// 修改某一个待办事项
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		// 删除某一个待办事项
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}
	return r
}
