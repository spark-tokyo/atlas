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

	// テストデータの作成を試みる
	// 例えば、Userエンティティが存在する場合、以下のようにテストデータを作成
	user, err := tx.User.Create().SetName("test user").SetAge(1).SetNickname("nickName").Save(ctx)
	if err != nil {
		t.Fatalf("Failed to create test data: %v", err)
	}

	// 作成したデータが正しく保存されたことを確認
	if user.Name != "test user" {
		t.Fatalf("Expected user name to be 'test user', got '%s'", user.Name)
	}
}
