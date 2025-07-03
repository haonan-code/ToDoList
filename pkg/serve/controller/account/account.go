package account

import (
	models2 "bubble/internal/model"
	"bubble/internal/utils"
	service "bubble/pkg/serve/service/account"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

func RegisterAcc(c *gin.Context) {
	var req models2.UserRegisterRequest
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
			c.JSON(http.StatusBadRequest, models2.CommonResponse{
				Status: http.StatusBadRequest,
				Msg:    "请求参数验证失败",
				Error:  strings.Join(errorMsgs, "; "), // 或者拼接 errorMsgs
			})
		} else {
			// 其他类型的 binding 错误，例如 JSON 格式错误
			c.JSON(http.StatusBadRequest, models2.CommonResponse{
				Status: http.StatusBadRequest,
				Msg:    "请求参数格式错误或不完整",
				Error:  err.Error(),
			})
		}
		return
	}
	// 调用服务层创建用户
	user, err := service.CreateUser(&req)
	if err != nil {
		// 根据服务层返回的错误类型，设置不同的HTTP状态码和消息
		switch {
		case errors.Is(err, service.ErrUsernameExists):
			c.JSON(http.StatusConflict, models2.CommonResponse{ // 409 Conflict
				Status: http.StatusConflict,
				Msg:    "用户名已存在",
				Error:  err.Error(),
			})
		case errors.Is(err, service.ErrEmailExists):
			c.JSON(http.StatusConflict, models2.CommonResponse{ // 409 Conflict
				Status: http.StatusConflict,
				Msg:    "邮箱已注册",
				Error:  err.Error(),
			})
		case errors.Is(err, service.ErrPasswordHash), errors.Is(err, service.ErrUserCreate):
			c.JSON(http.StatusInternalServerError, models2.CommonResponse{ // 500 Internal Server Error
				Status: http.StatusInternalServerError,
				Msg:    "服务器内部错误，用户创建失败",
				Error:  err.Error(),
			})
		default:
			c.JSON(http.StatusInternalServerError, models2.CommonResponse{ // 500 Internal Server Error
				Status: http.StatusInternalServerError,
				Msg:    "未知错误",
				Error:  err.Error(),
			})
		}
		return
	}

	// 注册成功，返回 201 Created
	c.JSON(http.StatusCreated, models2.CommonResponse{
		Status: http.StatusCreated,
		Msg:    "注册成功",
		Data: models2.UserRegisterResponseData{
			UserID:   user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
	})
}

func LoginAcc(c *gin.Context) {
	var req models2.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"msg":    "请求参数格式错误或不完整",
			"error":  err.Error(),
		})
		return
	}
	// 1. 验证用户凭据
	user, err := service.AuthenticateUser(req.Username, req.Password)
	if err != nil {
		statusCode := http.StatusInternalServerError
		msg := "服务器内部错误"
		if errors.Is(err, service.ErrInvalidCredentials) || errors.Is(err, service.ErrUserNotFound) {
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
	c.JSON(http.StatusOK, models2.AuthResponse{
		Status: 200,
		Msg:    "登录成功",
		Data: models2.LoginResponseData{
			Token:    token,
			Username: user.Username,
			UserID:   user.ID,
		},
	})
}

func GetAccount(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	user, err := service.GetMyInfo(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    "获取用户信息失败",
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models2.CommonResponse{
		Status: 200,
		Msg:    "获取我的信息成功",
		Data: models2.MyInfoResponseData{
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
	var input models2.UpdateUserInfo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := service.UpdateUserInfo(userID.(uint), &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "用户信息修改成功",
	})
}

func ResetPassword(c *gin.Context) {
	// 1. 解析请求体中的新旧密码
	var input models2.ChangePassword
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
	err := service.ChangePassword(userID.(uint), &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "密码修改成功",
	})
}

//func LogoutAccount(c *gin.Context) error {
//	if err := service.LogoutAcc(c); err != nil {
//		return c.JSON(http.StatusInternalServerError, vo.Fail(c, err, bizErr.New(bizErr.SERVER_ERR, err.Error())))
//	}
//
//	return c.JSON(http.StatusOK, vo.Success(c, "用户注销成功"))
//}
