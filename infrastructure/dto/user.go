package dto

import "github.com/TulioGuaraldoB/car-app/domain/entity"

type UserResponse struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type UserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func UserToResponse(user *entity.User) *UserResponse {
	return &UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
}

func CredentialsToUser(credentials *Credentials) *entity.User {
	return &entity.User{
		Email:    credentials.Email,
		Password: credentials.Password,
	}
}
