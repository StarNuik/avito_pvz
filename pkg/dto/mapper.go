package dto

import "github.com/starnuik/avito_pvz/pkg/entity"

func (req PostDummyLoginJSONRequestBody) ToUserRole() (entity.UserRole, error) {
	switch req.Role {
	case Employee:
		return entity.RoleEmployee, nil
	case Moderator:
		return entity.RoleModerator, nil
	default:
		return entity.UserRole(-1), entity.ErrCantParse
	}
}
