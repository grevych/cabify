package marketplace

type NotFoundError struct {
	message string
}

func NewNotFoundError(message string) *NotFoundError {
	return &NotFoundError{message}
}

func (err *NotFoundError) Error() string {
	return err.message
}

type NotCreatedError struct {
	message string
}

func NewNotCreatedError(message string) *NotCreatedError {
	return &NotCreatedError{message}
}

func (err *NotCreatedError) Error() string {
	return err.message
}

type NotDeletedError struct {
	message string
}

func NewNotDeletedError(message string) *NotDeletedError {
	return &NotDeletedError{message}
}

func (err *NotDeletedError) Error() string {
	return err.message
}

type NotUpdatedError struct {
	message string
}

func NewNotUpdatedError(message string) *NotUpdatedError {
	return &NotUpdatedError{message}
}

func (err *NotUpdatedError) Error() string {
	return err.message
}
