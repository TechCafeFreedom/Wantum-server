package user

import (
	"context"
	"errors"
	"net/http"
	"time"
	"wantum/pkg/domain/entity"
	"wantum/pkg/domain/repository"
	fileservice "wantum/pkg/domain/service/file"
	profileservice "wantum/pkg/domain/service/profile"
	userservice "wantum/pkg/domain/service/user"
	"wantum/pkg/tlog"
	"wantum/pkg/werrors"

	"google.golang.org/grpc/codes"
)

type Interactor interface {
	CreateNewUser(ctx context.Context, authID, userName, mail, name, bio, phone, place string, thumbnail []byte, birth, gender int) (*entity.User, error)
	GetAuthorizedUser(ctx context.Context, authID string) (*entity.User, error)
	GetAll(ctx context.Context) (entity.UserMap, error)
}

type intereractor struct {
	masterTxManager repository.MasterTxManager
	userService     userservice.Service
	profileService  profileservice.Service
	fileService     fileservice.Service
}

func New(masterTxManager repository.MasterTxManager, userService userservice.Service, profileService profileservice.Service, fileService fileservice.Service) Interactor {
	return &intereractor{
		masterTxManager: masterTxManager,
		userService:     userService,
		profileService:  profileService,
		fileService:     fileService,
	}
}

func (i *intereractor) CreateNewUser(ctx context.Context, authID, userName, mail, name, bio, phone, place string, thumbnail []byte, birth, gender int) (*entity.User, error) {
	if mail == "" {
		err := errors.New("mail is empty error")
		tlog.PrintErrorLogWithAuthID(authID, err)
		return nil, werrors.Newf(err, codes.InvalidArgument, http.StatusBadRequest, "メール情報は必須項目です。", "mail is required.")
	}

	birthDate := time.Unix(int64(birth), 0)

	var createdUser *entity.User
	err := i.masterTxManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
		// サムネイル画像のアップロード
		thumbnailURL, err := i.fileService.UploadImageToLocalFolder(thumbnail)
		if err != nil {
			return werrors.Stack(err)
		}

		// 新規ユーザ作成
		createdUser, err = i.userService.CreateNewUser(masterTx, authID, userName, mail)
		if err != nil {
			return werrors.Stack(err)
		}

		// 作成したユーザのプロフィール登録
		createdUser.Profile, err = i.profileService.CreateNewProfile(ctx, masterTx, createdUser.ID, name, thumbnailURL, bio, phone, place, &birthDate, gender)
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

func (i *intereractor) GetAuthorizedUser(ctx context.Context, authID string) (*entity.User, error) {
	var userData *entity.User
	var err error
	err = i.masterTxManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
		// ログイン済ユーザ情報の取得
		userData, err = i.userService.GetByAuthID(ctx, masterTx, authID)
		if err != nil {
			return werrors.Stack(err)
		}

		// ログイン済ユーザのプロフィール情報取得
		userProfile, err := i.profileService.GetByUserID(ctx, masterTx, userData.ID)
		if err != nil {
			return werrors.Stack(err)
		}
		userData.Profile = userProfile
		return nil
	})
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return userData, nil
}

func (i *intereractor) GetAll(ctx context.Context) (entity.UserMap, error) {
	userMap := make(entity.UserMap)
	err := i.masterTxManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
		// (管理者用)ユーザ全件取得
		userSlice, err := i.userService.GetAll(ctx, masterTx)
		if err != nil {
			return werrors.Stack(err)
		}

		for _, userData := range userSlice {
			userMap[userData.ID] = userData
		}

		// 取得したユーザそれぞれに紐づくプロフィール情報を取得
		userIDs := userMap.Keys(userMap)
		profileSlice, err := i.profileService.GetByUserIDs(ctx, masterTx, userIDs)
		if err != nil {
			return werrors.Stack(err)
		}

		// 取得したプロフィール情報のスライスをユーザのスライスにアサイン
		// TODO: 現状の設計だと、UserDataに対してProfile情報が必ず1つだけ存在しないとデータの整合性が保てない設計となっている。
		for _, profileData := range profileSlice {
			userMap[profileData.UserID].Profile = profileData
		}

		return nil
	})
	if err != nil {
		return nil, werrors.Stack(err)
	}
	return userMap, nil
}
