package users

import "gorm.io/gorm"

type User struct {
	gorm.Model
	DisplayName string `gorm:"size:128"`
	Email       string `gorm:"uniqueIndex"`
	Password    string
}
