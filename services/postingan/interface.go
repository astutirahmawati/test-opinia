package postingan

import (
	"opinia/entities"
)

type PostServiceInterface interface {
	CreatePost(user_id uint, request entities.Postingan) (entities.Postingan, error)
	GetbyID(id uint) (entities.Postingan, error)
	UpdatePost(user_id uint, post_id uint, request entities.Postingan) (entities.Postingan, error)
	DeletePost(id, user_id uint) error
}
