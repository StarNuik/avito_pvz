package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/starnuik/avito_pvz/pkg/token"
)

func (u *usecase) DeleteLastProduct(ctx context.Context, token token.Payload, pvzId uuid.UUID) error {
	panic("")
}
