package routes

import (
	"bubble/internal/middleware"
	"bubble/pkg/serve/controller/account"
	"github.com/gin-gonic/gin"
)

func RegisterAccountRoutes(r ...*gin.RouterGroup) {
	apiv1 := r[0]
	accountGroupV1 := apiv1.Group("/account")

	// 无需登录的接口
	accountGroupV1.POST("/register", account.RegisterAcc)
	accountGroupV1.POST("/login", account.LoginAcc)

	// 需要登录的接口统一挂载中间件
	authRequiredGroup := accountGroupV1.Group("")
	authRequiredGroup.Use(middleware.JWTAuthMiddleware())
	{
		authRequiredGroup.GET("/getAccount", account.GetAccount)
		authRequiredGroup.PUT("/updateAcc", account.UpdateUserInfo)
		authRequiredGroup.PUT("/resetPassword", account.ResetPassword)
		//authRequiredGroup.GET("/logout", account.LogoutAccount)
	}
}
