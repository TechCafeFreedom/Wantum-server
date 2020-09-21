SOURCE_FILE := $(notdir $(source))
SOURCE_DIR := $(dir $(source))
PROTOS_DIR := ./Wantum-ProtocolBuffer
MOCK_FILE := mock_${SOURCE_FILE}
MOCK_DIR := ${SOURCE_DIR}mock_$(lastword $(subst /, ,${SOURCE_DIR}))/
MOCK_TARGET := $(lastword $(subst /, ,${SOURCE_DIR}))

GOLINT_FILE_NAME := golangci-lint
GOLINT_FILE_PATH := $(shell ls $(GOPATH)/bin | grep ${GOLINT_FILE_NAME})

# gocli-lintパッケージがダウンロード済みであればダウンロードを実行しない
define golintExist
    $(ifneq (${GOLINT_FILE_PATH},${GOLINT_FILE_NAME}),GO111MODULE=on go get github.com/golangci/golangci-lint/cmd/golangci-lint)
endef

help: ## 使い方
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

protoc: ## protoファイルから自動生成
	clang-format -i ${PROTOS_DIR}/*.proto

	protoc \
            -I ${PROTOS_DIR} \
            --go_out=plugins=grpc:pkg/pb/ \
            ${PROTOS_DIR}/*.proto \

	protoc -I ${PROTOS_DIR} --doc_out=html,index.html:./Wantum-ProtocolBuffer ${PROTOS_DIR}/*.proto

proto-fmt: ## protoファイルのリフォーマット
	clang-format -i ${PROTOS_DIR}/*.proto

proto-doc: ## protoから生成するProtoドキュメントの生成
	protoc -I ${PROTOS_DIR} --doc_out=html,index.html:./Wantum-ProtocolBuffer ${PROTOS_DIR}/*.proto

mockgen: ## mockgenの実行
	# Usege: make mockgen source=<インターフェースの定義しているファイル>

	# mockgenのインストール
	go install github.com/golang/mock/mockgen

	# mockgenの実行
	mockgen -source ${SOURCE_DIR}${SOURCE_FILE} -destination ${MOCK_DIR}${MOCK_FILE}

wiregen: ## REST,gRPC両方のwire_gen.goを生成
	# google/wireのインストール
	GO111MODULE=off go get -u github.com/google/wire

	# REST wire genの実行
	wire gen cmd/rest/wire.go

	# gRPC wire genの実行
	wire gen cmd/grpc/wire.go

wirerest: ## REST wire_gen.goの生成
	# google/wireのインストール
	GO111MODULE=off go get -u github.com/google/wire

	# wire genの実行
	wire gen cmd/rest/wire.go

wiregrpc: ## gRPC wire_gen.goの生成
	# google/wireのインストール
	GO111MODULE=off go get -u github.com/google/wire

	# wire genの実行
	wire gen cmd/grpc/wire.go

test: ## testの実行
	go test -v ./...

lint: ## lintの実行
	# golangci-lintのインストール(既にパッケージがあれば実行されない)
	${golintExist}

	# pkg配下をチェック。設定は .golangci.yml に記載
	golangci-lint run

fmt: ## fmtの実行
	# goimportsのインストール
	GO111MODULE=off go get -u golang.org/x/tools/cmd/goimports

	# tidy, fmt, goimportsの実行
	go mod tidy -v
	gofmt -s -w pkg/
	goimports -w pkg/

fmt-lint: fmt lint ## fmtとlintの実行

run: ## APIをビルドせずに立ち上げるコマンド
	go run ./cmd

build: ## APIをビルドして立ち上げるコマンド
	go build -o binary/wantum ./cmd
	./binary/wantum

dev-up: ## 全コンテナの起動
	docker-compose -f docker-compose.yml up -d

dev-stop: ## 全コンテナを止める
	docker-compose stop

dev-down: ## 全コンテナを停止し、かつイメージを削除
	docker-compose down --rmi local --volumes

dev-api-watch: ## apiのログを監視
	docker-compose logs -f api

dev-api-log: ## apiのログを出力
	docker-compose logs api

dev-db-watch: ## dbのログを監視
	docker-compose logs -f db

dev-db-log: ## dbのログを出力
	docker-compose logs db

dev-db-init: ## 環境内のDB初期化
	chmod u+x init-mysql.sh
	./init-mysql.sh

dev-db-dump: ## mysqlDump
	mkdir -p ./db/mysql/dump
	docker-compose exec db /usr/bin/mysqldump -u root -proot wantum > ./db/mysql/dump/mysqldump.sql
