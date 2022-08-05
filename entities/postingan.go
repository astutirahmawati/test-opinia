package entities

import (
	"gorm.io/gorm"
)

type Postingan struct {
	gorm.Model
	Type_id     uint   `json:"type_id"`
	User_id     uint   `json:"user_id"`
	Title       string `json:"title" gorm:"type:varchar(161)"`
	Description string `json:"description" gorm:"type:varchar(161)"`
}
