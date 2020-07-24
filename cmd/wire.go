//+build wireinject

package main

import (
	userHandler "wantum/pkg/api/handler/user"
	userInteractor "wantum/pkg/api/usecase/user"
	"wantum/pkg/domain/repository"
	userSvc "wantum/pkg/domain/service/user"
	userRepo "wantum/pkg/infrastructure/mysql/user"

	"github.com/google/wire"
)

func InitUserAPI(masterTxManager repository.MasterTxManager) userHandler.Server {
	wire.Build(userRepo.New, userSvc.New, userInteractor.New, userHandler.New)

	return userHandler.Server{}
}
