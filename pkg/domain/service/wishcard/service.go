package wishcard

import (
	"context"
	"fmt"
	"net/http"
	"time"
	wishCardEntity "wantum/pkg/domain/entity/wishcard"
	"wantum/pkg/domain/repository"
	"wantum/pkg/domain/repository/place"
	"wantum/pkg/domain/repository/profile"
	"wantum/pkg/domain/repository/tag"
	"wantum/pkg/domain/repository/user"
	"wantum/pkg/domain/repository/wishcard"
	"wantum/pkg/domain/repository/wishcardtag"
	"wantum/pkg/werrors"

	"google.golang.org/grpc/codes"
)

type Service interface {
	Create(ctx context.Context, masterTx repository.MasterTx, activity, description string, date *time.Time, userID, categoryID, placeID int, tagsIDs []int) (*wishCardEntity.Entity, error)
	Update(ctx context.Context, masterTx repository.MasterTx, wishCardID int, activity, description string, date, doneAt *time.Time, userID, placeID int) (*wishCardEntity.Entity, error)
	UpdateActivity(ctx context.Context, masterTx repository.MasterTx, wishCardID int, activity string) (*wishCardEntity.Entity, error)
	UpdateDescription(ctx context.Context, masterTx repository.MasterTx, wishCardID int, description string) (*wishCardEntity.Entity, error)
	UpdateDate(ctx context.Context, masterTx repository.MasterTx, wishCardID int, date *time.Time) (*wishCardEntity.Entity, error)
	UpdateDoneAt(ctx context.Context, masterTx repository.MasterTx, wishCardID int, doneAt *time.Time) (*wishCardEntity.Entity, error)
	UpdateAuthor(ctx context.Context, masterTx repository.MasterTx, wishCardID, userID int) (*wishCardEntity.Entity, error)
	UpdatePlace(ctx context.Context, masterTx repository.MasterTx, wishCardID, placeID int) (*wishCardEntity.Entity, error)
	UpdateCategory(ctx context.Context, masterTx repository.MasterTx, wishCardID, categoryID int) (*wishCardEntity.Entity, error)
	UpdateWithCategoryID(ctx context.Context, masterTx repository.MasterTx, wishCardID int, activity, description string, date, doneAt *time.Time, userID, categoryID, placeID int, tagIDs []int) (*wishCardEntity.Entity, error)
	Delete(ctx context.Context, masterTx repository.MasterTx, wishCardID int) error
	UpDeleteFlag(ctx context.Context, masterTx repository.MasterTx, wishCardID int) (*wishCardEntity.Entity, error)
	DownDeleteFlag(ctx context.Context, masterTx repository.MasterTx, wishCardID int) (*wishCardEntity.Entity, error)
	GetByID(ctx context.Context, masterTx repository.MasterTx, wishCardID int) (*wishCardEntity.Entity, error)
	GetByIDs(ctx context.Context, masterTx repository.MasterTx, wishCardIDs []int) (wishCardEntity.EntitySlice, error)
	GetByCategoryID(ctx context.Context, masterTx repository.MasterTx, categoryID int) (wishCardEntity.EntitySlice, error)
	AddTags(ctx context.Context, masterTx repository.MasterTx, wishCardID int, tagIDs []int) (*wishCardEntity.Entity, error)
	DeleteTags(ctx context.Context, masterTx repository.MasterTx, wishCardID int, tagIDs []int) (*wishCardEntity.Entity, error)
}

type service struct {
	userRepository        user.Repository
	userProfileRepository profile.Repository
	wishCardRepository    wishcard.Repository
	placeRepository       place.Repository
	tagsRepository        tag.Repository
	wishCardTagRepository wishcardtag.Repository
}

func New(wcRepo wishcard.Repository, userRepo user.Repository, upRepo profile.Repository, placeRepo place.Repository, tagRepo tag.Repository, wctRepo wishcardtag.Repository) Service {
	return &service{
		wishCardRepository:    wcRepo,
		userRepository:        userRepo,
		userProfileRepository: upRepo,
		placeRepository:       placeRepo,
		tagsRepository:        tagRepo,
		wishCardTagRepository: wctRepo,
	}
}

