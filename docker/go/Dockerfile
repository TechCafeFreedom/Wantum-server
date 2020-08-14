FROM golang:1.12-alpine as build

WORKDIR /go/src/wantum

ADD . .

ENV GO111MODULE=on

RUN apk update \
    && apk add git \
    && GO111MODULE=off go get github.com/oxequa/realize \
    && go get -v ./... \
    && go build -o ./binary/app ./cmd

FROM alpine

WORKDIR /wantum

COPY --from=build /go/src/wantum .

RUN addgroup go \
  && adduser -D -G go go \
  && chown -R go:go /wantum/binary/app

CMD ["./binary/app"]
