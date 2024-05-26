
# 必要なライブラリをインストールしましょう コマンド ⇨ make install
install:
	go install github.com/google/wire/cmd/wire@v0.6.0
	go install github.com/99designs/gqlgen@v0.17.46

gqlgen:
	go run github.com/99designs/gqlgen generate

wire:
	cd api/app/di && wire gen -tags=wireinject

ent_gen:
	go generate ./ent

doc_up:
	docker network create atlas || true
	STAGE=local docker compose up -d --build

doc_down:
	docker compose down

doc_mysql:
	docker network create atlas || true
	docker compose up -d --build mysql

doc_res:
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