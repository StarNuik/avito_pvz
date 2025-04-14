package dto

import (
	"github.com/google/uuid"
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

func (pvzCity PvzCity) ToEntity() (entity.PvzCity, error) {
	switch pvzCity {
	case Москва:
		return entity.CityMoscow, nil
	case СанктПетербург:
		return entity.CitySaintPetersburg, nil
	case Казань:
		return entity.CityKazan, nil
	default:
		return entity.PvzCity(-1), entity.ErrCantParse
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

func FromPvzCity(city entity.PvzCity) PvzCity {
	switch city {
	case entity.CityMoscow:
		return Москва
	case entity.CitySaintPetersburg:
		return СанктПетербург
	case entity.CityKazan:
		return Казань
	default:
		panic("not implemented")
	}
}

func FromReceptionStatus(status entity.ReceptionStatus) ReceptionStatus {
	switch status {
	case entity.StatusInProgress:
		return InProgress
	case entity.StatusClosed:
		return Close
	default:
		panic("not implemented")
	}
}

func FromProductType(productType entity.ProductType) ProductType {
	switch productType {
	case entity.TypeFootwear:
		return Обувь
	case entity.TypeClothing:
		return Одежда
	case entity.TypeElectronics:
		return Электроника
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

func FromPvz(pvz entity.Pvz) PVZ {
	return PVZ{
		City:             FromPvzCity(pvz.City),
		Id:               &pvz.Id,
		RegistrationDate: &pvz.RegistrationDate,
	}
}

func FromReception(reception entity.Reception) Reception {
	return Reception{
		DateTime: reception.DateTime,
		Id:       &reception.Id,
		PvzId:    reception.PvzId,
		Status:   FromReceptionStatus(reception.Status),
	}
}

func FromProduct(product entity.Product) Product {
	return Product{
		DateTime:    &product.DateTime,
		Id:          &product.Id,
		ReceptionId: product.ReceptionId,
		Type:        FromProductType(product.Type),
	}
}

func FromPvzInfo(info entity.PvzInfo) []PVZListPVZ {
	remap := make(map[uuid.UUID]int)
	out := make([]PVZListPVZ, 0, len(info.Pvzs))

	for _, pvz := range info.Pvzs {
		remap[pvz.Id] = len(out)
		out = append(out, PVZListPVZ{
			Pvz:        FromPvz(pvz),
			Receptions: []PVZListReception{},
		})
	}

	for _, reception := range info.Receptions {
		pvzIdx := remap[reception.PvzId]

		dest := &out[pvzIdx].Receptions
		*dest = append(*dest, PVZListReception{
			Reception: FromReception(reception),
			Products:  []Product{},
		})

		remap[reception.Id] = len(*dest) - 1
	}

	for _, product := range info.Products {
		pvzIdx := remap[info.Receptions[product.ReceptionId].PvzId]
		recIdx := remap[product.ReceptionId]

		dest := &out[pvzIdx].Receptions[recIdx].Products
		*dest = append(*dest, FromProduct(product))
	}

	return out
}
