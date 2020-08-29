package model

import (
	"time"
	"wantum/pkg/domain/entity"
)

type PlaceModel struct {
	ID        int
	Name      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

type PlaceModelSlice []*PlaceModel

func ConvertToPlaceEntity(place *PlaceModel) *entity.Place {
	if place == nil {
		return nil
	}
	return &entity.Place{
		ID:        place.ID,
		Name:      place.Name,
		CreatedAt: place.CreatedAt,
		UpdatedAt: place.UpdatedAt,
		DeletedAt: place.DeletedAt,
	}
}

func ConvertToPlaceSliceEntity(places PlaceModelSlice) entity.PlaceSlice {
	if places == nil {
		return nil
	}
	res := make(entity.PlaceSlice, 0, len(places))
	for _, place := range places {
		res = append(res, ConvertToPlaceEntity(place))
	}
	return res
}
