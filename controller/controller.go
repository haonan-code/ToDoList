package controller

import (
	"bubble/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
 url --> controller  --> logic   --> model
请求  --> 控制器      --> 业务逻辑 --> 模型层的增删改查
*/
/*
控制器层：控制器层负责处理HTTP请求并进行业务逻辑处理。它通常会从请求中获取参数、
调用服务层进行数据操作、对返回的结果进行封装后返回给客户端。
*/

func IndexHandler(c *gin.Context) {

	c.HTML(http.StatusOK, "index.html", nil)

}

// Ping godoc
// @Summary 测试接口
// @Description 返回 pong
// @Tags 示例
// @Produce json
// @Success 200 {object} map[string]string
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// CreateTodo 创建一个新的待办事项
// @Summary 创建待办事项
// @Description 接收前端传来的 JSON，创建一个 Todo 项目
// @Tags Todo
// @Accept json
// @Produce json
// @Param todo body models.Todo true "待办事项内容"
// @Success 200 {object} models.TodoResponse "创建成功返回的结构体"
// @Failure 400 {object} models.ErrorResponse "请求参数错误"
// @Router /todo [post]
func CreateTodo(c *gin.Context) {
	// 前端页面填写待办事项 点击请求 会发请求到这里
	// 1. 从请求中把数据拿出来
	var todo models.Todo
	// BindJSON()用于从请求中获取JSON数据并将其绑定到指定的Go结构体变量&todo上
	if err := c.ShouldBind(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 2. 存入数据库
	if err := models.CreateATodo(&todo); err != nil {
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
// @Summary 查询所有待办事项
// @Description 返回给前端所有的 Todo 项目
// @Tags Todo
// @Produce json
// @Success 200 {object} models.TodoResponse "返回所有待办事项"
// @Failure 400 {object} models.ErrorResponse "请求参数错误"
// @Router /todo [get]
func GetTodoList(c *gin.Context) {
	// 查询todo这个表里的所有数据
	todoList, err := models.GetAllTodo()
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
// @Summary 修改待办事项
// @Description 根据 ID 更新待办事项的内容
// @Tags Todo
// @Accept json
// @Produce json
// @Param id path int true "待办事项id"
// @Param todo body models.Todo true "待办事项内容"
// @Success 200 {object} models.TodoResponse "修改成功返回的结构体"
// @Failure 400 {object} models.ErrorResponse "请求参数错误"
// @Router /todo/{id} [put]
func UpdateATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	todo, err := models.GetATodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.BindJSON(&todo)
	if err = models.UpdateATodo(todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		//c.JSON(http.StatusOK, todo)
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"msg":    "success",
			"data":   todo,
		})
	}
}

// DeleteATodo 删除一个待办事项
// @Summary 删除待办事项
// @Description 根据 ID 删除待办事项
// @Tags Todo
// @Accept json
// @Produce json
// @Param id path int true "待办事项id"
// @Success 200 {object} models.TodoResponse "删除成功返回的结构体"
// @Failure 400 {object} models.ErrorResponse "请求参数错误"
// @Router /todo/{id} [delete]
func DeleteATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	if err := models.DeleteATodo(id); err != nil {
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
