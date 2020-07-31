# Wantum-server

## 開発環境

- docker-compose 3.5
- containers
	- golang 1.12
	- mysql 5.7

ホットリロード対応。[oxequa/realize](https://github.com/oxequa/realize)を利用しています。

### 作業まわり
#### 起動 & 停止
- `make env-up`にて、全コンテナを起動
- `make env-stop`にて、全コンテナの停止
- `make env-down`にて、imageを含めたコンテナの停止&削除

#### 初期化
- 環境の初期化
	1. `make env-down`にて、全コンテナを停止&削除
	2. `make env-up`にて、起動
- DBの初期化
	1. `make env-db-init`にて、初期化

#### ログまわり
監視

- `make env-api-watch`にて、APIのログを監視
- `make env-db-watch`にて、DBのログを監視

出力。パイプなどで繋いで処理可。

- `make env-api-log`にて、APIのログを全取得
- `make env-db-log`にて、DBのログを全取得

### Makeコマンド

接頭辞に`env-`がついています。

```
env-up        環境をバックグラウンドで起動
env-stop      環境の停止
env-down      環境の停止&削除
env-api-watch apiのログを監視
env-api-log   apiのログを吐く
env-db-watch  DBのログを監視
env-db-log    DBのログを吐く
env-db-init   DBの初期化する。DBをDROPしてから再構築する
env-db-dump   DBをダンプする。出力先は /db/mysql/dump
```
