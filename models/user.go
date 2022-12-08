package models

import (
	"github.com/asaskevich/govalidator"
)

type User struct {
	ID        string `json: "userId" valid:"-"`
	FirstName string `json: "firstName" valid:"string"`
	LastName  string `json: "lastName" valid:"string"`
	Email     string `json: "email" valid:"string"`
	Contact   string `json: "contact" valid:"string"`
}

func Validate(user User) (bool, error) {
	// govalidator.SetFieldsRequiredByDefault(true)
	result, errValidate := govalidator.ValidateStruct(user)
	if errValidate != nil {
		return false, errValidate
	}
	return result, nil
}
