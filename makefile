## バージョンは以下に固定する
# Go -> 1.22.0
# gqlgen -> 0.17.46
# wire -> 0.6.0
# ent -> 0.13.1
# mysql -> 8.0.33
# docker -> 25.0.7

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

# INSERT INTO users (id, age, name, nickname, email) VALUES ('1', 25, 'John Doe', 'johnd', 'john.doe@example.com');

# run apiで実行したときは http://127.0.0.1:8080/
run_api:
	STAGE=local go run api/app/main.go

# !WARNING: Dockerのリセット dockerが起動しない場合
docker_prune:
	docker image prune -a

# entの定義を使ってローカルDBをセットアップする
setup_local_db:
	go run -mod=mod ./cmd/localdb/main.go

# マイグレーションファイルを作成する
# entのスキーマとデータベースの差分をマイグレーションファイルとして書き出す
# Dockerを起動し、ローカルDBをセットアップした後に実装する
create_migration_file:
	go run -mod=mod ./cmd/migration/main.go asdff


#  npm install -g npx
# source ~/.nvm/nvm.sh
#  nvm install 18.18.0

