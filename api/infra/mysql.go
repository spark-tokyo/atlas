package infra

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql" // MySQLドライバーのインポート 自動インポートされないので手動で追加

	"github.com/spark-tokyo/atlas/config"
	"github.com/spark-tokyo/atlas/ent"
)

type Ent struct {
	*ent.Client
}

func NewEnt(config *config.Config) (*Ent, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)

	// エンティティクライアントを作成
	client, err := Open(dsn)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	// defer client.Close()

	log.Println("Connected to MySQL database using ent")
	return &Ent{client}, nil
}

func Open(dsn string) (*ent.Client, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// データベースが利用可能になるまでリトライ
	for i := 0; i < 100; i++ {
		db, err = sql.Open("mysql", dsn)
		if err != nil {
			log.Printf("failed to connect to database, retrying in 2 seconds: %v", err)
			time.Sleep(2 * time.Second)
			continue
		}

		// データベースに Ping を送って接続確認
		if err = db.Ping(); err == nil {
			log.Println("Successfully connected to the database")
			break
		}

		log.Printf("failed to ping database, retrying in 2 seconds: %v", err)
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	drv := entsql.OpenDB("mysql", db)
	return ent.NewClient(ent.Driver(drv)), nil
}
