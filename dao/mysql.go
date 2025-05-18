package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var (
	DB *gorm.DB
)

func InitMySQL() (err error) {
	dsn := "root:20000406@(localhost)/todolist_db?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}
	// 测试与数据库的连接是否仍然 存在
	// DB()返回一个sql.DB类型的指针
	// Ping()方法用于测试与数据库的连接是否仍然存在
	return DB.DB().Ping()
}

func Close() {
	DB.Close()
}
