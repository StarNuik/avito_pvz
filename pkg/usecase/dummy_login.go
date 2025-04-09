package usecase

import (
	"github.com/starnuik/avito_pvz/pkg/entity"
	"github.com/starnuik/avito_pvz/pkg/token"
)

func (u *usecase) DummyLogin(userRole entity.UserRole) (token.Payload, error) {
	panic("")
}
