package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Name string `gorm:"type:varchar(150)"`
}
