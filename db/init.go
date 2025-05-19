package db

import (
	"bubble/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	DB *gorm.DB
)

func InitMySQL() (err error) {
	cfg := config.GetDBConfig()
	//dsn := "root:20000406@(localhost)/todolist_db?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.DBName)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		return nil
	}
	return sqlDB.Ping()
}

func Close() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Println("获取底层连接失败:", err)
		return
	}
	sqlDB.Close()
}
