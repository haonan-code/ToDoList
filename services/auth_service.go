package services

import (
	"bubble/db"
	"bubble/models"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials  = errors.New("用户名或密码错误")
	ErrUserNotFound        = errors.New("用户不存在")
	ErrGenerateTokenFailed = errors.New("生成认证令牌失败")
)

// AuthenticateUser 验证用户凭据
func AuthenticateUser(username, password string) (*models.User, error) {
	// 1. 使用 gorm 从数据库中查询用户数据
	user, err := db.GetUserByUsername(username)
	if err != nil {
		return nil, ErrUserNotFound
	}
	// 2. 验证密码
	// CompareHashAndPassword 接收哈希后的密码和明文密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, ErrInvalidCredentials // 密码不匹配
	}
	return user, nil
}
