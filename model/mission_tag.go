package model

import (
	"github.com/jinzhu/gorm"
)

type Mission_tag struct {
	gorm.Model
	Tag        string `gorm:"mediumtext;not null"`
	Mission_ID string `gorm:"varchar(45);not null"`
	Publisher  string `gorm:"varchar(45);not null"`
	Solver     string `gorm:"varchar(45);not null"`
	What       string `gorm:"varchar(45);not null"`
}
