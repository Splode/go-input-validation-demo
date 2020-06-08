package lead

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// Lead represents the lead form POST data.
type Lead struct {
	Name         string   `json:"name" validate:"required,ascii"`
	Email        string   `json:"email" validate:"email"`
	Organization string   `json:"organization"`
	Message      string   `json:"message"`
	Phone        string   `json:"phone"`
	Newsletter   bool     `json:"newsletter"`
	Products     []string `json:"products" validate:"dive,oneof=hardware software"`
}

func (lead *Lead) Validate() error {
	if err := validate.Struct(lead); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}
		return err
	}
	return nil
}
