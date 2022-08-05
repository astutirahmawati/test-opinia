package postingan

import "opinia/entities"

type PostinganRepositoryInterface interface {
	InsertPost(newPost entities.Postingan) (entities.Postingan, error)
	GetPostID(ID uint) (entities.Postingan, error)
	UpdatePost(ID uint, update entities.Postingan) (entities.Postingan, error)
	DeletePost(ID uint) error
}
