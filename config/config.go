package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

var DBConfig DatabaseConfig

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("读取配置失败: %w", err))
	}

	err = viper.Sub("database").Unmarshal(&DBConfig)
	if err != nil {
		panic(fmt.Errorf("解析数据库配置失败: %w", err))
	}
}
