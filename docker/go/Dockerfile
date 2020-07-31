FROM golang:1.12

WORKDIR /go/src/wantum

ADD . .

ENV GO111MODULE=on

RUN GO111MODULE=off go get github.com/oxequa/realize \
    && go get -v ./... \
    && go build -o ./binary/app ./cmd