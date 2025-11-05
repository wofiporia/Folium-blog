package db

import (
	"log"
	"os"

	"github.com/wofiporia/folium-backend-new/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		// 本地直跑默认连本机 MySQL（root/123456），数据库名 blog
		// Java 驱动使用 allowPublicKeyRetrieval，Go 驱动不支持该参数；
		// 若遇到认证问题，可为用户启用 mysql_native_password 或设置 allowNativePasswords=true。
		dsn = "root:123456@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local&allowNativePasswords=true"
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
