package pvztest

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/starnuik/avito_pvz/pkg/entity"
)

func FilledPvzInfo(t *testing.T) entity.PvzInfo {
	out := entity.PvzInfo{
		Pvzs:       make(map[uuid.UUID]entity.Pvz),
		Receptions: make(map[uuid.UUID]entity.Reception),
		Products:   make(map[uuid.UUID]entity.Product),
	}

	pvzs := []entity.Pvz{}
	for range 2 {
		pvzs = append(pvzs, entity.Pvz{
			Id:               NewUuid(t),
			RegistrationDate: time.Unix(0, 0),
		})
	}

	receptions := []entity.Reception{}
	for idx := range 4 {
		receptions = append(receptions, entity.Reception{
			Id:       NewUuid(t),
			PvzId:    pvzs[idx/2].Id,
			DateTime: time.Unix(0, 0),
		})
	}

	sec := int64(0)
	products := []entity.Product{}
	for idx := range 12 {
		sec += 100
		products = append(products, entity.Product{
			Id:          NewUuid(t),
			ReceptionId: receptions[idx/3].Id,
			DateTime:    time.Unix(sec, 0),
		})
	}

	for _, pvz := range pvzs {
		out.Pvzs[pvz.Id] = pvz
	}
	for _, r := range receptions {
		out.Receptions[r.Id] = r
	}
	for _, p := range products {
		out.Products[p.Id] = p
	}

	return out
}
