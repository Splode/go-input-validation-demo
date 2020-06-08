package lead

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/nyaruka/phonenumbers"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	if err := validate.RegisterValidation("phone", validatePhone); err != nil {
		fmt.Printf("unable to register custom validator: %s\n", err.Error())
	}
}

// Lead represents the lead form POST data.
type Lead struct {
	Name         string   `json:"name" validate:"required,ascii"`
	Email        string   `json:"email" validate:"email"`
	Organization string   `json:"organization"`
	Message      string   `json:"message"`
	Phone        string   `json:"phone" validate:"phone"`
	Newsletter   bool     `json:"newsletter"`
	Products     []string `json:"products" validate:"dive,oneof=hardware software"`
}

// Validate attempts to validate the lead's values.
func (lead *Lead) Validate() error {
	if err := validate.Struct(lead); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}
		return err
	}
	return nil
}

func validatePhone(fl validator.FieldLevel) bool {
	p, err := phonenumbers.Parse(fl.Field().String(), "US")
	if err != nil {
		return false
	}
	return phonenumbers.IsPossibleNumber(p)
}
