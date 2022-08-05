package entities

type AuthRequest struct {
	Email    string `form:"email" json:"email" validate:"required"`
	Password string `form:"password" json:"password" validate:"required"`
}

type CustomerAuthResponse struct {
	Token string           `json:"token"`
	User  CustomerResponse `json:"user"`
}

type InternalAuthResponse struct {
	Token string           `json:"token"`
	User  InternalResponse `json:"user"`
}
