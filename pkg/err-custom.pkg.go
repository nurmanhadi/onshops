package pkg

type ErrNotFound struct {
	Message string
}
type ErrBadRequest struct {
	Message string
}

func (ex *ErrNotFound) Error() string {
	return ex.Message
}
func (ex *ErrBadRequest) Error() string {
	return ex.Message
}
