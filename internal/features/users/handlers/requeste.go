package handlers

import "library/internal/features/users"

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ToModelUser(rr RegisterRequest) users.Users{
	return users.Users{
		Username: rr.Username,
		Email:    rr.Email,
		Password: rr.Password,
	}
}