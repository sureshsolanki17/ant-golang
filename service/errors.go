package service

import "fmt"

// CustomError defines a structured error type for better error handling
type CustomError struct {
	Code    string // A short error code (e.g., "INVALID_CONFIG", "API_ERROR")
	Message string // A detailed error message
	Err     error  // Optional: The original error (for wrapping)
}

// Error implements the error interface for CustomError
func (e *CustomError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("[%s] %s: %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

// WrapError creates a new CustomError with a wrapped error
func WrapError(code, message string, err error) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// Predefined Errors
var (
	ErrInvalidConfig = &CustomError{Code: "INVALID_CONFIG", Message: "Invalid configuration provided"}
	ErrTokenNotSet   = &CustomError{Code: "TOKEN_NOT_SET", Message: "Authorization token is not set"}
	ErrAPIRequest    = func(endpoint string, err error) *CustomError {
		return WrapError("API_REQUEST_FAILED", fmt.Sprintf("Failed to make API request to %s", endpoint), err)
	}
)
