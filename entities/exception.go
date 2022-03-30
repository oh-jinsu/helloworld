package entities

type Exception struct {
	err error
}

func NewException(err error) *Exception {
	return &Exception{err}
}

func (e *Exception) Occured() bool {
	return e.err != nil
}

func (e *Exception) Message() string {
	return e.err.Error()
}
