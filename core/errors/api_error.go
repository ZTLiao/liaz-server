package errors

type ApiError struct {
	Code    int
	Message string
}

func (e *ApiError) Error() string {
	return e.Message
}

func New(code int, message string) *ApiError {
	return &ApiError{
		Code:    code,
		Message: message,
	}
}
