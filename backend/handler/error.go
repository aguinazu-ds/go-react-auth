package handler

import (
	"fmt"
	"net/http"
)

type APIError struct {
	StatusCode int    `json:"statusCode"`
	Msg        string `json:"msg"`
}

func (e APIError) Error() string {
	return e.Msg
}

func NewAPIError(statusCode int, err error) APIError {
	return APIError{
		StatusCode: statusCode,
		Msg:        err.Error(),
	}
}

func InvalidJSON() APIError {
	return NewAPIError(http.StatusBadRequest, fmt.Errorf("invalid JSON"))
}
