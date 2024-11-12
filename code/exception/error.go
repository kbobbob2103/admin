package exception

import "fmt"

// ErrorCode is a custom type for error codes
type ErrorCode string

// Define your error codes
const (
	ErrCodeNotFound            ErrorCode = "ERR_USER_NOT_FOUND"
	ErrCodeUserAlreadyExisted            = "ERR_USER_ALREADY_EXISTED"
	ErrCodeDatabase            ErrorCode = "ERR_DATABASE_ERROR"
	ErrCodeBadRequest          ErrorCode = "ERR_BAD_REQUEST"
	ErrCodeApi                 ErrorCode = "ERR_API_ERROR"
	ErrCodeGrpc                ErrorCode = "ERR_GRPC_ERROR"
	ErrCodeUnknown             ErrorCode = "ERR_UNKNOWN_ERROR"
	ErrCodeExternalServiceDown ErrorCode = "ERR_EXTERNAL_SERVICE_DOWN"
)

// errorMessages maps error codes to user-friendly messages
var errorMessages = map[ErrorCode]string{
	ErrCodeNotFound:            "Data not found.",
	ErrCodeDatabase:            "เกิดข้อผิดพลาดบางอย่าง กรุณาลองใหม่อีกครั้ง.",
	ErrCodeGrpc:                "เกิดข้อผิดพลาดบางอย่าง กรุณาลองใหม่อีกครั้ง.",
	ErrCodeApi:                 "เกิดข้อผิดพลาดบางอย่าง กรุณาลองใหม่อีกครั้ง.",
	ErrCodeUnknown:             "An unexpected error occurred.",
	ErrCodeExternalServiceDown: "A dependent service is currently unavailable. Please try again later.",
}

// GetMessage returns the user-friendly message for the error code
func GetMessage(code ErrorCode) string {
	return errorMessages[code]
}

// AppError represents an application error with a code and an internal message.
type AppError struct {
	Code    ErrorCode
	Message string // Internal message for logging
}

// NewAppError creates a new AppError
func NewAppError(code ErrorCode, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

// Error returns the error message
func (e *AppError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}
