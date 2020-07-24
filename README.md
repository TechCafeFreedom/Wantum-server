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
