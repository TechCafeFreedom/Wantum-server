package tag

import (
	"context"
	"fmt"
	"net/http"
	"time"
	tagEntity "wantum/pkg/domain/entity/tag"
	"wantum/pkg/domain/repository"
	"wantum/pkg/domain/repository/tag"
	"wantum/pkg/werrors"

	"google.golang.org/grpc/codes"
)

type Service interface {
	Create(ctx context.Context, masterTx repository.MasterTx, name string) (*tagEntity.Entity, error)

	UpDeleteFlag(ctx context.Context, masterTx repository.MasterTx, tagID int) (*tagEntity.Entity, error)
	DownDeleteFlag(ctx context.Context, masterTx repository.MasterTx, tagID int) (*tagEntity.Entity, error)
	Delete(ctx context.Context, masterTx repository.MasterTx, tagID int) error

	GetByID(ctx context.Context, masterTx repository.MasterTx, tagID int) (*tagEntity.Entity, error)
	GetByName(ctx context.Context, masterTx repository.MasterTx, name string) (*tagEntity.Entity, error)
	GetByWishCardID(ctx context.Context, masterTx repository.MasterTx, wishCardID int) (tagEntity.EntitySlice, error)
	GetByMemoryID(ctx context.Context, masterTx repository.MasterTx, memoryID int) (tagEntity.EntitySlice, error)
}

type service struct {
	tagRepository tag.Repository
}

func New(repo tag.Repository) Service {
	return &service{
		tagRepository: repo,
	}
}

func (s *service) Create(ctx context.Context, masterTx repository.MasterTx, name string) (*tagEntity.Entity, error) {
	now := time.Now()
	tag := &tagEntity.Entity{
		Name:      name,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
	result, err := s.tagRepository.Insert(ctx, masterTx, tag)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	tag.ID = result
	return tag, nil
}

func (s *service) UpDeleteFlag(ctx context.Context, masterTx repository.MasterTx, tagID int) (*tagEntity.Entity, error) {
	tag, err := s.tagRepository.SelectByID(ctx, masterTx, tagID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	now := time.Now()
	tag.UpdatedAt = &now
	tag.DeletedAt = &now
	err = s.tagRepository.UpDeleteFlag(ctx, masterTx, tag)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return tag, nil
}

func (s *service) DownDeleteFlag(ctx context.Context, masterTx repository.MasterTx, tagID int) (*tagEntity.Entity, error) {
	tag, err := s.tagRepository.SelectByID(ctx, masterTx, tagID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	now := time.Now()
	tag.UpdatedAt = &now
	tag.DeletedAt = nil
	err = s.tagRepository.DownDeleteFlag(ctx, masterTx, tag)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return tag, nil
}

func (s *service) Delete(ctx context.Context, masterTx repository.MasterTx, tagID int) error {
	tag, err := s.tagRepository.SelectByID(ctx, masterTx, tagID)
	if err != nil {
		return werrors.Stack(err)
	}
	if tag.DeletedAt == nil {
		return werrors.Newf(
			fmt.Errorf("can't delete this data. this data did not up a delete flag. tagID=%v", tagID),
			codes.FailedPrecondition,
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

func (s *service) GetByID(ctx context.Context, masterTx repository.MasterTx, tagID int) (*tagEntity.Entity, error) {
	tag, err := s.tagRepository.SelectByID(ctx, masterTx, tagID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return tag, nil
}

func (s *service) GetByName(ctx context.Context, masterTx repository.MasterTx, name string) (*tagEntity.Entity, error) {
	tag, err := s.tagRepository.SelectByName(ctx, masterTx, name)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return tag, nil
}

func (s *service) GetByWishCardID(ctx context.Context, masterTx repository.MasterTx, wishCardID int) (tagEntity.EntitySlice, error) {
	tags, err := s.tagRepository.SelectByWishCardID(ctx, masterTx, wishCardID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return tags, nil
}

func (s *service) GetByMemoryID(ctx context.Context, masterTx repository.MasterTx, memoryID int) (tagEntity.EntitySlice, error) {
	tags, err := s.tagRepository.SelectByMemoryID(ctx, masterTx, memoryID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return tags, nil
}
