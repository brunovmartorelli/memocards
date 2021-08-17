package domain

type NotFoundError struct {
	Err error
}

func (n NotFoundError) Error() string {
	return n.Err.Error()
}

type UpdateError struct {
	Err error
}

func (u UpdateError) Error() string {
	return u.Err.Error()
}
