# MySQLの公式イメージをベースに使用
FROM mysql:8.2

# 初期化スクリプトをコピー
COPY ./docker/db/my.cnf /etc/mysql/conf.d/my.cnf
COPY ./docker/db/init.sh /docker-entrypoint-initdb.d

# スクリプトに実行権限を付与
RUN chmod +x /docker-entrypoint-initdb.d/init.sh

