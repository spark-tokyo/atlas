#!/usr/bin/env sh
set -e

MYSQL_ROOT_PASSWORD=${MYSQL_ROOT_PASSWORD}
MYSQL_DATABASE=${MYSQL_DATABASE}

# MySQL サーバーの起動を待機
until mysqladmin ping -h"localhost" -u root -p"${MYSQL_ROOT_PASSWORD}" --silent; do
    echo 'waiting for mysql to be connectable...'
    sleep 2
done

# データベース操作
mysql -uroot -p${MYSQL_ROOT_PASSWORD} <<-EOSQL
    DROP DATABASE IF EXISTS ${MYSQL_DATABASE};
    CREATE DATABASE IF NOT EXISTS ${MYSQL_DATABASE};
EOSQL
