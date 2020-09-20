package wishboard

import (
	"context"
	"errors"
	"wantum/pkg/domain/entity/wishboard"
	"wantum/pkg/domain/repository"
	fileservice "wantum/pkg/domain/service/file"
	userservice "wantum/pkg/domain/service/user"
	wishboardservice "wantum/pkg/domain/service/wishboard"
	"wantum/pkg/tlog"
	"wantum/pkg/werrors"
)

type Interactor interface {
	CreateNewWishBoard(ctx context.Context, authID, title string, backgroundImage []byte) (*wishboard.Entity, error)
	GetMyWishBoards(ctx context.Context, authID string) (wishboard.EntitySlice, error)
	GetWishBoard(ctx context.Context, wishBoardID int, authID string) (*wishboard.Entity, error)
	UpdateWishBoard(ctx context.Context, wishBoardID int, title, backgroundImageUrl string) (*wishboard.Entity, error)
	DeleteWishBoard(ctx context.Context, wishBoardID int) error
}

type interactor struct {
	masterTxManager  repository.MasterTxManager
	userService      userservice.Service
	wishBoardService wishboardservice.Service
	fileService      fileservice.Service
}

func New(masterTxManager repository.MasterTxManager, userService userservice.Service, wishBoardService wishboardservice.Service, fileService fileservice.Service) Interactor {
	return &interactor{
		masterTxManager:  masterTxManager,
		userService:      userService,
		wishBoardService: wishBoardService,
		fileService:      fileService,
	}
}

func (i *interactor) CreateNewWishBoard(ctx context.Context, authID, title string, backgroundImage []byte) (*wishboard.Entity, error) {
	if title == "" {
		err := errors.New("title is empty")
		tlog.PrintErrorLogWithCtx(ctx, err)
		return nil, werrors.FromConstant(err, werrors.BadRequest)
	}

	var b *wishboard.Entity
	err := i.masterTxManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
		// ログイン済ユーザ情報の取得
		u, err := i.userService.GetByAuthID(ctx, masterTx, authID)
		if err != nil {
			return werrors.Stack(err)
		}

		backgroundImageUrl, err := i.fileService.UploadImageToLocalFolder(backgroundImage)
		if err != nil {
			return werrors.Stack(err)
		}

		inviteUrl := "hoge" // karioki

		b, err = i.wishBoardService.Create(ctx, masterTx, title, backgroundImageUrl, inviteUrl, u.ID)
		if err != nil {
			return werrors.Stack(err)
		}

		return nil
	})
	if err != nil {
		return nil, werrors.Stack(err)
	}

	return b, nil
}

func (i *interactor) GetMyWishBoards(ctx context.Context, authID string) (wishboard.EntitySlice, error) {
	var bs wishboard.EntitySlice
	err := i.masterTxManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
		// ログイン済ユーザ情報の取得
		u, err := i.userService.GetByAuthID(ctx, masterTx, authID)
		if err != nil {
			return werrors.Stack(err)
		}

		bs, err = i.wishBoardService.GetByUserID(ctx, masterTx, u.ID)
		if err != nil {
			return werrors.Stack(err)
		}

		return nil
	})
	if err != nil {
		return nil, werrors.Stack(err)
	}

	return bs, nil
}

func (i *interactor) GetWishBoard(ctx context.Context, wishBoardID int, authID string) (*wishboard.Entity, error) {
	var b *wishboard.Entity
	err := i.masterTxManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
		// ログイン済ユーザ情報の取得
		u, err := i.userService.GetByAuthID(ctx, masterTx, authID)
		if err != nil {
			return werrors.Stack(err)
		}

		isMember, err := i.wishBoardService.UserBelongs(ctx, masterTx, u.ID, wishBoardID)
		if err != nil {
			return werrors.Stack(err)
		}
		if !isMember {
			err := errors.New("you don't belong to wish_board")
			tlog.PrintErrorLogWithCtx(ctx, err)
			return werrors.FromConstant(err, werrors.WishBoardPermissionDenied)
		}

		b, err = i.wishBoardService.GetByPK(ctx, masterTx, wishBoardID)
		if err != nil {
			return werrors.Stack(err)
		}

		// やりたいことカテゴリー・カードの取得

		return nil
	})
	if err != nil {
		return nil, werrors.Stack(err)
	}

	return b, nil
}
