package handlers

import (
	"auth-service/config"
	"auth-service/models"
	"encoding/json"
	"io"
	"net/http"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func decodeCreateUserRequest(w http.ResponseWriter, r *http.Request, user *models.User) error {
	maxBytes := 1048576 // one megabyte

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)

	err := dec.Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		msg := "Request body must only contain a single JSON object"
		http.Error(w, msg, http.StatusBadRequest)
	}

	return nil
}

func decodeLoginRequest(w http.ResponseWriter, r *http.Request, loginData *LoginRequest, app *config.Config) error {
	maxBytes := 1048576 // one megabyte

	// buf, err := io.ReadAll(r.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// app.Log.Info("Login Request", zap.Any("login body", buf))

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)

	if err := dec.Decode(&loginData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	// log.Println(*loginData)
	app.Log.Info("Login Request", zap.Any("login body", loginData))

	if err := dec.Decode(&struct{}{}); err != io.EOF {
		msg := "Request body must only contain a single JSON object"
		http.Error(w, msg, http.StatusBadRequest)
	}

	return nil
}

func encodeResponse(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}

func errorJSON(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload CreateUserResponse
	payload.Error = true
	payload.Message = err.Error()

	return encodeResponse(w, statusCode, payload)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
