FROM golang:1.14

WORKDIR /go/src/app

ADD . .

# realizeはgomodに対応していないらしい
RUN GO111MODULE=off go get github.com/oxequa/realize \
    && go get -v ./... \
    && go build -o ./binary/app ./cmd