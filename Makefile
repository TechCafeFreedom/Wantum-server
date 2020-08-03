dev-up: ## 全コンテナの起動
	docker-compose up -d

dev-stop: ## 全コンテナを止める
	docker-compose stop

dev-down: ## 全コンテナを停止し、かつイメージを削除
	docker-compose down --rmi local

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
