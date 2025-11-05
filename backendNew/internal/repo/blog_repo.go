package repo

import (
	"errors"

	"github.com/wofiporia/folium-backend-new/internal/db"
	"github.com/wofiporia/folium-backend-new/internal/model"
	"gorm.io/gorm"
)

func FindAllBlogs() ([]model.Blog, error) {
	var blogs []model.Blog
	if err := db.DB.Order("upload_date desc").Find(&blogs).Error; err != nil {
		return nil, err
	}
	return blogs, nil
}

func FindBlogByID(id uint) (*model.Blog, error) {
	var b model.Blog
	if err := db.DB.First(&b, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &b, nil
}

// CreateBlog 创建博客
func CreateBlog(blog *model.Blog) error {
	return db.DB.Create(blog).Error
}

// UpdateBlog 更新博客
func UpdateBlog(blog *model.Blog) error {
	return db.DB.Save(blog).Error
}

// DeleteBlog 删除博客
func DeleteBlog(id uint) error {
	return db.DB.Delete(&model.Blog{}, id).Error
}