func (s *service) Create(ctx context.Context, masterTx repository.MasterTx, activity, description string, date *time.Time, userID, categoryID, placeID int, tagIDs []int) (*wishCardEntity.Entity, error) {
	// get user
	author, err := s.userRepository.SelectByPK(ctx, masterTx, userID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	authorProfile, err := s.userProfileRepository.SelectByUserID(ctx, masterTx, userID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	author.Profile = authorProfile

	// get place
	place, err := s.placeRepository.SelectByID(ctx, masterTx, placeID)
	if err != nil {
		return nil, werrors.Stack(err)
	}

	// get tag
	tags, err := s.tagsRepository.SelectByIDs(ctx, masterTx, tagIDs)
	if err != nil {
		return nil, werrors.Stack(err)
	}

	now := time.Now()
	wishCard := &wishCardEntity.Entity{
		Author:      author,
		Activity:    activity,
		Description: description,
		Date:        date,
		CreatedAt:   &now,
		UpdatedAt:   &now,
		Place:       place,
		Tags:        tags,
	}
	newID, err := s.wishCardRepository.Insert(ctx, masterTx, wishCard, categoryID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	wishCard.ID = newID

	// create relation
	if err = s.wishCardTagRepository.BulkInsert(ctx, masterTx, newID, tagIDs); err != nil {
		return nil, werrors.Stack(err)
	}

	return wishCard, nil
}

func (s *service) Update(ctx context.Context, masterTx repository.MasterTx, wishCardID int, activity, description string, date, doneAt *time.Time, userID, placeID int) (*wishCardEntity.Entity, error) {
	wishCard, err := s.wishCardRepository.SelectByID(ctx, masterTx, wishCardID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	author, err := s.userRepository.SelectByPK(ctx, masterTx, wishCard.Author.ID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	profile, err := s.userProfileRepository.SelectByUserID(ctx, masterTx, wishCard.Author.ID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	author.Profile = profile
	place, err := s.placeRepository.SelectByID(ctx, masterTx, placeID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	tags, err := s.tagsRepository.SelectByWishCardID(ctx, masterTx, wishCardID)
	if err != nil {
		return nil, werrors.Stack(err)
	}

	now := time.Now()
	wishCard.Author = author // NOTE: 今後、authorの更新があるかも
	wishCard.Activity = activity
	wishCard.Description = description
	wishCard.Date = date
	wishCard.DoneAt = doneAt
	wishCard.Place = place
	wishCard.UpdatedAt = &now
	wishCard.Tags = tags

	if err = s.wishCardRepository.Update(ctx, masterTx, wishCard); err != nil {
		return nil, werrors.Stack(err)
	}

	return wishCard, nil
}

func (s *service) UpdateActivity(ctx context.Context, masterTx repository.MasterTx, wishCardID int, activity string) (*wishCardEntity.Entity, error) {
	wishCard, err := s.wishCardRepository.SelectByID(ctx, masterTx, wishCardID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	if wishCard == nil {
		return nil, werrors.Stack(werrors.WishCardNotFound)
	}
	now := time.Now()
	wishCard.UpdatedAt = &now
	wishCard.Activity = activity
	if err := s.wishCardRepository.UpdateActivity(ctx, masterTx, wishCardID, wishCard.Activity, wishCard.UpdatedAt); err != nil {
		return nil, werrors.Stack(err)
	}
	return wishCard, nil
}

func (s *service) UpdateDescription(ctx context.Context, masterTx repository.MasterTx, wishCardID int, description string) (*wishCardEntity.Entity, error) {
	wishCard, err := s.wishCardRepository.SelectByID(ctx, masterTx, wishCardID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	if wishCard == nil {
		return nil, werrors.Stack(werrors.WishCardNotFound)
	}
	now := time.Now()
	wishCard.UpdatedAt = &now
	wishCard.Description = description
	if err := s.wishCardRepository.UpdateDescription(ctx, masterTx, wishCardID, wishCard.Description, wishCard.UpdatedAt); err != nil {
		return nil, werrors.Stack(err)
	}
	return wishCard, nil
}

func (s *service) UpdateDate(ctx context.Context, masterTx repository.MasterTx, wishCardID int, date *time.Time) (*wishCardEntity.Entity, error) {
	wishCard, err := s.wishCardRepository.SelectByID(ctx, masterTx, wishCardID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	if wishCard == nil {
		return nil, werrors.Stack(werrors.WishCardNotFound)
	}
	now := time.Now()
	wishCard.UpdatedAt = &now
	wishCard.Date = date
	if err := s.wishCardRepository.UpdateDate(ctx, masterTx, wishCardID, wishCard.Date, wishCard.UpdatedAt); err != nil {
		return nil, werrors.Stack(err)
	}
	return wishCard, nil
}

func (s *service) UpdateDoneAt(ctx context.Context, masterTx repository.MasterTx, wishCardID int, doneAt *time.Time) (*wishCardEntity.Entity, error) {
	wishCard, err := s.wishCardRepository.SelectByID(ctx, masterTx, wishCardID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	if wishCard == nil {
		return nil, werrors.Stack(werrors.WishCardNotFound)
	}
	now := time.Now()
	wishCard.UpdatedAt = &now
	wishCard.DoneAt = doneAt
	if err := s.wishCardRepository.UpdateDoneAt(ctx, masterTx, wishCardID, wishCard.DoneAt, wishCard.UpdatedAt); err != nil {
		return nil, werrors.Stack(err)
	}
	return wishCard, nil
}

func (s *service) UpdateAuthor(ctx context.Context, masterTx repository.MasterTx, wishCardID, userID int) (*wishCardEntity.Entity, error) {
	wishCard, err := s.wishCardRepository.SelectByID(ctx, masterTx, wishCardID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	if wishCard == nil {
		return nil, werrors.Stack(werrors.WishCardNotFound)
	}
	now := time.Now()
	wishCard.UpdatedAt = &now
	wishCard.Author.ID = userID
	if err := s.wishCardRepository.UpdateUserID(ctx, masterTx, wishCardID, wishCard.Author.ID, wishCard.UpdatedAt); err != nil {
		return nil, werrors.Stack(err)
	}
	return wishCard, nil
}

func (s *service) UpdatePlace(ctx context.Context, masterTx repository.MasterTx, wishCardID, placeID int) (*wishCardEntity.Entity, error) {
	wishCard, err := s.wishCardRepository.SelectByID(ctx, masterTx, wishCardID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	if wishCard == nil {
		return nil, werrors.Stack(werrors.WishCardNotFound)
	}
	now := time.Now()
	wishCard.UpdatedAt = &now
	wishCard.Place.ID = placeID
	if err := s.wishCardRepository.UpdatePlaceID(ctx, masterTx, wishCardID, wishCard.Place.ID, wishCard.UpdatedAt); err != nil {
		return nil, werrors.Stack(err)
	}
	return wishCard, nil
}

func (s *service) UpdateCategory(ctx context.Context, masterTx repository.MasterTx, wishCardID, categoryID int) (*wishCardEntity.Entity, error) {
	wishCard, err := s.wishCardRepository.SelectByID(ctx, masterTx, wishCardID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	if wishCard == nil {
		return nil, werrors.Stack(werrors.WishCardNotFound)
	}
	now := time.Now()
	wishCard.UpdatedAt = &now
	if err := s.wishCardRepository.UpdateCategoryID(ctx, masterTx, wishCardID, categoryID, wishCard.UpdatedAt); err != nil {
		return nil, werrors.Stack(err)
	}
	return wishCard, nil
}

// WARNING: 空値があった時、元データが消滅する。
func (s *service) UpdateWithCategoryID(ctx context.Context, masterTx repository.MasterTx, wishCardID int, activity, description string, date, doneAt *time.Time, userID, categoryID, placeID int, tagIDs []int) (*wishCardEntity.Entity, error) {
	wishCard, err := s.wishCardRepository.SelectByID(ctx, masterTx, wishCardID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	author, err := s.userRepository.SelectByPK(ctx, masterTx, wishCard.Author.ID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	profile, err := s.userProfileRepository.SelectByUserID(ctx, masterTx, wishCard.Author.ID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	author.Profile = profile
	place, err := s.placeRepository.SelectByID(ctx, masterTx, placeID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	tags, err := s.tagsRepository.SelectByIDs(ctx, masterTx, tagIDs)
	if err != nil {
		return nil, werrors.Stack(err)
	}

	now := time.Now()
	wishCard.Author = author // NOTE: 今後、authorの更新があるかも
	wishCard.Activity = activity
	wishCard.Description = description
	wishCard.Date = date
	wishCard.DoneAt = doneAt
	wishCard.Place = place
	wishCard.UpdatedAt = &now
	wishCard.Tags = tags

	if err = s.wishCardRepository.UpdateWithCategoryID(ctx, masterTx, wishCard, categoryID); err != nil {
		return nil, werrors.Stack(err)
	}

	// regist tag
	if err = s.wishCardTagRepository.DeleteByWishCardID(ctx, masterTx, wishCardID); err != nil {
		return nil, werrors.Stack(err)
	}
	if err = s.wishCardTagRepository.BulkInsert(ctx, masterTx, wishCardID, tagIDs); err != nil {
		return nil, werrors.Stack(err)
	}
	return wishCard, nil
}

func (s *service) UpDeleteFlag(ctx context.Context, masterTx repository.MasterTx, wishCardID int) (*wishCardEntity.Entity, error) {
	wishCard, err := s.wishCardRepository.SelectByID(ctx, masterTx, wishCardID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	author, err := s.userRepository.SelectByPK(ctx, masterTx, wishCard.Author.ID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	profile, err := s.userProfileRepository.SelectByUserID(ctx, masterTx, wishCard.Author.ID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	author.Profile = profile
	place, err := s.placeRepository.SelectByID(ctx, masterTx, wishCard.Place.ID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	tags, err := s.tagsRepository.SelectByWishCardID(ctx, masterTx, wishCard.ID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	now := time.Now()
	wishCard.UpdatedAt = &now
	wishCard.DeletedAt = &now
	wishCard.Author = author
	wishCard.Place = place
	wishCard.Tags = tags

	if err = s.wishCardRepository.UpDeleteFlag(ctx, masterTx, wishCard.ID, wishCard.UpdatedAt, wishCard.DeletedAt); err != nil {
		return nil, werrors.Stack(err)
	}
	return wishCard, nil
}

func (s *service) DownDeleteFlag(ctx context.Context, masterTx repository.MasterTx, wishCardID int) (*wishCardEntity.Entity, error) {
	wishCard, err := s.wishCardRepository.SelectByID(ctx, masterTx, wishCardID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	author, err := s.userRepository.SelectByPK(ctx, masterTx, wishCard.Author.ID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	profile, err := s.userProfileRepository.SelectByUserID(ctx, masterTx, wishCard.Author.ID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	author.Profile = profile
	place, err := s.placeRepository.SelectByID(ctx, masterTx, wishCard.Place.ID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	tags, err := s.tagsRepository.SelectByWishCardID(ctx, masterTx, wishCard.ID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	now := time.Now()
	wishCard.UpdatedAt = &now
	wishCard.DeletedAt = nil
	wishCard.Author = author
	wishCard.Place = place
	wishCard.Tags = tags
	if err = s.wishCardRepository.DownDeleteFlag(ctx, masterTx, wishCard.ID, wishCard.UpdatedAt); err != nil {
		return nil, werrors.Stack(err)
	}
	return wishCard, nil
}

func (s *service) Delete(ctx context.Context, masterTx repository.MasterTx, wishCardID int) error {
	wishCard, err := s.wishCardRepository.SelectByID(ctx, masterTx, wishCardID)
	if err != nil {
		return werrors.Stack(err)
	}
	if wishCard.DeletedAt == nil {
		return werrors.Newf(
			fmt.Errorf("can't delete this data. this data did not up a delete flag. wishCardID=%v", wishCardID),
			codes.FailedPrecondition,
			http.StatusBadRequest,
			"このデータは削除できません",
			"could not delete this place",
		)
	}
	if err = s.wishCardRepository.Delete(ctx, masterTx, wishCard); err != nil {
		return werrors.Stack(err)
	}
	if err = s.wishCardTagRepository.DeleteByWishCardID(ctx, masterTx, wishCardID); err != nil {
		return werrors.Stack(err)
	}
	return nil
}

func (s *service) GetByID(ctx context.Context, masterTx repository.MasterTx, wishCardID int) (*wishCardEntity.Entity, error) {
	wishCard, err := s.wishCardRepository.SelectByID(ctx, masterTx, wishCardID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	author, err := s.userRepository.SelectByPK(ctx, masterTx, wishCard.Author.ID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	profile, err := s.userProfileRepository.SelectByUserID(ctx, masterTx, wishCard.Author.ID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	author.Profile = profile
	place, err := s.placeRepository.SelectByID(ctx, masterTx, wishCard.Place.ID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	tags, err := s.tagsRepository.SelectByWishCardID(ctx, masterTx, wishCard.ID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	wishCard.Author = author
	wishCard.Place = place
	wishCard.Tags = tags
	return wishCard, nil
}

func (s *service) GetByIDs(ctx context.Context, masterTx repository.MasterTx, wishCardIDs []int) (wishCardEntity.EntitySlice, error) {

	wishCards, err := s.wishCardRepository.SelectByIDs(ctx, masterTx, wishCardIDs)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	// OPTIMIZE: 絶対遅い
	for _, wishCard := range wishCards {
		author, err := s.userRepository.SelectByPK(ctx, masterTx, wishCard.Author.ID)
		if err != nil {
			return nil, werrors.Stack(err)
		}
		profile, err := s.userProfileRepository.SelectByUserID(ctx, masterTx, wishCard.Author.ID)
		if err != nil {
			return nil, werrors.Stack(err)
		}
		author.Profile = profile
		place, err := s.placeRepository.SelectByID(ctx, masterTx, wishCard.Place.ID)
		if err != nil {
			return nil, werrors.Stack(err)
		}
		tags, err := s.tagsRepository.SelectByWishCardID(ctx, masterTx, wishCard.ID)
		if err != nil {
			return nil, werrors.Stack(err)
		}
		wishCard.Author = author
		wishCard.Place = place
		wishCard.Tags = tags
	}
	return wishCards, nil
}

func (s *service) GetByCategoryID(ctx context.Context, masterTx repository.MasterTx, categoryID int) (wishCardEntity.EntitySlice, error) {
	wishCards, err := s.wishCardRepository.SelectByCategoryID(ctx, masterTx, categoryID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	// OPTIMIZE: 絶対遅い
	for _, wishCard := range wishCards {
		author, err := s.userRepository.SelectByPK(ctx, masterTx, wishCard.Author.ID)
		if err != nil {
			return nil, werrors.Stack(err)
		}
		profile, err := s.userProfileRepository.SelectByUserID(ctx, masterTx, wishCard.Author.ID)
		if err != nil {
			return nil, werrors.Stack(err)
		}
		author.Profile = profile
		place, err := s.placeRepository.SelectByID(ctx, masterTx, wishCard.Place.ID)
		if err != nil {
			return nil, werrors.Stack(err)
		}
		tags, err := s.tagsRepository.SelectByWishCardID(ctx, masterTx, wishCard.ID)
		if err != nil {
			return nil, werrors.Stack(err)
		}
		wishCard.Author = author
		wishCard.Place = place
		wishCard.Tags = tags
	}
	return wishCards, nil
}

func (s *service) AddTags(ctx context.Context, masterTx repository.MasterTx, wishCardID int, tagIDs []int) (*wishCardEntity.Entity, error) {
	if err := s.wishCardTagRepository.BulkInsert(ctx, masterTx, wishCardID, tagIDs); err != nil {
		return nil, err
	}

	wishCard, err := s.wishCardRepository.SelectByID(ctx, masterTx, wishCardID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	author, err := s.userRepository.SelectByPK(ctx, masterTx, wishCard.Author.ID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	profile, err := s.userProfileRepository.SelectByUserID(ctx, masterTx, wishCard.Author.ID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	author.Profile = profile
	place, err := s.placeRepository.SelectByID(ctx, masterTx, wishCard.Place.ID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	tags, err := s.tagsRepository.SelectByWishCardID(ctx, masterTx, wishCard.ID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	wishCard.Author = author
	wishCard.Place = place
	wishCard.Tags = tags
	return wishCard, nil
}

func (s *service) DeleteTags(ctx context.Context, masterTx repository.MasterTx, wishCardID int, tagIDs []int) (*wishCardEntity.Entity, error) {
	if err := s.wishCardTagRepository.DeleteByIDs(ctx, masterTx, wishCardID, tagIDs); err != nil {
		return nil, werrors.Stack(err)
	}

	wishCard, err := s.wishCardRepository.SelectByID(ctx, masterTx, wishCardID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	author, err := s.userRepository.SelectByPK(ctx, masterTx, wishCard.Author.ID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	profile, err := s.userProfileRepository.SelectByUserID(ctx, masterTx, wishCard.Author.ID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	author.Profile = profile
	place, err := s.placeRepository.SelectByID(ctx, masterTx, wishCard.Place.ID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	tags, err := s.tagsRepository.SelectByWishCardID(ctx, masterTx, wishCard.ID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	wishCard.Author = author
	wishCard.Place = place
	wishCard.Tags = tags
	return wishCard, nil
}
