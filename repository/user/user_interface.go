package user

import "opinia/entities"

type UserRepositoryInterface interface {
	InsertUser(newUser entities.User) (entities.User, error)
	GetUserID(ID uint) (entities.User, error)
	UpdateUser(ID uint, update entities.User) (entities.User, error)
	DeleteUser(ID uint) error
	FindByUser(value string) (entities.User, error)
	Login(email string, password string) (entities.User, error)
}
