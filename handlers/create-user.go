package handlers

import (
	"auth-service/config"
	"auth-service/models"
	"auth-service/services"
	"net/http"

	"go.uber.org/zap"
)

func CreateHandlerFunc(app *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()

		// decoding the request payload

		var user models.User

		requestPaylod := CreateUserRequest{
			user: &user,
		}

		if err := decodeCreateUserRequest(w, r, &user); err != nil {
			app.Log.Error("Error in decoding JSON : ", zap.Error(err))
			_ = errorJSON(w, err, http.StatusBadRequest)
			return
		}

		app.Log.Info("User:", zap.Any("user struct", user))

		if _, errValidate := models.Validate(&user); errValidate != nil {
			app.Log.Error("Validate Error", zap.Error(errValidate))
			_ = errorJSON(w, errValidate, http.StatusBadRequest)
			return
		}

		// encrypting password

		if hashString, err := HashPassword(requestPaylod.user.Password); err != nil {
			app.Log.Error("Not able to encrypt password", zap.Error(err))
		} else {
			requestPaylod.user.Password = hashString
		}

		// calling the required service

		service := services.NewService(app.Repo, app.Log)
		repsonse, err := service.CreateUser(ctx, requestPaylod.user)

		if err != nil {
			app.Log.Error("Some error occured: ", zap.Error(err))
			_ = errorJSON(w, err, http.StatusBadRequest)
			return
		}

		// sending the response back to the client

		app.Log.Info("Response: ", zap.Any("response", repsonse))
		payload := CreateUserResponse{
			Data: repsonse,
		}

		err = encodeResponse(w, http.StatusAccepted, payload)

		if err != nil {
			app.Log.Error("Some error occured: ", zap.Error(err))
			_ = errorJSON(w, err, http.StatusBadRequest)
			return
		}
	}
}
