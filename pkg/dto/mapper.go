package dto

import (
	"github.com/oapi-codegen/runtime/types"
	"github.com/starnuik/avito_pvz/pkg/entity"
)

func (userRole UserRole) ToEntity() (entity.UserRole, error) {
	switch userRole {
	case Employee:
		return entity.RoleEmployee, nil
	case Moderator:
		return entity.RoleModerator, nil
	default:
		return entity.UserRole(-1), entity.ErrCantParse
	}
}

func FromUserRole(userRole entity.UserRole) UserRole {
	switch userRole {
	case entity.RoleEmployee:
		return Employee
	case entity.RoleModerator:
		return Moderator
	default:
		panic("not implemented")
	}
}

func FromUser(user entity.User) User {
	return User{
		Email: types.Email(user.Email),
		Id:    &user.Id,
		Role:  FromUserRole(user.Role),
	}
}
