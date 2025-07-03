package model

// User 定义用户模型，gorm:"primaryKey" 表示 ID 是主键
type User struct {
	ID       uint   `json:"user_id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"-" gorm:"not null"` // `json:"-"` 表示不序列化到 JSON
	Email    string `json:"email" gorm:"unique;not null"`
	// CreatedAt time.Time `json:"created_at"` // 可选：记录创建时间
	// UpdatedAt time.Time `json:"updated_at"` // 可选：记录更新时间
	Todos []Todo `gorm:"foreignKey:UserID"`
}

// UserRegisterRequest 定义用户注册请求体
type UserRegisterRequest struct {
	Username string `json:"username" binding:"required,min=6,max=30"` // 必填，用户名长度6-30
	Password string `json:"password" binding:"required,min=8"`        // 必填，密码最小8位
	Email    string `json:"email" binding:"required,email"`           // 必填，邮箱格式
}

// 用于用户更新个人信息（可选字段）
type UpdateUserInfo struct {
	Username *string `json:"username,omitempty"`
	Email    *string `json:"email,omitempty"`
}

type ChangePassword struct {
	OldPassword     string `json:"old_password" binding:"required"`     // 必填，旧密码
	NewPassword     string `json:"new_password" binding:"required"`     // 新密码
	ConfirmPassword string `json:"confirm_password" binding:"required"` // 新密码确认（可选，视是否后端校验）
}

// UserRegisterResponseData 用于注册成功响应的数据部分
type UserRegisterResponseData struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// CommonResponse 统一的成功/失败响应结构体
type CommonResponse struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data,omitempty"`  // Data 字段可选，omitempty表示如果为空则不序列化
	Error  string      `json:"error,omitempty"` // 错误信息
}

type MyInfoResponseData struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
