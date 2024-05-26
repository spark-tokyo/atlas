FROM golang:1.22.0-alpine3.19 as local

RUN apk add --no-cache ca-certificates git

# Air および Delve のインストール（開発用）
RUN go install github.com/cosmtrek/air@v1.51.0
RUN go install github.com/go-delve/delve/cmd/dlv@v1.22.0

COPY ./docker/server/init.sh /usr/local/bin/init.sh

# スクリプトに実行権限を付与
RUN chmod +x /usr/local/bin/init.sh

WORKDIR /app

COPY . .

# init.sh スクリプトをエントリーポイントとして設定
ENTRYPOINT ["/usr/local/bin/init.sh"]
