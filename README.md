# Wantum-server

## 開発環境

- docker-compose 3.5
- containers
	- golang 1.12
	- mysql 5.7

ホットリロード対応。[oxequa/realize](https://github.com/oxequa/realize)を利用しています。

### 作業まわり
#### 起動 & 停止
- `make dev-up`にて、全コンテナを構築&起動
- `make dev-stop`にて、全コンテナの停止
- `make dev-down`にて、imageを含めたコンテナの停止&削除

#### 初期化
- 環境の初期化
	1. `make dev-down`にて、全コンテナを停止&削除
	2. `make dev-up`にて、起動
- DBの初期化
	1. `make dev-db-init`にて、データベースの初期化

#### ログまわり
監視

- `make dev-api-watch`にて、APIのログを監視
- `make dev-db-watch`にて、DBのログを監視

出力。パイプなどで繋いで処理可。

- `make dev-api-log`にて、APIのログを全取得
- `make dev-db-log`にて、DBのログを全取得

### Makeコマンド

接頭辞に`dev-`がついています。

```
dev-up        環境をバックグラウンドで構築&起動
dev-stop      環境の停止
dev-down      環境の停止&削除
dev-api-watch apiのログを監視
dev-api-log   apiのログを吐く
dev-db-watch  DBのログを監視
dev-db-log    DBのログを吐く
dev-db-init   DBの初期化する。DBをDROPしてから再構築する
dev-db-dump   DBをダンプする。出力先は /db/mysql/dump。ディレクトリは自動で作成されます。
```
