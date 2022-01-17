package errors

import "net/http"

type RestErr struct {
	Message string `json:"message"`
	Status  int64  `json:"status"`
	Error   string `json:"error"`
}

func NewBadResuestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}
