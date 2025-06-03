package models

// 用于登录请求体
type LoginRequest struct {
	Username string `json:"username" binding:"required"` // binding:"required" 表示该字段是必填的
	Password string `json:"password" binding:"required"`
}

// 用于登录成功响应
type LoginResponseData struct {
	Token    string `json:"token"`
	Username string `json:"username"`
	UserID   uint   `json:"user_id"`
}

type AuthResponse struct {
	Status int               `json:"status"`
	Msg    string            `json:"msg"`
	Data   LoginResponseData `json:"data"`
}

type MyInfoResponseData struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
