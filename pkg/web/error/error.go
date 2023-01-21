package web

import "fmt"

type WebError struct {
	Status  int
	Message string
}

func (e *WebError) StatusCode() int {
	return e.Status
}

func (e *WebError) Error() string {
	return fmt.Sprintf("%d: %s", e.Status, e.Message)
}

func NewWebError(status int, message string) *WebError {
	return &WebError{
		Status:  status,
		Message: message,
	}
}
