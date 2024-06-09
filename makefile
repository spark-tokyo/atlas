
# 必要なライブラリをインストールしましょう コマンド ⇨ make install
install:
	go install github.com/google/wire/cmd/wire@v0.6.0
	go install github.com/99designs/gqlgen@v0.17.46

gqlgen:
	go run github.com/99designs/gqlgen generate

wire:
	cd api/app/di && wire gen -tags=wireinject

entgen:
	go generate ./ent

docker_up:
	docker network create atlas || true
	STAGE=local docker compose up -d --build

docker_down:
	docker compose down

docker_mysql:
	docker network create atlas || true
	docker compose up -d --build mysql

docker_res:
	docker network create atlas || true
	docker compose down
	STAGE=local docker compose up -d --build

connect_db:
	docker-compose exec mysql mysql -uroot -p $(MYSQL_DATABASE)

# run apiで実行したときは http://127.0.0.1:8080/
run_api:
	STAGE=local go run api/app/main.go

# !WARNING: Dockerのリセット dockerが起動しない場合
docker_prune:
	docker image prune -a

# ローカルDBをセットアップする
setup_local_db:
	go run -mod=mod ./cmd/localdb/main.go

# マイグレーションファイルを作成する (DBを更新したときに差分としてファイルを作成する)
create_migration_file:
	go run -mod=mod ./cmd/migration/main.go asdff

