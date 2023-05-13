package exception

type InterfaceError struct {
	Error string
}

func NewInterfaceError(err string) InterfaceError {
	return InterfaceError{
		Error: err,
	}
}
