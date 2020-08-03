env-up: ## 全コンテナの起動
	docker-compose up -d

env-stop: ## 全コンテナを止める
	docker-compose stop

env-down: ## 全コンテナを停止し、かつイメージを削除
	docker-compose down --rmi local

env-api-watch: ## apiのログを監視
	docker-compose logs -f api

env-api-log: ## apiのログを出力
	docker-compose logs api

env-db-watch: ## dbのログを監視
	docker-compose logs -f db

env-db-log: ## dbのログを出力
	docker-compose logs log

env-db-init: ## 環境内のDB初期化
	chmod u+x init-mysql.sh
	./init-mysql.sh

env-db-dump: ## mysqlDump
	mkdir -p ./db/mysql/dump
	docker-compose exec db /usr/bin/mysqldump -u root -proot wantum > /dev/null > ./db/mysql/dump/mysqldump.sql
