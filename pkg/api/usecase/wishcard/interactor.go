package wishcard

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	tagEntity "wantum/pkg/domain/entity/tag"
	wishCardEntity "wantum/pkg/domain/entity/wishcard"
	"wantum/pkg/domain/repository"
	"wantum/pkg/domain/service/place"
	"wantum/pkg/domain/service/profile"
	"wantum/pkg/domain/service/tag"
	"wantum/pkg/domain/service/user"
	"wantum/pkg/domain/service/wishcard"
	"wantum/pkg/tlog"
	"wantum/pkg/werrors"

	"google.golang.org/grpc/codes"
)

type Interactor interface {
	CreateNewWishCard(ctx context.Context, userID, categoryID int, activity, description, place string, date *time.Time, tags []string) (*wishCardEntity.Entity, error)
	UpdateActivity(ctx context.Context, userID, wishCardID int, activity string) (*wishCardEntity.Entity, error)
	UpdateDescription(ctx context.Context, userID, wishCardID int, description string) (*wishCardEntity.Entity, error)
	UpdatePlace(ctx context.Context, userID, wishCardID int, place string) (*wishCardEntity.Entity, error)
	UpdateDate(ctx context.Context, userID, wishCardID int, date *time.Time) (*wishCardEntity.Entity, error)
	UpdateWishCardWithCategoryID(ctx context.Context, wishCardID, userID int, activity, description, place string, date, doneAt *time.Time, categoryID int, tags []string) (*wishCardEntity.Entity, error)
	DeleteWishCardByID(ctx context.Context, wishCardID int) error
	GetByID(ctx context.Context, wishCardID int) (*wishCardEntity.Entity, error)
	GetByCategoryID(ctx context.Context, categoryID int) (wishCardEntity.EntitySlice, error)
	AddTags(ctx context.Context, userID, wishCardID int, tags []string) (*wishCardEntity.Entity, error)
	DeleteTags(ctx context.Context, userID, wishCardID int, tagIDs []int) (*wishCardEntity.Entity, error)
}

type interactor struct {
	masterTxManager repository.MasterTxManager
	wishCardService wishcard.Service
	userService     user.Service
	profileService  profile.Service
	tagService      tag.Service
	placeService    place.Service
}

func New(masterTxManager repository.MasterTxManager, wishCardService wishcard.Service, userService user.Service, profileService profile.Service, tagService tag.Service, placeService place.Service) Interactor {
	return &interactor{
		masterTxManager: masterTxManager,
		wishCardService: wishCardService,
		userService:     userService,
		profileService:  profileService,
		tagService:      tagService,
		placeService:    placeService,
	}
}

// TODO: 先に取得した値ボンボン詰め込んで大丈夫か...?
// TODO: upd系の編集権限があるか

func (i *interactor) CreateNewWishCard(ctx context.Context, userID, categoryID int, activity, description, place string, date *time.Time, tags []string) (*wishCardEntity.Entity, error) {
	// validation
	var err error
	if err = validateActivity(activity); err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.Stack(err)
	}
	if err = validateDescription(description); err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.Stack(err)
	}
	if err = validatePlace(place); err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.Stack(err)
	}
	if err = validateDate(date); err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.Stack(err)
	}
	if err = validateTags(tags); err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.Stack(err)
	}

	// create new entity
	var newWishCard *wishCardEntity.Entity
	err = i.masterTxManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
		_author, err := i.userService.GetByPK(ctx, masterTx, userID)
		if err != nil {
			return werrors.Stack(err)
		}
		if _author == nil {
			return werrors.Stack(werrors.UserNotFound)
		}
		_profile, err := i.profileService.GetByUserID(ctx, masterTx, _author.ID)
		if err != nil {
			return werrors.Stack(err)
		}

		_place, err := i.placeService.Create(ctx, masterTx, place)
		if err != nil {
			// TODO: placeがすでにあったら無限に増えてしまう
			return werrors.Stack(err)
		}

		// TODO: きたない
		tagIDs := make([]int, 0, len(tags))
		_tags := make(tagEntity.EntitySlice, 0, len(tags))
		for _, tagName := range tags {
			var tag *tagEntity.Entity
			tag, err = i.tagService.GetByName(ctx, masterTx, tagName)
			if err != nil {
				return werrors.Stack(err)
			}
			if tag == nil {
				tag, err = i.tagService.Create(ctx, masterTx, tagName)
				if err != nil {
					return werrors.Stack(err)
				}
			}
			_tags = append(_tags, tag)
			tagIDs = append(tagIDs, tag.ID)
		}

		// TODO: カテゴリが存在するか確認する

		newWishCard, err = i.wishCardService.Create(ctx, masterTx, activity, description, date, _author.ID, categoryID, _place.ID, tagIDs)
		if err != nil {
			return werrors.Stack(err)
		}
		newWishCard.Author = _author
		newWishCard.Author.Profile = _profile
		newWishCard.Place = _place
		newWishCard.Tags = _tags

		return nil
	})
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return newWishCard, nil
}

