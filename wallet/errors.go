package wallet

type ErrHttp struct {
	error      string
	statusCode int
}

// NewErrHttp creates ErrHttp. We need it to be able to change response status codes
func NewErrHttp(error string, statusCode int) ErrHttp {
	return ErrHttp{
		error:      error,
		statusCode: statusCode,
	}
}

func (e ErrHttp) StatusCode() int {
	return e.statusCode
}

func (e ErrHttp) Error() string {
	return e.error
}
