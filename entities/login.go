package entities

type LoginRequest struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

type LoginResponse struct {
	UserID uint   `json:"user_id"`
	Name   string `json:"name"`
	Role   string `json:"role"`
	Token  string `json:"token"`
}
