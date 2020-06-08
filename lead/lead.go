package lead

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// Lead represents the lead form POST data.
type Lead struct {
	Name         string `json:"name" validate:"required"`
	Email        string `json:"email"`
	Organization string `json:"organization"`
	Message      string `json:"message"`
	Referrer     string `json:"referrer"`
}

func (lead *Lead) Validate() (e []error) {
	if err := validate.Struct(&lead); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			e = append(e, errors.New(err.Tag()))
		}
		return e
	}
	return nil
}
