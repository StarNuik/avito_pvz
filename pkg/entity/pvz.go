package entity

import (
	"time"

	"github.com/google/uuid"
)

type PvzCity int

type Pvz struct {
	Id               uuid.UUID
	RegistrationDate time.Time
	City             PvzCity
}
