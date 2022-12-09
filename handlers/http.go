package handlers

import "auth-service/models"

// CreateUserRequest holds the request parameters for the Create User method.
type CreateUserRequest struct {
	user models.User
}

// CreateUserResponse holds the response values for the Create User method.
type CreateUserResponse struct {
	Error   bool        `json:"error,omitEmpty"`
	Message string      `json:"message,omitEmpty"`
	Data    interface{} `json:"data,omitEmpty"`
}

type LoginRequest struct {
	email    string
	passowrd string
}
