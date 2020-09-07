package user

import (
	"encoding/json"
	"net/http"
	"wantum/pkg/api/request"
	"wantum/pkg/api/request/reqbody"
	"wantum/pkg/api/response"
	userinteractor "wantum/pkg/api/usecase/user"
	"wantum/pkg/api/wcontext"
	"wantum/pkg/tlog"
	"wantum/pkg/werrors"
)

type Server struct {
	userInteractor userinteractor.Interactor
}

func New(userInteractor userinteractor.Interactor) Server {
	return Server{userInteractor: userInteractor}
}

func (s *Server) CreateNewUser(w http.ResponseWriter, r *http.Request) {
	authID, err := wcontext.GetAuthIDFromContext(r.Context())
	if err != nil {
		response.Error(w, r, werrors.FromConstant(err, werrors.AuthFail))
	}

	buf, err := request.BodyToBuffer(r)
	if err != nil {
		response.Error(w, r, werrors.Stack(err))
		return
	}

	var reqBody reqbody.UserCreate
	if err := json.Unmarshal(buf.Bytes(), &reqBody); err != nil {
		tlog.PrintErrorLogWithCtx(r.Context(), err)
		response.Error(w, r, werrors.FromConstant(err, werrors.BadRequest))
		return
	}

	createdUser, err := s.userInteractor.CreateNewUser(
		r.Context(),
		authID,
		reqBody.UserName,
		reqBody.Mail,
		reqBody.Name,
		reqBody.Bio,
		reqBody.Phone,
		reqBody.Place,
		nil,
		reqBody.Birth,
		reqBody.Gender,
	)
	if err != nil {
		response.Error(w, r, werrors.Stack(err))
		return
	}

	response.JSON(w, r, response.ConvertToUserResponse(createdUser))
}

func (s *Server) GetAuthorizedUser(w http.ResponseWriter, r *http.Request) {
	authID, err := wcontext.GetAuthIDFromContext(r.Context())
	if err != nil {
		response.Error(w, r, werrors.FromConstant(err, werrors.AuthFail))
	}

	user, err := s.userInteractor.GetAuthorizedUser(r.Context(), authID)
	if err != nil {
		response.Error(w, r, werrors.Stack(err))
		return
	}

	response.JSON(w, r, response.ConvertToUserResponse(user))
}

func (s *Server) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := s.userInteractor.GetAll(r.Context())
	if err != nil {
		response.Error(w, r, werrors.Stack(err))
		return
	}

	response.JSON(w, r, response.ConvertToUsersResponse(users))
}
