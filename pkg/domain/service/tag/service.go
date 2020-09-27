package tag

import (
	"context"
	"time"
	tagEntity "wantum/pkg/domain/entity/tag"
	"wantum/pkg/domain/repository"
	"wantum/pkg/domain/repository/tag"
	"wantum/pkg/werrors"
)

type Service interface {
	Create(ctx context.Context, masterTx repository.MasterTx, name string) (*tagEntity.Entity, error)
	Delete(ctx context.Context, masterTx repository.MasterTx, tagID int) (*tagEntity.Entity, error)
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
	newID, err := s.tagRepository.Insert(ctx, masterTx, tag)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	tag.ID = newID
	return tag, nil
}

func (s *service) Delete(ctx context.Context, masterTx repository.MasterTx, tagID int) (*tagEntity.Entity, error) {
	tag, err := s.tagRepository.SelectByID(ctx, masterTx, tagID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	now := time.Now()
	tag.UpdatedAt = &now
	tag.DeletedAt = &now
	if err = s.tagRepository.UpDeleteFlag(ctx, masterTx, tag.ID, tag.UpdatedAt, tag.DeletedAt); err != nil {
		return nil, werrors.Stack(err)
	}
	return tag, nil
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
