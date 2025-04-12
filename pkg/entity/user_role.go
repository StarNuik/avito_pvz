package entity

import "fmt"

type UserRole int

const (
	RoleModerator UserRole = iota
	RoleEmployee
)

func ParseUserRole(str string) (UserRole, error) {
	switch str {
	case "moderator":
		return RoleModerator, nil
	case "employee":
		return RoleEmployee, nil
	default:
		return UserRole(-1), fmt.Errorf("UserRole: %w", ErrCantParse)
	}
}
