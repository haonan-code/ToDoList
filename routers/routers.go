package routers

import (
	"bubble/controller"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 告诉gin框架模板文件引用的静态文件去哪里找
	r.Static("/static", "static")
	// 告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")

	r.GET("/", controller.IndexHandler)
	r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 注册认证相关的路由
	userGroup := r.Group("/auth")
	{
		userGroup.POST("/register", controller.Register)       // 用户注册
		userGroup.POST("/login", controller.Login)             // 用户登录
		userGroup.GET("/me", controller.GetMyInfo)             // 获取当前用户
		userGroup.PUT("/me", controller.UpdateUserInfo)        // 更新当前用户信息
		userGroup.PUT("/changepwd", controller.ChangePassword) //  修改密码
	}

	// TODO：添加用户管理模块
	//adminGroup := r.Group("/admin")
	//{
	//	adminGroup.GET("/users", controller.GetUserList)       // 查询所有用户（管理端）
	//	adminGroup.GET("/users/:id", controller.GetUserInfo)   // 查看单个用户信息（管理端）
	//	adminGroup.DELETE("/users/:id", controller.DeleteUser) // 删除用户（管理员权限）
	//}

	// 注册todo相关的路由
	v1Group := r.Group("v1")
	{
		v1Group.GET("/ping", controller.Ping)
		v1Group.POST("/todo", controller.CreateTodo)
		v1Group.GET("/todo", controller.GetTodoList)
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}
	return r
}
