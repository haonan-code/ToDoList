package db

import (
	"bubble/internal/global"
	models2 "bubble/internal/model"
	"log"
)

// autoMigrate 执行数据库表结构自动迁移
func autoMigrate() {
	if global.DB == nil {
		log.Fatal("数据库初始化失败，无法执行自动迁移...")
	}

	err := global.DB.AutoMigrate(&models2.User{}, &models2.Todo{})
	if err != nil {
		log.Fatalf("数据库自动迁移失败: %v", err)
	}

	log.Println("数据库自动迁移成功...")
	//global.SysLog.Infof("数据库自动迁移成功...")
}
