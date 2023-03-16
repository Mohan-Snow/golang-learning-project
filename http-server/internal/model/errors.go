package model

import (
	"fmt"
)

type GeneralError struct {
	Message   string
	ErrorCode int
}

type ExternalServiceError struct {
	Message    string
	ServiceUrl string
	ErrorCode  int
}

type UserDataError struct {
	Message   string
	Data      string
	ErrorCode int
}

func (g *GeneralError) Error() string {
	return fmt.Sprintf("Error Message: %s\nErrorCode: %d", g.Message, g.ErrorCode)
}

func (e *ExternalServiceError) Error() string {
	return fmt.Sprintf("Error Message: %s\nServiceUrl: %s\nErrorCode: %d", e.Message, e.ServiceUrl, e.ErrorCode)
}

func (u *UserDataError) Error() string {
	return fmt.Sprintf("Error Message: %s\nData: %s\nErrorCode: %d", u.Message, u.Data, u.ErrorCode)
}
