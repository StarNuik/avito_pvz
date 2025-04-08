package entity

import (
	"time"

	"github.com/google/uuid"
)

type ReceptionStatus int

type Reception struct {
	Id       uuid.UUID
	PvzId    uuid.UUID
	DateTime time.Time
	Status   ReceptionStatus
}
