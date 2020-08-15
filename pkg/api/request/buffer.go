package request

import (
	"bytes"
	"io"
	"net/http"
	"wantum/pkg/tlog"
	"wantum/pkg/werrors"
)

func BodyToBuffer(w http.ResponseWriter, r *http.Request) (*bytes.Buffer, error) {
	body := r.Body
	defer body.Close()

	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, body); err != nil {
		tlog.PrintErrorLogWithCtx(r.Context(), err)
		return nil, werrors.Wrapf(err, http.StatusBadRequest, "リクエストされたユーザ情報が空でした", "requested user data is empty")
	}
	return buf, nil
}
