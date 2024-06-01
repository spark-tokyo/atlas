package main

import (
	"context"
	"fmt"
	"log"
	"os"

	atlasMigrate "ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"

	"github.com/spark-tokyo/atlas/ent/migrate"
)

func main() {
	dataSourceName := fmt.Sprintf("mysql://%s:%s@%s/%s?charset=utf8&parseTime=True", "sample_user", "sample_password", "localhost:3306", "sample_db")
	ctx := context.Background()

	// Create a local migration directory able to understand Atlas migration file format for replay.
	dir, err := atlasMigrate.NewLocalDir("ent/migrate/migrations")
	if err != nil {
		log.Fatalf("failed creating atlas migration directory: %v", err)
	}

	// Migrate diff options.
	opts := []schema.MigrateOption{
		schema.WithDir(dir),                          // provide migration directory
		schema.WithMigrationMode(schema.ModeInspect), // provide migration mode
		schema.WithDialect(dialect.MySQL),            // Ent dialect to use
		schema.WithFormatter(atlasMigrate.DefaultFormatter),
	}

	// Generate migrations using Atlas support for MySQL (note the Ent dialect option passed above).
	err = migrate.NamedDiff(ctx, dataSourceName, os.Args[1], opts...)
	if err != nil {
		log.Fatalf("failed generating migration file: %v", err)
	}
}
