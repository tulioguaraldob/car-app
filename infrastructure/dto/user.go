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
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type Credentials struct {
	Username string `json:"username"`
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

func RequestToUser(req *UserRequest) *entity.User {
	return &entity.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Username:  req.Username,
		Password:  req.Password,
	}
}

func CredentialsToUser(credentials *Credentials) *entity.User {
	return &entity.User{
		Username: credentials.Username,
		Password: credentials.Password,
	}
}
