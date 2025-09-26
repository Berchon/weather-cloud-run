package model

type CustomError struct {
	StatusCode int    `json:"status_code"`
	Err        string `json:"message"`
}

func NewCustomError(statusCode int, message string) *CustomError {
	return &CustomError{
		StatusCode: statusCode,
		Err:        message,
	}
}

func (e *CustomError) Error() string {
	return e.Err
}
