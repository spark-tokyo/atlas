package test

import (
	"context"
	"testing"

	_ "github.com/go-sql-driver/mysql" // MySQLドライバーのインポート 自動インポートされないので手動で追加
	_ "github.com/mattn/go-sqlite3"
)

func TestNewTestClient(t *testing.T) {
	ctx := context.Background()
	client := NewTestClient(ctx, t)

	// クライアントがnilではないことを確認
	if client == nil {
		t.Fatal("Expected non-nil client")
	}

	defer client.Client.Close()

	// データベースに接続できることを確認
	err := client.Client.Schema.Create(ctx)
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
}

func TestFetchTestReadWriteTransaction(t *testing.T) {
	ctx := context.Background()
	tx, cleanup := FetchTestReadWriteTransaction(ctx, t)
	defer cleanup()

	// トランザクションがnilではないことを確認
	if tx == nil {
		t.Fatal("Expected non-nil transaction")
	}
}
