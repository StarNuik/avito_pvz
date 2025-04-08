package entity

import (
	"time"

	"github.com/google/uuid"
)

type Reception struct {
	Id       uuid.UUID
	PvzId    uuid.UUID
	DateTime time.Time
	Status   string
}
