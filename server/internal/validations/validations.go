package validations

import (
	"github.com/go-playground/validator/v10"
)

// ErrorResponse represents the structure for validation error responses
type ErrorResponse struct {
	Error       bool        `json:"error"`
	FailedField string      `json:"failed_field"`
	Tag         string      `json:"tag"`
	Value       interface{} `json:"value"`
}

// Validator interface defines the contract for validation
type Validator interface {
	Validate(data interface{}) []ErrorResponse
}

// xValidator implements the Validator interface using a global validator instance
type xValidator struct {
	validator *validator.Validate
}

// GlobalErrorHandlerResp represents the structure for global error handling responses
type GlobalErrorHandlerResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// Create a global validator instance
var globalValidator = &xValidator{
	validator: validator.New(validator.WithRequiredStructEnabled()),
}

// Validate validates the given data using the global validator instance
func (v *xValidator) Validate(data interface{}) []ErrorResponse {
	validationErrors := []ErrorResponse{}

	errs := v.validator.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			// Create an ErrorResponse for each validation error
			var elem ErrorResponse

			elem.FailedField = err.Field() // Export struct field name
			elem.Tag = err.Tag()           // Export struct tag
			elem.Value = err.Value()       // Export field value
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}

// GetGlobalValidator returns the global validator instance
func GetGlobalValidator() Validator {
	return globalValidator
}
