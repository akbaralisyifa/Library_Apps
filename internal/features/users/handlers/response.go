package handlers

import "library/internal/features/users"

type LoginResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role	 string `json:"role"`
	Token    string `json:"token"`
}

func ToLoginResponse(val users.Users, token string) LoginResponse{
	return LoginResponse{
		ID: val.ID,
		Username: val.Username,
		Email: val.Email,
		Role: val.Role,
		Token: token,
	}
}
