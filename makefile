
# 必要なライブラリをインストールしましょう コマンド ⇨ make install
install:
	go install github.com/google/wire/cmd/wire@v0.6.0
	go install github.com/99designs/gqlgen@v0.17.46

gqlgen:
	go run github.com/99designs/gqlgen generate
