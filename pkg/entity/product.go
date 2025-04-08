package entity

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	Id          uuid.UUID
	DateTime    time.Time
	ReceptionId uuid.UUID
	Type        string
}
