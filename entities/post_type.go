package entities

import (
	"gorm.io/gorm"
)

type PostType struct {
	gorm.Model
	Jenis string `json:"jenis" gorm:"type:varchar(161)"`
	Type  string `json:"type" gorm:"type:varchar(161)"`
}
