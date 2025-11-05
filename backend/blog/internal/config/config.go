package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config 应用配置结构体
type Config struct {
	AppPort    string
	JWTSecret  string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

// LoadConfig 加载环境变量配置
func LoadConfig() *Config {
	// 加载项目根目录与后端目录的环境文件（本地优先，不存在则忽略错误）
	_ = godotenv.Load("../../.env.local")
	_ = godotenv.Load("../../.env")
	_ = godotenv.Load(".env.local")
	if err := godotenv.Load(); err != nil {
		// 没有 .env 时不报错，仅提示一次
		if !os.IsNotExist(err) {
			log.Println("warn: load .env failed:", err)
		}
	}

	return &Config{
		AppPort:    getEnv("APP_PORT", "8081"),
		JWTSecret:  getEnv("JWT_SECRET", "WfP@2025!Blog#Secure*Key&Salt123456789"),
		DBHost:     getEnv("DB_HOST", "127.0.0.1"),
		DBPort:     getEnv("DB_PORT", "3306"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "123456"),
		DBName:     getEnv("DB_NAME", "blog"),
	}
}

// getEnv 获取环境变量，如果不存在则返回默认值
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
