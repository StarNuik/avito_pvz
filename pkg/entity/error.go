package entity

import "fmt"

var (
	ErrNotAuthorized  = fmt.Errorf("not authorized to access this resource")
	ErrNotFound       = fmt.Errorf("not found")
	ErrIncorrectLogin = fmt.Errorf("incorrect email or password")
)
