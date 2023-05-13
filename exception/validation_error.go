package exception

type ValidationError struct {
	Error []string
}

func NewValidationError(err []string) ValidationError {
	return ValidationError{
		Error: err,
	}
}
