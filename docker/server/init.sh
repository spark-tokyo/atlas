#!/bin/sh
set -e

STAGE=${STAGE}

# 初期化作業
echo "Running initialization tasks..."
# STAGE=local の場合はサーバーをリッスンしない
# ローカル実行時はコンテナのサーバーだと起動が遅くなるので、別で make run_api でサーバーを起動する
# dev以上はクラウド内のため コンテナを使用
if [ "$STAGE" = "local" ]; then
  echo "STAGE is local. Skipping server start."
  exit 0
fi

echo "Starting server..."

# サーバーをリッスン
exec go run api/app/main.go
