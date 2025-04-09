package entity

import "fmt"

var (
	ErrUnauthorized   = fmt.Errorf("not authorized to access this resource")
	ErrNotFound       = fmt.Errorf("not found")
	ErrIncorrectLogin = fmt.Errorf("incorrect email or password")
	ErrAlreadyExists  = fmt.Errorf("resource already exists")
)
