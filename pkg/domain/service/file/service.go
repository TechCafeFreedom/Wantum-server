package file

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"wantum/pkg/werrors"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
)

type Service interface {
	UploadImageToLocalFolder(uploadedFile []byte) (string, error)
}

type service struct {
}

func New() Service {
	return &service{}
}

/**
UploadImageToLocalFolder ローカルの files/ 配下にアップロードされた画像を保存する
成功したらファイルのパスを返す
*/
func (s *service) UploadImageToLocalFolder(uploadedFile []byte) (string, error) {
	folderName := "files"
	fileName, err := uuid.NewRandom() // ファイル名はUUIDで一意な名前を生成する
	if err != nil {
		errMessageJP := "アップロードされたファイルの名前生成に失敗しました"
		errMessageEn := "Error occurred when file name creating."
		return "", werrors.Newf(err, codes.Internal, http.StatusInternalServerError, errMessageJP, errMessageEn)
	}

	if err := os.MkdirAll(folderName, 0777); err != nil {
		return "", err
	}
	file, err := os.Create(filepath.Join("files", fmt.Sprintf("%s.png", fileName)))
	if err != nil {
		return "", werrors.FromConstant(err, werrors.ServerError)
	}
	defer file.Close()

	if _, err := file.Write(uploadedFile); err != nil {
		return "", werrors.FromConstant(err, werrors.ServerError)
	}
	return fmt.Sprintf("%s/%s.png", folderName, fileName), nil
}
