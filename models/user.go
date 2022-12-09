package models

import (
	"github.com/asaskevich/govalidator"
)

type User struct {
	ID        string `json:"userId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Contact   string `json:"contact"`
}

func Validate(user *User) (bool, error) {
	result, errValidate := govalidator.ValidateStruct(user)
	if errValidate != nil {
		return false, errValidate
	}
	return result, nil
}
