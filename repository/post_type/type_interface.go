package post_type

import "opinia/entities"

type TypeRepositoryInterface interface {
	InsertType(newType entities.PostType) (entities.PostType, error)
	GetTypeID(ID uint) (entities.PostType, error)
	UpdateType(ID uint, update entities.PostType) (entities.PostType, error)
	DeleteType(ID uint) error
}