func (i *interactor) UpdateWishCardWithCategoryID(ctx context.Context, wishCardID, userID int, activity, description, place string, date, doneAt *time.Time, categoryID int, tags []string) (*wishCardEntity.Entity, error) {
	var err error
	if err = validateActivity(activity); err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.Stack(err)
	}
	if err = validateDescription(description); err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.Stack(err)
	}
	if err = validatePlace(place); err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.Stack(err)
	}
	if err = validateDate(date); err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.Stack(err)
	}
	if err = validateTags(tags); err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.Stack(err)
	}

	var wishCard *wishCardEntity.Entity
	err = i.masterTxManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
		_author, err := i.userService.GetByPK(ctx, masterTx, userID)
		if err != nil {
			return werrors.Stack(err)
		}
		if _author == nil {
			return werrors.Stack(werrors.UserNotFound)
		}
		_profile, err := i.profileService.GetByUserID(ctx, masterTx, _author.ID)
		if err != nil {
			return werrors.Stack(err)
		}

		_place, err := i.placeService.Create(ctx, masterTx, place)
		if err != nil {
			// TODO: placeがすでにあったら無限に増えてしまう
			return werrors.Stack(err)
		}

		tagIDs := make([]int, 0, len(tags))
		_tags := make(tagEntity.EntitySlice, 0, len(tags))
		for _, tagName := range tags {
			var tag *tagEntity.Entity
			tag, err = i.tagService.GetByName(ctx, masterTx, tagName)
			if err != nil {
				return werrors.Stack(err)
			}
			if tag == nil {
				tag, err = i.tagService.Create(ctx, masterTx, tagName)
				if err != nil {
					return werrors.Stack(err)
				}
			}
			_tags = append(_tags, tag)
			tagIDs = append(tagIDs, tag.ID)
		}

		wishCard, err = i.wishCardService.UpdateWithCategoryID(ctx, masterTx, wishCardID, activity, description, date, doneAt, userID, categoryID, _place.ID, tagIDs)
		if err != nil {
			return werrors.Stack(err)
		}
		wishCard.Author = _author
		wishCard.Author.Profile = _profile
		wishCard.Place = _place
		wishCard.Tags = _tags
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
		wishCard.Author, err = i.userService.GetByPK(ctx, masterTx, wishCard.Author.ID)
		if err != nil {
			return werrors.Stack(err)
		}
		// QUESTION: author == nilだった時？
		wishCard.Author.Profile, err = i.profileService.GetByUserID(ctx, masterTx, wishCard.Author.ID)
		if err != nil {
			return werrors.Stack(err)
		}
		wishCard.Place, err = i.placeService.GetByID(ctx, masterTx, wishCard.Place.ID)
		if err != nil {
			return werrors.Stack(err)
		}
		wishCard.Tags, err = i.tagService.GetByWishCardID(ctx, masterTx, wishCard.ID)
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
		// OPTIMIZE: でら遅い...?
		for _, wishCard := range wishCards {
			wishCard.Author, err = i.userService.GetByPK(ctx, masterTx, wishCard.Author.ID)
			if err != nil {
				return werrors.Stack(err)
			}
			// QUESTION: author == nilだった時？
			wishCard.Author.Profile, err = i.profileService.GetByUserID(ctx, masterTx, wishCard.Author.ID)
			if err != nil {
				return werrors.Stack(err)
			}
			wishCard.Place, err = i.placeService.GetByID(ctx, masterTx, wishCard.Place.ID)
			if err != nil {
				return werrors.Stack(err)
			}
			wishCard.Tags, err = i.tagService.GetByWishCardID(ctx, masterTx, wishCard.ID)
			if err != nil {
				return werrors.Stack(err)
			}
		}
		return nil
	})
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return wishCards, nil
}

