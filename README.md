# Wantum-server
Wantum（Androidアプリ）サーバサイド

## Documentation
- [esa](https://tech-cafe-freedom.esa.io/)
- [jira](https://wantum.atlassian.net/jira/software/projects/WAN/boards/1)
- [clientリポジトリ](https://github.com/TechCafeFreedom/Wantum-android)
- [ZenHub](https://app.zenhub.com/workspaces/wantum----5f12ff5c6ec353000f8ca0cb/board?epics:settings=epicsWithSubtasks&repos=277304125,278130353,280759607&showPRs=false)

## ブランチルール
- デフォルトブランチ＝`develop`
- 新規ブランチ作成時
  - `feat/[issue番号]/[タスクの内容（迷えばissueのタイトル)]`
- `develop`=開発環境
- `master`=プロダクション環境
- `master`や`develop`へのforced pushは🆖
- `Squash and merge`のみ許可。コミット履歴をきれいにまとめる。

## 初回起動（セットアップ）
1. `.env_example`をコピーして、`.env`ファイルを作成
2. 自分の環境（MySQL）に合わせて環境変数を書き換える
3. メンバーから`wantum-firebase-adminsdk-cz9e4-4c4789f0f4.json`(Firebaseの認証情報）をもらい、プロジェクトのルートディレクトリに配置する
4. `db/mysql/ddl/ddl.sql`をローカルのMySQLにてRUNする
5. `make run`コマンドでサーバが立ち上がる

## Makeコマンド
```shell script
help                           使い方
wiregen                        wire_gen.goの生成
test                           testの実行
lint                           lintの実行
fmt                            fmtの実行
fmt-lint                       fmtとlintの実行
run                            APIをビルドせずに立ち上げるコマンド
build                          APIをビルドして立ち上げるコマンド
```

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

``` shell script
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
