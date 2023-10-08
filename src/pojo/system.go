package pojo

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name string `json:"name" gorm:"name"`
}
