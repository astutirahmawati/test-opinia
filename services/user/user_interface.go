package user

import (
	"opinia/entities"
)

type UserServiceInterface interface {
	CreateUser(internalRequest entities.CreateUserRequest) (entities.User, error)
	GetbyID(id uint) (entities.CustomerResponse, error)
	UpdateCustomer(id uint, customerRequest entities.UpdateCustomerRequest) (entities.CustomerResponse, error)
	DeleteCustomer(id uint) error
}
