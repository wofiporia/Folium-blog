package main

import (
	"fmt"
	"log"

	"github.com/wofiporia/folium-backend-new/internal/db"
	"github.com/wofiporia/folium-backend-new/internal/repo"
)

func main() {
	// 初始化数据库连接
	db.Init()

	// 检查是否已存在管理员
	exists, err := repo.CheckAdminExists("admin")
	if err != nil {
		log.Fatal("检查管理员是否存在时出错:", err)
	}

	if exists {
		fmt.Println("管理员 'admin' 已存在")
		return
	}

	// 创建默认管理员
	admin, err := repo.CreateAdmin("admin", "admin123")
	if err != nil {
		log.Fatal("创建管理员时出错:", err)
	}

	fmt.Printf("成功创建管理员: ID=%d, Username=%s\n", admin.ID, admin.Username)
	fmt.Println("默认用户名: admin")
	fmt.Println("默认密码: admin123")
}

