package repo

import (
	"errors"

	"github.com/wofiporia/folium-backend-new/internal/db"
	"github.com/wofiporia/folium-backend-new/internal/model"
	"gorm.io/gorm"
)

// FindAdminByUsername 根据用户名查找管理员
func FindAdminByUsername(username string) (*model.Admin, error) {
	var admin model.Admin
	if err := db.DB.Where("username = ?", username).First(&admin).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // 用户不存在
		}
		return nil, err
	}
	return &admin, nil
}

// CreateAdmin 创建管理员（用于初始化）
func CreateAdmin(username, password string) (*model.Admin, error) {
	admin := &model.Admin{
		Username: username,
		Password: password,
	}
	if err := db.DB.Create(admin).Error; err != nil {
		return nil, err
	}
	return admin, nil
}

// CheckAdminExists 检查管理员是否存在
func CheckAdminExists(username string) (bool, error) {
	var count int64
	if err := db.DB.Model(&model.Admin{}).Where("username = ?", username).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

