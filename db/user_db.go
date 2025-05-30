package db

import (
	"bubble/models"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"log"
)

func CheckUserExistsByUsername(username string) bool {
	err := DB.Model(&models.User{}).Where("username = ?", username).First(&models.User{}).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)

}

func CheckUserExistsByEmail(email string) bool {
	err := DB.Model(&models.User{}).Where("email = ?", email).First(&models.User{}).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

func CreateUser(user *models.User) error {
	result := DB.Create(user)
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
