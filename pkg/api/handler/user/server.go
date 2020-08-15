package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"wantum/pkg/api/request/reqbody"
	"wantum/pkg/api/response"
	userinteractor "wantum/pkg/api/usecase/user"
	"wantum/pkg/constants"
	"wantum/pkg/werrors"
)

type Server struct {
	userInteractor userinteractor.Interactor
}

func New(userInteractor userinteractor.Interactor) Server {
	return Server{userInteractor: userInteractor}
}

func (s *Server) CreateNewUser(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	defer body.Close()

	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, body); err != nil {
		response.Error(w, r, werrors.Stack(err))
		return
	}

	var reqBody reqbody.UserCreate
	if err := json.Unmarshal(buf.Bytes(), &reqBody); err != nil {
		response.Error(w, r, werrors.Stack(err))
		return
	}

	authID, ok := r.Context().Value(constants.AuthCtxKey).(string)
	if !ok {
		errMessageJP := "不正なユーザからのアクセスをブロックしました。"
		errMessageEN := "The content blocked because user is not certified."
		response.Error(w, r, werrors.Newf(errors.New("コンテキストのUIDキャストでエラーが発生しました。"), http.StatusUnauthorized, errMessageJP, errMessageEN))
		return
	}

	createdUser, err := s.userInteractor.CreateNewUser(
		r.Context(),
		authID,
		reqBody.UserName,
		reqBody.Mail,
		reqBody.Name,
		reqBody.Thumbnail,
		reqBody.Bio,
		reqBody.Phone,
		reqBody.Place,
		reqBody.Birth,
		reqBody.Gender,
	)
	if err != nil {
		response.Error(w, r, werrors.Stack(err))
		return
	}

	response.JSON(w, r, response.ConvertToUserResponse(createdUser))
}

func (s *Server) GetUserProfile(w http.ResponseWriter, r *http.Request) {
	uid, ok := r.Context().Value(constants.AuthCtxKey).(string)
	if !ok {
		errMessageJP := "不正なユーザからのアクセスをブロックしました。"
		errMessageEN := "The content blocked because user is not certified."
		response.Error(w, r, werrors.Newf(errors.New("コンテキストのUIDキャストでエラーが発生しました。"), http.StatusUnauthorized, errMessageJP, errMessageEN))
		return
	}

	user, err := s.userInteractor.GetUserProfile(r.Context(), uid)
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
