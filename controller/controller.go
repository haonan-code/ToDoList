package controller

import (
	"bubble/models"
	"bubble/services"
	"bubble/utils"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func IndexHandler(c *gin.Context) {

	c.HTML(http.StatusOK, "index.html", nil)

}

// Ping godoc
//
//	@Summary		测试接口
//	@Description	返回 pong
//	@Tags			示例
//	@Produce		json
//	@Success		200	{object}	map[string]string
//	@Router			/ping [get]
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// Register 用户注册
//
//	@Summary		注册一个新用户
//	@Description	接收前端传来的 JSON 用户信息，创建一个新用户
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			user	body		models.user				true	"注册用户信息"
//	@Success		200		{object}	map[string]interface{}	"创建成功返回的结构体"
//	@Failure		400		{object}	map[string]interface{}	"请求参数错误"
//	@Router			/register [post]
func Register(c *gin.Context) {
	var req models.UserRegisterRequest
	// ShouldBindJSON 确保只解析 JSON 请求体，并根据 binding tag 验证
	if err := c.ShouldBindJSON(&req); err != nil {
		// 处理 Gin 的 binding 错误，可以更友好的返回验证失败信息
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			// 如果是 validator 验证错误，可以循环遍历并返回详细信息
			var errorMsgs []string
			for _, e := range errs {
				errorMsgs = append(errorMsgs, e.Field()+" 字段验证失败: "+e.Tag())
			}
			c.JSON(http.StatusBadRequest, models.CommonResponse{
				Status: http.StatusBadRequest,
				Msg:    "请求参数验证失败",
				Error:  strings.Join(errorMsgs, "; "), // 或者拼接 errorMsgs
			})
		} else {
			// 其他类型的 binding 错误，例如 JSON 格式错误
			c.JSON(http.StatusBadRequest, models.CommonResponse{
				Status: http.StatusBadRequest,
				Msg:    "请求参数格式错误或不完整",
				Error:  err.Error(),
			})
		}
		return
	}
	// 调用服务层创建用户
	user, err := services.CreateUser(&req)
	if err != nil {
		// 根据服务层返回的错误类型，设置不同的HTTP状态码和消息
		switch {
		case errors.Is(err, services.ErrUsernameExists):
			c.JSON(http.StatusConflict, models.CommonResponse{ // 409 Conflict
				Status: http.StatusConflict,
				Msg:    "用户名已存在",
				Error:  err.Error(),
			})
		case errors.Is(err, services.ErrEmailExists):
			c.JSON(http.StatusConflict, models.CommonResponse{ // 409 Conflict
				Status: http.StatusConflict,
				Msg:    "邮箱已注册",
				Error:  err.Error(),
			})
		case errors.Is(err, services.ErrPasswordHash), errors.Is(err, services.ErrUserCreate):
			c.JSON(http.StatusInternalServerError, models.CommonResponse{ // 500 Internal Server Error
				Status: http.StatusInternalServerError,
				Msg:    "服务器内部错误，用户创建失败",
				Error:  err.Error(),
			})
		default:
			c.JSON(http.StatusInternalServerError, models.CommonResponse{ // 500 Internal Server Error
				Status: http.StatusInternalServerError,
				Msg:    "未知错误",
				Error:  err.Error(),
			})
		}
		return
	}

	// 注册成功，返回 201 Created
	c.JSON(http.StatusCreated, models.CommonResponse{
		Status: http.StatusCreated,
		Msg:    "注册成功",
		Data: models.UserRegisterResponseData{
			UserID:   user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
	})
}

// Login 用户登录
//
//	@Summary      用户登录
//	@Description   接收用户凭据，验证并返回 JWT 令牌
//	@Tags        Auth
//	@Accept          json
//	@Produce      json
//	@Param       user   body      models.LoginRequest    true   "用户登录凭据"
//	@Success      200       {object}   models.AuthResponse       "登录成功"
//	@Failure      400       {object}   map[string]interface{}   "请求参数错误"
//	@Failure      401       {object}   map[string]interface{}   "用户名或密码错误"
//	@Failure      500       {object}   map[string]interface{}   "服务器内部错误"
//	@Router          /login [post]
func Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"msg":    "请求参数格式错误或不完整",
			"error":  err.Error(),
		})
		return
	}
	// 1. 验证用户凭据
	user, err := services.AuthenticateUser(req.Username, req.Password)
	if err != nil {
		statusCode := http.StatusInternalServerError
		msg := "服务器内部错误"
		if errors.Is(err, services.ErrInvalidCredentials) || errors.Is(err, services.ErrUserNotFound) {
			statusCode = http.StatusUnauthorized // 401 Unauthorized
			msg = "用户名或密码错误"
		}
		c.JSON(statusCode, gin.H{
			"status": statusCode,
			"msg":    msg,
			"error":  err.Error(),
		})
		return
	}
	// 2. 生成 JWT 令牌
	token, err := utils.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    "生成认证令牌失败",
			"error":  err.Error(),
		})
		return
	}
	// 3. 返回成功响应
	c.JSON(http.StatusOK, models.AuthResponse{
		Status: 200,
		Msg:    "登录成功",
		Data: models.LoginResponseData{
			Token:    token,
			Username: user.Username,
			UserID:   user.ID,
		},
	})
}

func GetMyInfo(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	user, err := services.GetMyInfo(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    "获取用户信息失败",
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.CommonResponse{
		Status: 200,
		Msg:    "获取我的信息成功",
		Data: models.MyInfoResponseData{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
	})
}

func UpdateUserInfo(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	var input models.UpdateUserInfo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := services.UpdateUserInfo(userID.(uint), &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "用户信息修改成功",
	})
}

func ChangePassword(c *gin.Context) {
	// 1. 解析请求体中的新旧密码
	var input models.ChangePassword
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数不正确"})
		return
	}
	// 2. 验证新密码和确认密码是否一致
	if input.NewPassword != input.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "两次新密码输入不一致"})
		return
	}
	// 3. 获取 JWT 中的 id 信息
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	// 4. 调用服务层方法更改密码
	err := services.ChangePassword(userID.(uint), &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "密码修改成功",
	})
}

// CreateTodo 创建一个新的待办事项
//
//	@Summary		创建待办事项
//	@Description	接收前端传来的 JSON，创建一个 Todo 项目
//	@Tags			Todo
//	@Accept			json
//	@Produce		json
//	@Param			todo	body		models.Todo				true	"待办事项内容"
//	@Success		200		{object}	map[string]interface{}		"创建成功返回的结构体"
//	@Failure		400		{object}	map[string]interface{}	"请求参数错误"
//	@Router			/todo [post]
func CreateTodo(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var todo models.Todo
	if err := c.ShouldBind(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.CreateATodo(userID.(uint), &todo); err != nil {
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
func GetTodoList(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	// 查询当前用户的所有待办事项
	todoList, err := services.GetAllTodo(userID.(uint))
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
//	@Param			todo	body		models.Todo				true	"待办事项内容"
//	@Success		200		{object}	map[string]interface{}		"修改成功返回的结构体"
//	@Failure		400		{object}	map[string]interface{}	"请求参数错误"
//	@Router			/todo/{id} [put]
func UpdateATodo(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	//id, ok := c.Params.Get("id")
	//if !ok {
	//	c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
	//	return
	//}
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
	var input models.UpdateTodoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todo, err := services.UpdateTodo(userID.(uint), &input)
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
func DeleteATodo(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	if err := services.DeleteATodo(userID.(uint)); err != nil {
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
