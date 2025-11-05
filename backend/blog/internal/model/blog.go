package model

import "time"

type Blog struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Title      string    `json:"title"`
	Content    string    `gorm:"type:longtext" json:"content"`
	UploadDate time.Time `gorm:"column:upload_date" json:"uploadDate"`
	UpdateDate time.Time `gorm:"column:update_date" json:"updateDate"`
}

// TableName aligns with existing Java schema (table name: blog)
func (Blog) TableName() string { return "blog" }
