package model

import (
	"time"
	"wantum/pkg/domain/entity"
)

type TagModel struct {
	ID        int
	Name      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

type TagModelSlice []*TagModel

func ConvertToTagEntity(tag *TagModel) *entity.Tag {
	if tag == nil {
		return nil
	}
	return &entity.Tag{
		ID:        tag.ID,
		Name:      tag.Name,
		CreatedAt: tag.CreatedAt,
		UpdatedAt: tag.UpdatedAt,
		DeletedAt: tag.DeletedAt,
	}
}

func ConvertToTagSliceEntity(tags TagModelSlice) entity.TagSlice {
	if tags == nil {
		return nil
	}
	result := make(entity.TagSlice, 0, len(tags))
	for _, tag := range tags {
		result = append(result, ConvertToTagEntity(tag))
	}
	return result
}
