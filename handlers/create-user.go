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

		err := decodeCreateUserRequest(w, r, &user)

		requestPaylod := CreateUserRequest{
			user: user,
		}

		_, errValidate := models.Validate(&user)

		app.Log.Info("User:", zap.Any("user struct", user))

		if errValidate != nil {
			app.Log.Error("Validate Error", zap.Error(errValidate))
			_ = errorJSON(w, errValidate, http.StatusBadRequest)
			return
		}

		if err != nil {
			app.Log.Error("Error in decoding JSON : ", zap.Error(err))
			_ = errorJSON(w, err, http.StatusBadRequest)
			return
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
