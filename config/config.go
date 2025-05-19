package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type DBConfig struct {
	User     string
	Password string
	Host     string
	DBName   string
}

func LoadEnv() {
	err := godotenv.Load("settings.env")
	if err != nil {
		log.Println("❗ 未找到 .env 文件，使用系统环境变量")
	}
}

func GetDBConfig() DBConfig {
	LoadEnv()
	return DBConfig{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		DBName:   os.Getenv("DB_NAME"),
	}
}
