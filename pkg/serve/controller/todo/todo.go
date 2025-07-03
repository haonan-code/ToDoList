package todo

import (
	"bubble/internal/model"
	service "bubble/pkg/serve/service/todo"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateTodo 创建一个新的待办事项
//
//	@Summary		创建待办事项
//	@Description	接收前端传来的 JSON，创建一个 Todo 项目
//	@Tags			Todo
//	@Accept			json
//	@Produce		json
//	@Param			todo	body		model.Todo				true	"待办事项内容"
//	@Success		200		{object}	map[string]interface{}		"创建成功返回的结构体"
//	@Failure		400		{object}	map[string]interface{}	"请求参数错误"
//	@Router			/todo [post]
func CreateOneTodo(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var todo model.Todo
	if err := c.ShouldBind(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := service.CreateATodo(userID.(uint), &todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		// 直接返回结构体todo，返回的格式与定义的结构体格式一致
		//c.JSON(http.StatusOK, todo)
		// 返回自定义构建的json结构体
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"msg":    "success",
			"data":   todo,
		})
	}
}

// GetTodoList 查询所有待办事项
//
//	@Summary		查询所有待办事项
//	@Description	返回给前端所有的 Todo 项目
//	@Tags			Todo
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}	"返回所有待办事项"
//	@Failure		400	{object}	map[string]interface{}	"请求参数错误"
//	@Router			/todo [get]
func GetALLTodos(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	// 查询当前用户的所有待办事项
	todoList, err := service.GetAllTodo(userID.(uint))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		//c.JSON(http.StatusOK, todoList)
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"msg":    "success",
			"data":   todoList,
		})
	}
}

// UpdateATodo 修改一个待办事项
//
//	@Summary		修改待办事项
//	@Description	根据 ID 更新待办事项的内容
//	@Tags			Todo
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int						true	"待办事项id"
//	@Param			todo	body		model.Todo				true	"待办事项内容"
//	@Success		200		{object}	map[string]interface{}		"修改成功返回的结构体"
//	@Failure		400		{object}	map[string]interface{}	"请求参数错误"
//	@Router			/todo/{id} [put]
func UpdateOneTodo(c *gin.Context) {
	_, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	//todo, err := services.GetATodo(id)
	//if err != nil {
	//	c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	//	return
	//}
	//if err = c.BindJSON(&todo); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}
	//if err = services.UpdateATodo(todo); err != nil {
	//	c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	//} else {
	//	//c.JSON(http.StatusOK, todo)
	//	c.JSON(http.StatusOK, gin.H{
	//		"status": 200,
	//		"msg":    "success",
	//		"data":   todo,
	//	})
	//}
	//id, ok := c.Params.Get("id")
	//if !ok {
	//	c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
	//	return
	//}
	var input model.UpdateTodoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todo, err := service.UpdateTodo(id, &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "success",
		"data":   todo,
	})

}

// DeleteATodo 删除一个待办事项
//
//	@Summary		删除待办事项
//	@Description	根据 ID 删除待办事项
//	@Tags			Todo
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int						true	"待办事项id"
//	@Success		200	{object}	map[string]interface{}		"删除成功返回的结构体"
//	@Failure		400	{object}	map[string]interface{}	"请求参数错误"
//	@Router			/todo/{id} [delete]
func DeleteOneTodo(c *gin.Context) {
	_, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	if err := service.DeleteATodo(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		//c.JSON(http.StatusOK, gin.H{id: "deleted"})
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"msg":    "success",
			"data":   struct{ ID string }{ID: "deleted"},
		})
	}
}
