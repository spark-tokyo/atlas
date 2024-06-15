package main

import (
	"context"
	"log"
	"os"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"github.com/spark-tokyo/atlas/ent"
)

/*
データベースをセットアップする。 動作確認などで必要になる
make local_db
中身を見たい場合は make connect_db で見れます
*/
func main() {
	// .envファイルを読み込む
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// 環境変数の読み込み
	mysqlUser := os.Getenv("MYSQL_ROOT")
	mysqlPassword := os.Getenv("MYSQL_ROOT_PASSWORD")
	mysqlHost := os.Getenv("HOST")
	mysqlPort := os.Getenv("MYSQL_PORT")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")

	// DSNの構築
	dataSourceName := mysqlUser + ":" + mysqlPassword + "@tcp(" + mysqlHost + ":" + mysqlPort + ")/" + mysqlDatabase + "?charset=utf8mb4&parseTime=True&loc=Local"
	// データベース接続設定
	db, err := sql.Open(dialect.MySQL, dataSourceName)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer db.Close()

	// クライアントの作成
	client := ent.NewClient(ent.Driver(db))
	defer client.Close()

	// コンテキストの作成
	ctx := context.Background()

	// スキーマの作成
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	log.Println("database successfully created")
}
