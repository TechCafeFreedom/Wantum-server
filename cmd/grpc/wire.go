//+build wireinject

package rest

import (
	userHandler "wantum/pkg/api/handler/grpc/user"
	userInteractor "wantum/pkg/api/usecase/user"
	"wantum/pkg/domain/repository"
	fileSvc "wantum/pkg/domain/service/file"
	profileSvc "wantum/pkg/domain/service/profile"
	userSvc "wantum/pkg/domain/service/user"
	profileRepo "wantum/pkg/infrastructure/mysql/profile"
	userRepo "wantum/pkg/infrastructure/mysql/user"

	"github.com/google/wire"
)

func InitUserAPI(masterTxManager repository.MasterTxManager) userHandler.Server {
	wire.Build(userRepo.New, profileRepo.New, fileSvc.New, profileSvc.New, userSvc.New, userInteractor.New, userHandler.New)

	return userHandler.Server{}
}
