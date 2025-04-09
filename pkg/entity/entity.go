package entity

import (
	"time"

	"github.com/google/uuid"
)

type ProductType int

type Product struct {
	Id          uuid.UUID
	DateTime    time.Time
	ReceptionId uuid.UUID
	Type        ProductType
}

type PvzCity int

type Pvz struct {
	Id               uuid.UUID
	RegistrationDate time.Time
	City             PvzCity
}

type ReceptionStatus int

type Reception struct {
	Id       uuid.UUID
	PvzId    uuid.UUID
	DateTime time.Time
	Status   ReceptionStatus
}

type UserRole int

const (
	RoleModerator UserRole = iota
	RoleEmployee
)

type User struct {
	Id           uuid.UUID
	Email        string
	Role         UserRole
	PasswordHash []byte
}

type PvzInfo []struct {
	Pvz        Pvz
	Receptions []struct {
		Reception Reception
		Products  []Product
	}
}
