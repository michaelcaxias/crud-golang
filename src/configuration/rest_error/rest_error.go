package rest_error

import "net/http"

type RestErr struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Status  int      `json:"status"`
	Causes  []Causes `json:"causes"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *RestErr) Error() string {
	return r.Message
}

func NewRestErr(message, err string, status int, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Status:  status,
		Causes:  causes,
	}
}

func ThrowBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Status:  http.StatusBadRequest,
	}
}

func ThrowRequestValidationError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Status:  http.StatusBadRequest,
		Causes:  causes,
	}
}

func ThrowInternalServerError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "internal_server_error",
		Status:  http.StatusInternalServerError,
	}
}

func ThrowNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "not_found",
		Status:  http.StatusNotFound,
	}
}

func ThrowUnauthorizedError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "unathorized",
		Status:  http.StatusUnauthorized,
	}
}
