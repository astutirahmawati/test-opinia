package entities

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Email     string      `json:"email" gorm:"column:email;unique"`
	Name      string      `json:"name" gorm:"type:varchar(161)"`
	Password  string      `json:"password" gorm:"column:password;unique"`
	Gender    string      `json:"gender" gorm:"default:female"`
	Phone     string      `json:"phone" gorm:"type:varchar(161);unique"`
	Role      string      `json:"role" gorm:"type:varchar(161);default:customer"`
	Postingan []Postingan `gorm:"foreingkey:user_id; references: id"`
}

type CreateUserRequest struct {
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required"`
	Name     string `form:"name" validate:"required"`
	Gender   string `form:"gender" validate:"required"`
	Phone    string `form:"phone" validate:"required"`
}

type CustomerResponse struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Gender    string    `json:"gender"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type UpdateCustomerRequest struct {
	Email    string `form:"email"`
	Password string `form:"password"`
	Name     string `form:"name"`
	Phone    string `form:"phone" validate:"required"`
}
type UpdateInternalRequest struct {
	Email    string `form:"email"`
	Password string `form:"password"`
	Name     string `form:"name"`

	Phone string `form:"phone"`
}

type InternalResponse struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Gender    string    `json:"gender"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
