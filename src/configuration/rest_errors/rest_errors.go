package rest_error

import "net/http"

type RestErr struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Status  int64    `json:"status"`
	Causes  []Causes `json:"causes"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *RestErr) Error() string {
	return r.Message
}

func NewRestErr(message, err string, status int64, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Status:  status,
		Causes:  causes,
	}
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Bad Request",
		Status:  http.StatusBadRequest,
	}
}

func NewRequestValidationError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Bad Request",
		Status:  http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Internal Server Error",
		Status:  http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Not Found",
		Status:  http.StatusNotFound,
	}
}

func NewUnauthorizedError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Not Authorized",
		Status:  http.StatusUnauthorized,
	}
}
