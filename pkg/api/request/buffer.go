package request

import (
	"bytes"
	"io"
	"net/http"
	"wantum/pkg/tlog"
	"wantum/pkg/werrors"
)

func BodyToBuffer(r *http.Request) (*bytes.Buffer, error) {
	body := r.Body
	defer body.Close()

	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, body); err != nil {
		tlog.PrintErrorLogWithCtx(r.Context(), err)
		return nil, werrors.FromConstant(err, werrors.BadRequest)
	}
	return buf, nil
}
