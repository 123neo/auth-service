package handlers

import (
	"auth-service/config"
	"auth-service/models"
	"auth-service/services"
	"context"

	"github.com/asaskevich/govalidator"
	"go.uber.org/zap"
)

// CreateUserRequest holds the request parameters for the Create User method.
type CreateUserRequest struct {
	user *models.User
}

// CreateUserResponse holds the response values for the Create User method.
type CreateUserResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type LoginRequest struct {
	email    string `valid:"required"`
	password string `valid:"required"`
}

func ValidateLoginRequest(ctx context.Context, app *config.Config, loginPayload LoginRequest) bool {

	service := services.NewService(app.Repo, app.Log)
	count, err := service.GetUserByEmail(ctx, loginPayload.email)
	if err != nil {
		app.Log.Fatal("Get User by Email:", zap.Error(err))
	}

	if count == 0 {
		return false
	}

	if !govalidator.IsEmail(loginPayload.email) {
		return false
	}

	if len(loginPayload.password) == 0 {
		return false
	}

	return true
}
