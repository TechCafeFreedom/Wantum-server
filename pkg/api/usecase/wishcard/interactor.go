package wishcard

import (
	"context"
	"time"
	"wantum/pkg/domain/entity"
	"wantum/pkg/domain/repository"
	"wantum/pkg/domain/service/place"
	"wantum/pkg/domain/service/tag"
	"wantum/pkg/domain/service/wishcard"
	"wantum/pkg/domain/service/wishcardtag"
	"wantum/pkg/werrors"
)

type Interactor interface {
	CreateNewWishCard(ctx context.Context, userID int, activity, description, place string, date *time.Time, categoryID int, tags []string) (*entity.WishCard, error)
	UpdateWishCard(ctx context.Context, wishCardID, userID int, activity, description, place string, date, doneAt *time.Time, categoryID int, tags []string) (*entity.WishCard, error)
	DeleteWishCardByID(ctx context.Context, wishCardID int) error
	GetByID(ctx context.Context, wishCardID int) (*entity.WishCard, error)
	GetByCategoryID(ctx context.Context, categoryID int) (entity.WishCardSlice, error)
}

type interactor struct {
	masterTxManager      repository.MasterTxManager
	wishCardService      wishcard.Service
	tagService           tag.Service
	placeService         place.Service
	wishCardsTagsService wishcardtag.Service
}

func New(masterTxManager repository.MasterTxManager, wishCardService wishcard.Service, tagService tag.Service, placeService place.Service, wishCardsTagsService wishcardtag.Service) Interactor {
	return &interactor{
		masterTxManager:      masterTxManager,
		wishCardService:      wishCardService,
		tagService:           tagService,
		placeService:         placeService,
		wishCardsTagsService: wishCardsTagsService,
	}
}

func (i *interactor) CreateNewWishCard(ctx context.Context, userID int, activity, description, place string, date *time.Time, categoryID int, tags []string) (*entity.WishCard, error) {

	var newWishCard *entity.WishCard
	var err error
	err = i.masterTxManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
		place, err := i.placeService.Create(ctx, masterTx, place)
		if err != nil {
			// TODO: placeがすでにあったら辛いね...
			return werrors.Stack(err)
		}
		var tagIDs []int
		for _, tagName := range tags {
			var tag *entity.Tag
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

		err = i.wishCardsTagsService.CreateMultipleRelation(ctx, masterTx, newWishCard.ID, tagIDs)
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

// TODO: serviceに色々以降
func (i *interactor) UpdateWishCard(ctx context.Context, wishCardID, userID int, activity, description, place string, date, doneAt *time.Time, categoryID int, tags []string) (*entity.WishCard, error) {
	var wishCard *entity.WishCard
	var err error
	err = i.masterTxManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
		place, err := i.placeService.Create(ctx, masterTx, place)
		if err != nil {
			// TODO: placeがすでにあったら辛いね...
			// bynameで撮ってきて、なかったらcreateかな...ぐぬぬ
			return werrors.Stack(err)
		}
		wishCard, err = i.wishCardService.Update(ctx, masterTx, wishCardID, activity, description, date, doneAt, userID, categoryID, place.ID)
		if err != nil {
			return werrors.Stack(err)
		}
		wishCard.Place = place

		var tagIDs []int
		for _, tagName := range tags {
			var tag *entity.Tag
			tag, _ = i.tagService.GetByName(ctx, masterTx, tagName)
			if tag == nil {
				tag, err = i.tagService.Create(ctx, masterTx, tagName)
				if err != nil {
					return werrors.Stack(err)
				}
			}
			tagIDs = append(tagIDs, tag.ID)
			wishCard.Tags = append(wishCard.Tags, tag)
		}

		err = i.wishCardsTagsService.DeleteByWishCardID(ctx, masterTx, wishCard.ID)
		if err != nil {
			return werrors.Stack(err)
		}
		err = i.wishCardsTagsService.CreateMultipleTags(ctx, masterTx, wishCard.ID, tagIDs)
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

// TODO: serviceに色々以降
func (i *interactor) DeleteWishCardByID(ctx context.Context, wishCardID int) error {
	var err error
	err = i.masterTxManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
		err = i.wishCardService.Delete(ctx, masterTx, wishCardID)
		if err != nil {
			return werrors.Stack(err)
		}

		err = i.wishCardsTagsService.DeleteByWishCardID(ctx, masterTx, wishCardID)
		if err != nil {
			return werrors.Stack(err)
		}
		return nil
	})
	if err != nil {
		return werrors.Stack(err)
	}
	return nil
}

// TODO: serviceに色々以降
func (i *interactor) GetByID(ctx context.Context, wishCardID int) (*entity.WishCard, error) {
	var wishCard *entity.WishCard
	var err error
	err = i.masterTxManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
		wishCard, err = i.wishCardService.GetByID(ctx, masterTx, wishCardID)
		if err != nil {
			return werrors.Stack(err)
		}
		place, err := i.placeService.GetByID(ctx, masterTx, wishCard.Place.ID)
		if err != nil {
			return werrors.Stack(err)
		}
		wishCard.Place = place
		tags, err := i.tagService.GetByWishCardID(ctx, masterTx, wishCardID)
		if err != nil {
			return werrors.Stack(err)
		}
		wishCard.Tags = tags
		return nil
	})
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return wishCard, nil
}

// TODO: serviceに色々以降
func (i *interactor) GetByCategoryID(ctx context.Context, categoryID int) (entity.WishCardSlice, error) {
	var wishCards entity.WishCardSlice
	var err error
	err = i.masterTxManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
		wishCards, err = i.wishCardService.GetByCategoryID(ctx, masterTx, categoryID)
		if err != nil {
			return werrors.Stack(err)
		}
		for _, wishCard := range wishCards {
			place, err := i.placeService.GetByID(ctx, masterTx, wishCard.Place.ID)
			if err != nil {
				return werrors.Stack(err)
			}
			wishCard.Place = place
			tags, err := i.tagService.GetByWishCardID(ctx, masterTx, wishCard.ID)
			if err != nil {
				return werrors.Stack(err)
			}
			wishCard.Tags = tags
		}
		return nil
	})
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return wishCards, nil
}
