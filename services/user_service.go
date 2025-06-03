package services

import (
	"bubble/db"
	"bubble/models"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// 定义一些业务错误
var (
	ErrUsernameExists = errors.New("用户名已存在")
	ErrEmailExists    = errors.New("邮箱已注册")
	ErrPasswordHash   = errors.New("密码哈希失败")
	ErrUserCreate     = errors.New("用户创建失败")
)

// CreateUser 在数据库中创建新用户
func CreateUser(userReq *models.UserRegisterRequest) (*models.User, error) {
	// 1. 检查用户名是否已存在 (这里需要与你的数据库层交互)
	if db.CheckUserExistsByUsername(userReq.Username) {
		return nil, ErrUsernameExists
	}
	if db.CheckUserExistsByEmail(userReq.Email) {
		return nil, ErrEmailExists
	}
	// 2. 哈希密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, ErrPasswordHash
	}
	// 3. 构建用户模型
	newUser := &models.User{
		Username: userReq.Username,
		Password: string(hashedPassword),
		Email:    userReq.Email,
	}
	// 4. 将用户保存到数据库 (这里需要与你的数据库层交互)
	if err := db.CreateUser(newUser); err != nil {
		return nil, ErrUserCreate
	}
	return newUser, nil
}

// GetMyInfo 查询当前用户的信息
func GetMyInfo(userID uint) (*models.User, error) {
	user, err := db.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUserInfo(id string, input *models.UpdateUserInfo) (models.User, error) {
	var user models.User
	if err := db.DB.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, errors.New("用户未找到")
		}
		// 其他查询错误（连接、权限等）
		return user, err
	}

	// 使用结构体 Updates：只更新非 nil 字段
	if err := db.DB.Model(&user).Updates(input).Error; err != nil {
		return user, err
	}
	return user, nil
}
