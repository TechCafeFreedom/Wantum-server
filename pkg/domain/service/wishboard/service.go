package wishboard

import (
	"context"
	"time"
	"wantum/pkg/domain/entity/wishboard"
	"wantum/pkg/domain/repository"
	userwishboardrepository "wantum/pkg/domain/repository/userwishboard"
	wishboardrepository "wantum/pkg/domain/repository/wishboard"
	"wantum/pkg/werrors"
)

type Service interface {
	Create(ctx context.Context, masterTx repository.MasterTx, title, backgroundImageURL, inviteURL string, userID int) (*wishboard.Entity, error)
	GetByPK(ctx context.Context, masterTx repository.MasterTx, wishBoardID int) (*wishboard.Entity, error)
	GetByOwner(ctx context.Context, masterTx repository.MasterTx, userID int) (wishboard.EntitySlice, error)
	GetByMember(ctx context.Context, masterTx repository.MasterTx, userID int) (wishboard.EntitySlice, error)
	IsUserMember(ctx context.Context, masterTx repository.MasterTx, userID, wishBoardID int) (bool, error)
	UpdateTitle(ctx context.Context, masterTx repository.MasterTx, wishBoardID int, title string) error
	UpdateBackgroundImageURL(ctx context.Context, masterTx repository.MasterTx, wishBoardID int, backgroundImageURL string) error
	Delete(ctx context.Context, masterTx repository.MasterTx, wishBoardID int) error
}

type service struct {
	wishBoardRepository     wishboardrepository.Repository
	userWishBoardRepository userwishboardrepository.Repository
}

func New(wishBoardRepository wishboardrepository.Repository, userWishBoardRepository userwishboardrepository.Repository) Service {
	return &service{
		wishBoardRepository:     wishBoardRepository,
		userWishBoardRepository: userWishBoardRepository,
	}
}

func (s *service) Create(ctx context.Context, masterTx repository.MasterTx, title, backgroundImageURL, inviteURL string, userID int) (*wishboard.Entity, error) {
	// 現在時刻を取得
	now := time.Now()

	b := &wishboard.Entity{
		Title:              title,
		BackgroundImageURL: backgroundImageURL,
		InviteURL:          inviteURL,
		UserID:             userID,
		CreatedAt:          &now,
		UpdatedAt:          &now,
	}

	// WishBoardの新規作成
	b, err := s.wishBoardRepository.Insert(ctx, masterTx, b)
	if err != nil {
		return nil, werrors.Stack(err)
	}

	// UserとWishBoardのリレーションを作成
	err = s.userWishBoardRepository.Insert(ctx, masterTx, userID, b.ID)
	if err != nil {
		return nil, werrors.Stack(err)
	}

	return b, nil
}

func (s *service) GetByPK(ctx context.Context, masterTx repository.MasterTx, wishBoardID int) (*wishboard.Entity, error) {
	// WishBoardを主キーから取得
	b, err := s.wishBoardRepository.SelectByPK(ctx, masterTx, wishBoardID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return b, nil
}

func (s *service) GetByOwner(ctx context.Context, masterTx repository.MasterTx, userID int) (wishboard.EntitySlice, error) {
	// Userが所有するWishBoard一覧を取得（招待されているだけのものは含まない）
	bs, err := s.wishBoardRepository.SelectByUserID(ctx, masterTx, userID)
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return bs, err
}

func (s *service) GetByMember(ctx context.Context, masterTx repository.MasterTx, userID int) (wishboard.EntitySlice, error) {
	// Userが所属するWishBoardのIDをリストで取得（招待されているものも含む）
	wishBoardIDs, err := s.userWishBoardRepository.SelectByUserID(ctx, masterTx, userID)
	if err != nil {
		return nil, werrors.Stack(err)
	}

	// 存在しないなら空で返す
	if len(wishBoardIDs) == 0 {
		return nil, nil
	}

	// IDのリストをもとにWishBoardを複数取得
	bs, err := s.wishBoardRepository.SelectByPKs(ctx, masterTx, wishBoardIDs)
	if err != nil {
		return nil, werrors.Stack(err)
	}

	return bs, nil
}

func (s *service) IsUserMember(ctx context.Context, masterTx repository.MasterTx, userID, wishBoardID int) (bool, error) {
	// UserとWishBoardの間にリレーションはあるか？
	exists, err := s.userWishBoardRepository.Exists(ctx, masterTx, userID, wishBoardID)
	if err != nil {
		return false, werrors.Stack(err)
	}
	return exists, nil
}

func (s *service) UpdateTitle(ctx context.Context, masterTx repository.MasterTx, wishBoardID int, title string) error {
	// 現在時刻を取得
	now := time.Now()

	// WishBoardのタイトルを更新
	err := s.wishBoardRepository.UpdateTitle(ctx, masterTx, wishBoardID, title, &now)
	if err != nil {
		return werrors.Stack(err)
	}
	return nil
}

func (s *service) UpdateBackgroundImageURL(ctx context.Context, masterTx repository.MasterTx, wishBoardID int, backgroundImageURL string) error {
	// 現在時刻を取得
	now := time.Now()

	// WishBoardの背景画像URLを更新
	err := s.wishBoardRepository.UpdateBackgroundImageURL(ctx, masterTx, wishBoardID, backgroundImageURL, &now)
	if err != nil {
		return werrors.Stack(err)
	}
	return nil
}

func (s *service) Delete(ctx context.Context, masterTx repository.MasterTx, wishBoardID int) error {
	// 現在時刻を取得
	now := time.Now()

	b := &wishboard.Entity{ID: wishBoardID, UpdatedAt: &now, DeletedAt: &now}

	// WishBoardの削除
	err := s.wishBoardRepository.Delete(ctx, masterTx, b)
	if err != nil {
		return werrors.Stack(err)
	}
	return nil
}
