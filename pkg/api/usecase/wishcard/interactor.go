package wishcard

import (
	"context"
	"time"
	tagEntity "wantum/pkg/domain/entity/tag"
	wishCardEntity "wantum/pkg/domain/entity/wishcard"
	"wantum/pkg/domain/repository"
	"wantum/pkg/domain/service/place"
	"wantum/pkg/domain/service/tag"
	"wantum/pkg/domain/service/wishcard"
	"wantum/pkg/werrors"
)

type Interactor interface {
	CreateNewWishCard(ctx context.Context, userID, categoryID int, activity, description, place string, date *time.Time, tags []string) (*wishCardEntity.Entity, error)
	UpdateWishCard(ctx context.Context, wishCardID, userID int, activity, description, place string, date, doneAt *time.Time, categoryID int, tags []string) (*wishCardEntity.Entity, error)
	DeleteWishCardByID(ctx context.Context, wishCardID int) error
	GetByID(ctx context.Context, wishCardID int) (*wishCardEntity.Entity, error)
	GetByCategoryID(ctx context.Context, categoryID int) (wishCardEntity.EntitySlice, error)
}

type interactor struct {
	masterTxManager repository.MasterTxManager
	wishCardService wishcard.Service
	tagService      tag.Service
	placeService    place.Service
}

func New(masterTxManager repository.MasterTxManager, wishCardService wishcard.Service, tagService tag.Service, placeService place.Service) Interactor {
	return &interactor{
		masterTxManager: masterTxManager,
		wishCardService: wishCardService,
		tagService:      tagService,
		placeService:    placeService,
	}
}

func (i *interactor) CreateNewWishCard(ctx context.Context, userID, categoryID int, activity, description, place string, date *time.Time, tags []string) (*wishCardEntity.Entity, error) {

	var newWishCard *wishCardEntity.Entity
	err := i.masterTxManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
		place, err := i.placeService.Create(ctx, masterTx, place)
		if err != nil {
			// TODO: placeがすでにあったら無限に増えてしまう
			return werrors.Stack(err)
		}
		tagIDs := make([]int, 0, len(tags))
		for _, tagName := range tags {
			var tag *tagEntity.Entity
			tag, _ = i.tagService.GetByName(ctx, masterTx, tagName)
			if tag == nil {
				tag, err = i.tagService.Create(ctx, masterTx, tagName)
				if err != nil {
					return werrors.Stack(err)
				}
			}
			tagIDs = append(tagIDs, tag.ID)
		}

		newWishCard, err = i.wishCardService.Create(ctx, masterTx, activity, description, date, userID, categoryID, place.ID, tagIDs)
		if err != nil {
			return werrors.Stack(err)
		}

		return nil
	})
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return newWishCard, nil
}

func (i *interactor) UpdateWishCard(ctx context.Context, wishCardID, userID int, activity, description, place string, date, doneAt *time.Time, categoryID int, tags []string) (*wishCardEntity.Entity, error) {
	var wishCard *wishCardEntity.Entity
	err := i.masterTxManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
		place, err := i.placeService.Create(ctx, masterTx, place)
		if err != nil {
			// TODO: placeがすでにあったら無限に増えてしまう
			return werrors.Stack(err)
		}
		tagIDs := make([]int, 0, len(tags))
		for _, tagName := range tags {
			var tag *tagEntity.Entity
			tag, _ = i.tagService.GetByName(ctx, masterTx, tagName)
			if tag == nil {
				tag, err = i.tagService.Create(ctx, masterTx, tagName)
				if err != nil {
					return werrors.Stack(err)
				}
			}
			tagIDs = append(tagIDs, tag.ID)
		}
		wishCard, err = i.wishCardService.Update(ctx, masterTx, wishCardID, activity, description, date, doneAt, userID, categoryID, place.ID, tagIDs)
		if err != nil {
			return werrors.Stack(err)
		}
		return nil
	})
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return wishCard, nil
}

func (i *interactor) DeleteWishCardByID(ctx context.Context, wishCardID int) error {
	var err error
	err = i.masterTxManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
		if err = i.wishCardService.Delete(ctx, masterTx, wishCardID); err != nil {
			return werrors.Stack(err)
		}

		return nil
	})
	if err != nil {
		return werrors.Stack(err)
	}
	return nil
}

func (i *interactor) GetByID(ctx context.Context, wishCardID int) (*wishCardEntity.Entity, error) {
	var wishCard *wishCardEntity.Entity
	var err error
	err = i.masterTxManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
		wishCard, err = i.wishCardService.GetByID(ctx, masterTx, wishCardID)
		if err != nil {
			return werrors.Stack(err)
		}
		return nil
	})
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return wishCard, nil
}

func (i *interactor) GetByCategoryID(ctx context.Context, categoryID int) (wishCardEntity.EntitySlice, error) {
	var wishCards wishCardEntity.EntitySlice
	var err error
	err = i.masterTxManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
		wishCards, err = i.wishCardService.GetByCategoryID(ctx, masterTx, categoryID)
		if err != nil {
			return werrors.Stack(err)
		}
		return nil
	})
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return wishCards, nil
}
