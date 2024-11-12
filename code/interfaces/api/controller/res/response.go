package res

import (
	"admin/microservice/exception"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BusinessCode int

type Status struct {
	ErrCode     string `json:"err_code"`
	Code        int    `json:"code"`
	Description string `json:"description"`
}

type Response struct {
	Status Status      `json:"status"`
	Data   interface{} `json:"data"`
}

// HandleError takes an AppError, logs it, and sends a user-friendly error response
func HandleError(c *gin.Context, err error, objectType string) {
	var appErr *exception.AppError
	ok := errors.As(err, &appErr)
	if !ok {
		err = exception.NewAppError(exception.ErrCodeUnknown, err.Error())
	}

	// Prepare the user-friendly error response
	status := Status{
		ErrCode:     string(err.(*exception.AppError).Code),
		Code:        2000,
		Description: exception.GetMessage(err.(*exception.AppError).Code),
	}
	fmt.Println(err.(*exception.AppError).Error())

	// Set the HTTP status code based on the error code
	var statusCode int
	switch err.(*exception.AppError).Code {
	case exception.ErrCodeNotFound:
		statusCode = http.StatusOK
	case exception.ErrCodeUserAlreadyExisted:
		statusCode = http.StatusOK
	case exception.ErrCodeBadRequest:
		statusCode = http.StatusBadRequest
	case exception.ErrCodeApi:
		statusCode = http.StatusBadRequest
	case exception.ErrCodeDatabase, exception.ErrCodeUnknown, exception.ErrCodeGrpc:
		statusCode = http.StatusInternalServerError
	case exception.ErrCodeExternalServiceDown:
		statusCode = http.StatusServiceUnavailable
	default:
		statusCode = http.StatusInternalServerError
	}
	var response Response
	if objectType == "object" {
		response.Data = struct{}{}
	} else if objectType == "slice" {
		response.Data = []struct{}{}
	}
	response.Status = status
	c.AbortWithStatusJSON(statusCode, response)
}

func HandleSuccess(
	c *gin.Context,
	result interface{},
	command string,
) {
	status := Status{
		ErrCode:     "",
		Code:        1000,
		Description: "Success",
	}
	response := Response{
		Status: status,
		Data:   result,
	}
	statusCode := http.StatusOK
	if command == "created" {
		statusCode = http.StatusCreated
	}
	c.JSON(statusCode, response)
}
