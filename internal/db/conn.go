package db

import (
	"bubble/configs"
	"bubble/internal/global"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// New 初始化数据库连接
func New() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBConfig.User,
		config.DBConfig.Password,
		config.DBConfig.Host,
		config.DBConfig.Port,
		config.DBConfig.Name,
	)

	var err error
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("数据库连接失败: %w", err))
	}

	log.Println("✅ 数据库连接成功")

	autoMigrate()
}
