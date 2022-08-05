package auth

import (
	"opinia/entities"
)

type AuthServiceInterface interface {
	Login(AuthReq entities.AuthRequest) (entities.LoginResponse, error)
}
