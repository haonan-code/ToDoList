package mapper

import (
	"bubble/internal/global"
	"bubble/internal/model"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"log"
)

func CheckUserExistsByUsername(username string) bool {
	err := global.DB.Model(&model.User{}).Where("username = ?", username).First(&model.User{}).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)

}

func CheckUserExistsByEmail(email string) bool {
	err := global.DB.Model(&model.User{}).Where("email = ?", email).First(&model.User{}).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

func CreateUser(user *model.User) error {
	result := global.DB.Create(user)
	if result.Error != nil {
		if mysqlErr, ok := result.Error.(*mysql.MySQLError); ok {
			// 处理特定的 MySQL 错误
			switch mysqlErr.Number {
			case 1062: // 唯一键错误
				return fmt.Errorf("用户名或邮箱已存在")
			default:
				return fmt.Errorf("数据库错误: %w", mysqlErr)
			}
		}
		return fmt.Errorf("创建用户失败: %w", result.Error)
	}

	// 可以添加创建成功的日志
	log.Printf("成功创建用户: %s", user.Username)
	return nil
}

func GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	err := global.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("用户 %s 不存在", username)
		}
		return nil, fmt.Errorf("查询用户失败: %w", err)
	}
	return &user, nil
}

func GetUserByID(userID uint) (*model.User, error) {
	var user model.User
	err := global.DB.First(&user, userID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("用户 ID %d 不存在", userID)
		}
		return nil, fmt.Errorf("查询用户失败: %w", err)
	}
	return &user, nil
}
