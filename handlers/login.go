package handlers

import (
	"auth-service/config"
	"net/http"

	"go.uber.org/zap"
)

func LoginHandlerFunc(app *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// ctx := r.Context()

		var loginData LoginRequest

		if err := decodeLoginRequest(w, r, &loginData); err != nil {
			app.Log.Error("Validate Error", zap.Error(err))
			_ = errorJSON(w, err, http.StatusBadRequest)
		}

		app.Log.Info("Login Request", zap.Any("login body", loginData))

	}
}
