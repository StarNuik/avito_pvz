package entity

import "github.com/google/uuid"

type UserRole int

const (
	RoleModerator UserRole = iota
	RoleClient
)

type User struct {
	Id           uuid.UUID
	Email        string
	Role         UserRole
	PasswordHash []byte
}