func (i *interactor) UpdateActivity(ctx context.Context, userID, wishCardID int, activity string) (*wishCardEntity.Entity, error) {
	var err error
	if err = validateActivity(activity); err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.Stack(err)
	}

	var wishCard = new(wishCardEntity.Entity)
	err = i.masterTxManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
		wishCard, err = i.wishCardService.UpdateActivity(ctx, masterTx, wishCardID, activity)
		if err != nil {
			return werrors.Stack(err)
		}
		wishCard.Author, err = i.userService.GetByPK(ctx, masterTx, wishCard.Author.ID)
		if err != nil {
			return werrors.Stack(err)
		}
		// QUESTION: author == nilだった時？
		wishCard.Author.Profile, err = i.profileService.GetByUserID(ctx, masterTx, wishCard.Author.ID)
		if err != nil {
			return werrors.Stack(err)
		}
		wishCard.Place, err = i.placeService.GetByID(ctx, masterTx, wishCard.Place.ID)
		if err != nil {
			return werrors.Stack(err)
		}
		wishCard.Tags, err = i.tagService.GetByWishCardID(ctx, masterTx, wishCard.ID)
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

func (i *interactor) UpdateDescription(ctx context.Context, userID, wishCardID int, description string) (*wishCardEntity.Entity, error) {
	var err error
	if err = validateDescription(description); err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.Stack(err)
	}

	var wishCard = new(wishCardEntity.Entity)
	err = i.masterTxManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
		wishCard, err = i.wishCardService.UpdateDescription(ctx, masterTx, wishCardID, description)
		if err != nil {
			return werrors.Stack(err)
		}
		wishCard.Author, err = i.userService.GetByPK(ctx, masterTx, wishCard.Author.ID)
		if err != nil {
			return werrors.Stack(err)
		}
		// QUESTION: author == nilだった時？
		wishCard.Author.Profile, err = i.profileService.GetByUserID(ctx, masterTx, wishCard.Author.ID)
		if err != nil {
			return werrors.Stack(err)
		}
		wishCard.Place, err = i.placeService.GetByID(ctx, masterTx, wishCard.Place.ID)
		if err != nil {
			return werrors.Stack(err)
		}
		wishCard.Tags, err = i.tagService.GetByWishCardID(ctx, masterTx, wishCard.ID)
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

func (i *interactor) UpdatePlace(ctx context.Context, userID, wishCardID int, place string) (*wishCardEntity.Entity, error) {
	var err error
	if err = validatePlace(place); err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.Stack(err)
	}

	var wishCard = new(wishCardEntity.Entity)
	err = i.masterTxManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
		wishCard.Place, err = i.placeService.Create(ctx, masterTx, place)
		if err != nil {
			// TODO: placeがすでにあったら無限に増えてしまう
			return werrors.Stack(err)
		}
		wishCard, err = i.wishCardService.UpdatePlace(ctx, masterTx, wishCardID, wishCard.Place.ID)
		if err != nil {
			return werrors.Stack(err)
		}
		wishCard.Author, err = i.userService.GetByPK(ctx, masterTx, wishCard.Author.ID)
		if err != nil {
			return werrors.Stack(err)
		}
		// QUESTION: author == nilだった時？
		wishCard.Author.Profile, err = i.profileService.GetByUserID(ctx, masterTx, wishCard.Author.ID)
		if err != nil {
			return werrors.Stack(err)
		}
		wishCard.Tags, err = i.tagService.GetByWishCardID(ctx, masterTx, wishCard.ID)
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

