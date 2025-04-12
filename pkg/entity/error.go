package entity

import (
	"fmt"
)

var (
	ErrUnauthorized    = fmt.Errorf("not authorized to access this resource")
	ErrNotFound        = fmt.Errorf("not found")
	ErrIncorrectLogin  = fmt.Errorf("incorrect email or password")
	ErrAlreadyExists   = fmt.Errorf("resource already exists")
	ErrReceptionClosed = fmt.Errorf("resource is already closed")
	ErrCantParse       = fmt.Errorf("cant parse")
	ErrInternal        = fmt.Errorf("internal error")
)

// TODO hide inner error
func InternalError(location string, inner error) error {
	return fmt.Errorf("%w: %s: %w", ErrInternal, location, inner)
}
