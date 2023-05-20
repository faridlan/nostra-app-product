package exception

type InterfaceErrorUnauth struct {
	Error string
}

func NewInterfaceErrorUnauth(err string) InterfaceErrorUnauth {
	return InterfaceErrorUnauth{
		Error: err,
	}
}