func (i *interactor) UpdateDate(ctx context.Context, userID, wishCardID int, date *time.Time) (*wishCardEntity.Entity, error) {
	var err error
	if err = validateDate(date); err != nil {
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.Stack(err)
	}

	var wishCard = new(wishCardEntity.Entity)
	err = i.masterTxManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
		wishCard, err = i.wishCardService.UpdateDate(ctx, masterTx, wishCardID, date)
		if err != nil {
			return werrors.Stack(err)
		}
		wishCard.Author, err = i.userService.GetByPK(ctx, masterTx, wishCard.Author.ID)
		if err != nil {
			return werrors.Stack(err)
		}
		// QUESTION: author == nilだった時？
		wishCard.Author.Profile, err = i.profileService.GetByUserID(ctx, masterTx, wishCard.Author.ID)
		if err != nil {
			return werrors.Stack(err)
		}
		wishCard.Place, err = i.placeService.GetByID(ctx, masterTx, wishCard.Place.ID)
		if err != nil {
			return werrors.Stack(err)
		}
		wishCard.Tags, err = i.tagService.GetByWishCardID(ctx, masterTx, wishCard.ID)
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

func (i *interactor) AddTags(ctx context.Context, userID, wishCardID int, tags []string) (*wishCardEntity.Entity, error) {
	var wishCard = new(wishCardEntity.Entity)
	var err error
	err = i.masterTxManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
		// create or get tagIDs
		tagIDs := make([]int, 0, len(tags))
		for _, tagName := range tags {
			var tag *tagEntity.Entity
			tag, err = i.tagService.GetByName(ctx, masterTx, tagName)
			if err != nil {
				return werrors.Stack(err)
			}
			if tag == nil {
				tag, err = i.tagService.Create(ctx, masterTx, tagName)
				if err != nil {
					return werrors.Stack(err)
				}
			}
			tagIDs = append(tagIDs, tag.ID)
		}
		// add relation
		wishCard, err = i.wishCardService.AddTags(ctx, masterTx, wishCardID, tagIDs)
		if err != nil {
			return werrors.Stack(err)
		}
		wishCard.Author, err = i.userService.GetByPK(ctx, masterTx, wishCard.Author.ID)
		if err != nil {
			return werrors.Stack(err)
		}
		// QUESTION: author == nilだった時？
		wishCard.Author.Profile, err = i.profileService.GetByUserID(ctx, masterTx, wishCard.Author.ID)
		if err != nil {
			return werrors.Stack(err)
		}
		wishCard.Place, err = i.placeService.GetByID(ctx, masterTx, wishCard.Place.ID)
		if err != nil {
			return werrors.Stack(err)
		}
		wishCard.Tags, err = i.tagService.GetByWishCardID(ctx, masterTx, wishCard.ID)
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

func (i *interactor) DeleteTags(ctx context.Context, userID, wishCardID int, tagIDs []int) (*wishCardEntity.Entity, error) {
	var wishCard = new(wishCardEntity.Entity)
	var err error
	err = i.masterTxManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
		// delete relation
		wishCard, err = i.wishCardService.DeleteTags(ctx, masterTx, wishCardID, tagIDs)
		if err != nil {
			return werrors.Stack(err)
		}
		wishCard.Author, err = i.userService.GetByPK(ctx, masterTx, wishCard.Author.ID)
		if err != nil {
			return werrors.Stack(err)
		}
		// QUESTION: author == nilだった時？
		wishCard.Author.Profile, err = i.profileService.GetByUserID(ctx, masterTx, wishCard.Author.ID)
		if err != nil {
			return werrors.Stack(err)
		}
		wishCard.Place, err = i.placeService.GetByID(ctx, masterTx, wishCard.Place.ID)
		if err != nil {
			return werrors.Stack(err)
		}
		wishCard.Tags, err = i.tagService.GetByWishCardID(ctx, masterTx, wishCard.ID)
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

func validateActivity(activity string) error {
	if activity == "" {
		err := errors.New("activity is empty error")
		return werrors.Newf(err, codes.InvalidArgument, http.StatusBadRequest, "「やりたいこと」は必須項目です。", "activity is required.")
	}
	if len(activity) > 50 {
		err := errors.New("activity is too long error")
		return werrors.Newf(err, codes.InvalidArgument, http.StatusBadRequest, "「やりたいこと」が長すぎます。", "activity is too long.")
	}
	return nil
}

func validateDescription(description string) error {
	if len(description) > 100 {
		err := errors.New("description is too long error")
		return werrors.Newf(err, codes.InvalidArgument, http.StatusBadRequest, "「詳細」が長すぎます。", "description is too long.")
	}
	return nil
}

func validatePlace(place string) error {
	if place == "" {
		err := errors.New("place is empty error")
		return werrors.Newf(err, codes.InvalidArgument, http.StatusBadRequest, "「場所」は必須項目です。", "place is required.")
	}
	if len(place) > 200 {
		err := errors.New("place is too long error")
		return werrors.Newf(err, codes.InvalidArgument, http.StatusBadRequest, "「場所」が長すぎます。", "place is too long.")
	}
	return nil
}

func validateDate(date *time.Time) error {
	if date == nil {
		err := errors.New("date is empty error")
		return werrors.Newf(err, codes.InvalidArgument, http.StatusBadRequest, "「日付」は必須項目です。", "date is required.")
	}
	if date.Before(time.Now()) {
		err := errors.New("date is in the past")
		return werrors.Newf(err, codes.InvalidArgument, http.StatusBadRequest, "過去の「日付」は指定できません。", "date is in the past.")
	}
	return nil
}

func validateTags(tags []string) error {
	for _, tag := range tags {
		if err := validateTag(tag); err != nil {
			return err
		}
	}
	return nil
}

func validateTag(tag string) error {
	if tag == "" {
		err := errors.New("tag is invalid error")
		return werrors.Newf(
			err,
			codes.InvalidArgument,
			http.StatusBadRequest,
			fmt.Sprintf("「%s」は無効なタグです。", tag),
			fmt.Sprintf("「%s」is invalid.", tag),
		)
	}
	if len(tag) > 100 {
		err := errors.New("tag is too long error")
		return werrors.Newf(err, codes.InvalidArgument, http.StatusBadRequest, "「タグ」が長すぎます。", "tag is too long.")
	}
	return nil
}
