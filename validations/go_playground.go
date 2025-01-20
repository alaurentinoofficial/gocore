package validators

import (
	"strings"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type PlagroundValidator struct {
	validate *validator.Validate
}

func NewStructValidator() *PlagroundValidator {
	return &PlagroundValidator{
		validate: validator.New(),
	}
}

func (v *PlagroundValidator) Register(name string, errorMessage string, validation validator.Func) {
	errorMessages[name] = errorMessage
	v.validate.RegisterValidation(name, validation)
}

func (v *PlagroundValidator) Validate(value any) *FormErrors {
	if errs := v.validate.Struct(value); errs != nil {
		return ConvertToValidationFormErrors(errs.(validator.ValidationErrors))
	}

	return nil
}

func CleanNamespace(namespace string) string {
	result := strings.SplitN(namespace, ".", 2)
	return result[len(result)-1]
}

func ConvertToValidationFormErrors(validationErrors validator.ValidationErrors) *FormErrors {
	var response FormErrors

	for _, errorField := range validationErrors {
		response = append(response,
			NewFieldError(
				CleanNamespace(errorField.Namespace()),
				errorField.Tag(),
				errorField.Value()),
		)
	}

	return &response
}

var (
	ErrInternalServer   = errors.New("Internal server")
	ErrBadRequest       = errors.New("Bad request")
	ErrNotFound         = errors.New("Not found")
	ErrAlreadyRegistred = errors.New("Already registred")
	ErrForbidden        = errors.New("Forbidden")
	ErrMethodNotAllowed = errors.New("Method not allowed")
)

type FieldError struct {
	Name    string  `json:"name"`
	Code    string  `json:"code"`
	Message *string `json:"message"`
	Source  string  `json:"path"`
	Value   any     `json:"value"`
}

type FormErrors []FieldError

func (e FormErrors) AddNamespace(ns string) {
	for i := 0; i < len(e); i++ {
		e[i].Name = fmt.Sprintf("%s.%s", ns, e[i].Name)
	}
}

func (e FormErrors) Error() string {
	errorJson, _ := json.MarshalIndent(e, "", "\t")
	return string(errorJson)
}

var errorsParser = map[string]string{}

var errorMessages = map[string]string{
	"AlreadyExists": "Value already exists",
}

func NewFieldError(name string, code string, value any) FieldError {
	if tmpCode, ok := errorsParser[code]; ok {
		code = tmpCode
	}

	message, ok := errorMessages[code]
	var nullMessage *string = nil
	if ok {
		nullMessage = &message
	}

	return FieldError{Code: code, Source: "body", Name: name, Message: nullMessage, Value: value}
}

var (
	AlreadyExists = func(name string, value any) FieldError {
		return NewFieldError(name, "AlreadyExists", value)
	}

	NotFound = func(name string, value any) FieldError {
		return NewFieldError(name, "NotFound", value)
	}
)
