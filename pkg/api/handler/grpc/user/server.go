package user

import (
	"context"
	userinteractor "wantum/pkg/api/usecase/user"
	"wantum/pkg/api/wcontext"
	"wantum/pkg/pb"
	"wantum/pkg/werrors"

	"github.com/golang/protobuf/ptypes/empty"
)

type Server struct {
	userInteractor userinteractor.Interactor
}

func New(userInteractor userinteractor.Interactor) Server {
	return Server{userInteractor: userInteractor}
}

func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
	authID, err := wcontext.GetAuthIDFromContext(ctx)
	if err != nil {
		return nil, werrors.FromConstant(err, werrors.AuthFail)
	}
	email, err := wcontext.GetEmailFromContext(ctx)
	if err != nil {
		email = ""
	}

	createdUser, err := s.userInteractor.CreateNewUser(
		ctx,
		authID,
		req.UserName,
		email,
		req.Name,
		req.Bio,
		req.Phone,
		req.Place,
		req.Thumbnail,
		int(req.Birth),
		int(req.Gender),
	)
	if err != nil {
		return nil, werrors.Stack(err)
	}

	return &pb.User{
		UserId:    int64(createdUser.ID),
		Name:      createdUser.Profile.Name,
		UserName:  createdUser.UserName,
		Thumbnail: createdUser.Profile.Thumbnail,
		Bio:       createdUser.Profile.Bio,
		Gender:    pb.GenderType(createdUser.Profile.Gender),
		Place:     createdUser.Profile.Place,
		Birth:     createdUser.Profile.Birth.Unix(),
	}, nil
}

func (s *Server) GetMyProfile(ctx context.Context, req *empty.Empty) (*pb.User, error) {
	authID, err := wcontext.GetAuthIDFromContext(ctx)
	if err != nil {
		return nil, werrors.FromConstant(err, werrors.AuthFail)
	}

	userData, err := s.userInteractor.GetAuthorizedUser(ctx, authID)
	if err != nil {
		return nil, werrors.Stack(err)
	}

	return &pb.User{
		UserId:    int64(userData.ID),
		Name:      userData.Profile.Name,
		UserName:  userData.UserName,
		Thumbnail: userData.Profile.Thumbnail,
		Bio:       userData.Profile.Bio,
		Gender:    pb.GenderType(userData.Profile.Gender),
		Place:     userData.Profile.Place,
		Birth:     userData.Profile.Birth.Unix(),
	}, nil
}

func (s *Server) GetUserProfile(ctx context.Context, request *pb.GetUserRequest) (*pb.User, error) {
	panic("implement me")
}

func (s *Server) UpdateUserProfile(ctx context.Context, request *pb.UpdateUserProfileRequest) (*pb.User, error) {
	panic("implement me")
}
