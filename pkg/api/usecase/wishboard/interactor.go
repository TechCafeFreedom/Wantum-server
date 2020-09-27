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
	UpdateTitle(ctx context.Context, wishBoardID int, title, authID string) error
	UpdateBackgroundImage(ctx context.Context, wishBoardID int, backgroundImage []byte, authID string) error
	DeleteWishBoard(ctx context.Context, wishBoardID int, authID string) error
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
	if err := validateTitle(ctx, title); err != nil {
		return nil, werrors.Stack(err)
	}

	var wishBoardEntity *wishboard.Entity
	err := i.masterTxManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
		userEntity, err := i.userService.GetByAuthID(ctx, masterTx, authID)
		if err != nil {
			return werrors.Stack(err)
		}

		// 背景画像を保存し、URLを取得
		backgroundImageURL, err := i.fileService.UploadImageToLocalFolder(backgroundImage)
		if err != nil {
			return werrors.Stack(err)
		}

		// TODO: 招待URLの自動生成
		inviteURL := "hoge"

		wishBoardEntity, err = i.wishBoardService.Create(ctx, masterTx, title, backgroundImageURL, inviteURL, userEntity.ID)
		if err != nil {
			return werrors.Stack(err)
		}

		return nil
	})
	if err != nil {
		return nil, werrors.Stack(err)
	}

	return wishBoardEntity, nil
}

func (i *interactor) GetMyWishBoards(ctx context.Context, authID string) (wishboard.EntitySlice, error) {
	var wishBoardSlice wishboard.EntitySlice
	err := i.masterTxManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
		userEntity, err := i.userService.GetByAuthID(ctx, masterTx, authID)
		if err != nil {
			return werrors.Stack(err)
		}

		// 自分が所属しているWishBoardのリストを取得
		wishBoardSlice, err = i.wishBoardService.GetMyBoards(ctx, masterTx, userEntity.ID)
		if err != nil {
			return werrors.Stack(err)
		}

		return nil
	})
	if err != nil {
		return nil, werrors.Stack(err)
	}

	return wishBoardSlice, nil
}

func (i *interactor) UpdateTitle(ctx context.Context, wishBoardID int, title, authID string) error {
	if err := validateTitle(ctx, title); err != nil {
		return werrors.Stack(err)
	}

	err := i.masterTxManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
		userEntity, err := i.userService.GetByAuthID(ctx, masterTx, authID)
		if err != nil {
			return werrors.Stack(err)
		}

		// WishBoardが存在するか確認
		wishBoardEntity, err := i.wishBoardService.GetByPK(ctx, masterTx, wishBoardID)
		if err != nil {
			return werrors.Stack(err)
		}

		// ユーザがWishBoardのメンバーでなければPermissionDenied
		isMember, err := i.wishBoardService.IsMember(ctx, masterTx, userEntity.ID, wishBoardEntity.ID)
		if err != nil {
			return werrors.Stack(err)
		}
		if !isMember {
			err := errors.New("Error occurred when update board title. cause: permission denied")
			tlog.PrintErrorLogWithCtx(ctx, err)
			return werrors.FromConstant(err, werrors.WishBoardPermissionDenied)
		}

		if err := i.wishBoardService.UpdateTitle(ctx, masterTx, wishBoardEntity.ID, title); err != nil {
			return werrors.Stack(err)
		}
		return nil
	})
	if err != nil {
		return werrors.Stack(err)
	}

	return nil
}

func (i *interactor) UpdateBackgroundImage(ctx context.Context, wishBoardID int, backgroundImage []byte, authID string) error {
	err := i.masterTxManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
		userEntity, err := i.userService.GetByAuthID(ctx, masterTx, authID)
		if err != nil {
			return werrors.Stack(err)
		}

		// WishBoardが存在するか確認
		wishBoardEntity, err := i.wishBoardService.GetByPK(ctx, masterTx, wishBoardID)
		if err != nil {
			return werrors.Stack(err)
		}

		// ユーザがWishBoardのメンバーでなければPermissionDenied
		isMember, err := i.wishBoardService.IsMember(ctx, masterTx, userEntity.ID, wishBoardEntity.ID)
		if err != nil {
			return werrors.Stack(err)
		}
		if !isMember {
			err := errors.New("Error occurred when update board background image. cause: permission denied")
			tlog.PrintErrorLogWithCtx(ctx, err)
			return werrors.FromConstant(err, werrors.WishBoardPermissionDenied)
		}

		// 背景画像を保存し、URLを取得
		backgroundImageURL, err := i.fileService.UploadImageToLocalFolder(backgroundImage)
		if err != nil {
			return werrors.Stack(err)
		}

		if err := i.wishBoardService.UpdateBackgroundImageURL(ctx, masterTx, wishBoardEntity.ID, backgroundImageURL); err != nil {
			return werrors.Stack(err)
		}

		return nil
	})
	if err != nil {
		return werrors.Stack(err)
	}

	return nil
}

func (i *interactor) DeleteWishBoard(ctx context.Context, wishBoardID int, authID string) error {
	err := i.masterTxManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
		userEntity, err := i.userService.GetByAuthID(ctx, masterTx, authID)
		if err != nil {
			return werrors.Stack(err)
		}

		// WishBoardが存在するか確認
		wishBoardEntity, err := i.wishBoardService.GetByPK(ctx, masterTx, wishBoardID)
		if err != nil {
			return werrors.Stack(err)
		}

		// ユーザがWishBoardのメンバーでなければPermissionDenied
		isMember, err := i.wishBoardService.IsMember(ctx, masterTx, userEntity.ID, wishBoardEntity.ID)
		if err != nil {
			return werrors.Stack(err)
		}
		if !isMember {
			err := errors.New("Error occurred when delete board. cause: permission denied")
			tlog.PrintErrorLogWithCtx(ctx, err)
			return werrors.FromConstant(err, werrors.WishBoardPermissionDenied)
		}

		if err := i.wishBoardService.Delete(ctx, masterTx, wishBoardEntity.ID); err != nil {
			return werrors.Stack(err)
		}

		return nil
	})
	if err != nil {
		return werrors.Stack(err)
	}

	return nil
}

func validateTitle(ctx context.Context, title string) error {
	if title == "" {
		err := errors.New("Error occurred when board title validation.")
		tlog.PrintErrorLogWithCtx(ctx, err)
		return werrors.FromConstant(err, werrors.BadRequest)
	}

	return nil
}
