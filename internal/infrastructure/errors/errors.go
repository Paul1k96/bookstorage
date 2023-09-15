package errors

const (
	NoError = iota
	InternalError
	GeneralError
)

const (
	BookServiceGetBooksByAuthorErr = 1000 + iota
	BookServiceGetAuthorsByBookErr
)
