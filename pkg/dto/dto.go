package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/starnuik/avito_pvz/pkg/entity"
)

type ErrorResponse struct {
	Message string
}

// Auth
type DummyLoginRequest struct {
	Role entity.UserRole
}

type RegisterRequest struct {
	Email    string
	Password string
	Role     entity.UserRole
}

type RegisterResponse struct {
	Id    uuid.UUID
	Email string
	Role  entity.UserRole
}

type LoginRequest struct {
	Email    string
	Password string
}

// Create
type PvzRequest struct {
	Id               *uuid.UUID
	RegistrationDate *time.Time
	City             entity.PvzCity
}

type PvzResponse struct {
	Id               uuid.UUID
	RegistrationDate time.Time
	City             entity.PvzCity
}

type ReceptionRequest struct {
	PvzId uuid.UUID
}

type ReceptionResponse struct {
	Id       uuid.UUID
	DateTime time.Time
	PvzId    uuid.UUID
	Status   entity.ReceptionStatus
}

type ProductRequest struct {
	Type  entity.ProductType
	PvzId uuid.UUID
}

type ProductResponse struct {
	Id          uuid.UUID
	DateTime    time.Time
	Type        entity.ProductType
	ReceptionId uuid.UUID
}

// Read
// type GetPvzRequest struct {
// 	StartDate time.Time
// 	EndDate   time.Time
// 	Page      *int
// 	Limit     *int
// }

type GetPvzResponse []struct {
	Pvz        PvzResponse
	Receptions []struct {
		Reception ReceptionResponse
		Products  []ProductResponse
	}
}
