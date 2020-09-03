package tag

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"wantum/pkg/domain/entity"
	"wantum/pkg/domain/repository"
	"wantum/pkg/domain/repository/tag"
	"wantum/pkg/infrastructure/mysql/model"
	"wantum/pkg/werrors"
)

type Service interface {
	Create(ctx context.Context, masterTx repository.MasterTx, name string) (*entity.Tag, error)

	UpDeleteFlag(ctx context.Context, masterTx repository.MasterTx, tagID int) (*entity.Tag, error)
	DownDeleteFlag(ctx context.Context, masterTx repository.MasterTx, tagID int) (*entity.Tag, error)
	Delete(ctx context.Context, masterTx repository.MasterTx, tagID int) error

	GetByID(ctx context.Context, masterTx repository.MasterTx, tagID int) (*entity.Tag, error)
	GetByWishCardID(ctx context.Context, masterTx repository.MasterTx, wishCardID int) (entity.TagSlice, error)
	GetByMemoryID(ctx context.Context, masterTx repository.MasterTx, memoryID int) (entity.TagSlice, error)
}

type service struct {
	tagRepository tag.Repository
}

func New(repo tag.Repository) Service {
	return &service{
		tagRepository: repo,
	}
}

func (s *service) Create(ctx context.Context, masterTx repository.MasterTx, name string) (*entity.Tag, error) {
	createdAt := time.Now()
	tag := &model.TagModel{
		Name:      name,
		CreatedAt: &createdAt,
		UpdatedAt: &createdAt,
	}
	result, err := s.tagRepository.Insert(ctx, masterTx, tag)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	tag.ID = result
	return model.ConvertToTagEntity(tag), nil
}

func (s *service) UpDeleteFlag(ctx context.Context, masterTx repository.MasterTx, tagID int) (*entity.Tag, error) {
	tag, err := s.tagRepository.SelectByID(ctx, masterTx, tagID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	updatedAt := time.Now()
	tag.UpdatedAt = &updatedAt
	tag.DeletedAt = &updatedAt
	err = s.tagRepository.UpDeleteFlag(ctx, masterTx, tag)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return model.ConvertToTagEntity(tag), nil
}

func (s *service) DownDeleteFlag(ctx context.Context, masterTx repository.MasterTx, tagID int) (*entity.Tag, error) {
	tag, err := s.tagRepository.SelectByID(ctx, masterTx, tagID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	updatedAt := time.Now()
	tag.UpdatedAt = &updatedAt
	tag.DeletedAt = nil
	err = s.tagRepository.DownDeleteFlag(ctx, masterTx, tag)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return model.ConvertToTagEntity(tag), nil
}

func (s *service) Delete(ctx context.Context, masterTx repository.MasterTx, tagID int) error {
	tag, err := s.tagRepository.SelectByID(ctx, masterTx, tagID)
	if err != nil {
		return werrors.Stack(err)
	}
	if tag.DeletedAt == nil {
		return werrors.Newf(
			fmt.Errorf("can't delete this data. this data did not up a delete flag. tagID=%v", tagID),
			http.StatusBadRequest,
			"このデータは削除できません",
			"could not delete this place",
		)
	}
	err = s.tagRepository.Delete(ctx, masterTx, tagID)
	if err != nil {
		return werrors.Stack(err)
	}
	return nil
}

func (s *service) GetByID(ctx context.Context, masterTx repository.MasterTx, tagID int) (*entity.Tag, error) {
	tag, err := s.tagRepository.SelectByID(ctx, masterTx, tagID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return model.ConvertToTagEntity(tag), nil
}

func (s *service) GetByWishCardID(ctx context.Context, masterTx repository.MasterTx, wishCardID int) (entity.TagSlice, error) {
	tags, err := s.tagRepository.SelectByWishCardID(ctx, masterTx, wishCardID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return model.ConvertToTagSliceEntity(tags), nil
}

func (s *service) GetByMemoryID(ctx context.Context, masterTx repository.MasterTx, memoryID int) (entity.TagSlice, error) {
	tags, err := s.tagRepository.SelectByMemoryID(ctx, masterTx, memoryID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return model.ConvertToTagSliceEntity(tags), nil
}
