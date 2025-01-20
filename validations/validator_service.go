package validators

import "github.com/go-playground/validator/v10"

type ValidatorService interface {
	Register(name string, errorMessage string, validation validator.Func)
	Validate(value any)
}
