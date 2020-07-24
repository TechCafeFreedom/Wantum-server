package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"wantum/pkg/api/middleware"
	"wantum/pkg/api/request/reqbody"
	"wantum/pkg/api/response"
	userinteractor "wantum/pkg/api/usecase/user"
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

	uid, ok := r.Context().Value(middleware.AuthCtxKey).(string)
	if !ok {
		errMessageJP := "不正なユーザからのアクセスをブロックしました。"
		errMessageEN := "The content blocked because user is not certified."
		response.Error(w, r, werrors.Newf(errors.New("コンテキストのUIDキャストでエラーが発生しました。"), http.StatusUnauthorized, errMessageJP, errMessageEN))
		return
	}

	if err := s.userInteractor.CreateNewUser(r.Context(), uid, reqBody.Name, reqBody.Thumbnail); err != nil {
		response.Error(w, r, werrors.Stack(err))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) GetUserProfile(w http.ResponseWriter, r *http.Request) {
	uid, ok := r.Context().Value(middleware.AuthCtxKey).(string)
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

	response.JSON(w, r, user)
}

func (s *Server) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := s.userInteractor.GetAll(r.Context())
	if err != nil {
		response.Error(w, r, werrors.Stack(err))
		return
	}

	response.JSON(w, r, users)
}
