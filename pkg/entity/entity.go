package entity

import (
	"time"

	"github.com/google/uuid"
)

type UserRole int
type PvzCity int
type ReceptionStatus int
type ProductType int

const (
	RoleModerator UserRole = iota
	RoleEmployee
)

const (
	CityMoscow PvzCity = iota
	CitySaintPetersburg
	CityKazan
)

const (
	StatusInProgress ReceptionStatus = iota
	StatusClosed
)

type Product struct {
	Id          uuid.UUID
	DateTime    time.Time
	ReceptionId uuid.UUID
	Type        ProductType
}

type Pvz struct {
	Id               uuid.UUID
	RegistrationDate time.Time
	City             PvzCity
}

type Reception struct {
	Id       uuid.UUID
	PvzId    uuid.UUID
	DateTime time.Time
	Status   ReceptionStatus
}

type User struct {
	Id           uuid.UUID
	Email        string
	Role         UserRole
	PasswordHash []byte
}

type PvzInfo struct {
	Pvzs       map[uuid.UUID]Pvz
	Receptions map[uuid.UUID]Reception
	Products   map[uuid.UUID]Product
}
