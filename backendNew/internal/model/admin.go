package model

import "time"

type Admin struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"uniqueIndex;not null;size:50" json:"username"`
	Password  string    `gorm:"not null;size:255" json:"-"` // 不返回密码字段
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// TableName sets the table name for the Admin model
func (Admin) TableName() string {
	return "admin"
}
