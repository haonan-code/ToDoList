package account

import (
	"bubble/internal/global"
	"bubble/internal/model"
	"bubble/pkg/serve/mapper"
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

var (
	ErrInvalidCredentials  = errors.New("用户名或密码错误")
	ErrUserNotFound        = errors.New("用户不存在")
	ErrGenerateTokenFailed = errors.New("生成认证令牌失败")
)

// CreateUser 在数据库中创建新用户
func CreateUser(userReq *model.UserRegisterRequest) (*model.User, error) {
	// 1. 检查用户名是否已存在 (这里需要与你的数据库层交互)
	if mapper.CheckUserExistsByUsername(userReq.Username) {
		return nil, ErrUsernameExists
	}
	if mapper.CheckUserExistsByEmail(userReq.Email) {
		return nil, ErrEmailExists
	}
	// 2. 哈希密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, ErrPasswordHash
	}
	// 3. 构建用户模型
	newUser := &model.User{
		Username: userReq.Username,
		Password: string(hashedPassword),
		Email:    userReq.Email,
	}
	// 4. 将用户保存到数据库 (这里需要与你的数据库层交互)
	if err := mapper.CreateUser(newUser); err != nil {
		return nil, ErrUserCreate
	}
	return newUser, nil
}

func AuthenticateUser(username, password string) (*model.User, error) {
	// 1. 使用 gorm 从数据库中查询用户数据
	user, err := mapper.GetUserByUsername(username)
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

func GetMyInfo(userID uint) (*model.User, error) {
	user, err := mapper.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUserInfo(id uint, input *model.UpdateUserInfo) error {
	var user model.User

	//// 1.先查再改
	//if err := db.DB.First(&user, id).Error; err != nil {
	//	if errors.Is(err, gorm.ErrRecordNotFound) {
	//		return user, errors.New("用户未找到")
	//	}
	//	// 其他查询错误（连接、权限等）
	//	return user, err
	//}
	//// 使用结构体 Updates：只更新非 nil 字段
	//if err := db.DB.Model(&user).Updates(input).Error; err != nil {
	//	return user, err
	//}
	//return user, nil

	// 2.直接改
	if err := global.DB.Model(&user).Where("id = ?", id).Updates(input).Error; err != nil {
		return err
	}
	return nil
}

func ChangePassword(id uint, input *model.ChangePassword) error {
	var user model.User
	if err := global.DB.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("用户不存在")
		}
		return err
	}
	// 验证旧密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.OldPassword)); err != nil {
		return errors.New("旧密码不正确")
	}
	// 检查新密码与旧密码是否相同
	if input.OldPassword == input.NewPassword {
		return errors.New("新密码不能和旧密码相同")
	}
	// 哈希新密码并保存
	newHashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("新密码哈希失败")
	}

	user.Password = string(newHashedPassword)
	if err := global.DB.Save(&user).Error; err != nil {
		return ErrUserCreate
	}
	return nil
}
