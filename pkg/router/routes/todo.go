package routes

import (
	"bubble/internal/middleware"
	"bubble/pkg/serve/controller/todo"
	"github.com/gin-gonic/gin"
)

func RegisterTodoRoutes(r ...*gin.RouterGroup) {
	apiv1 := r[0]
	todoGroupV1 := apiv1.Group("/todo")
	//todoGroupV1.GET("getOneTodo", todo.GetOneTodo)
	todoGroupV1.GET("getAllTodos", todo.GetALLTodos)

	authRequiredGroup := todoGroupV1.Group("")
	authRequiredGroup.Use(middleware.JWTAuthMiddleware())
	{
		authRequiredGroup.POST("/createOneTodo", todo.CreateOneTodo)
		authRequiredGroup.PUT("/updateOneTodo", todo.UpdateOneTodo)
		authRequiredGroup.DELETE("/deleteOneTodo", todo.DeleteOneTodo)
	}

}
