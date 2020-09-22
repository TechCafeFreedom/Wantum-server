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
	UpdateTitle(ctx context.Context, wishBoardID int, title, authID string) error
	UpdateBackgroundImage(ctx context.Context, wishBoardID int, backgroundImage []byte, authID string) error
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
	// 空のタイトルは許容しない
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

		// 背景画像を保存し、URLを取得
		backgroundImageUrl, err := i.fileService.UploadImageToLocalFolder(backgroundImage)
		if err != nil {
			return werrors.Stack(err)
		}

		// TODO: 招待URLの自動生成
		inviteUrl := "hoge" // karioki

		// WishBoardの新規作成
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

		// 自分が所属しているWishBoardのリストを取得
		bs, err = i.wishBoardService.GetByMember(ctx, masterTx, u.ID)
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

		// ユーザがWishBoardのメンバーでなければPermissionDenied
		isMember, err := i.wishBoardService.UserBelongs(ctx, masterTx, u.ID, wishBoardID)
		if err != nil {
			return werrors.Stack(err)
		}
		if !isMember {
			err := errors.New("you don't belong to wish_board")
			tlog.PrintErrorLogWithCtx(ctx, err)
			return werrors.FromConstant(err, werrors.WishBoardPermissionDenied)
		}

		// WishBoardを取得
		b, err = i.wishBoardService.GetByPK(ctx, masterTx, wishBoardID)
		if err != nil {
			return werrors.Stack(err)
		}

		// TODO: WishCategory, WishCardの取得

		return nil
	})
	if err != nil {
		return nil, werrors.Stack(err)
	}

	return b, nil
}

func (i *interactor) UpdateTitle(ctx context.Context, wishBoardID int, title, authID string) error {
	// 空のタイトルは許容しない
	if title == "" {
		err := errors.New("title is empty")
		tlog.PrintErrorLogWithCtx(ctx, err)
		return werrors.FromConstant(err, werrors.BadRequest)
	}

	err := i.masterTxManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
		// ログイン済ユーザ情報の取得
		u, err := i.userService.GetByAuthID(ctx, masterTx, authID)
		if err != nil {
			return werrors.Stack(err)
		}

		// WishBoardが存在するか確認
		b, err := i.wishBoardService.GetByPK(ctx, masterTx, wishBoardID)
		if err != nil {
			return werrors.Stack(err)
		}

		// ユーザがWishBoardのメンバーでなければPermissionDenied
		isMember, err := i.wishBoardService.UserBelongs(ctx, masterTx, u.ID, b.ID)
		if err != nil {
			return werrors.Stack(err)
		}
		if !isMember {
			err := errors.New("you don't belong to wish_board")
			tlog.PrintErrorLogWithCtx(ctx, err)
			return werrors.FromConstant(err, werrors.WishBoardPermissionDenied)
		}

		// タイトル更新
		err = i.wishBoardService.UpdateTitle(ctx, masterTx, b.ID, title)
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

func (i *interactor) UpdateBackgroundImage(ctx context.Context, wishBoardID int, backgroundImage []byte, authID string) error {
	err := i.masterTxManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
		// ログイン済ユーザ情報の取得
		u, err := i.userService.GetByAuthID(ctx, masterTx, authID)
		if err != nil {
			return werrors.Stack(err)
		}

		// WishBoardが存在するか確認
		b, err := i.wishBoardService.GetByPK(ctx, masterTx, wishBoardID)
		if err != nil {
			return werrors.Stack(err)
		}

		// ユーザがWishBoardのメンバーでなければPermissionDenied
		isMember, err := i.wishBoardService.UserBelongs(ctx, masterTx, u.ID, b.ID)
		if err != nil {
			return werrors.Stack(err)
		}
		if !isMember {
			err := errors.New("you don't belong to wish_board")
			tlog.PrintErrorLogWithCtx(ctx, err)
			return werrors.FromConstant(err, werrors.WishBoardPermissionDenied)
		}

		// 背景画像を保存し、URLを取得
		backgroundImageUrl, err := i.fileService.UploadImageToLocalFolder(backgroundImage)
		if err != nil {
			return werrors.Stack(err)
		}

		// 背景画像URLの更新
		err = i.wishBoardService.UpdateBackgroundImageUrl(ctx, masterTx, b.ID, backgroundImageUrl)
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
