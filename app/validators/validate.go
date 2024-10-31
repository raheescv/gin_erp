// app/validators/validate.go
package validators

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct(data interface{}) []string {
	var errors []string
	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, err.Error())
		}
	}
	return errors
}
