package entity

import "github.com/google/uuid"

type UserRole string

type User struct {
	Id           uuid.UUID
	Email        string
	Role         string
	PasswordHash []byte
}
