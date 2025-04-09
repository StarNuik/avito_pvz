package usecase

import (
	"github.com/google/uuid"
	"github.com/starnuik/avito_pvz/pkg/entity"
	"github.com/starnuik/avito_pvz/pkg/token"
)

func (u *usecase) CloseLastReception(token token.Payload, pvzId uuid.UUID) (entity.Reception, error) {
	panic("")
}
