package services

import (
	"bubble/models"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials  = error.New("用户名或密码错误")
	ErrUserNotFound        = error.New("用户不存在")
	ErrGenerateTokenFailed = error.New("生成认证令牌失败")
)

// AuthenticateUser 验证用户凭据
func AuthenticateUser(username, password string) (*models.User, error) {
	// TODO 从数据库查询用户
	// TODO 使用 gorm 从数据库中查询用户
	user := &models.User{}

	if username == "testuser" {
		user.ID = 1
		user.Username = "testuser"
		hashedPassword, _ := bcrype.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
		user.Password = string(hashedPassword)
		user.Email = "test@example.com"
	} else {
		return nil, ErrUserNotFound
	}

	// 验证密码
	// CompareHashAndPassword 接收哈希后的密码和明文密码
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, ErrInvalidCredentials // 密码不匹配
	}
	return user, nil
}
