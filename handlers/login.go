package handlers

import (
	"auth-service/config"
	"net/http"

	"go.uber.org/zap"
)

func LoginHandlerFunc(app *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var loginData LoginRequest

		if err := decodeLoginRequest(w, r, &loginData, app); err != nil {
			app.Log.Error("Decoding to JSON Error", zap.Error(err))
			if errorResponse := errorJSON(w, err, http.StatusBadRequest); errorResponse != nil {
				app.Log.Error("Error in sending error response", zap.Error(errorResponse))
			}
		}

		app.Log.Info("Login Request", zap.Any("login body", loginData))

		response := ValidateLoginRequest(ctx, app, loginData)

		// sending the response back to the client

		app.Log.Info("Response: ", zap.Any("response", response))
		payload := CreateUserResponse{
			Data: "Data Validated",
		}

		err := encodeResponse(w, http.StatusAccepted, payload)

		if err != nil {
			app.Log.Error("Some error occured: ", zap.Error(err))
			_ = errorJSON(w, err, http.StatusBadRequest)
			return
		}

	}
}
