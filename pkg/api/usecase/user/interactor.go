package user

import (
	"context"
	"errors"
	"net/http"
	"wantum/pkg/domain/entity"
	"wantum/pkg/domain/repository"
	profileservice "wantum/pkg/domain/service/profile"
	userservice "wantum/pkg/domain/service/user"
	"wantum/pkg/werrors"
)

type Interactor interface {
	CreateNewUser(ctx context.Context, authID, userName, mail, name, thumbnail, bio, phone, place, birth string, gender int) (*entity.User, error)
	GetUserProfile(ctx context.Context, authID string) (*entity.User, error)
	GetAll(ctx context.Context) (entity.UserSlice, error)
}

type intereractor struct {
	masterTxManager repository.MasterTxManager
	userService     userservice.Service
	profileService  profileservice.Service
}

func New(masterTxManager repository.MasterTxManager, userService userservice.Service, profileService profileservice.Service) Interactor {
	return &intereractor{
		masterTxManager: masterTxManager,
		userService:     userService,
		profileService:  profileService,
	}
}

func (i *intereractor) CreateNewUser(ctx context.Context, authID, userName, mail, name, thumbnail, bio, phone, place, birth string, gender int) (*entity.User, error) {
	var createdUser *entity.User
	var err error
	err = i.masterTxManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
		// 新規ユーザ作成
		createdUser, err = i.userService.CreateNewUser(masterTx, authID, userName, mail)
		if err != nil {
			return werrors.Stack(err)
		}

		// 作成したユーザのプロフィール登録
		createdUser.Profile, err = i.profileService.CreateNewProfile(ctx, masterTx, createdUser.ID, name, thumbnail, bio, phone, place, birth, gender)
		if err != nil {
			return werrors.Stack(err)
		}

		return nil
	})
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return createdUser, nil
}

func (i *intereractor) GetUserProfile(ctx context.Context, authID string) (*entity.User, error) {
	var userData *entity.User
	var err error
	err = i.masterTxManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
		// ログイン済ユーザのプロフィール情報取得
		userData, err = i.userService.GetByAuthID(ctx, masterTx, authID)
		if err != nil {
			return werrors.Stack(err)
		}
		return nil
	})
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return userData, nil
}

func (i *intereractor) GetAll(ctx context.Context) (entity.UserSlice, error) {
	var userSlice entity.UserSlice
	var err error
	err = i.masterTxManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
		// (管理者用)ユーザ全件取得
		userSlice, err = i.userService.GetAll(ctx, masterTx)
		if err != nil {
			return werrors.Stack(err)
		}
		return nil
	})
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return userSlice, nil
}
