package users

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID          int64
	DisplayName string `gorm:"size:128"`
	Email       string `gorm:"uniqueIndex"`
	Password    string
	Avatar      string `gorm:"size:256"`
}
