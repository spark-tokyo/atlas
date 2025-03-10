###### このファイルの解説 ######
# makefile はスクリプトのショートカットを作成するファイルです。
# install: となっている場合、make install で下に書いてあるプログラムが順に実行されていきます。
###### このファイルの解説 ######

###### バージョン管理について ######
## gqlgenとwireのバージョンは最新にする(コードがエラーになる場合はcursorでチャットに聞きながら解決して見てください)
# latest は最新という意味です

# make installで最新バージョンをインストール
# gqlgen -> latest
# wire -> latest

# エラーになる場合更新すればいい(その場合@latestにしてインストール)
# ent -> 0.13.1
# バージョンを固定する(dockerは動かない場合docker desktopをアップデートしてみてください)
# mysql -> 8.0.33
# docker -> 25.0.7
###### バージョン管理について ######

# 必要なライブラリをインストールしましょう コマンド ⇨ make install
install:
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/99designs/gqlgen@latest


# gqlgenのスケルトン(プログラムを書くための土台)を生成する。
gqlgen:
	go run github.com/99designs/gqlgen generate

# クリーンアーキテクチャの依存関係を解決する
wire:
	cd api/app/di && wire gen -tags=wireinject

# DBのテーブルを定義するための ORM
entgen:
	go generate ./ent

# dockerの起動
docker_up:
	docker network create atlas || true
	STAGE=local docker-compose up -d --build

# dockerの終了
docker_down:
	docker-compose down

# docker内でmysqlの起動
docker_mysql:
	docker network create atlas || true
	docker-compose up -d --build mysql

# dockerを再起動
docker_res:
	docker network create atlas || true
	docker-compose down
	STAGE=local docker-compose up -d --build


# entの定義を使ってDocker内にDBをセットアップする(dockerを起動した後です)
setup_local_db:
	go run -mod=mod ./cmd/localdb/main.go

# マイグレーションファイルを作成する
# entのスキーマとデータベースの差分をマイグレーションファイルとして書き出す
# Dockerを起動し、ローカルDBをセットアップした後に実装する
create_migration_file:
	go run -mod=mod ./cmd/migration/main.go asdff

# ターミナル上でdockerのmysqlに接続
###### docker composeが使えない場合 ######
# 下記のコマンドで docker-compose をダウンロードしてくだい。V2の docker compose がある人はそちらでも大丈夫です
# brew install docker-compose
###### docker composeが使えない場合 ######
connect_db:
	docker-compose exec mysql mysql -uroot -p atlas

# サンプルデータが欲しい場合は、docker内のmysqlに接続した状態で以下のSQLを実行してください。
# INSERT INTO users (id, age, name, nickname, email) VALUES ('1', 25, 'John Doe', 'johnd', 'john.doe@example.com');

# run apiで実行した後 http://127.0.0.1:8080/ に接続
run_api:
	STAGE=local go run api/app/main.go

# !WARNING: Dockerのリセット dockerが起動しない場合
docker_prune:
	docker image prune -a




#  npm install -g npx
# source ~/.nvm/nvm.sh
#  nvm install 18.18.0

