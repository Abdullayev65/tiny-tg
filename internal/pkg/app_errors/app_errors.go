package app_errors

func newErr(msg string, status int) *AppError {
	return &AppError{msg, status}
}

type AppError struct {
	Message string
	Status  int
}

func (err *AppError) Error() string {
	return err.Message
}
