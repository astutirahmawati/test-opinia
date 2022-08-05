package auth

import (
	"opinia/delivery/helpers"
	middleware "opinia/delivery/middlewares"
	"opinia/entities"
	userRepository "opinia/repository/user"
)

type AuthService struct {
	userRepo userRepository.UserRepositoryInterface
}

func NewAuthService(userRepo userRepository.UserRepositoryInterface) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

/*
 * Auth Service - Login
 * -------------------------------
 * Mencari user berdasarkan ID
 */
func (as AuthService) Login(authReq entities.AuthRequest) (entities.LoginResponse, error) {

	// Get user by username via repository
	user, err := as.userRepo.FindByUser(authReq.Email)
	if err != nil {
		return entities.LoginResponse{}, err
	}

	// Verify password
	if !helpers.CheckPasswordHash(authReq.Password, user.Password) {
		return entities.LoginResponse{}, err
	}

	if user.Role != "customer" {

		// Create token
		token, err := middleware.CreateToken(int(user.ID), user.Name, user.Role)
		if err != nil {
			return entities.LoginResponse{}, err
		}
		response := entities.LoginResponse{UserID: user.ID, Name: user.Name, Token: token, Role: user.Role}
		return response, nil
	}

	// Create token
	token, err := middleware.CreateToken(int(user.ID), user.Name, "customer")
	if err != nil {
		return entities.LoginResponse{}, err
	}

	response := entities.LoginResponse{UserID: user.ID, Name: user.Name, Token: token, Role: user.Role}
	return response, nil
}
