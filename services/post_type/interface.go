package post_type

import (
	"opinia/entities"
)

type TypeServiceInterface interface {
	CreateType(user_id uint, request entities.PostType) (entities.PostType, error)
	GetbyID(id uint) (entities.PostType, error)
	UpdateType(type_id uint, request entities.PostType) (entities.PostType, error)
	DeleteType(user_id, id uint) error
}
