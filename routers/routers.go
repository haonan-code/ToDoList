package routers

import (
	"bubble/controller"
	"bubble/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 告诉gin框架模板文件引用的静态文件去哪里找
	r.Static("/static", "static")
	// 告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")

	r.GET("/", controller.IndexHandler)
	r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 注册用户管理相关的路由
	userGroup := r.Group("/auth")
	{
		// 公共接口（无需认证）
		userGroup.POST("/register", controller.Register) // 用户注册
		userGroup.POST("/login", controller.Login)       // 用户登录

		// 需要认证的接口
		authRequired := userGroup.Group("")
		authRequired.Use(middleware.JWTAuthMiddleware())
		{
			authRequired.GET("/me", controller.GetMyInfo)             // 获取当前用户
			authRequired.PUT("/me", controller.UpdateUserInfo)        // 更新当前用户信息
			authRequired.PUT("/changepwd", controller.ChangePassword) //  修改密码
		}
	}

	// TODO 添加用户管理模块
	//adminGroup := r.Group("/admin")
	//{
	//	adminGroup.GET("/users", controller.GetUserList)       // 查询所有用户（管理端）
	//	adminGroup.GET("/users/:id", controller.GetUserInfo)   // 查看单个用户信息（管理端）
	//	adminGroup.DELETE("/users/:id", controller.DeleteUser) // 删除用户（管理员权限）
	//}

	// 注册待办事项相关的路由
	v1Group := r.Group("v1")
	v1Group.Use(middleware.JWTAuthMiddleware()) // 添加JWT认证中间件
	{
		v1Group.POST("/todo", controller.CreateTodo)
		v1Group.GET("/todo", controller.GetTodoList)
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}
	return r
}
