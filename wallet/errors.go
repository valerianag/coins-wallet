package wallet

type ErrHttp struct {
	error      string
	statusCode int
}

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
