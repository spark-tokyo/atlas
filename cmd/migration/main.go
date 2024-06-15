package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	_ "ariga.io/atlas/sql/mysql"
	"ariga.io/atlas/sql/sqltool"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"github.com/spark-tokyo/atlas/ent/migrate"
)

/*
マイグレーションファイルを生成するコマンド
dbを更新したときにローカルで手動で実行する
コマンド= make migrate
※makefileを参照
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

	// dataSourceNameの構築
	dataSourceName := fmt.Sprintf("mysql://%s:%s@%s:%s/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlUser, mysqlPassword, mysqlHost, mysqlPort, mysqlDatabase)

	// Create a local migration directory able to understand Atlas migration file format for replay.
	dir, err := sqltool.NewGolangMigrateDir("./migrations")
	if err != nil {
		log.Fatalf("failed creating atlas migration directory: %v", err)
	}

	// Migrate diff options.
	opts := []schema.MigrateOption{
		schema.WithDir(dir), // provide migration directory
		schema.WithIndent(" "),
		schema.WithMigrationMode(schema.ModeInspect), // provide migration mode
		schema.WithDialect(dialect.MySQL),            // Ent dialect to use
		schema.WithFormatter(sqltool.GolangMigrateFormatter),
		schema.WithDropColumn(false),
		schema.WithDropIndex(true),
	}

	// アトラスがサポートするMySQLを使用してマイグレーションを生成する
	err = migrate.NamedDiff(
		context.Background(),
		dataSourceName,
		time.Now().Format("20060102150405"),
		opts...)
	if err != nil {
		log.Fatalf("failed generating migration file: %v", err)
	}
	log.Println("migration successfully created")
}
