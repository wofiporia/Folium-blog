package db

import (
	"fmt"
	"log"
	"os"

	"github.com/wofiporia/folium-backend-new/internal/config"
	"github.com/wofiporia/folium-backend-new/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	// 1) 优先使用完整 DSN（便于一次性覆盖）
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		// 2) 使用分段变量（来自根 .env/.env.local），由 config 统一管理
		cfg := config.LoadConfig()
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&allowNativePasswords=true",
			cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
		)
	}
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB = conn

	// Auto migrate models
	err = DB.AutoMigrate(&model.Blog{}, &model.Admin{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// 创建默认管理员账户（如果不存在）
	createDefaultAdmin()
}

// createDefaultAdmin 创建默认管理员账户
func createDefaultAdmin() {
	var count int64
	DB.Model(&model.Admin{}).Count(&count)

	if count == 0 {
		admin := &model.Admin{
			Username: "admin",
			Password: "admin123",
		}

		if err := DB.Create(admin).Error; err != nil {
			log.Printf("创建默认管理员失败: %v", err)
		} else {
			log.Println("已创建默认管理员账户: admin/admin123")
		}
	}
}
