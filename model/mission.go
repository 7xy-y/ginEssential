package model

import (
	"github.com/jinzhu/gorm"
)

type Mission struct {
	gorm.Model
	File     string `gorm:"varchar(100);not null"`
	Tag      string `gorm:"varchar(45);not null"`
	Username string `gorm:"varchar(45);not null"`
}
