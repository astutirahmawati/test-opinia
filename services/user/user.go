package user

import (
	"opinia/delivery/helpers"
	"opinia/entities"
	userRepository "opinia/repository/user"

	"github.com/jinzhu/copier"
)

type UserService struct {
	userRepo userRepository.UserRepositoryInterface
}

func NewUserService(repository userRepository.UserRepositoryInterface) *UserService {
	return &UserService{
		userRepo: repository,
	}
}

func (us *UserService) CreateUser(internalRequest entities.CreateUserRequest) (entities.User, error) {

	user := entities.User{}
	copier.Copy(&user, &internalRequest)

	// Insert ke sistem melewati repository
	respond, err := us.userRepo.InsertUser(user)
	if err != nil {
		return user, err
	}

	return respond, nil
}

func (us *UserService) GetbyID(id uint) (entities.CustomerResponse, error) {

	user, err := us.userRepo.GetUserID(id)
	if err != nil {
		return entities.CustomerResponse{}, err
	}
	userRes := entities.CustomerResponse{}
	copier.Copy(&userRes, &user)

	return userRes, err
}

func (us *UserService) UpdateCustomer(id uint, customerRequest entities.UpdateCustomerRequest) (entities.CustomerResponse, error) {

	// Get user by ID via repository
	user, err := us.userRepo.GetUserID(id)
	if err != nil {
		return entities.CustomerResponse{}, err
	}

	copier.CopyWithOption(&user, &customerRequest, copier.Option{IgnoreEmpty: true, DeepCopy: true})

	// Hanya hash password jika password juga diganti (tidak kosong)
	if customerRequest.Password != "" {
		hashedPassword, _ := helpers.HashPassword(user.Password)
		user.Password = hashedPassword
	}

	// Update via repository
	user, err = us.userRepo.UpdateUser(id, user)
	// Konversi user domain menjadi user response
	userRes := entities.CustomerResponse{}
	copier.Copy(&userRes, &user)

	return userRes, err
}

func (us *UserService) DeleteCustomer(id uint) error {

	// Delete via repository
	err := us.userRepo.DeleteUser(id)
	return err
}
