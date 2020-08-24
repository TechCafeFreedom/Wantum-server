//+build wireinject

package main

import (
	userHandler "wantum/pkg/api/handler/user"
	userInteractor "wantum/pkg/api/usecase/user"
	"wantum/pkg/domain/repository"
	profileSvc "wantum/pkg/domain/service/profile"
	userSvc "wantum/pkg/domain/service/user"
	profileRepo "wantum/pkg/infrastructure/mysql/profile"
	userRepo "wantum/pkg/infrastructure/mysql/user"

	"github.com/google/wire"
)

func InitUserAPI(masterTxManager repository.MasterTxManager) userHandler.Server {
	wire.Build(userRepo.New, profileRepo.New, profileSvc.New, userSvc.New, userInteractor.New, userHandler.New)

	return userHandler.Server{}
}
